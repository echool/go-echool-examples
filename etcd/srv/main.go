package main

import (
	"context"
	"log"
	"strconv"

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
	for i := 1; i <= 3; i++ {
		go func(n int) {
			startServer(n)
		}(i)
	}
	select {}
}

func startServer(count int) {
	serviceAddr := ":2000" + strconv.Itoa(count)
	conf := config.EtcdSrvConf{
		ServiceName:    serviceName,
		ServiceAddress: serviceAddr,
		Endpoints:      etcdAddrs,
		Auth:           config.Auth{},
	}

	demo := &RegionHandlerServer{Address: serviceAddr}
	register, err := sd.NewRegister(conf)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := echool.NewServer(register, func(server *grpc.Server) {
		knowing.RegisterRegionHandlerServer(server, demo)
	})

	log.Fatal(grpcServer.Run())
}

type RegionHandlerServer struct {
	Address string
}

func (s *RegionHandlerServer) GetListenAudio(ctx context.Context, r *knowing.FindRequest) (*knowing.HasOptionResponse, error) {
	has := []*knowing.HasOption(nil)
	for _, t := range r.Tokens {
		has = append(has, &knowing.HasOption{Token: t + " FROM " + s.Address, Listen: 1})
	}
	res := &knowing.HasOptionResponse{
		Items: has,
	}
	return res, nil
}
