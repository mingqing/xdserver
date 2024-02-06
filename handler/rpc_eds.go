package handler

import (
	"context"
	"errors"
	"io"
	"time"

	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	srvendpointv3 "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/mingqing/xdserver/modeler"
	"google.golang.org/protobuf/types/known/anypb"
)

// FetchEndpoints xx
func (m *Microservice) FetchEndpoints(ctx context.Context, req *discoveryv3.DiscoveryRequest) (*discoveryv3.DiscoveryResponse, error) {
	result := &discoveryv3.DiscoveryResponse{
		Resources: make([]*any.Any, 0),
	}

	ds := modeler.DSCluster{}
	cla := ds.DemoClusterLoadAssignment()

	any1, err := anypb.New(cla)
	if err != nil {
		return result, err
	}

	result.Resources = append(result.Resources, any1)

	m.logger.Infof("node id: %v", req.Node.Id)

	return result, nil
}

// StreamEndpoints xx
// 一般不建议自定义实现服务端，而是复用：github.com/envoyproxy/go-control-plane/pkg/server
func (m *Microservice) StreamEndpoints(stream srvendpointv3.EndpointDiscoveryService_StreamEndpointsServer) error {
	m.logger.Infof("stream endpoints")

	result := &discoveryv3.DiscoveryResponse{
		VersionInfo: "1",
		Resources:   make([]*any.Any, 0),
		// 必须设置且与 `Resources[_]` 中的 type 一致
		TypeUrl: resource.EndpointType,
		Nonce:   "1",
	}

	ds := modeler.DSCluster{}
	cla := ds.DemoClusterLoadAssignment()

	any1, _ := anypb.New(cla)
	result.Resources = append(result.Resources, any1)

	for {
		r, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				m.logger.Infof("stream close")
			} else {
				m.logger.Errorf("stream err: %v", err)
			}

			break
		}

		// m.logger.Infof("node id: %v", r.Node.Id)

		err = stream.Send(result)
		if err != nil {
			m.logger.Errorf("stream send err: %v", err)
		}

		x := jsonpb.Marshaler{}
		rawData, _ := x.MarshalToString(result)
		m.logger.Debugf("%v", rawData)

		m.logger.Infof("node id: %v, type_url: %v, version: %v, nonce: %v",
			r.Node.Id, r.TypeUrl, r.VersionInfo, r.ResourceNames)

		time.Sleep(3600 * time.Second)
	}

	return nil
}

// DeltaEndpoints xx
func (m *Microservice) DeltaEndpoints(stream srvendpointv3.EndpointDiscoveryService_DeltaEndpointsServer) error {
	m.logger.Infof("delta endpoints")

	return nil
}
