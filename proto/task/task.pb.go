// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/task/task.proto

package task

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TaskResponse_Code int32

const (
	TaskResponse_failed   TaskResponse_Code = 0
	TaskResponse_success  TaskResponse_Code = 1
	TaskResponse_excepted TaskResponse_Code = 2
)

var TaskResponse_Code_name = map[int32]string{
	0: "failed",
	1: "success",
	2: "excepted",
}

var TaskResponse_Code_value = map[string]int32{
	"failed":   0,
	"success":  1,
	"excepted": 2,
}

func (x TaskResponse_Code) String() string {
	return proto.EnumName(TaskResponse_Code_name, int32(x))
}

func (TaskResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f05afca300cf24a4, []int{3, 0}
}

type PlanRequest struct {
	PlanID               string   `protobuf:"bytes,1,opt,name=PlanID,proto3" json:"PlanID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanRequest) Reset()         { *m = PlanRequest{} }
func (m *PlanRequest) String() string { return proto.CompactTextString(m) }
func (*PlanRequest) ProtoMessage()    {}
func (*PlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05afca300cf24a4, []int{0}
}

func (m *PlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanRequest.Unmarshal(m, b)
}
func (m *PlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanRequest.Marshal(b, m, deterministic)
}
func (m *PlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanRequest.Merge(m, src)
}
func (m *PlanRequest) XXX_Size() int {
	return xxx_messageInfo_PlanRequest.Size(m)
}
func (m *PlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlanRequest proto.InternalMessageInfo

func (m *PlanRequest) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

type PlanResponse struct {
	PlanID               string   `protobuf:"bytes,1,opt,name=PlanID,proto3" json:"PlanID,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanResponse) Reset()         { *m = PlanResponse{} }
func (m *PlanResponse) String() string { return proto.CompactTextString(m) }
func (*PlanResponse) ProtoMessage()    {}
func (*PlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05afca300cf24a4, []int{1}
}

func (m *PlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanResponse.Unmarshal(m, b)
}
func (m *PlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanResponse.Marshal(b, m, deterministic)
}
func (m *PlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanResponse.Merge(m, src)
}
func (m *PlanResponse) XXX_Size() int {
	return xxx_messageInfo_PlanResponse.Size(m)
}
func (m *PlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlanResponse proto.InternalMessageInfo

func (m *PlanResponse) GetPlanID() string {
	if m != nil {
		return m.PlanID
	}
	return ""
}

func (m *PlanResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type TaskRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	NameReq              string   `protobuf:"bytes,2,opt,name=nameReq,proto3" json:"nameReq,omitempty"`
	AgeReq               int32    `protobuf:"varint,3,opt,name=ageReq,proto3" json:"ageReq,omitempty"`
	LocalTime            string   `protobuf:"bytes,4,opt,name=localTime,proto3" json:"localTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05afca300cf24a4, []int{2}
}

func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (m *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(m, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *TaskRequest) GetNameReq() string {
	if m != nil {
		return m.NameReq
	}
	return ""
}

func (m *TaskRequest) GetAgeReq() int32 {
	if m != nil {
		return m.AgeReq
	}
	return 0
}

func (m *TaskRequest) GetLocalTime() string {
	if m != nil {
		return m.LocalTime
	}
	return ""
}

type TaskResponse struct {
	Response             string            `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	NameRes              string            `protobuf:"bytes,2,opt,name=nameRes,proto3" json:"nameRes,omitempty"`
	AgeRes               int32             `protobuf:"varint,3,opt,name=ageRes,proto3" json:"ageRes,omitempty"`
	LocalAndServerTime   string            `protobuf:"bytes,4,opt,name=localAndServerTime,proto3" json:"localAndServerTime,omitempty"`
	Code                 TaskResponse_Code `protobuf:"varint,5,opt,name=code,proto3,enum=task.TaskResponse_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05afca300cf24a4, []int{3}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *TaskResponse) GetNameRes() string {
	if m != nil {
		return m.NameRes
	}
	return ""
}

func (m *TaskResponse) GetAgeRes() int32 {
	if m != nil {
		return m.AgeRes
	}
	return 0
}

func (m *TaskResponse) GetLocalAndServerTime() string {
	if m != nil {
		return m.LocalAndServerTime
	}
	return ""
}

func (m *TaskResponse) GetCode() TaskResponse_Code {
	if m != nil {
		return m.Code
	}
	return TaskResponse_failed
}

func init() {
	proto.RegisterEnum("task.TaskResponse_Code", TaskResponse_Code_name, TaskResponse_Code_value)
	proto.RegisterType((*PlanRequest)(nil), "task.PlanRequest")
	proto.RegisterType((*PlanResponse)(nil), "task.PlanResponse")
	proto.RegisterType((*TaskRequest)(nil), "task.TaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "task.TaskResponse")
}

func init() { proto.RegisterFile("api/proto/task/task.proto", fileDescriptor_f05afca300cf24a4) }

var fileDescriptor_f05afca300cf24a4 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x29, 0x96, 0x02, 0x03, 0x31, 0x38, 0x07, 0xad, 0xc4, 0x03, 0xd9, 0xc4, 0xa4, 0x89,
	0xb1, 0x44, 0x8c, 0x17, 0x6f, 0x46, 0x2e, 0xde, 0x4c, 0xe1, 0x05, 0xd6, 0xed, 0xa8, 0x0d, 0xa5,
	0x2d, 0xbb, 0x45, 0x7d, 0x64, 0x1f, 0xc3, 0xec, 0x76, 0x0b, 0x25, 0xea, 0x85, 0xcc, 0xcf, 0xcc,
	0xdf, 0x6f, 0xf2, 0xef, 0xc0, 0x39, 0x2f, 0x92, 0x69, 0x21, 0xf3, 0x32, 0x9f, 0x96, 0x5c, 0xad,
	0xcc, 0x4f, 0x68, 0x34, 0xba, 0xba, 0x66, 0x97, 0x30, 0x78, 0x4e, 0x79, 0x16, 0xd1, 0x66, 0x4b,
	0xaa, 0xc4, 0x53, 0xf0, 0xb4, 0x7c, 0x9a, 0xfb, 0xce, 0xc4, 0x09, 0xfa, 0x91, 0x55, 0xec, 0x1e,
	0x86, 0xd5, 0x98, 0x2a, 0xf2, 0x4c, 0xd1, 0x7f, 0x73, 0x88, 0xe0, 0x3e, 0xe6, 0x31, 0xf9, 0xed,
	0x89, 0x13, 0x74, 0x22, 0x53, 0xb3, 0x4f, 0x18, 0x2c, 0xb9, 0x5a, 0xd5, 0x08, 0x1f, 0xba, 0xb2,
	0x2a, 0xad, 0xb7, 0x96, 0xba, 0x93, 0xf1, 0x35, 0x45, 0xb4, 0x31, 0xfe, 0x7e, 0x54, 0x4b, 0x8d,
	0xe3, 0x6f, 0xa6, 0x71, 0x64, 0x3e, 0x6c, 0x15, 0x5e, 0x40, 0x3f, 0xcd, 0x05, 0x4f, 0x97, 0xc9,
	0x9a, 0x7c, 0xd7, 0x78, 0xf6, 0x7f, 0xb0, 0x6f, 0x07, 0x86, 0x15, 0xd9, 0x6e, 0x3d, 0x86, 0x9e,
	0xb4, 0xb5, 0x65, 0xef, 0xf4, 0x1e, 0xae, 0x0e, 0xe1, 0x6a, 0x07, 0x57, 0x07, 0x70, 0x85, 0x21,
	0xa0, 0x61, 0x3d, 0x64, 0xf1, 0x82, 0xe4, 0x07, 0xc9, 0xc6, 0x16, 0x7f, 0x74, 0xf0, 0x0a, 0x5c,
	0xa1, 0xb3, 0xe9, 0x4c, 0x9c, 0xe0, 0x78, 0x76, 0x16, 0x9a, 0xb7, 0x68, 0xee, 0x17, 0xea, 0xb8,
	0x22, 0x33, 0xc4, 0xae, 0xab, 0x20, 0x11, 0xc0, 0x7b, 0xe5, 0x49, 0x4a, 0xf1, 0xa8, 0x85, 0x03,
	0xe8, 0xaa, 0xad, 0x10, 0xa4, 0xd4, 0xc8, 0xc1, 0x21, 0xf4, 0xe8, 0x4b, 0x50, 0x51, 0x52, 0x3c,
	0x6a, 0xcf, 0x6c, 0xc6, 0x9a, 0x96, 0x08, 0xc2, 0x1b, 0xf0, 0x16, 0xc4, 0xa5, 0x78, 0xc7, 0x93,
	0x26, 0xc6, 0xc4, 0x3c, 0xc6, 0xdf, 0x64, 0xd6, 0xc2, 0x3b, 0x00, 0xfd, 0x86, 0x73, 0x2a, 0x79,
	0x92, 0xd6, 0xb6, 0xc6, 0x69, 0xd4, 0xb6, 0xe6, 0x19, 0xb0, 0xd6, 0x8b, 0x67, 0x8e, 0xe9, 0xf6,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x00, 0x53, 0x5a, 0x3e, 0x69, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	Search(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	PlanDetail(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Search(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) PlanDetail(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error) {
	out := new(PlanResponse)
	err := c.cc.Invoke(ctx, "/task.TaskService/PlanDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	Search(context.Context, *TaskRequest) (*TaskResponse, error)
	PlanDetail(context.Context, *PlanRequest) (*PlanResponse, error)
}

// UnimplementedTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (*UnimplementedTaskServiceServer) Search(ctx context.Context, req *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedTaskServiceServer) PlanDetail(ctx context.Context, req *PlanRequest) (*PlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlanDetail not implemented")
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Search(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_PlanDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).PlanDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskService/PlanDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).PlanDetail(ctx, req.(*PlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _TaskService_Search_Handler,
		},
		{
			MethodName: "PlanDetail",
			Handler:    _TaskService_PlanDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/task/task.proto",
}
