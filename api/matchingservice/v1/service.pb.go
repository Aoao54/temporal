// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/matchingservice/v1/service.proto

package matchingservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("temporal/server/api/matchingservice/v1/service.proto", fileDescriptor_1a5c83076e651916)
}

var fileDescriptor_1a5c83076e651916 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x3d, 0x6f, 0x13, 0x31,
	0x18, 0xc7, 0xcf, 0x0b, 0x83, 0x25, 0x54, 0xf5, 0x04, 0x42, 0x54, 0xc2, 0x42, 0x0c, 0x1d, 0x2f,
	0x2a, 0xb0, 0xd1, 0x02, 0x69, 0x52, 0x42, 0xa5, 0x56, 0x4d, 0x81, 0x80, 0xc4, 0x82, 0x9c, 0xbb,
	0x87, 0x60, 0xf5, 0x72, 0x3e, 0x6c, 0x5f, 0x50, 0x36, 0x3e, 0x01, 0x62, 0x60, 0x62, 0x45, 0x42,
	0x0c, 0x4c, 0x7c, 0x00, 0x56, 0x18, 0x33, 0xb6, 0x1b, 0xb9, 0x2c, 0x8c, 0xfd, 0x08, 0xe8, 0x9a,
	0xd8, 0x79, 0x0f, 0xce, 0xcb, 0x96, 0x5c, 0xfc, 0xff, 0x3d, 0xbf, 0xbf, 0x62, 0x9f, 0x8c, 0xef,
	0x2a, 0xa8, 0xc7, 0x5c, 0xd0, 0x30, 0x27, 0x41, 0x34, 0x40, 0xe4, 0x68, 0xcc, 0x72, 0x75, 0xaa,
	0xfc, 0x37, 0x2c, 0xaa, 0x65, 0x8f, 0x98, 0x0f, 0xb9, 0xc6, 0x56, 0xae, 0xf7, 0xd1, 0x8b, 0x05,
	0x57, 0xdc, 0xdd, 0xd4, 0x29, 0xaf, 0x9b, 0xf2, 0x68, 0xcc, 0xbc, 0x91, 0x94, 0xd7, 0xd8, 0xda,
	0xd8, 0xb1, 0xa4, 0x0b, 0x78, 0x9b, 0x80, 0x54, 0xaf, 0x04, 0xc8, 0x98, 0x47, 0xb2, 0x37, 0xe6,
	0xf6, 0xd9, 0x3a, 0x5e, 0x3b, 0xec, 0xad, 0x7e, 0xda, 0x5d, 0xed, 0x7e, 0x45, 0xf8, 0x6a, 0x99,
	0x87, 0xe1, 0x0b, 0x2e, 0x4e, 0x5e, 0x87, 0xfc, 0xdd, 0x33, 0x2a, 0x4f, 0x8e, 0x13, 0x48, 0xc0,
	0x2d, 0x7a, 0x76, 0x56, 0xde, 0xc4, 0xf8, 0x93, 0xae, 0xc2, 0xc6, 0xde, 0x92, 0x94, 0x6e, 0x81,
	0x5b, 0x8e, 0x11, 0xcd, 0xfb, 0x8a, 0x35, 0x98, 0x6a, 0x2e, 0x28, 0x3a, 0x16, 0x5f, 0x48, 0x74,
	0x02, 0xc5, 0x88, 0x7e, 0x42, 0x78, 0x2d, 0x1f, 0x04, 0x83, 0x5d, 0xdc, 0xfb, 0xb6, 0xf0, 0x91,
	0xa0, 0x96, 0x7b, 0xb0, 0x70, 0x7e, 0x54, 0x6b, 0xd0, 0x7c, 0x2e, 0xad, 0xc1, 0xe0, 0x22, 0x5a,
	0xc3, 0x79, 0xa3, 0xf5, 0x01, 0xe1, 0xcb, 0xc7, 0x09, 0x88, 0xa6, 0xd6, 0x76, 0xb7, 0x6d, 0xa1,
	0x43, 0x31, 0xad, 0xb4, 0xb3, 0x60, 0xda, 0x08, 0xfd, 0x40, 0xf8, 0x7a, 0xf7, 0x6b, 0x70, 0xb1,
	0x24, 0xf3, 0x2d, 0xf0, 0x7a, 0x1c, 0x82, 0x82, 0xc0, 0x7d, 0x6c, 0x8b, 0x9f, 0x8a, 0xd0, 0xa2,
	0xfb, 0x2b, 0x20, 0x0d, 0x1d, 0x8e, 0x02, 0x8d, 0x7c, 0x08, 0x8f, 0x12, 0x25, 0x15, 0x8d, 0x02,
	0x16, 0xd5, 0xb2, 0x8d, 0x6a, 0x7f, 0x38, 0x26, 0xc6, 0xe7, 0x3e, 0x1c, 0x53, 0x28, 0x46, 0xf4,
	0x33, 0xc2, 0xeb, 0x45, 0x90, 0xbe, 0x60, 0x55, 0xe8, 0x9f, 0xe0, 0x87, 0xb6, 0xf8, 0xb1, 0xa8,
	0x16, 0xcc, 0x2f, 0x41, 0x30, 0x72, 0xdf, 0x11, 0xbe, 0x76, 0xc0, 0xa4, 0x32, 0xbf, 0x95, 0xa9,
	0x50, 0x4c, 0x31, 0x1e, 0x49, 0xf7, 0x91, 0xed, 0x80, 0x29, 0x00, 0x2d, 0x5a, 0x5a, 0x9a, 0x63,
	0x74, 0x7f, 0x21, 0x7c, 0xb3, 0x12, 0x07, 0x54, 0x41, 0xb6, 0x8d, 0x41, 0xec, 0x26, 0x2c, 0x0c,
	0xf6, 0x83, 0x6c, 0x7f, 0x50, 0xc5, 0xaa, 0x2c, 0x64, 0xaa, 0xe9, 0x1e, 0xd9, 0xce, 0xfb, 0x1f,
	0x49, 0x17, 0x28, 0xaf, 0x0e, 0x68, 0x9a, 0xfc, 0x44, 0xf8, 0x46, 0x09, 0xd4, 0x8c, 0x1a, 0x07,
	0xb6, 0x53, 0x67, 0x62, 0x74, 0x87, 0xc3, 0x15, 0xd1, 0x4c, 0x81, 0x2f, 0x08, 0x5f, 0x29, 0x41,
	0xff, 0xff, 0xaa, 0x48, 0x10, 0x45, 0xaa, 0xa8, 0x5b, 0x98, 0x63, 0xd2, 0x58, 0x5a, 0xeb, 0x16,
	0x97, 0x83, 0x18, 0xcb, 0x33, 0x84, 0x37, 0xf3, 0x71, 0x1c, 0x36, 0x27, 0x2c, 0x8a, 0x43, 0xe6,
	0xd3, 0x6c, 0x87, 0xed, 0x35, 0x20, 0x52, 0x6e, 0xc5, 0xfa, 0xcd, 0x6e, 0xc5, 0xd3, 0x4d, 0x9e,
	0xaf, 0x1a, 0xab, 0xbb, 0xed, 0x8a, 0x56, 0x9b, 0x38, 0xa7, 0x6d, 0xe2, 0x9c, 0xb7, 0x09, 0x7a,
	0x9f, 0x12, 0xf4, 0x2d, 0x25, 0xe8, 0x77, 0x4a, 0x50, 0x2b, 0x25, 0xe8, 0x4f, 0x4a, 0xd0, 0xdf,
	0x94, 0x38, 0xe7, 0x29, 0x41, 0x1f, 0x3b, 0xc4, 0x69, 0x75, 0x88, 0x73, 0xda, 0x21, 0xce, 0xcb,
	0xed, 0x1a, 0xef, 0x1b, 0x31, 0x3e, 0xfb, 0x5a, 0x75, 0x6f, 0xe4, 0x51, 0xf5, 0xd2, 0xc5, 0xb5,
	0xea, 0xce, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xea, 0x8a, 0xdd, 0xf5, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MatchingServiceClient is the client API for MatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatchingServiceClient interface {
	// PollWorkflowTaskQueue is called by frontend to process WorkflowTask from a specific task queue.  A
	// WorkflowTask is dispatched to callers for active workflow executions, with pending workflow tasks.
	PollWorkflowTaskQueue(ctx context.Context, in *PollWorkflowTaskQueueRequest, opts ...grpc.CallOption) (*PollWorkflowTaskQueueResponse, error)
	// PollActivityTaskQueue is called by frontend to process ActivityTask from a specific task queue.  ActivityTask
	// is dispatched to callers whenever a ScheduleTask command is made for a workflow execution.
	PollActivityTaskQueue(ctx context.Context, in *PollActivityTaskQueueRequest, opts ...grpc.CallOption) (*PollActivityTaskQueueResponse, error)
	// AddWorkflowTask is called by the history service when a workflow task is scheduled, so that it can be dispatched
	// by the MatchingEngine.
	AddWorkflowTask(ctx context.Context, in *AddWorkflowTaskRequest, opts ...grpc.CallOption) (*AddWorkflowTaskResponse, error)
	// AddActivityTask is called by the history service when a workflow task is scheduled, so that it can be dispatched
	// by the MatchingEngine.
	AddActivityTask(ctx context.Context, in *AddActivityTaskRequest, opts ...grpc.CallOption) (*AddActivityTaskResponse, error)
	// QueryWorkflow is called by frontend to query a workflow.
	QueryWorkflow(ctx context.Context, in *QueryWorkflowRequest, opts ...grpc.CallOption) (*QueryWorkflowResponse, error)
	// RespondQueryTaskCompleted is called by frontend to respond query completed.
	RespondQueryTaskCompleted(ctx context.Context, in *RespondQueryTaskCompletedRequest, opts ...grpc.CallOption) (*RespondQueryTaskCompletedResponse, error)
	// CancelOutstandingPoll is called by frontend to unblock long polls on matching for zombie pollers.
	// Our rpc stack does not support context propagation, so when a client connection goes away frontend sees
	// cancellation of context for that handler, but any corresponding calls (long-poll) to matching service does not
	// see the cancellation propagated so it can unblock corresponding long-polls on its end.  This results is tasks
	// being dispatched to zombie pollers in this situation.  This API is added so every time frontend makes a long-poll
	// api call to matching it passes in a pollerId and then calls this API when it detects client connection is closed
	// to unblock long polls for this poller and prevent tasks being sent to these zombie pollers.
	CancelOutstandingPoll(ctx context.Context, in *CancelOutstandingPollRequest, opts ...grpc.CallOption) (*CancelOutstandingPollResponse, error)
	// DescribeTaskQueue returns information about the target task queue, right now this API returns the
	// pollers which polled this task queue in last few minutes.
	DescribeTaskQueue(ctx context.Context, in *DescribeTaskQueueRequest, opts ...grpc.CallOption) (*DescribeTaskQueueResponse, error)
	// ListTaskQueuePartitions returns a map of partitionKey and hostAddress for a task queue.
	ListTaskQueuePartitions(ctx context.Context, in *ListTaskQueuePartitionsRequest, opts ...grpc.CallOption) (*ListTaskQueuePartitionsResponse, error)
	// (-- api-linter: core::0134::response-message-name=disabled
	//     aip.dev/not-precedent: UpdateWorkerBuildIdOrdering RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0134::method-signature=disabled
	//     aip.dev/not-precedent: UpdateWorkerBuildIdOrdering RPC doesn't follow Google API format. --)
	UpdateWorkerBuildIdCompatibility(ctx context.Context, in *UpdateWorkerBuildIdCompatibilityRequest, opts ...grpc.CallOption) (*UpdateWorkerBuildIdCompatibilityResponse, error)
	GetWorkerBuildIdCompatibility(ctx context.Context, in *GetWorkerBuildIdCompatibilityRequest, opts ...grpc.CallOption) (*GetWorkerBuildIdCompatibilityResponse, error)
	// Fetch user data for a task queue, this request should always be routed to the node holding the root partition of the workflow task queue.
	GetTaskQueueUserData(ctx context.Context, in *GetTaskQueueUserDataRequest, opts ...grpc.CallOption) (*GetTaskQueueUserDataResponse, error)
	// Apply a user data replication event.
	ApplyTaskQueueUserDataReplicationEvent(ctx context.Context, in *ApplyTaskQueueUserDataReplicationEventRequest, opts ...grpc.CallOption) (*ApplyTaskQueueUserDataReplicationEventResponse, error)
}

type matchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchingServiceClient(cc grpc.ClientConnInterface) MatchingServiceClient {
	return &matchingServiceClient{cc}
}

func (c *matchingServiceClient) PollWorkflowTaskQueue(ctx context.Context, in *PollWorkflowTaskQueueRequest, opts ...grpc.CallOption) (*PollWorkflowTaskQueueResponse, error) {
	out := new(PollWorkflowTaskQueueResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/PollWorkflowTaskQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) PollActivityTaskQueue(ctx context.Context, in *PollActivityTaskQueueRequest, opts ...grpc.CallOption) (*PollActivityTaskQueueResponse, error) {
	out := new(PollActivityTaskQueueResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/PollActivityTaskQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) AddWorkflowTask(ctx context.Context, in *AddWorkflowTaskRequest, opts ...grpc.CallOption) (*AddWorkflowTaskResponse, error) {
	out := new(AddWorkflowTaskResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/AddWorkflowTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) AddActivityTask(ctx context.Context, in *AddActivityTaskRequest, opts ...grpc.CallOption) (*AddActivityTaskResponse, error) {
	out := new(AddActivityTaskResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/AddActivityTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) QueryWorkflow(ctx context.Context, in *QueryWorkflowRequest, opts ...grpc.CallOption) (*QueryWorkflowResponse, error) {
	out := new(QueryWorkflowResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/QueryWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) RespondQueryTaskCompleted(ctx context.Context, in *RespondQueryTaskCompletedRequest, opts ...grpc.CallOption) (*RespondQueryTaskCompletedResponse, error) {
	out := new(RespondQueryTaskCompletedResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/RespondQueryTaskCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) CancelOutstandingPoll(ctx context.Context, in *CancelOutstandingPollRequest, opts ...grpc.CallOption) (*CancelOutstandingPollResponse, error) {
	out := new(CancelOutstandingPollResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/CancelOutstandingPoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) DescribeTaskQueue(ctx context.Context, in *DescribeTaskQueueRequest, opts ...grpc.CallOption) (*DescribeTaskQueueResponse, error) {
	out := new(DescribeTaskQueueResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/DescribeTaskQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) ListTaskQueuePartitions(ctx context.Context, in *ListTaskQueuePartitionsRequest, opts ...grpc.CallOption) (*ListTaskQueuePartitionsResponse, error) {
	out := new(ListTaskQueuePartitionsResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/ListTaskQueuePartitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) UpdateWorkerBuildIdCompatibility(ctx context.Context, in *UpdateWorkerBuildIdCompatibilityRequest, opts ...grpc.CallOption) (*UpdateWorkerBuildIdCompatibilityResponse, error) {
	out := new(UpdateWorkerBuildIdCompatibilityResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/UpdateWorkerBuildIdCompatibility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) GetWorkerBuildIdCompatibility(ctx context.Context, in *GetWorkerBuildIdCompatibilityRequest, opts ...grpc.CallOption) (*GetWorkerBuildIdCompatibilityResponse, error) {
	out := new(GetWorkerBuildIdCompatibilityResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/GetWorkerBuildIdCompatibility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) GetTaskQueueUserData(ctx context.Context, in *GetTaskQueueUserDataRequest, opts ...grpc.CallOption) (*GetTaskQueueUserDataResponse, error) {
	out := new(GetTaskQueueUserDataResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/GetTaskQueueUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingServiceClient) ApplyTaskQueueUserDataReplicationEvent(ctx context.Context, in *ApplyTaskQueueUserDataReplicationEventRequest, opts ...grpc.CallOption) (*ApplyTaskQueueUserDataReplicationEventResponse, error) {
	out := new(ApplyTaskQueueUserDataReplicationEventResponse)
	err := c.cc.Invoke(ctx, "/temporal.server.api.matchingservice.v1.MatchingService/ApplyTaskQueueUserDataReplicationEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchingServiceServer is the server API for MatchingService service.
type MatchingServiceServer interface {
	// PollWorkflowTaskQueue is called by frontend to process WorkflowTask from a specific task queue.  A
	// WorkflowTask is dispatched to callers for active workflow executions, with pending workflow tasks.
	PollWorkflowTaskQueue(context.Context, *PollWorkflowTaskQueueRequest) (*PollWorkflowTaskQueueResponse, error)
	// PollActivityTaskQueue is called by frontend to process ActivityTask from a specific task queue.  ActivityTask
	// is dispatched to callers whenever a ScheduleTask command is made for a workflow execution.
	PollActivityTaskQueue(context.Context, *PollActivityTaskQueueRequest) (*PollActivityTaskQueueResponse, error)
	// AddWorkflowTask is called by the history service when a workflow task is scheduled, so that it can be dispatched
	// by the MatchingEngine.
	AddWorkflowTask(context.Context, *AddWorkflowTaskRequest) (*AddWorkflowTaskResponse, error)
	// AddActivityTask is called by the history service when a workflow task is scheduled, so that it can be dispatched
	// by the MatchingEngine.
	AddActivityTask(context.Context, *AddActivityTaskRequest) (*AddActivityTaskResponse, error)
	// QueryWorkflow is called by frontend to query a workflow.
	QueryWorkflow(context.Context, *QueryWorkflowRequest) (*QueryWorkflowResponse, error)
	// RespondQueryTaskCompleted is called by frontend to respond query completed.
	RespondQueryTaskCompleted(context.Context, *RespondQueryTaskCompletedRequest) (*RespondQueryTaskCompletedResponse, error)
	// CancelOutstandingPoll is called by frontend to unblock long polls on matching for zombie pollers.
	// Our rpc stack does not support context propagation, so when a client connection goes away frontend sees
	// cancellation of context for that handler, but any corresponding calls (long-poll) to matching service does not
	// see the cancellation propagated so it can unblock corresponding long-polls on its end.  This results is tasks
	// being dispatched to zombie pollers in this situation.  This API is added so every time frontend makes a long-poll
	// api call to matching it passes in a pollerId and then calls this API when it detects client connection is closed
	// to unblock long polls for this poller and prevent tasks being sent to these zombie pollers.
	CancelOutstandingPoll(context.Context, *CancelOutstandingPollRequest) (*CancelOutstandingPollResponse, error)
	// DescribeTaskQueue returns information about the target task queue, right now this API returns the
	// pollers which polled this task queue in last few minutes.
	DescribeTaskQueue(context.Context, *DescribeTaskQueueRequest) (*DescribeTaskQueueResponse, error)
	// ListTaskQueuePartitions returns a map of partitionKey and hostAddress for a task queue.
	ListTaskQueuePartitions(context.Context, *ListTaskQueuePartitionsRequest) (*ListTaskQueuePartitionsResponse, error)
	// (-- api-linter: core::0134::response-message-name=disabled
	//     aip.dev/not-precedent: UpdateWorkerBuildIdOrdering RPC doesn't follow Google API format. --)
	// (-- api-linter: core::0134::method-signature=disabled
	//     aip.dev/not-precedent: UpdateWorkerBuildIdOrdering RPC doesn't follow Google API format. --)
	UpdateWorkerBuildIdCompatibility(context.Context, *UpdateWorkerBuildIdCompatibilityRequest) (*UpdateWorkerBuildIdCompatibilityResponse, error)
	GetWorkerBuildIdCompatibility(context.Context, *GetWorkerBuildIdCompatibilityRequest) (*GetWorkerBuildIdCompatibilityResponse, error)
	// Fetch user data for a task queue, this request should always be routed to the node holding the root partition of the workflow task queue.
	GetTaskQueueUserData(context.Context, *GetTaskQueueUserDataRequest) (*GetTaskQueueUserDataResponse, error)
	// Apply a user data replication event.
	ApplyTaskQueueUserDataReplicationEvent(context.Context, *ApplyTaskQueueUserDataReplicationEventRequest) (*ApplyTaskQueueUserDataReplicationEventResponse, error)
}

// UnimplementedMatchingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMatchingServiceServer struct {
}

func (*UnimplementedMatchingServiceServer) PollWorkflowTaskQueue(ctx context.Context, req *PollWorkflowTaskQueueRequest) (*PollWorkflowTaskQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PollWorkflowTaskQueue not implemented")
}
func (*UnimplementedMatchingServiceServer) PollActivityTaskQueue(ctx context.Context, req *PollActivityTaskQueueRequest) (*PollActivityTaskQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PollActivityTaskQueue not implemented")
}
func (*UnimplementedMatchingServiceServer) AddWorkflowTask(ctx context.Context, req *AddWorkflowTaskRequest) (*AddWorkflowTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWorkflowTask not implemented")
}
func (*UnimplementedMatchingServiceServer) AddActivityTask(ctx context.Context, req *AddActivityTaskRequest) (*AddActivityTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivityTask not implemented")
}
func (*UnimplementedMatchingServiceServer) QueryWorkflow(ctx context.Context, req *QueryWorkflowRequest) (*QueryWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWorkflow not implemented")
}
func (*UnimplementedMatchingServiceServer) RespondQueryTaskCompleted(ctx context.Context, req *RespondQueryTaskCompletedRequest) (*RespondQueryTaskCompletedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondQueryTaskCompleted not implemented")
}
func (*UnimplementedMatchingServiceServer) CancelOutstandingPoll(ctx context.Context, req *CancelOutstandingPollRequest) (*CancelOutstandingPollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOutstandingPoll not implemented")
}
func (*UnimplementedMatchingServiceServer) DescribeTaskQueue(ctx context.Context, req *DescribeTaskQueueRequest) (*DescribeTaskQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTaskQueue not implemented")
}
func (*UnimplementedMatchingServiceServer) ListTaskQueuePartitions(ctx context.Context, req *ListTaskQueuePartitionsRequest) (*ListTaskQueuePartitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskQueuePartitions not implemented")
}
func (*UnimplementedMatchingServiceServer) UpdateWorkerBuildIdCompatibility(ctx context.Context, req *UpdateWorkerBuildIdCompatibilityRequest) (*UpdateWorkerBuildIdCompatibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkerBuildIdCompatibility not implemented")
}
func (*UnimplementedMatchingServiceServer) GetWorkerBuildIdCompatibility(ctx context.Context, req *GetWorkerBuildIdCompatibilityRequest) (*GetWorkerBuildIdCompatibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerBuildIdCompatibility not implemented")
}
func (*UnimplementedMatchingServiceServer) GetTaskQueueUserData(ctx context.Context, req *GetTaskQueueUserDataRequest) (*GetTaskQueueUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskQueueUserData not implemented")
}
func (*UnimplementedMatchingServiceServer) ApplyTaskQueueUserDataReplicationEvent(ctx context.Context, req *ApplyTaskQueueUserDataReplicationEventRequest) (*ApplyTaskQueueUserDataReplicationEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyTaskQueueUserDataReplicationEvent not implemented")
}

func RegisterMatchingServiceServer(s *grpc.Server, srv MatchingServiceServer) {
	s.RegisterService(&_MatchingService_serviceDesc, srv)
}

func _MatchingService_PollWorkflowTaskQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollWorkflowTaskQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).PollWorkflowTaskQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/PollWorkflowTaskQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).PollWorkflowTaskQueue(ctx, req.(*PollWorkflowTaskQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_PollActivityTaskQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollActivityTaskQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).PollActivityTaskQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/PollActivityTaskQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).PollActivityTaskQueue(ctx, req.(*PollActivityTaskQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_AddWorkflowTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWorkflowTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).AddWorkflowTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/AddWorkflowTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).AddWorkflowTask(ctx, req.(*AddWorkflowTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_AddActivityTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).AddActivityTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/AddActivityTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).AddActivityTask(ctx, req.(*AddActivityTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_QueryWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).QueryWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/QueryWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).QueryWorkflow(ctx, req.(*QueryWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_RespondQueryTaskCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondQueryTaskCompletedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).RespondQueryTaskCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/RespondQueryTaskCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).RespondQueryTaskCompleted(ctx, req.(*RespondQueryTaskCompletedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_CancelOutstandingPoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOutstandingPollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).CancelOutstandingPoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/CancelOutstandingPoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).CancelOutstandingPoll(ctx, req.(*CancelOutstandingPollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_DescribeTaskQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTaskQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).DescribeTaskQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/DescribeTaskQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).DescribeTaskQueue(ctx, req.(*DescribeTaskQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_ListTaskQueuePartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskQueuePartitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).ListTaskQueuePartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/ListTaskQueuePartitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).ListTaskQueuePartitions(ctx, req.(*ListTaskQueuePartitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_UpdateWorkerBuildIdCompatibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkerBuildIdCompatibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).UpdateWorkerBuildIdCompatibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/UpdateWorkerBuildIdCompatibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).UpdateWorkerBuildIdCompatibility(ctx, req.(*UpdateWorkerBuildIdCompatibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_GetWorkerBuildIdCompatibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerBuildIdCompatibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).GetWorkerBuildIdCompatibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/GetWorkerBuildIdCompatibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).GetWorkerBuildIdCompatibility(ctx, req.(*GetWorkerBuildIdCompatibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_GetTaskQueueUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskQueueUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).GetTaskQueueUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/GetTaskQueueUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).GetTaskQueueUserData(ctx, req.(*GetTaskQueueUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingService_ApplyTaskQueueUserDataReplicationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyTaskQueueUserDataReplicationEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingServiceServer).ApplyTaskQueueUserDataReplicationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/temporal.server.api.matchingservice.v1.MatchingService/ApplyTaskQueueUserDataReplicationEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingServiceServer).ApplyTaskQueueUserDataReplicationEvent(ctx, req.(*ApplyTaskQueueUserDataReplicationEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MatchingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.server.api.matchingservice.v1.MatchingService",
	HandlerType: (*MatchingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PollWorkflowTaskQueue",
			Handler:    _MatchingService_PollWorkflowTaskQueue_Handler,
		},
		{
			MethodName: "PollActivityTaskQueue",
			Handler:    _MatchingService_PollActivityTaskQueue_Handler,
		},
		{
			MethodName: "AddWorkflowTask",
			Handler:    _MatchingService_AddWorkflowTask_Handler,
		},
		{
			MethodName: "AddActivityTask",
			Handler:    _MatchingService_AddActivityTask_Handler,
		},
		{
			MethodName: "QueryWorkflow",
			Handler:    _MatchingService_QueryWorkflow_Handler,
		},
		{
			MethodName: "RespondQueryTaskCompleted",
			Handler:    _MatchingService_RespondQueryTaskCompleted_Handler,
		},
		{
			MethodName: "CancelOutstandingPoll",
			Handler:    _MatchingService_CancelOutstandingPoll_Handler,
		},
		{
			MethodName: "DescribeTaskQueue",
			Handler:    _MatchingService_DescribeTaskQueue_Handler,
		},
		{
			MethodName: "ListTaskQueuePartitions",
			Handler:    _MatchingService_ListTaskQueuePartitions_Handler,
		},
		{
			MethodName: "UpdateWorkerBuildIdCompatibility",
			Handler:    _MatchingService_UpdateWorkerBuildIdCompatibility_Handler,
		},
		{
			MethodName: "GetWorkerBuildIdCompatibility",
			Handler:    _MatchingService_GetWorkerBuildIdCompatibility_Handler,
		},
		{
			MethodName: "GetTaskQueueUserData",
			Handler:    _MatchingService_GetTaskQueueUserData_Handler,
		},
		{
			MethodName: "ApplyTaskQueueUserDataReplicationEvent",
			Handler:    _MatchingService_ApplyTaskQueueUserDataReplicationEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/server/api/matchingservice/v1/service.proto",
}
