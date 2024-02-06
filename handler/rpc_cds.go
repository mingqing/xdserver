package handler

import (
	"context"

	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/mingqing/xdserver/modeler"
)

// FetchClusters xx
func (m *Microservice) FetchClusters(ctx context.Context, req *discoveryv3.DiscoveryRequest) (*discoveryv3.DiscoveryResponse, error) {
	result := &discoveryv3.DiscoveryResponse{
		Resources: make([]*any.Any, 0),
	}

	ds := modeler.DSCluster{}
	cls := ds.DemoCluster()

	for _, v := range cls {
		any1, err := anypb.New(v)
		if err != nil {
			return result, err
		}

		result.Resources = append(result.Resources, any1)
	}

	m.logger.Infof("node id: %v", req.Node.Id)

	return result, nil
}
