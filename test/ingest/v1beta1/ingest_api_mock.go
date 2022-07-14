// Copyright (c) 2022 IndyKite
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/indykite/jarvis-sdk-go/gen/indykite/ingest/v1beta1 (interfaces: IngestAPIClient,IngestAPI_StreamRecordsClient)

// Package ingest is a generated GoMock package.
package ingest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"

	ingestv1beta1 "github.com/indykite/jarvis-sdk-go/gen/indykite/ingest/v1beta1"
)

// MockIngestAPIClient is a mock of IngestAPIClient interface.
type MockIngestAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIngestAPIClientMockRecorder
}

// MockIngestAPIClientMockRecorder is the mock recorder for MockIngestAPIClient.
type MockIngestAPIClientMockRecorder struct {
	mock *MockIngestAPIClient
}

// NewMockIngestAPIClient creates a new mock instance.
func NewMockIngestAPIClient(ctrl *gomock.Controller) *MockIngestAPIClient {
	mock := &MockIngestAPIClient{ctrl: ctrl}
	mock.recorder = &MockIngestAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestAPIClient) EXPECT() *MockIngestAPIClientMockRecorder {
	return m.recorder
}

// StreamRecords mocks base method.
func (m *MockIngestAPIClient) StreamRecords(arg0 context.Context, arg1 ...grpc.CallOption) (ingestv1beta1.IngestAPI_StreamRecordsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamRecords", varargs...)
	ret0, _ := ret[0].(ingestv1beta1.IngestAPI_StreamRecordsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamRecords indicates an expected call of StreamRecords.
func (mr *MockIngestAPIClientMockRecorder) StreamRecords(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamRecords", reflect.TypeOf((*MockIngestAPIClient)(nil).StreamRecords), varargs...)
}

// MockIngestAPI_StreamRecordsClient is a mock of IngestAPI_StreamRecordsClient interface.
type MockIngestAPI_StreamRecordsClient struct {
	ctrl     *gomock.Controller
	recorder *MockIngestAPI_StreamRecordsClientMockRecorder
}

// MockIngestAPI_StreamRecordsClientMockRecorder is the mock recorder for MockIngestAPI_StreamRecordsClient.
type MockIngestAPI_StreamRecordsClientMockRecorder struct {
	mock *MockIngestAPI_StreamRecordsClient
}

// NewMockIngestAPI_StreamRecordsClient creates a new mock instance.
func NewMockIngestAPI_StreamRecordsClient(ctrl *gomock.Controller) *MockIngestAPI_StreamRecordsClient {
	mock := &MockIngestAPI_StreamRecordsClient{ctrl: ctrl}
	mock.recorder = &MockIngestAPI_StreamRecordsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIngestAPI_StreamRecordsClient) EXPECT() *MockIngestAPI_StreamRecordsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) Recv() (*ingestv1beta1.StreamRecordsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ingestv1beta1.StreamRecordsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) Send(arg0 *ingestv1beta1.StreamRecordsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockIngestAPI_StreamRecordsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockIngestAPI_StreamRecordsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockIngestAPI_StreamRecordsClient)(nil).Trailer))
}
