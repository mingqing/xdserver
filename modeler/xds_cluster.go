package modeler

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
)

// DSCluster xx
type DSCluster struct {
}

// DemoClusterLoadAssignment xx
func (x *DSCluster) DemoClusterLoadAssignment() *endpointv3.ClusterLoadAssignment {
	cla := &endpointv3.ClusterLoadAssignment{
		ClusterName: "demo-eds-v1",
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

	return cla
}
