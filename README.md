# xdserver.v1.mingqing

## Description

// TODO(user): Add simple overview of use/purpose

用于 xds 服务的示例服务端。

## Getting Started

// TODO(user): Add quick start

注意 `scripts/genproto.sh` 以来的 proto 文件。

```text
protoc \
    -I ./ \
    -I /Users/coding/go/src/github.com/cncf/xds \
    -I /Users/coding/go/src/github.com/envoyproxy/envoy/api \
    -I /Users/coding/go/src/github.com/bufbuild/protoc-gen-validate \
```

1. 通过 rest http api
2. 通过 grpc 继承 go-control
3. 通过 grpc 自定义

## EDS 实现

###  方式一：使用 go-control-plane 服务

```text
pkg.go.dev/github.com/envoyproxy/go-control-plane/pkg/server/v3#NewServer
```

在 `handler/register.go` 新增以下内容：

```text
// Register 用于服务启动前环境准备
func (m *Microservice) Register(ctx context.Context) error {
    ......
    
	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, nil)
	endpointservice.RegisterEndpointDiscoveryServiceServer(m.server.Server(), server)
	
	// 具体服务发现逻辑
	go testSnapshot(snapshotCache)
	
	......
}	
```

### 方式二：使用自定义 HTTP 服务

在 `api/mingqing/xdserver/v1/microservice.proto` 新增以下内容：

```text
// 该微服务支持的 RPC 方法定义
service MingqingXdserver {
  ......
  
  // eds rest 接口定义
  rpc FetchEndpoints(envoy.service.discovery.v3.DiscoveryRequest) returns (envoy.service.discovery.v3.DiscoveryResponse) {}
  
  ......
}
```

定义 http rest 接口，在 `api/mingqing/xdserver/v1/microservice.gateway.yaml`：

```text
type: google.api.Service
config_version: 3

http:
  rules:
  ......
  - selector: default.api.mingqing.xdserver.v1.MingqingXdserver.FetchEndpoints
    post: "/api/v3/discovery:endpoints"
    body: "*"
  ......
```

实现具体服务发现逻辑，在 `handler/rpc_eds.go`：

```text
// FetchEndpoints xx
func (m *Microservice) FetchEndpoints(ctx context.Context, req *discoveryv3.DiscoveryRequest) (*discoveryv3.DiscoveryResponse, error) {
    result := &discoveryv3.DiscoveryResponse{
        Resources: make([]*any.Any, 0),
    }
    
    ......

    return result, nil
}
```

### 方式三：使用自定义 gRPC 服务

在 `handler/register.go` 新增以下内容：

```text
// Register 用于服务启动前环境准备
func (m *Microservice) Register(ctx context.Context) error {
    ......
    
	endpointservice.RegisterEndpointDiscoveryServiceServer(m.server.Server(), m)
	
	......
}	
```

注意代码不要与方案一重复。

在 `handler/rpc_eds.go` 添加以下代码，以实现 eds 服务定义：

```text
func (m *Microservice) StreamEndpoints(stream srvendpointv3.EndpointDiscoveryService_StreamEndpointsServer) error {
    ......
    
    return nil
}

func (m *Microservice) DeltaEndpoints(stream srvendpointv3.EndpointDiscoveryService_DeltaEndpointsServer) error {
    ......

    return nil
}
```
