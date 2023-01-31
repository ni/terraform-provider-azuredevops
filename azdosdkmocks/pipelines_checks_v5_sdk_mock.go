// Code generated by MockGen. DO NOT EDIT.
// Source: C:\dev\terraform-provider-azuredevops\terraform-provider-azuredevops-ni-pre\vendor\github.com\microsoft\azure-devops-go-api\azuredevops\pipelineschecks\client.go

// Package mock_pipelineschecks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pipelineschecks "github.com/microsoft/azure-devops-go-api/azuredevops/pipelineschecks"
)

// PipelinesChecksClientV5 is a mock of Client interface.
type PipelinesChecksClientV5 struct {
	ctrl     *gomock.Controller
	recorder *PipelinesChecksClientV5MockRecorder
}

// PipelinesChecksClientV5MockRecorder is the mock recorder for PipelinesChecksClientV5.
type PipelinesChecksClientV5MockRecorder struct {
	mock *PipelinesChecksClientV5
}

// NewPipelinesChecksClientV5 creates a new mock instance.
func NewPipelinesChecksClientV5(ctrl *gomock.Controller) *PipelinesChecksClientV5 {
	mock := &PipelinesChecksClientV5{ctrl: ctrl}
	mock.recorder = &PipelinesChecksClientV5MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PipelinesChecksClientV5) EXPECT() *PipelinesChecksClientV5MockRecorder {
	return m.recorder
}

// AddCheckConfiguration mocks base method.
func (m *PipelinesChecksClientV5) AddCheckConfiguration(arg0 context.Context, arg1 pipelineschecks.AddCheckConfigurationArgs) (*pipelineschecks.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCheckConfiguration indicates an expected call of AddCheckConfiguration.
func (mr *PipelinesChecksClientV5MockRecorder) AddCheckConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCheckConfiguration", reflect.TypeOf((*PipelinesChecksClientV5)(nil).AddCheckConfiguration), arg0, arg1)
}

// DeleteCheckConfiguration mocks base method.
func (m *PipelinesChecksClientV5) DeleteCheckConfiguration(arg0 context.Context, arg1 pipelineschecks.DeleteCheckConfigurationArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCheckConfiguration indicates an expected call of DeleteCheckConfiguration.
func (mr *PipelinesChecksClientV5MockRecorder) DeleteCheckConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCheckConfiguration", reflect.TypeOf((*PipelinesChecksClientV5)(nil).DeleteCheckConfiguration), arg0, arg1)
}

// EvaluateCheckSuite mocks base method.
func (m *PipelinesChecksClientV5) EvaluateCheckSuite(arg0 context.Context, arg1 pipelineschecks.EvaluateCheckSuiteArgs) (*pipelineschecks.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvaluateCheckSuite", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateCheckSuite indicates an expected call of EvaluateCheckSuite.
func (mr *PipelinesChecksClientV5MockRecorder) EvaluateCheckSuite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateCheckSuite", reflect.TypeOf((*PipelinesChecksClientV5)(nil).EvaluateCheckSuite), arg0, arg1)
}

// GetCheckConfiguration mocks base method.
func (m *PipelinesChecksClientV5) GetCheckConfiguration(arg0 context.Context, arg1 pipelineschecks.GetCheckConfigurationArgs) (*pipelineschecks.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckConfiguration indicates an expected call of GetCheckConfiguration.
func (mr *PipelinesChecksClientV5MockRecorder) GetCheckConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckConfiguration", reflect.TypeOf((*PipelinesChecksClientV5)(nil).GetCheckConfiguration), arg0, arg1)
}

// GetCheckConfigurationsOnResource mocks base method.
func (m *PipelinesChecksClientV5) GetCheckConfigurationsOnResource(arg0 context.Context, arg1 pipelineschecks.GetCheckConfigurationsOnResourceArgs) (*[]pipelineschecks.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckConfigurationsOnResource", arg0, arg1)
	ret0, _ := ret[0].(*[]pipelineschecks.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckConfigurationsOnResource indicates an expected call of GetCheckConfigurationsOnResource.
func (mr *PipelinesChecksClientV5MockRecorder) GetCheckConfigurationsOnResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckConfigurationsOnResource", reflect.TypeOf((*PipelinesChecksClientV5)(nil).GetCheckConfigurationsOnResource), arg0, arg1)
}

// GetCheckSuite mocks base method.
func (m *PipelinesChecksClientV5) GetCheckSuite(arg0 context.Context, arg1 pipelineschecks.GetCheckSuiteArgs) (*pipelineschecks.CheckSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSuite", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSuite indicates an expected call of GetCheckSuite.
func (mr *PipelinesChecksClientV5MockRecorder) GetCheckSuite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSuite", reflect.TypeOf((*PipelinesChecksClientV5)(nil).GetCheckSuite), arg0, arg1)
}

// UpdateCheckConfiguration mocks base method.
func (m *PipelinesChecksClientV5) UpdateCheckConfiguration(arg0 context.Context, arg1 pipelineschecks.UpdateCheckConfigurationArgs) (*pipelineschecks.CheckConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCheckConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*pipelineschecks.CheckConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCheckConfiguration indicates an expected call of UpdateCheckConfiguration.
func (mr *PipelinesChecksClientV5MockRecorder) UpdateCheckConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCheckConfiguration", reflect.TypeOf((*PipelinesChecksClientV5)(nil).UpdateCheckConfiguration), arg0, arg1)
}
