// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.18.1
// source: heartbeat.proto

package heads

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type AckIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AckIn) Reset() {
	*x = AckIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckIn) ProtoMessage() {}

func (x *AckIn) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckIn.ProtoReflect.Descriptor instead.
func (*AckIn) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *AckIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x68, 0x65, 0x61, 0x64, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x05, 0x41, 0x63, 0x6b, 0x49, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x2e, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x49,
	0x6e, 0x1a, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_proto_rawDescData = file_heartbeat_proto_rawDesc
)

func file_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heartbeat_proto_rawDescData)
	})
	return file_heartbeat_proto_rawDescData
}

var file_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heartbeat_proto_goTypes = []interface{}{
	(*AckIn)(nil), // 0: heads.AckIn
	(*Empty)(nil), // 1: heads.Empty
}
var file_heartbeat_proto_depIdxs = []int32{
	0, // 0: heads.Heartbeat.Ack:input_type -> heads.AckIn
	1, // 1: heads.Heartbeat.Ack:output_type -> heads.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heartbeat_proto_init() }
func file_heartbeat_proto_init() {
	if File_heartbeat_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_proto = out.File
	file_heartbeat_proto_rawDesc = nil
	file_heartbeat_proto_goTypes = nil
	file_heartbeat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	Ack(ctx context.Context, in *AckIn, opts ...grpc.CallOption) (*Empty, error)
}

type heartbeatClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatClient(cc grpc.ClientConnInterface) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Ack(ctx context.Context, in *AckIn, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/heads.Heartbeat/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	Ack(context.Context, *AckIn) (*Empty, error)
}

// UnimplementedHeartbeatServer can be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServer struct {
}

func (*UnimplementedHeartbeatServer) Ack(context.Context, *AckIn) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heads.Heartbeat/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Ack(ctx, req.(*AckIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heads.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ack",
			Handler:    _Heartbeat_Ack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}