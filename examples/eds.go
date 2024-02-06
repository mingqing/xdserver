package main

import (
	"context"
	"fmt"
	endpointservice "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	xds "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
	"net"
)

type EnvoyCluster struct {
	name      string
	port      uint32
	endpoints []string
}

func main() {
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":10081")

	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, nil)
	endpointservice.RegisterEndpointDiscoveryServiceServer(grpcServer, server)

	IDs := snapshotCache.GetStatusKeys()
	for _, id := range IDs {
		fmt.Println("id: ", id)
	}

	edsServiceData := map[string]*EnvoyCluster{}
	edsEndpoints := make([]types.Resource, len(edsServiceData))

	snapshot, err := cache.NewSnapshot("v1", map[resource.Type][]types.Resource{
		resource.EndpointType: edsEndpoints,
	})
	if err != nil {
		fmt.Println("cache new snapshot err: ", err)
		return
	}

	err = snapshotCache.SetSnapshot(context.Background(), "node1", snapshot)
	if err != nil {
		fmt.Println("snapshot cache set err: ", err)
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("%v", err)
	}
}
