// Code generated by protoc-gen-go. DO NOT EDIT.
// source: q.proto

/*
Package qubit is a generated protocol buffer package.

It is generated from these files:
	q.proto

It has these top-level messages:
	Q
	QsimRequest
	ApplyRequest
	ControlledRequest
	QsimResponse
*/
package qubit

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import qubit1 "."
import math1 "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Q struct {
	Qubit *qubit1.Qubit `protobuf:"bytes,1,opt,name=qubit" json:"qubit,omitempty"`
}

func (m *Q) Reset()                    { *m = Q{} }
func (m *Q) String() string            { return proto.CompactTextString(m) }
func (*Q) ProtoMessage()               {}
func (*Q) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Q) GetQubit() *qubit1.Qubit {
	if m != nil {
		return m.Qubit
	}
	return nil
}

type QsimRequest struct {
	Q []*Q `protobuf:"bytes,1,rep,name=q" json:"q,omitempty"`
}

func (m *QsimRequest) Reset()                    { *m = QsimRequest{} }
func (m *QsimRequest) String() string            { return proto.CompactTextString(m) }
func (*QsimRequest) ProtoMessage()               {}
func (*QsimRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QsimRequest) GetQ() []*Q {
	if m != nil {
		return m.Q
	}
	return nil
}

type ApplyRequest struct {
	M *math1.Matrix `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	Q []*Q          `protobuf:"bytes,2,rep,name=q" json:"q,omitempty"`
}

func (m *ApplyRequest) Reset()                    { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()               {}
func (*ApplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplyRequest) GetM() *math1.Matrix {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *ApplyRequest) GetQ() []*Q {
	if m != nil {
		return m.Q
	}
	return nil
}

type ControlledRequest struct {
	Control []*Q  `protobuf:"bytes,1,rep,name=control" json:"control,omitempty"`
	Target  *Q    `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	K       int64 `protobuf:"varint,3,opt,name=k" json:"k,omitempty"`
}

func (m *ControlledRequest) Reset()                    { *m = ControlledRequest{} }
func (m *ControlledRequest) String() string            { return proto.CompactTextString(m) }
func (*ControlledRequest) ProtoMessage()               {}
func (*ControlledRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ControlledRequest) GetControl() []*Q {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *ControlledRequest) GetTarget() *Q {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ControlledRequest) GetK() int64 {
	if m != nil {
		return m.K
	}
	return 0
}

type QsimResponse struct {
	Q *Q `protobuf:"bytes,1,opt,name=q" json:"q,omitempty"`
}

func (m *QsimResponse) Reset()                    { *m = QsimResponse{} }
func (m *QsimResponse) String() string            { return proto.CompactTextString(m) }
func (*QsimResponse) ProtoMessage()               {}
func (*QsimResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QsimResponse) GetQ() *Q {
	if m != nil {
		return m.Q
	}
	return nil
}

func init() {
	proto.RegisterType((*Q)(nil), "qubit.Q")
	proto.RegisterType((*QsimRequest)(nil), "qubit.QsimRequest")
	proto.RegisterType((*ApplyRequest)(nil), "qubit.ApplyRequest")
	proto.RegisterType((*ControlledRequest)(nil), "qubit.ControlledRequest")
	proto.RegisterType((*QsimResponse)(nil), "qubit.QsimResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Qsim service

type QsimClient interface {
	H(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	X(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	Y(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	Z(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	S(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	T(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	ControlledR(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	ControlledZ(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	ControlledNOT(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error)
	QFT(ctx context.Context, in *qubit1.Empty, opts ...grpc.CallOption) (*QsimResponse, error)
	InverseQFT(ctx context.Context, in *qubit1.Empty, opts ...grpc.CallOption) (*QsimResponse, error)
}

type qsimClient struct {
	cc *grpc.ClientConn
}

func NewQsimClient(cc *grpc.ClientConn) QsimClient {
	return &qsimClient{cc}
}

func (c *qsimClient) H(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/H", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) X(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/X", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) Y(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/Y", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) Z(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/Z", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) S(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/S", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) T(ctx context.Context, in *QsimRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/T", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) ControlledR(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/ControlledR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) ControlledZ(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/ControlledZ", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) ControlledNOT(ctx context.Context, in *ControlledRequest, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/ControlledNOT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) QFT(ctx context.Context, in *qubit1.Empty, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/QFT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qsimClient) InverseQFT(ctx context.Context, in *qubit1.Empty, opts ...grpc.CallOption) (*QsimResponse, error) {
	out := new(QsimResponse)
	err := grpc.Invoke(ctx, "/qubit.Qsim/InverseQFT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Qsim service

type QsimServer interface {
	H(context.Context, *QsimRequest) (*QsimResponse, error)
	X(context.Context, *QsimRequest) (*QsimResponse, error)
	Y(context.Context, *QsimRequest) (*QsimResponse, error)
	Z(context.Context, *QsimRequest) (*QsimResponse, error)
	S(context.Context, *QsimRequest) (*QsimResponse, error)
	T(context.Context, *QsimRequest) (*QsimResponse, error)
	Apply(context.Context, *ApplyRequest) (*QsimResponse, error)
	ControlledR(context.Context, *ControlledRequest) (*QsimResponse, error)
	ControlledZ(context.Context, *ControlledRequest) (*QsimResponse, error)
	ControlledNOT(context.Context, *ControlledRequest) (*QsimResponse, error)
	QFT(context.Context, *qubit1.Empty) (*QsimResponse, error)
	InverseQFT(context.Context, *qubit1.Empty) (*QsimResponse, error)
}

func RegisterQsimServer(s *grpc.Server, srv QsimServer) {
	s.RegisterService(&_Qsim_serviceDesc, srv)
}

func _Qsim_H_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).H(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/H",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).H(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_X_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).X(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/X",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).X(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_Y_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).Y(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/Y",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).Y(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_Z_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).Z(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/Z",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).Z(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/S",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).S(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_T_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QsimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).T(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/T",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).T(ctx, req.(*QsimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_ControlledR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).ControlledR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/ControlledR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).ControlledR(ctx, req.(*ControlledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_ControlledZ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).ControlledZ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/ControlledZ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).ControlledZ(ctx, req.(*ControlledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_ControlledNOT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).ControlledNOT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/ControlledNOT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).ControlledNOT(ctx, req.(*ControlledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_QFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(qubit1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).QFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/QFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).QFT(ctx, req.(*qubit1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qsim_InverseQFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(qubit1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QsimServer).InverseQFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qubit.Qsim/InverseQFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QsimServer).InverseQFT(ctx, req.(*qubit1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Qsim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qubit.Qsim",
	HandlerType: (*QsimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "H",
			Handler:    _Qsim_H_Handler,
		},
		{
			MethodName: "X",
			Handler:    _Qsim_X_Handler,
		},
		{
			MethodName: "Y",
			Handler:    _Qsim_Y_Handler,
		},
		{
			MethodName: "Z",
			Handler:    _Qsim_Z_Handler,
		},
		{
			MethodName: "S",
			Handler:    _Qsim_S_Handler,
		},
		{
			MethodName: "T",
			Handler:    _Qsim_T_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Qsim_Apply_Handler,
		},
		{
			MethodName: "ControlledR",
			Handler:    _Qsim_ControlledR_Handler,
		},
		{
			MethodName: "ControlledZ",
			Handler:    _Qsim_ControlledZ_Handler,
		},
		{
			MethodName: "ControlledNOT",
			Handler:    _Qsim_ControlledNOT_Handler,
		},
		{
			MethodName: "QFT",
			Handler:    _Qsim_QFT_Handler,
		},
		{
			MethodName: "InverseQFT",
			Handler:    _Qsim_InverseQFT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "q.proto",
}

func init() { proto.RegisterFile("q.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcb, 0x4f, 0xf2, 0x40,
	0x14, 0xc5, 0xbf, 0x4b, 0x79, 0x7c, 0xb9, 0xad, 0x31, 0x8c, 0x89, 0x69, 0xea, 0xa6, 0x69, 0xa2,
	0x12, 0x17, 0x25, 0x82, 0x6b, 0xa3, 0x18, 0x8d, 0x2e, 0x7c, 0xb4, 0xb2, 0x50, 0x76, 0x05, 0x27,
	0xd0, 0x40, 0xe9, 0x63, 0x06, 0x23, 0xfe, 0xed, 0x2e, 0x4c, 0x3b, 0x8c, 0xa0, 0x84, 0x84, 0xd9,
	0x4c, 0x93, 0x73, 0x7f, 0xe7, 0xdc, 0xce, 0xe3, 0x62, 0x2d, 0x75, 0x93, 0x2c, 0xe6, 0x31, 0xa9,
	0xa4, 0xb3, 0x7e, 0xc8, 0x2d, 0xbd, 0xf8, 0x08, 0xcd, 0xda, 0x8d, 0x02, 0x3e, 0x6a, 0xe6, 0x8b,
	0x10, 0x9c, 0x63, 0x04, 0x8f, 0x38, 0x28, 0x58, 0x13, 0x6c, 0x68, 0xe8, 0x2d, 0xc3, 0x15, 0x16,
	0x2f, 0x5f, 0x7d, 0x51, 0x72, 0x0e, 0x51, 0xf7, 0x58, 0x18, 0xf9, 0x34, 0x9d, 0x51, 0xc6, 0xc9,
	0x3e, 0x42, 0x6a, 0x82, 0xad, 0x35, 0xf4, 0xd6, 0x7f, 0x89, 0xfb, 0x90, 0x3a, 0x1d, 0x34, 0x2e,
	0x93, 0x64, 0x32, 0x97, 0x9c, 0x85, 0x10, 0xfd, 0xc4, 0x16, 0x7d, 0xef, 0x03, 0x9e, 0x85, 0x1f,
	0x3e, 0x44, 0x22, 0xa3, 0xb4, 0x9e, 0x31, 0xc4, 0xfa, 0x55, 0x3c, 0xe5, 0x59, 0x3c, 0x99, 0xd0,
	0x37, 0x19, 0xe4, 0x60, 0x6d, 0x20, 0xc4, 0xb5, 0xb6, 0xb2, 0x40, 0x6c, 0xac, 0xf2, 0x20, 0x1b,
	0x52, 0x6e, 0x96, 0x8a, 0x8e, 0x4b, 0x64, 0xa1, 0x13, 0x03, 0x61, 0x6c, 0x6a, 0x36, 0x34, 0x34,
	0x1f, 0xc6, 0xce, 0x11, 0x1a, 0x62, 0x4f, 0x2c, 0x89, 0xa7, 0x8c, 0xca, 0x4d, 0xc1, 0x9f, 0x1f,
	0x6a, 0x7d, 0x95, 0xb1, 0x9c, 0x83, 0xc4, 0x45, 0xb8, 0x25, 0x44, 0x96, 0x96, 0xc7, 0x61, 0xed,
	0xfd, 0xd2, 0x44, 0x9c, 0xf3, 0x2f, 0xe7, 0x5f, 0x14, 0xf9, 0x57, 0x45, 0xbe, 0xa7, 0xc8, 0x3f,
	0x2b, 0xf2, 0x5d, 0x15, 0xbe, 0x8d, 0x95, 0xe2, 0xf6, 0x89, 0xac, 0xaf, 0xbe, 0x85, 0x4d, 0xa6,
	0x73, 0xd4, 0x57, 0xae, 0x9b, 0x98, 0x0b, 0x6a, 0xed, 0x09, 0x6c, 0xe5, 0xef, 0xa9, 0xfb, 0x2f,
	0x70, 0x67, 0xc9, 0x3e, 0x3c, 0x76, 0xd5, 0x13, 0x4e, 0x50, 0xf3, 0x6e, 0xba, 0x44, 0xce, 0xcd,
	0x75, 0x94, 0xf0, 0xf9, 0x26, 0xf6, 0x14, 0xf1, 0x6e, 0xfa, 0x4e, 0x33, 0x46, 0xb7, 0xb5, 0x74,
	0x0e, 0xb0, 0x3e, 0x88, 0x23, 0x77, 0x18, 0xf2, 0xd1, 0xac, 0xdf, 0x64, 0xc1, 0xe8, 0x8c, 0x7e,
	0x76, 0xca, 0x29, 0x0b, 0xa3, 0x27, 0xe8, 0x57, 0x8b, 0x39, 0x6e, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x59, 0x2c, 0x34, 0xf7, 0x03, 0x00, 0x00,
}
