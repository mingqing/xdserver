package modeler

import (
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
)

// DSCluster xx
type DSCluster struct {
}

// DemoRoute xx
// envoy.config.route.v3.RouteConfiguration
func (x *DSCluster) DemoRoute() []*routev3.RouteConfiguration {
	result := make([]*routev3.RouteConfiguration, 0)

	r1 := &routev3.RouteConfiguration{
		Name: "test-route-1",
		VirtualHosts: []*routev3.VirtualHost{
			{
				Name:    "host1",
				Domains: []string{"*"},
				Routes: []*routev3.Route{
					{
						Match: &routev3.RouteMatch{
							PathSpecifier: &routev3.RouteMatch_Prefix{
								Prefix: "/",
							},
						},
						Action: &routev3.Route_Route{
							Route: &routev3.RouteAction{
								ClusterSpecifier: &routev3.RouteAction_Cluster{Cluster: "demo-eds-v1"},
							},
						},
					},
				},
			},
		},
	}

	result = append(result, r1)

	return result
}

// DemoCluster xx
// envoy.config.cluster.v3.Cluster
func (x *DSCluster) DemoCluster() []*clusterv3.Cluster {
	result := make([]*clusterv3.Cluster, 0)

	c1 := &clusterv3.Cluster{
		Name:                 "demo-eds-v1",
		DnsLookupFamily:      clusterv3.Cluster_V4_ONLY,
		ClusterDiscoveryType: &clusterv3.Cluster_Type{Type: clusterv3.Cluster_EDS},
		EdsClusterConfig: &clusterv3.Cluster_EdsClusterConfig{
			EdsConfig: &corev3.ConfigSource{
				ResourceApiVersion: corev3.ApiVersion_V3,
				ConfigSourceSpecifier: &corev3.ConfigSource_ApiConfigSource{
					ApiConfigSource: &corev3.ApiConfigSource{
						ApiType:             corev3.ApiConfigSource_REST,
						TransportApiVersion: corev3.ApiVersion_V3,
						ClusterNames:        []string{"xds-cluster"},
					},
				},
			},
		},
	}

	c2 := &clusterv3.Cluster{
		Name:                 "xds-cluster",
		DnsLookupFamily:      clusterv3.Cluster_V4_ONLY,
		ClusterDiscoveryType: &clusterv3.Cluster_Type{Type: clusterv3.Cluster_STATIC},
		LoadAssignment: &endpointv3.ClusterLoadAssignment{
			ClusterName: "xds-cluster",
			Endpoints:   []*endpointv3.LocalityLbEndpoints{},
		},
	}
	c2.LoadAssignment.Endpoints = append(c2.LoadAssignment.Endpoints, &endpointv3.LocalityLbEndpoints{
		LbEndpoints: []*endpointv3.LbEndpoint{{
			HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
				Endpoint: &endpointv3.Endpoint{
					Address: &corev3.Address{
						Address: &corev3.Address_SocketAddress{
							SocketAddress: &corev3.SocketAddress{
								Protocol:      corev3.SocketAddress_TCP,
								Address:       "10.5.17.224",
								PortSpecifier: &corev3.SocketAddress_PortValue{PortValue: 80},
							},
						},
					},
				},
			},
		}},
	})

	result = append(result, c1)

	return result
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
