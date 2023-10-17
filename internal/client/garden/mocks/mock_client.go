// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardenctl-v2/internal/client/garden (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1alpha1 "github.com/gardener/gardener/pkg/apis/operations/v1alpha1"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CheckUserRoles mocks base method.
func (m *MockClient) CheckUserRoles(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserRoles", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserRoles indicates an expected call of CheckUserRoles.
func (mr *MockClientMockRecorder) CheckUserRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserRoles", reflect.TypeOf((*MockClient)(nil).CheckUserRoles), arg0)
}

// CurrentUser mocks base method.
func (m *MockClient) CurrentUser(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentUser", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentUser indicates an expected call of CurrentUser.
func (mr *MockClientMockRecorder) CurrentUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentUser", reflect.TypeOf((*MockClient)(nil).CurrentUser), arg0)
}

// FindShoot mocks base method.
func (m *MockClient) FindShoot(arg0 context.Context, arg1 ...client.ListOption) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindShoot", varargs...)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindShoot indicates an expected call of FindShoot.
func (mr *MockClientMockRecorder) FindShoot(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindShoot", reflect.TypeOf((*MockClient)(nil).FindShoot), varargs...)
}

// GetCloudProfile mocks base method.
func (m *MockClient) GetCloudProfile(arg0 context.Context, arg1 string) (*v1beta1.CloudProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudProfile", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.CloudProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudProfile indicates an expected call of GetCloudProfile.
func (mr *MockClientMockRecorder) GetCloudProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudProfile", reflect.TypeOf((*MockClient)(nil).GetCloudProfile), arg0, arg1)
}

// GetConfigMap mocks base method.
func (m *MockClient) GetConfigMap(arg0 context.Context, arg1, arg2 string) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap.
func (mr *MockClientMockRecorder) GetConfigMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockClient)(nil).GetConfigMap), arg0, arg1, arg2)
}

// GetNamespace mocks base method.
func (m *MockClient) GetNamespace(arg0 context.Context, arg1 string) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0, arg1)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockClientMockRecorder) GetNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockClient)(nil).GetNamespace), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockClient) GetProject(arg0 context.Context, arg1 string) (*v1beta1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockClientMockRecorder) GetProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockClient)(nil).GetProject), arg0, arg1)
}

// GetProjectByNamespace mocks base method.
func (m *MockClient) GetProjectByNamespace(arg0 context.Context, arg1 string) (*v1beta1.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectByNamespace", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectByNamespace indicates an expected call of GetProjectByNamespace.
func (mr *MockClientMockRecorder) GetProjectByNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByNamespace", reflect.TypeOf((*MockClient)(nil).GetProjectByNamespace), arg0, arg1)
}

// GetSecret mocks base method.
func (m *MockClient) GetSecret(arg0 context.Context, arg1, arg2 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockClientMockRecorder) GetSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockClient)(nil).GetSecret), arg0, arg1, arg2)
}

// GetSecretBinding mocks base method.
func (m *MockClient) GetSecretBinding(arg0 context.Context, arg1, arg2 string) (*v1beta1.SecretBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBinding", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.SecretBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBinding indicates an expected call of GetSecretBinding.
func (mr *MockClientMockRecorder) GetSecretBinding(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBinding", reflect.TypeOf((*MockClient)(nil).GetSecretBinding), arg0, arg1, arg2)
}

// GetSeed mocks base method.
func (m *MockClient) GetSeed(arg0 context.Context, arg1 string) (*v1beta1.Seed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeed", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Seed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeed indicates an expected call of GetSeed.
func (mr *MockClientMockRecorder) GetSeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeed", reflect.TypeOf((*MockClient)(nil).GetSeed), arg0, arg1)
}

// GetSeedClientConfig mocks base method.
func (m *MockClient) GetSeedClientConfig(arg0 context.Context, arg1 string) (clientcmd.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedClientConfig", arg0, arg1)
	ret0, _ := ret[0].(clientcmd.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeedClientConfig indicates an expected call of GetSeedClientConfig.
func (mr *MockClientMockRecorder) GetSeedClientConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedClientConfig", reflect.TypeOf((*MockClient)(nil).GetSeedClientConfig), arg0, arg1)
}

// GetShoot mocks base method.
func (m *MockClient) GetShoot(arg0 context.Context, arg1, arg2 string) (*v1beta1.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShoot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShoot indicates an expected call of GetShoot.
func (mr *MockClientMockRecorder) GetShoot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShoot", reflect.TypeOf((*MockClient)(nil).GetShoot), arg0, arg1, arg2)
}

// GetShootClientConfig mocks base method.
func (m *MockClient) GetShootClientConfig(arg0 context.Context, arg1, arg2 string) (clientcmd.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShootClientConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(clientcmd.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShootClientConfig indicates an expected call of GetShootClientConfig.
func (mr *MockClientMockRecorder) GetShootClientConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShootClientConfig", reflect.TypeOf((*MockClient)(nil).GetShootClientConfig), arg0, arg1, arg2)
}

// GetShootOfManagedSeed mocks base method.
func (m *MockClient) GetShootOfManagedSeed(arg0 context.Context, arg1 string) (*v1alpha10.Shoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShootOfManagedSeed", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha10.Shoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShootOfManagedSeed indicates an expected call of GetShootOfManagedSeed.
func (mr *MockClientMockRecorder) GetShootOfManagedSeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShootOfManagedSeed", reflect.TypeOf((*MockClient)(nil).GetShootOfManagedSeed), arg0, arg1)
}

// ListBastions mocks base method.
func (m *MockClient) ListBastions(arg0 context.Context, arg1 ...client.ListOption) (*v1alpha1.BastionList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBastions", varargs...)
	ret0, _ := ret[0].(*v1alpha1.BastionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBastions indicates an expected call of ListBastions.
func (mr *MockClientMockRecorder) ListBastions(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBastions", reflect.TypeOf((*MockClient)(nil).ListBastions), varargs...)
}

// ListProjects mocks base method.
func (m *MockClient) ListProjects(arg0 context.Context, arg1 ...client.ListOption) (*v1beta1.ProjectList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*v1beta1.ProjectList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockClientMockRecorder) ListProjects(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockClient)(nil).ListProjects), varargs...)
}

// ListSeeds mocks base method.
func (m *MockClient) ListSeeds(arg0 context.Context, arg1 ...client.ListOption) (*v1beta1.SeedList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSeeds", varargs...)
	ret0, _ := ret[0].(*v1beta1.SeedList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSeeds indicates an expected call of ListSeeds.
func (mr *MockClientMockRecorder) ListSeeds(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSeeds", reflect.TypeOf((*MockClient)(nil).ListSeeds), varargs...)
}

// ListShoots mocks base method.
func (m *MockClient) ListShoots(arg0 context.Context, arg1 ...client.ListOption) (*v1beta1.ShootList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListShoots", varargs...)
	ret0, _ := ret[0].(*v1beta1.ShootList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShoots indicates an expected call of ListShoots.
func (mr *MockClientMockRecorder) ListShoots(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShoots", reflect.TypeOf((*MockClient)(nil).ListShoots), varargs...)
}

// PatchBastion mocks base method.
func (m *MockClient) PatchBastion(arg0 context.Context, arg1, arg2 *v1alpha1.Bastion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchBastion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchBastion indicates an expected call of PatchBastion.
func (mr *MockClientMockRecorder) PatchBastion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchBastion", reflect.TypeOf((*MockClient)(nil).PatchBastion), arg0, arg1, arg2)
}

// RuntimeClient mocks base method.
func (m *MockClient) RuntimeClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuntimeClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// RuntimeClient indicates an expected call of RuntimeClient.
func (mr *MockClientMockRecorder) RuntimeClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuntimeClient", reflect.TypeOf((*MockClient)(nil).RuntimeClient))
}
