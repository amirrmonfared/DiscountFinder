// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amirrmonfared/DiscountFinder/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddOnSalePrice mocks base method.
func (m *MockStore) AddOnSalePrice(arg0 context.Context, arg1 db.AddOnSalePriceParams) (db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOnSalePrice", arg0, arg1)
	ret0, _ := ret[0].(db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOnSalePrice indicates an expected call of AddOnSalePrice.
func (mr *MockStoreMockRecorder) AddOnSalePrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOnSalePrice", reflect.TypeOf((*MockStore)(nil).AddOnSalePrice), arg0, arg1)
}

// AddSecondPrice mocks base method.
func (m *MockStore) AddSecondPrice(arg0 context.Context, arg1 db.AddSecondPriceParams) (db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecondPrice", arg0, arg1)
	ret0, _ := ret[0].(db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSecondPrice indicates an expected call of AddSecondPrice.
func (mr *MockStoreMockRecorder) AddSecondPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecondPrice", reflect.TypeOf((*MockStore)(nil).AddSecondPrice), arg0, arg1)
}

// AddfirstProductPrice mocks base method.
func (m *MockStore) AddfirstProductPrice(arg0 context.Context, arg1 db.AddfirstProductPriceParams) (db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddfirstProductPrice", arg0, arg1)
	ret0, _ := ret[0].(db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddfirstProductPrice indicates an expected call of AddfirstProductPrice.
func (mr *MockStoreMockRecorder) AddfirstProductPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddfirstProductPrice", reflect.TypeOf((*MockStore)(nil).AddfirstProductPrice), arg0, arg1)
}

// CreateFirstProduct mocks base method.
func (m *MockStore) CreateFirstProduct(arg0 context.Context, arg1 db.CreateFirstProductParams) (db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFirstProduct", arg0, arg1)
	ret0, _ := ret[0].(db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFirstProduct indicates an expected call of CreateFirstProduct.
func (mr *MockStoreMockRecorder) CreateFirstProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFirstProduct", reflect.TypeOf((*MockStore)(nil).CreateFirstProduct), arg0, arg1)
}

// CreateOnSale mocks base method.
func (m *MockStore) CreateOnSale(arg0 context.Context, arg1 db.CreateOnSaleParams) (db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOnSale", arg0, arg1)
	ret0, _ := ret[0].(db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOnSale indicates an expected call of CreateOnSale.
func (mr *MockStoreMockRecorder) CreateOnSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOnSale", reflect.TypeOf((*MockStore)(nil).CreateOnSale), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.CreateProductResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.CreateProductResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateSecond mocks base method.
func (m *MockStore) CreateSecond(arg0 context.Context, arg1 db.CreateSecondParams) (db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecond", arg0, arg1)
	ret0, _ := ret[0].(db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecond indicates an expected call of CreateSecond.
func (mr *MockStoreMockRecorder) CreateSecond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecond", reflect.TypeOf((*MockStore)(nil).CreateSecond), arg0, arg1)
}

// DeleteFirstProduct mocks base method.
func (m *MockStore) DeleteFirstProduct(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFirstProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFirstProduct indicates an expected call of DeleteFirstProduct.
func (mr *MockStoreMockRecorder) DeleteFirstProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFirstProduct", reflect.TypeOf((*MockStore)(nil).DeleteFirstProduct), arg0, arg1)
}

// DeleteOnSale mocks base method.
func (m *MockStore) DeleteOnSale(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOnSale", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOnSale indicates an expected call of DeleteOnSale.
func (mr *MockStoreMockRecorder) DeleteOnSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOnSale", reflect.TypeOf((*MockStore)(nil).DeleteOnSale), arg0, arg1)
}

// DeleteSecond mocks base method.
func (m *MockStore) DeleteSecond(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecond", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecond indicates an expected call of DeleteSecond.
func (mr *MockStoreMockRecorder) DeleteSecond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecond", reflect.TypeOf((*MockStore)(nil).DeleteSecond), arg0, arg1)
}

// GetFirstProduct mocks base method.
func (m *MockStore) GetFirstProduct(arg0 context.Context, arg1 int64) (db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstProduct", arg0, arg1)
	ret0, _ := ret[0].(db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstProduct indicates an expected call of GetFirstProduct.
func (mr *MockStoreMockRecorder) GetFirstProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstProduct", reflect.TypeOf((*MockStore)(nil).GetFirstProduct), arg0, arg1)
}

// GetFirstProductForUpdate mocks base method.
func (m *MockStore) GetFirstProductForUpdate(arg0 context.Context, arg1 int64) (db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstProductForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstProductForUpdate indicates an expected call of GetFirstProductForUpdate.
func (mr *MockStoreMockRecorder) GetFirstProductForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstProductForUpdate", reflect.TypeOf((*MockStore)(nil).GetFirstProductForUpdate), arg0, arg1)
}

// GetLengthOfFirst mocks base method.
func (m *MockStore) GetLengthOfFirst(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLengthOfFirst", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLengthOfFirst indicates an expected call of GetLengthOfFirst.
func (mr *MockStoreMockRecorder) GetLengthOfFirst(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLengthOfFirst", reflect.TypeOf((*MockStore)(nil).GetLengthOfFirst), arg0)
}

// GetOnSale mocks base method.
func (m *MockStore) GetOnSale(arg0 context.Context, arg1 int64) (db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnSale", arg0, arg1)
	ret0, _ := ret[0].(db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnSale indicates an expected call of GetOnSale.
func (mr *MockStoreMockRecorder) GetOnSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnSale", reflect.TypeOf((*MockStore)(nil).GetOnSale), arg0, arg1)
}

// GetOnSaleForUpdate mocks base method.
func (m *MockStore) GetOnSaleForUpdate(arg0 context.Context, arg1 int64) (db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnSaleForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnSaleForUpdate indicates an expected call of GetOnSaleForUpdate.
func (mr *MockStoreMockRecorder) GetOnSaleForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnSaleForUpdate", reflect.TypeOf((*MockStore)(nil).GetOnSaleForUpdate), arg0, arg1)
}

// GetSecond mocks base method.
func (m *MockStore) GetSecond(arg0 context.Context, arg1 int64) (db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecond", arg0, arg1)
	ret0, _ := ret[0].(db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecond indicates an expected call of GetSecond.
func (mr *MockStoreMockRecorder) GetSecond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecond", reflect.TypeOf((*MockStore)(nil).GetSecond), arg0, arg1)
}

// GetSecondForUpdate mocks base method.
func (m *MockStore) GetSecondForUpdate(arg0 context.Context, arg1 int64) (db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecondForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecondForUpdate indicates an expected call of GetSecondForUpdate.
func (mr *MockStoreMockRecorder) GetSecondForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecondForUpdate", reflect.TypeOf((*MockStore)(nil).GetSecondForUpdate), arg0, arg1)
}

// ListFirstProduct mocks base method.
func (m *MockStore) ListFirstProduct(arg0 context.Context, arg1 db.ListFirstProductParams) ([]db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFirstProduct", arg0, arg1)
	ret0, _ := ret[0].([]db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFirstProduct indicates an expected call of ListFirstProduct.
func (mr *MockStoreMockRecorder) ListFirstProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFirstProduct", reflect.TypeOf((*MockStore)(nil).ListFirstProduct), arg0, arg1)
}

// ListOnSale mocks base method.
func (m *MockStore) ListOnSale(arg0 context.Context, arg1 db.ListOnSaleParams) ([]db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOnSale", arg0, arg1)
	ret0, _ := ret[0].([]db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOnSale indicates an expected call of ListOnSale.
func (mr *MockStoreMockRecorder) ListOnSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOnSale", reflect.TypeOf((*MockStore)(nil).ListOnSale), arg0, arg1)
}

// ListSecond mocks base method.
func (m *MockStore) ListSecond(arg0 context.Context, arg1 db.ListSecondParams) ([]db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecond", arg0, arg1)
	ret0, _ := ret[0].([]db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecond indicates an expected call of ListSecond.
func (mr *MockStoreMockRecorder) ListSecond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecond", reflect.TypeOf((*MockStore)(nil).ListSecond), arg0, arg1)
}

// UpdateFirstProduct mocks base method.
func (m *MockStore) UpdateFirstProduct(arg0 context.Context, arg1 db.UpdateFirstProductParams) (db.First, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFirstProduct", arg0, arg1)
	ret0, _ := ret[0].(db.First)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFirstProduct indicates an expected call of UpdateFirstProduct.
func (mr *MockStoreMockRecorder) UpdateFirstProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFirstProduct", reflect.TypeOf((*MockStore)(nil).UpdateFirstProduct), arg0, arg1)
}

// UpdateOnSale mocks base method.
func (m *MockStore) UpdateOnSale(arg0 context.Context, arg1 db.UpdateOnSaleParams) (db.OnSale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOnSale", arg0, arg1)
	ret0, _ := ret[0].(db.OnSale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOnSale indicates an expected call of UpdateOnSale.
func (mr *MockStoreMockRecorder) UpdateOnSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOnSale", reflect.TypeOf((*MockStore)(nil).UpdateOnSale), arg0, arg1)
}

// UpdateSecond mocks base method.
func (m *MockStore) UpdateSecond(arg0 context.Context, arg1 db.UpdateSecondParams) (db.Second, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecond", arg0, arg1)
	ret0, _ := ret[0].(db.Second)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecond indicates an expected call of UpdateSecond.
func (mr *MockStoreMockRecorder) UpdateSecond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecond", reflect.TypeOf((*MockStore)(nil).UpdateSecond), arg0, arg1)
}
