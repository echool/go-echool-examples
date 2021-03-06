package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/echool/go-echool-examples/proto/message"

	"github.com/echool/go-echool"
	"github.com/echool/go-echool-examples/logger"
	"github.com/echool/go-echool/config"
	sd "github.com/echool/go-echool/discovery/consul"
	"google.golang.org/grpc/credentials"
)

var (
	serviceName = "message"
	serviceAddr = ":6789"
	tlsCertfile = "../../conf/server.crt"
	tlsKeyfile  = "../../conf/server.key"

	consulAddr  = "192.168.0.213:8500"
	consulToken = ""
)

func main() {

	creds, err := credentials.NewServerTLSFromFile(tlsCertfile, tlsKeyfile)
	if err != nil {
		panic(err)
	}

	conf := config.ConsulSrvConf{
		ServiceName:    serviceName,
		ServiceAddress: serviceAddr,
		Address:        consulAddr,
		Token:          consulToken,
	}
	register, err := sd.NewRegister(conf)
	if err != nil {
		panic(err)
	}

	grpcServer := echool.NewServer(register, func(server *grpc.Server) {
		message.RegisterMessSrvServer(server, new(MessSrv))
	})

	log.Fatal(grpcServer.Run(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(UnaryServerInterceptor),
	))
}

// UnaryServerInterceptor 自定义拦截器
// 或者使用: https://github.com/grpc-ecosystem/go-grpc-middleware
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)

	// 记录每个服务方法的执行耗时
	logger.Log.Info(
		fmt.Sprintf("%16s  %s",
			time.Since(start),
			info.FullMethod),
	)
	return resp, err
}

type MessSrv struct{}

var users []*message.UserInfo

func (*MessSrv) Send(ctx context.Context, req *message.SendRequest) (*message.SendReply, error) {
	time.Sleep(time.Millisecond * 1200) //mock timeout

	if req.Userinfo.Userid <= 0 {
		return nil, fmt.Errorf("userid必须大于0")
	}

	users = append(users, req.Userinfo)

	resp := new(message.SendReply)
	resp.Result = true
	return resp, nil
}

func (*MessSrv) Receive(ctx context.Context, req *message.Empty) (*message.ReceiveReply, error) {
	resp := new(message.ReceiveReply)
	resp.Userinfo = users
	return resp, nil
}
