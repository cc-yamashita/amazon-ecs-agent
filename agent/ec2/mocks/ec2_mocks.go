// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/ec2 (interfaces: EC2MetadataClient,HttpClient)

package mock_ec2

import (
	ec2 "github.com/aws/amazon-ecs-agent/agent/ec2"
	ec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	gomock "github.com/golang/mock/gomock"
)

// Mock of EC2MetadataClient interface
type MockEC2MetadataClient struct {
	ctrl     *gomock.Controller
	recorder *_MockEC2MetadataClientRecorder
}

// Recorder for MockEC2MetadataClient (not exported)
type _MockEC2MetadataClientRecorder struct {
	mock *MockEC2MetadataClient
}

func NewMockEC2MetadataClient(ctrl *gomock.Controller) *MockEC2MetadataClient {
	mock := &MockEC2MetadataClient{ctrl: ctrl}
	mock.recorder = &_MockEC2MetadataClientRecorder{mock}
	return mock
}

func (_m *MockEC2MetadataClient) EXPECT() *_MockEC2MetadataClientRecorder {
	return _m.recorder
}

func (_m *MockEC2MetadataClient) DefaultCredentials() (*ec2.RoleCredentials, error) {
	ret := _m.ctrl.Call(_m, "DefaultCredentials")
	ret0, _ := ret[0].(*ec2.RoleCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) DefaultCredentials() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultCredentials")
}

func (_m *MockEC2MetadataClient) GetDynamicData(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetDynamicData", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) GetDynamicData(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDynamicData", arg0)
}

func (_m *MockEC2MetadataClient) GetMetadata(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetadata", arg0)
}

func (_m *MockEC2MetadataClient) InstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	ret := _m.ctrl.Call(_m, "InstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) InstanceIdentityDocument() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstanceIdentityDocument")
}

// Mock of HttpClient interface
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *_MockHttpClientRecorder
}

// Recorder for MockHttpClient (not exported)
type _MockHttpClientRecorder struct {
	mock *MockHttpClient
}

func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &_MockHttpClientRecorder{mock}
	return mock
}

func (_m *MockHttpClient) EXPECT() *_MockHttpClientRecorder {
	return _m.recorder
}

func (_m *MockHttpClient) GetDynamicData(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetDynamicData", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHttpClientRecorder) GetDynamicData(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDynamicData", arg0)
}

func (_m *MockHttpClient) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	ret := _m.ctrl.Call(_m, "GetInstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHttpClientRecorder) GetInstanceIdentityDocument() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstanceIdentityDocument")
}

func (_m *MockHttpClient) GetMetadata(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHttpClientRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetadata", arg0)
}
