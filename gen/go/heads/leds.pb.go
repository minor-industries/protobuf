// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.18.1
// source: leds.proto

package heads

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type RunIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NewStartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=new_start_time,json=newStartTime,proto3" json:"new_start_time,omitempty"`
}

func (x *RunIn) Reset() {
	*x = RunIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunIn) ProtoMessage() {}

func (x *RunIn) ProtoReflect() protoreflect.Message {
	mi := &file_leds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunIn.ProtoReflect.Descriptor instead.
func (*RunIn) Descriptor() ([]byte, []int) {
	return file_leds_proto_rawDescGZIP(), []int{0}
}

func (x *RunIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RunIn) GetNewStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NewStartTime
	}
	return nil
}

type SetScaleIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scale float64 `protobuf:"fixed64,1,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *SetScaleIn) Reset() {
	*x = SetScaleIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScaleIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScaleIn) ProtoMessage() {}

func (x *SetScaleIn) ProtoReflect() protoreflect.Message {
	mi := &file_leds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScaleIn.ProtoReflect.Descriptor instead.
func (*SetScaleIn) Descriptor() ([]byte, []int) {
	return file_leds_proto_rawDescGZIP(), []int{1}
}

func (x *SetScaleIn) GetScale() float64 {
	if x != nil {
		return x.Scale
	}
	return 0
}

var File_leds_proto protoreflect.FileDescriptor

var file_leds_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x05, 0x52, 0x75, 0x6e, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x22, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x32, 0x7f, 0x0a, 0x04, 0x6c, 0x65, 0x64, 0x73, 0x12, 0x21, 0x0a,
	0x03, 0x72, 0x75, 0x6e, 0x12, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x52, 0x75, 0x6e,
	0x49, 0x6e, 0x1a, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x26, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x49, 0x6e, 0x1a, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leds_proto_rawDescOnce sync.Once
	file_leds_proto_rawDescData = file_leds_proto_rawDesc
)

func file_leds_proto_rawDescGZIP() []byte {
	file_leds_proto_rawDescOnce.Do(func() {
		file_leds_proto_rawDescData = protoimpl.X.CompressGZIP(file_leds_proto_rawDescData)
	})
	return file_leds_proto_rawDescData
}

var file_leds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_leds_proto_goTypes = []interface{}{
	(*RunIn)(nil),                 // 0: heads.RunIn
	(*SetScaleIn)(nil),            // 1: heads.SetScaleIn
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Empty)(nil),                 // 3: heads.Empty
	(*Event)(nil),                 // 4: heads.Event
}
var file_leds_proto_depIdxs = []int32{
	2, // 0: heads.RunIn.new_start_time:type_name -> google.protobuf.Timestamp
	0, // 1: heads.leds.run:input_type -> heads.RunIn
	3, // 2: heads.leds.events:input_type -> heads.Empty
	1, // 3: heads.leds.set_scale:input_type -> heads.SetScaleIn
	3, // 4: heads.leds.run:output_type -> heads.Empty
	4, // 5: heads.leds.events:output_type -> heads.Event
	3, // 6: heads.leds.set_scale:output_type -> heads.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_leds_proto_init() }
func file_leds_proto_init() {
	if File_leds_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_leds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunIn); i {
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
		file_leds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScaleIn); i {
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
			RawDescriptor: file_leds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_leds_proto_goTypes,
		DependencyIndexes: file_leds_proto_depIdxs,
		MessageInfos:      file_leds_proto_msgTypes,
	}.Build()
	File_leds_proto = out.File
	file_leds_proto_rawDesc = nil
	file_leds_proto_goTypes = nil
	file_leds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LedsClient is the client API for Leds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LedsClient interface {
	Run(ctx context.Context, in *RunIn, opts ...grpc.CallOption) (*Empty, error)
	Events(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Leds_EventsClient, error)
	SetScale(ctx context.Context, in *SetScaleIn, opts ...grpc.CallOption) (*Empty, error)
}

type ledsClient struct {
	cc grpc.ClientConnInterface
}

func NewLedsClient(cc grpc.ClientConnInterface) LedsClient {
	return &ledsClient{cc}
}

func (c *ledsClient) Run(ctx context.Context, in *RunIn, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/heads.leds/run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ledsClient) Events(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Leds_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Leds_serviceDesc.Streams[0], "/heads.leds/events", opts...)
	if err != nil {
		return nil, err
	}
	x := &ledsEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Leds_EventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type ledsEventsClient struct {
	grpc.ClientStream
}

func (x *ledsEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ledsClient) SetScale(ctx context.Context, in *SetScaleIn, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/heads.leds/set_scale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LedsServer is the server API for Leds service.
type LedsServer interface {
	Run(context.Context, *RunIn) (*Empty, error)
	Events(*Empty, Leds_EventsServer) error
	SetScale(context.Context, *SetScaleIn) (*Empty, error)
}

// UnimplementedLedsServer can be embedded to have forward compatible implementations.
type UnimplementedLedsServer struct {
}

func (*UnimplementedLedsServer) Run(context.Context, *RunIn) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedLedsServer) Events(*Empty, Leds_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (*UnimplementedLedsServer) SetScale(context.Context, *SetScaleIn) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScale not implemented")
}

func RegisterLedsServer(s *grpc.Server, srv LedsServer) {
	s.RegisterService(&_Leds_serviceDesc, srv)
}

func _Leds_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedsServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heads.leds/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedsServer).Run(ctx, req.(*RunIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leds_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LedsServer).Events(m, &ledsEventsServer{stream})
}

type Leds_EventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type ledsEventsServer struct {
	grpc.ServerStream
}

func (x *ledsEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Leds_SetScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetScaleIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LedsServer).SetScale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heads.leds/SetScale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LedsServer).SetScale(ctx, req.(*SetScaleIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Leds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heads.leds",
	HandlerType: (*LedsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "run",
			Handler:    _Leds_Run_Handler,
		},
		{
			MethodName: "set_scale",
			Handler:    _Leds_SetScale_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "events",
			Handler:       _Leds_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "leds.proto",
}
