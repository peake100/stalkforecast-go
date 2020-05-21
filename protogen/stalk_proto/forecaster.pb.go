// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: stalk_proto/forecaster.proto

package stalkproto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_stalk_proto_forecaster_proto protoreflect.FileDescriptor

var file_stalk_proto_forecaster_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f,
	0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5d, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x6c, 0x6b, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0e,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x31, 0x30, 0x30, 0x2f,
	0x73, 0x74, 0x61, 0x6c, 0x6b, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d,
	0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6c, 0x6b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_stalk_proto_forecaster_proto_goTypes = []interface{}{
	(*Ticker)(nil),   // 0: proto.Ticker
	(*Forecast)(nil), // 1: proto.Forecast
}
var file_stalk_proto_forecaster_proto_depIdxs = []int32{
	0, // 0: proto.StalkForecaster.ForecastPrices:input_type -> proto.Ticker
	1, // 1: proto.StalkForecaster.ForecastPrices:output_type -> proto.Forecast
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stalk_proto_forecaster_proto_init() }
func file_stalk_proto_forecaster_proto_init() {
	if File_stalk_proto_forecaster_proto != nil {
		return
	}
	file_stalk_proto_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stalk_proto_forecaster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stalk_proto_forecaster_proto_goTypes,
		DependencyIndexes: file_stalk_proto_forecaster_proto_depIdxs,
	}.Build()
	File_stalk_proto_forecaster_proto = out.File
	file_stalk_proto_forecaster_proto_rawDesc = nil
	file_stalk_proto_forecaster_proto_goTypes = nil
	file_stalk_proto_forecaster_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StalkForecasterClient is the client API for StalkForecaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StalkForecasterClient interface {
	// Returns a turnip price forecast for supplied price ticker.
	ForecastPrices(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Forecast, error)
}

type stalkForecasterClient struct {
	cc grpc.ClientConnInterface
}

func NewStalkForecasterClient(cc grpc.ClientConnInterface) StalkForecasterClient {
	return &stalkForecasterClient{cc}
}

func (c *stalkForecasterClient) ForecastPrices(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Forecast, error) {
	out := new(Forecast)
	err := c.cc.Invoke(ctx, "/proto.StalkForecaster/ForecastPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StalkForecasterServer is the server API for StalkForecaster service.
type StalkForecasterServer interface {
	// Returns a turnip price forecast for supplied price ticker.
	ForecastPrices(context.Context, *Ticker) (*Forecast, error)
}

// UnimplementedStalkForecasterServer can be embedded to have forward compatible implementations.
type UnimplementedStalkForecasterServer struct {
}

func (*UnimplementedStalkForecasterServer) ForecastPrices(context.Context, *Ticker) (*Forecast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForecastPrices not implemented")
}

func RegisterStalkForecasterServer(s *grpc.Server, srv StalkForecasterServer) {
	s.RegisterService(&_StalkForecaster_serviceDesc, srv)
}

func _StalkForecaster_ForecastPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StalkForecasterServer).ForecastPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StalkForecaster/ForecastPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StalkForecasterServer).ForecastPrices(ctx, req.(*Ticker))
	}
	return interceptor(ctx, in, info, handler)
}

var _StalkForecaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StalkForecaster",
	HandlerType: (*StalkForecasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForecastPrices",
			Handler:    _StalkForecaster_ForecastPrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stalk_proto/forecaster.proto",
}