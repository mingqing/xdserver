package main

import (
	"context"
	"fmt"
	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	srvendpointv3 "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/health"
)

var serviceConfig = "{\"loadBalancingPolicy\": \"round_robin\",\"healthCheckConfig\": {\"serviceName\": \"\"}}"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var err error
	cc, err := grpcConn(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = grpcEDS(ctx, cc); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("grpc eds check ok")
	}
}

func grpcConn(ctx context.Context) (*grpc.ClientConn, error) {
	// TODO; 仅用于测试，配置为该服务对外的地址
	cc, err := grpc.DialContext(ctx, "127.0.0.1:10081",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig))
	if err != nil {
		return nil, err
	}

	return cc, err
}

func grpcEDS(ctx context.Context, cc *grpc.ClientConn) error {
	c := srvendpointv3.NewEndpointDiscoveryServiceClient(cc)
	esc, err := c.StreamEndpoints(ctx)
	if err != nil {
		return err
	}

	md, err := esc.Header()
	if err != nil {
		return err
	}

	for k, v := range md {
		fmt.Println("k:", k, "v:", v)
	}

	/*
		x, err := esc.Recv()
		if err != nil {
			return err
		}
	*/

	req := &discoveryv3.DiscoveryRequest{}
	if err := esc.Send(req); err != nil {
		return fmt.Errorf("set dis request err: %v", err)
	}

	return nil
}
