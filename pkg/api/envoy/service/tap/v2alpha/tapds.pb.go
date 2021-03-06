// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tapds.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/datawire/ambassador/pkg/api/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// [#not-implemented-hide:] A tap resource is essentially a tap configuration with a name
// The filter TapDS config references this name.
type TapResource struct {
	// The name of the tap configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tap config to apply
	Config               *TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eef591ebf2a5317, []int{0}
}
func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return m.Size()
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v2alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/tapds.proto", fileDescriptor_4eef591ebf2a5317)
}

var fileDescriptor_4eef591ebf2a5317 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0xce, 0xd3, 0x30,
	0x14, 0xc5, 0x3f, 0x87, 0x4f, 0x45, 0xb8, 0x43, 0xab, 0x30, 0xd0, 0xa6, 0x25, 0x54, 0xa5, 0x40,
	0xd5, 0xc1, 0x41, 0x61, 0xab, 0x98, 0x42, 0xc5, 0x88, 0xaa, 0x34, 0x0b, 0x13, 0xba, 0x4d, 0x4c,
	0x6b, 0x29, 0xb1, 0x4d, 0xec, 0x44, 0x74, 0x41, 0x88, 0x89, 0x9d, 0xb7, 0xe1, 0x09, 0x3a, 0x82,
	0xe0, 0x01, 0x50, 0xc5, 0x53, 0x30, 0xa1, 0xfc, 0x69, 0xd5, 0x82, 0xc2, 0xc4, 0x66, 0xf9, 0xfc,
	0xce, 0xb9, 0xb9, 0xf1, 0xc1, 0x0f, 0x28, 0xcf, 0xc5, 0xce, 0x51, 0x34, 0xcd, 0x59, 0x48, 0x1d,
	0x0d, 0xd2, 0xc9, 0x5d, 0x88, 0xe5, 0x16, 0x8a, 0x73, 0xa4, 0x88, 0x4c, 0x85, 0x16, 0x66, 0xbf,
	0xc4, 0x48, 0x8d, 0x11, 0x0d, 0x92, 0xd4, 0x98, 0x35, 0xac, 0x12, 0x40, 0x32, 0x27, 0x77, 0x9d,
	0x88, 0xa9, 0x50, 0xe4, 0x34, 0xdd, 0x55, 0x46, 0xeb, 0x61, 0x73, 0x7e, 0x28, 0x92, 0x44, 0xf0,
	0x9a, 0x1b, 0x6e, 0x84, 0xd8, 0xc4, 0xb4, 0x8c, 0x01, 0xce, 0x85, 0x06, 0xcd, 0x04, 0xaf, 0xc7,
	0x5b, 0x77, 0xb3, 0x48, 0xc2, 0xf9, 0xbd, 0xa3, 0x34, 0xe8, 0xec, 0x28, 0xdf, 0xc9, 0x21, 0x66,
	0x11, 0x68, 0xea, 0x1c, 0x0f, 0x95, 0x30, 0xde, 0xe2, 0x76, 0x00, 0xd2, 0xa7, 0x4a, 0x64, 0x69,
	0x48, 0xcd, 0x01, 0xbe, 0xe6, 0x90, 0xd0, 0x1e, 0x1a, 0xa1, 0xe9, 0x2d, 0xef, 0xe6, 0x2f, 0xef,
	0x3a, 0x35, 0x46, 0xc8, 0x2f, 0x2f, 0xcd, 0xa7, 0xb8, 0x15, 0x0a, 0xfe, 0x9a, 0x6d, 0x7a, 0xc6,
	0x08, 0x4d, 0xdb, 0xee, 0x84, 0x34, 0xee, 0x4c, 0x02, 0x90, 0xcf, 0x4a, 0xd6, 0xaf, 0x3d, 0xee,
	0x77, 0x03, 0xdf, 0x0e, 0x40, 0x2e, 0x8e, 0xeb, 0xaf, 0x2a, 0x97, 0xf9, 0x12, 0x77, 0x57, 0x3a,
	0xa5, 0x90, 0x9c, 0x2c, 0xca, 0xb4, 0xeb, 0x64, 0x90, 0x8c, 0xe4, 0x2e, 0x39, 0x79, 0x7c, 0xfa,
	0x26, 0xa3, 0x4a, 0x5b, 0xf7, 0x1a, 0x75, 0x25, 0x05, 0x57, 0x74, 0x7c, 0x35, 0x45, 0x8f, 0x91,
	0xb9, 0xc6, 0x9d, 0x05, 0x8d, 0x35, 0x9c, 0x25, 0xdf, 0xff, 0xc3, 0x59, 0xc8, 0x7f, 0xc5, 0x4f,
	0xfe, 0x0d, 0x5d, 0xcc, 0x78, 0x87, 0x3b, 0xcf, 0xa9, 0x0e, 0xb7, 0xff, 0xf3, 0xeb, 0x67, 0x1f,
	0xbe, 0xfd, 0xfc, 0x64, 0x0c, 0xc6, 0xfd, 0x8b, 0xbe, 0xcc, 0x35, 0xc8, 0x57, 0xd5, 0xcf, 0x54,
	0x25, 0x70, 0x63, 0x8e, 0x66, 0xde, 0x8b, 0xfd, 0xc1, 0x46, 0x5f, 0x0e, 0x36, 0xfa, 0x71, 0xb0,
	0xd1, 0xe7, 0xf7, 0xfb, 0xaf, 0x2d, 0xa3, 0x7b, 0x85, 0x1f, 0x31, 0x51, 0x0d, 0x91, 0xa9, 0x78,
	0xbb, 0x6b, 0x7e, 0x27, 0x0f, 0x07, 0x45, 0x87, 0x97, 0x45, 0x17, 0x96, 0xe8, 0x23, 0x42, 0xeb,
	0x56, 0xd9, 0x8b, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x94, 0x0d, 0xd3, 0xfb, 0xf7, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v2alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(srv TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(srv TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tapds.proto",
}

func (m *TapResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapResource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapResource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTapds(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTapds(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTapds(dAtA []byte, offset int, v uint64) int {
	offset -= sovTapds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TapResource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTapds(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTapds(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTapds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTapds(x uint64) (n int) {
	return sovTapds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TapResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTapds
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TapResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTapds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTapds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTapds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTapds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTapds
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTapds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &TapConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTapds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTapds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTapds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTapds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTapds
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTapds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTapds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTapds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTapds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTapds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTapds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTapds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTapds = fmt.Errorf("proto: unexpected end of group")
)
