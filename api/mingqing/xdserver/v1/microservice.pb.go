// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/mingqing/xdserver/v1/microservice.proto

package xdserverv1

import (
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	v1 "github.com/grpc-kit/pkg/api/known/status/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_mingqing_xdserver_v1_microservice_proto protoreflect.FileDescriptor

var file_api_mingqing_xdserver_v1_microservice_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67, 0x2f, 0x78,
	0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x6e, 0x67, 0x71,
	0x69, 0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x23, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x03, 0x0a, 0x10, 0x4d, 0x69, 0x6e,
	0x67, 0x71, 0x69, 0x6e, 0x67, 0x58, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x74, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x30, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x2d, 0x2e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69,
	0x6e, 0x67, 0x2e, 0x78, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69, 0x6e,
	0x67, 0x2e, 0x78, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x0e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2c,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a,
	0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67, 0x2f, 0x78, 0x64, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x71, 0x69, 0x6e, 0x67,
	0x2f, 0x78, 0x64, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x78, 0x64, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_mingqing_xdserver_v1_microservice_proto_goTypes = []interface{}{
	(*v1.HealthCheckRequest)(nil),  // 0: grpc_kit.api.known.status.v1.HealthCheckRequest
	(*DemoRequest)(nil),            // 1: default.api.mingqing.xdserver.v1.DemoRequest
	(*v3.DiscoveryRequest)(nil),    // 2: envoy.service.discovery.v3.DiscoveryRequest
	(*v1.HealthCheckResponse)(nil), // 3: grpc_kit.api.known.status.v1.HealthCheckResponse
	(*DemoResponse)(nil),           // 4: default.api.mingqing.xdserver.v1.DemoResponse
	(*v3.DiscoveryResponse)(nil),   // 5: envoy.service.discovery.v3.DiscoveryResponse
}
var file_api_mingqing_xdserver_v1_microservice_proto_depIdxs = []int32{
	0, // 0: default.api.mingqing.xdserver.v1.MingqingXdserver.HealthCheck:input_type -> grpc_kit.api.known.status.v1.HealthCheckRequest
	1, // 1: default.api.mingqing.xdserver.v1.MingqingXdserver.Demo:input_type -> default.api.mingqing.xdserver.v1.DemoRequest
	2, // 2: default.api.mingqing.xdserver.v1.MingqingXdserver.FetchEndpoints:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	2, // 3: default.api.mingqing.xdserver.v1.MingqingXdserver.StreamEndpoints:input_type -> envoy.service.discovery.v3.DiscoveryRequest
	3, // 4: default.api.mingqing.xdserver.v1.MingqingXdserver.HealthCheck:output_type -> grpc_kit.api.known.status.v1.HealthCheckResponse
	4, // 5: default.api.mingqing.xdserver.v1.MingqingXdserver.Demo:output_type -> default.api.mingqing.xdserver.v1.DemoResponse
	5, // 6: default.api.mingqing.xdserver.v1.MingqingXdserver.FetchEndpoints:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	5, // 7: default.api.mingqing.xdserver.v1.MingqingXdserver.StreamEndpoints:output_type -> envoy.service.discovery.v3.DiscoveryResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_mingqing_xdserver_v1_microservice_proto_init() }
func file_api_mingqing_xdserver_v1_microservice_proto_init() {
	if File_api_mingqing_xdserver_v1_microservice_proto != nil {
		return
	}
	file_api_mingqing_xdserver_v1_demo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_mingqing_xdserver_v1_microservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_mingqing_xdserver_v1_microservice_proto_goTypes,
		DependencyIndexes: file_api_mingqing_xdserver_v1_microservice_proto_depIdxs,
	}.Build()
	File_api_mingqing_xdserver_v1_microservice_proto = out.File
	file_api_mingqing_xdserver_v1_microservice_proto_rawDesc = nil
	file_api_mingqing_xdserver_v1_microservice_proto_goTypes = nil
	file_api_mingqing_xdserver_v1_microservice_proto_depIdxs = nil
}
