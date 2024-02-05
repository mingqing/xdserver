package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/anypb"

	pb "github.com/mingqing/xdserver/api/mingqing/xdserver/v1"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

// FetchEndpoints xx
func (m *Microservice) FetchEndpoints(ctx context.Context, req *discoveryv3.DiscoveryRequest) (*discoveryv3.DiscoveryResponse, error) {
	result := &discoveryv3.DiscoveryResponse{
		Resources: make([]*any.Any, 0),
	}

	cla := &endpointv3.ClusterLoadAssignment{
		ClusterName: "oneops-syncds-v1",
		Endpoints:   make([]*endpointv3.LocalityLbEndpoints, 0),
	}

	edp1 := &endpointv3.LbEndpoint{
		HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
			Endpoint: &endpointv3.Endpoint{Address: &corev3.Address{
				Address: &corev3.Address_SocketAddress{
					SocketAddress: &corev3.SocketAddress{
						Address:       "192.168.0.1",
						PortSpecifier: &corev3.SocketAddress_PortValue{PortValue: 8000},
					},
				},
			}},
		},
	}

	edp2 := &endpointv3.LbEndpoint{
		HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
			Endpoint: &endpointv3.Endpoint{Address: &corev3.Address{
				Address: &corev3.Address_SocketAddress{
					SocketAddress: &corev3.SocketAddress{
						Address:       "192.168.0.1",
						PortSpecifier: &corev3.SocketAddress_PortValue{PortValue: 8080},
					},
				},
			}},
		},
	}

	cla.Endpoints = append(cla.Endpoints, &endpointv3.LocalityLbEndpoints{
		LbEndpoints: []*endpointv3.LbEndpoint{edp1, edp2},
	})

	any1, err := anypb.New(cla)
	if err != nil {
		return result, err
	}

	result.Resources = append(result.Resources, any1)

	m.logger.Infof("node id: %v", req.Node.Id)

	return result, nil
}

// StreamEndpoints xx
func (m *Microservice) StreamEndpoints(stream pb.MingqingXdserver_StreamEndpointsServer) error {
	r, err := stream.Recv()
	if err != nil {
		return err
	}

	m.logger.Infof("node id: %v", r.Node.Id)

	return nil
}
