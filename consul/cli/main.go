package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/echool/go-echool-examples/proto/message"

	"google.golang.org/grpc"

	"github.com/echool/go-echool"
	"github.com/echool/go-echool/config"
	sd "github.com/echool/go-echool/discovery/consul"
	"google.golang.org/grpc/credentials"
)

var (
	serviceName = "message"
	tlsCertfile = "../../conf/server.crt"

	consulAddr  = "192.168.0.213:8500"
	consulToken = ""
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(tlsCertfile, serviceName)
	if err != nil {
		log.Fatal(err)
	}

	conf := config.ConsulCliConf{
		ServiceName: serviceName,
		Address:     consulAddr,
		Token:       consulToken,
	}
	target := sd.NewResolver(conf)

	grpcClient := echool.NewClient(target)
	conn, err := grpcClient.GetConnection(grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	cli := message.NewMessSrvClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// call Send
		sendRs, err := cli.Send(ctx, &message.SendRequest{
			Userinfo: &message.UserInfo{
				Userid:   1,
				Username: "ferrmin",
				Intro:    strconv.FormatInt(time.Now().UnixNano(), 10),
			},
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(sendRs.Result)

		// call Receive
		recRs, err := cli.Receive(ctx, &message.Empty{})
		if err != nil {
			log.Fatalln(err)
		}

		for _, item := range recRs.Userinfo {
			fmt.Println(item.Userid, item.Username, item.Intro)
		}

		time.Sleep(time.Second * 3)
	}
}
