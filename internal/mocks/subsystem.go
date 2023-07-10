// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/subsystem (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/base-org/pessimism/internal/api/models"
	core "github.com/base-org/pessimism/internal/core"
	invariant "github.com/base-org/pessimism/internal/engine/invariant"
	gomock "github.com/golang/mock/gomock"
)

// SubManager is a mock of Manager interface.
type SubManager struct {
	ctrl     *gomock.Controller
	recorder *SubManagerMockRecorder
}

// SubManagerMockRecorder is the mock recorder for SubManager.
type SubManagerMockRecorder struct {
	mock *SubManager
}

// NewSubManager creates a new mock instance.
func NewSubManager(ctrl *gomock.Controller) *SubManager {
	mock := &SubManager{ctrl: ctrl}
	mock.recorder = &SubManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *SubManager) EXPECT() *SubManagerMockRecorder {
	return m.recorder
}

// BuildDeployCfg mocks base method.
func (m *SubManager) BuildDeployCfg(arg0 *core.PipelineConfig, arg1 *core.SessionConfig) (*invariant.DeployConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildDeployCfg", arg0, arg1)
	ret0, _ := ret[0].(*invariant.DeployConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildDeployCfg indicates an expected call of BuildDeployCfg.
func (mr *SubManagerMockRecorder) BuildDeployCfg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDeployCfg", reflect.TypeOf((*SubManager)(nil).BuildDeployCfg), arg0, arg1)
}

// BuildPipelineCfg mocks base method.
func (m *SubManager) BuildPipelineCfg(arg0 *models.InvRequestParams) (*core.PipelineConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildPipelineCfg", arg0)
	ret0, _ := ret[0].(*core.PipelineConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildPipelineCfg indicates an expected call of BuildPipelineCfg.
func (mr *SubManagerMockRecorder) BuildPipelineCfg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildPipelineCfg", reflect.TypeOf((*SubManager)(nil).BuildPipelineCfg), arg0)
}

// RunInvSession mocks base method.
func (m *SubManager) RunInvSession(arg0 *invariant.DeployConfig) (core.SUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInvSession", arg0)
	ret0, _ := ret[0].(core.SUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInvSession indicates an expected call of RunInvSession.
func (mr *SubManagerMockRecorder) RunInvSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInvSession", reflect.TypeOf((*SubManager)(nil).RunInvSession), arg0)
}

// Shutdown mocks base method.
func (m *SubManager) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *SubManagerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*SubManager)(nil).Shutdown))
}

// StartEventRoutines mocks base method.
func (m *SubManager) StartEventRoutines(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartEventRoutines", arg0)
}

// StartEventRoutines indicates an expected call of StartEventRoutines.
func (mr *SubManagerMockRecorder) StartEventRoutines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEventRoutines", reflect.TypeOf((*SubManager)(nil).StartEventRoutines), arg0)
}