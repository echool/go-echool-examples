package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/echool/go-echool-examples/proto/knowing"

	"github.com/echool/go-echool"
	"github.com/echool/go-echool/config"
	sd "github.com/echool/go-echool/discovery/etcd"
)

var (
	serviceName = "knowing"
	etcdAddrs   = []string{"192.168.0.213:2379"}
)

func main() {
	conf := config.EtcdCliConf{
		ServiceName: serviceName,
		Endpoints:   etcdAddrs,
		Auth:        config.Auth{},
	}
	target := sd.NewResolver(conf)

	grpcClient := echool.NewClient(target)
	conn, err := grpcClient.GetConnection(grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	regionHandlerClient := knowing.NewRegionHandlerClient(conn)

	for {
		res, err := regionHandlerClient.GetListenAudio(
			context.Background(),
			&knowing.FindRequest{Tokens: []string{"a_"}},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
		time.Sleep(2 * time.Second) // 每两次请求一次knowing的rpc服务
	}
}
