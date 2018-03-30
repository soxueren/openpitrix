// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTaskRequest struct {
	X          *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NodeId     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Target     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
}

func (m *CreateTaskRequest) Reset()                    { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()               {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *CreateTaskRequest) GetX() *google_protobuf.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateTaskRequest) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *google_protobuf.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *google_protobuf.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *CreateTaskResponse) Reset()                    { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()               {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *CreateTaskResponse) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type Task struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode  *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor   *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner      *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target     *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId     *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *Task) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *google_protobuf.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *google_protobuf.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *google_protobuf.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string                     `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId    []string                     `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Executor *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Status   []string                     `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset uint32 `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeTasksRequest) Reset()                    { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()               {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet    []*Task `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
}

func (m *DescribeTasksResponse) Reset()                    { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()               {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xe3, 0x36, 0x63, 0x82, 0x60, 0xd5, 0x82, 0x15, 0x95, 0x36, 0x58, 0x1c, 0xca,
	0x4f, 0x13, 0x48, 0x8b, 0x84, 0x5a, 0x71, 0x28, 0xe5, 0x92, 0x03, 0x97, 0xb4, 0x70, 0xe0, 0x12,
	0x6d, 0xec, 0x89, 0xd9, 0x36, 0xf5, 0xba, 0xbb, 0xeb, 0xa6, 0x47, 0xc4, 0x01, 0x89, 0x6b, 0x79,
	0x16, 0x5e, 0x81, 0x17, 0xe0, 0xc8, 0x95, 0x07, 0x41, 0xbb, 0x76, 0x1a, 0xd3, 0x50, 0x70, 0x40,
	0xe2, 0x14, 0xed, 0xec, 0xf7, 0x79, 0x76, 0xe6, 0xfb, 0x66, 0x02, 0xa0, 0xa8, 0x3c, 0x6c, 0xc6,
	0x82, 0x2b, 0x4e, 0x80, 0xc7, 0x18, 0xc5, 0x4c, 0x09, 0x76, 0x5a, 0x5f, 0x09, 0x39, 0x0f, 0x87,
	0xd8, 0x32, 0x37, 0xfd, 0x64, 0xd0, 0x1a, 0x09, 0x1a, 0xc7, 0x28, 0x64, 0x8a, 0xad, 0xaf, 0x5e,
	0xbc, 0x57, 0xec, 0x08, 0xa5, 0xa2, 0x47, 0x71, 0x06, 0x58, 0xce, 0x00, 0x34, 0x66, 0x2d, 0x1a,
	0x45, 0x5c, 0x51, 0xc5, 0x78, 0x34, 0xa6, 0x3f, 0x34, 0x3f, 0xfe, 0x7a, 0x88, 0xd1, 0xba, 0x1c,
	0xd1, 0x30, 0x44, 0xd1, 0xe2, 0xb1, 0x41, 0x4c, 0xa3, 0xbd, 0x6f, 0x25, 0xb8, 0xb1, 0x2b, 0x90,
	0x2a, 0xdc, 0xa7, 0xf2, 0xb0, 0x8b, 0xc7, 0x09, 0x4a, 0x45, 0xee, 0x81, 0xd5, 0x73, 0xad, 0x86,
	0xb5, 0xe6, 0xb4, 0x97, 0x9b, 0x69, 0xb6, 0xe6, 0xf8, 0x39, 0xcd, 0x3d, 0x25, 0x58, 0x14, 0xbe,
	0xa6, 0xc3, 0x04, 0xbb, 0x57, 0xc8, 0x06, 0xd8, 0x07, 0xbc, 0xdf, 0x63, 0x81, 0x5b, 0x2a, 0x80,
	0xaf, 0x1c, 0xf0, 0x7e, 0x27, 0x20, 0x4f, 0x60, 0x3e, 0xe2, 0x01, 0x6a, 0x56, 0xb9, 0x00, 0xcb,
	0xd6, 0xe0, 0x4e, 0x40, 0x36, 0xc1, 0x56, 0x54, 0x84, 0xa8, 0xdc, 0xb9, 0x22, 0xac, 0x14, 0x4b,
	0x9e, 0x81, 0xa3, 0x95, 0xe8, 0x51, 0x5f, 0x17, 0xee, 0x56, 0x0a, 0x50, 0x8d, 0x74, 0x3b, 0x06,
	0x4f, 0xb6, 0xa0, 0x1a, 0x30, 0x81, 0xbe, 0x62, 0x27, 0xe8, 0xda, 0x05, 0xc8, 0x13, 0xb8, 0xf7,
	0xce, 0x02, 0x92, 0xef, 0xae, 0x8c, 0x79, 0x24, 0x51, 0x97, 0x6f, 0x5e, 0xc4, 0x82, 0x42, 0x4d,
	0xb6, 0x35, 0xb8, 0x13, 0xfc, 0x55, 0xab, 0xbd, 0xcf, 0x15, 0x98, 0xd3, 0xc9, 0xff, 0x67, 0xd2,
	0x8b, 0x2d, 0x2f, 0xcf, 0xd8, 0xf2, 0x4d, 0xb0, 0xa5, 0xa2, 0x2a, 0x91, 0xc5, 0x74, 0x4e, 0xb1,
	0x64, 0x1b, 0x00, 0x85, 0xe0, 0xa2, 0xe7, 0xf3, 0x00, 0x2f, 0x95, 0xf9, 0x55, 0x27, 0x52, 0x1b,
	0xed, 0x4c, 0x29, 0x83, 0xdf, 0xe5, 0x01, 0xfe, 0x8b, 0xca, 0xe4, 0x29, 0x2c, 0xe0, 0x29, 0xfa,
	0x89, 0xe2, 0xc2, 0x9d, 0x2f, 0x40, 0x3d, 0x47, 0x93, 0x36, 0x54, 0xf8, 0x28, 0x42, 0xe1, 0x2e,
	0x14, 0xe9, 0xad, 0x81, 0xe6, 0x86, 0xa0, 0x3a, 0xc3, 0x10, 0xe4, 0x26, 0x0e, 0x66, 0x98, 0xb8,
	0x6d, 0x70, 0x7c, 0xe3, 0xdf, 0x9e, 0x5e, 0x42, 0xae, 0x63, 0xa8, 0xf5, 0x29, 0xea, 0xfe, 0x78,
	0x43, 0x75, 0x21, 0x85, 0xeb, 0x80, 0x26, 0xa7, 0xd2, 0xa4, 0xe4, 0xab, 0x7f, 0x26, 0xa7, 0x70,
	0x1d, 0xf0, 0xbe, 0x58, 0xb0, 0xf8, 0x02, 0xa5, 0x2f, 0x58, 0xdf, 0x0c, 0x8f, 0x1c, 0xef, 0xa6,
	0x5b, 0x79, 0x1f, 0x97, 0xd7, 0xaa, 0xe7, 0x4e, 0x5d, 0xca, 0x39, 0x55, 0xc7, 0x33, 0x2f, 0xe6,
	0xd5, 0x29, 0xcf, 0xa4, 0xce, 0xcd, 0x9c, 0x0d, 0x4d, 0xa2, 0xcc, 0x68, 0x8b, 0x50, 0x19, 0xb2,
	0x23, 0xa6, 0x8c, 0xc7, 0x6a, 0xdd, 0xf4, 0xa0, 0xd1, 0x7c, 0x30, 0x90, 0xa8, 0x8c, 0x7d, 0x6a,
	0xdd, 0xec, 0xe4, 0x21, 0x2c, 0x5d, 0xa8, 0x23, 0xdb, 0x02, 0xab, 0xe0, 0x28, 0xae, 0xe8, 0xb0,
	0xe7, 0xf3, 0x24, 0x52, 0x66, 0x28, 0x6b, 0x5d, 0x30, 0xa1, 0x5d, 0x1d, 0x21, 0x0f, 0x60, 0xc1,
	0x54, 0xaa, 0xbf, 0xa9, 0x4b, 0x72, 0xda, 0xd7, 0x9b, 0x93, 0xff, 0x91, 0xa6, 0x59, 0x29, 0xa6,
	0x17, 0x7b, 0xa8, 0xda, 0x1f, 0x4b, 0xe0, 0xe8, 0xc8, 0x4b, 0x1a, 0xd1, 0x10, 0x05, 0x39, 0x06,
	0x98, 0x6c, 0x1e, 0x72, 0x3b, 0x4f, 0x9c, 0xda, 0xf7, 0xf5, 0x95, 0xcb, 0xae, 0xd3, 0xa7, 0x7a,
	0x77, 0xcf, 0x76, 0x6a, 0x24, 0x73, 0x42, 0x43, 0x67, 0x7c, 0xff, 0xf5, 0xfb, 0xa7, 0xd2, 0x35,
	0xaf, 0xda, 0x3a, 0x79, 0xdc, 0xd2, 0x67, 0xb9, 0x65, 0xdd, 0x27, 0x1f, 0x2c, 0xa8, 0xfd, 0x54,
	0x2a, 0x69, 0xe4, 0xbf, 0xfb, 0x2b, 0x35, 0xeb, 0x77, 0x7e, 0x83, 0xc8, 0x92, 0x3f, 0x3a, 0xdb,
	0x59, 0x26, 0xf5, 0x20, 0xbb, 0x33, 0xe9, 0x65, 0x63, 0xc4, 0xd4, 0xdb, 0xc6, 0x80, 0x0d, 0x15,
	0x0a, 0xf3, 0x16, 0x87, 0x4c, 0xde, 0xf2, 0x7c, 0xee, 0x4d, 0x29, 0xee, 0xf7, 0x6d, 0x23, 0xef,
	0xc6, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x3e, 0x8f, 0xec, 0x88, 0x07, 0x00, 0x00,
}
