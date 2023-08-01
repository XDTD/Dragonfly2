// Code generated by MockGen. DO NOT EDIT.
// Source: objectstorage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	objectstorage "github.com/XDTD/Dragonfly2/pkg/objectstorage"
	gomock "github.com/golang/mock/gomock"
)

// MockObjectStorage is a mock of ObjectStorage interface.
type MockObjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStorageMockRecorder
}

// MockObjectStorageMockRecorder is the mock recorder for MockObjectStorage.
type MockObjectStorageMockRecorder struct {
	mock *MockObjectStorage
}

// NewMockObjectStorage creates a new mock instance.
func NewMockObjectStorage(ctrl *gomock.Controller) *MockObjectStorage {
	mock := &MockObjectStorage{ctrl: ctrl}
	mock.recorder = &MockObjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStorage) EXPECT() *MockObjectStorageMockRecorder {
	return m.recorder
}

// CopyObject mocks base method.
func (m *MockObjectStorage) CopyObject(ctx context.Context, bucketName, sourceObjectKey, destinationObjectKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyObject", ctx, bucketName, sourceObjectKey, destinationObjectKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyObject indicates an expected call of CopyObject.
func (mr *MockObjectStorageMockRecorder) CopyObject(ctx, bucketName, sourceObjectKey, destinationObjectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyObject", reflect.TypeOf((*MockObjectStorage)(nil).CopyObject), ctx, bucketName, sourceObjectKey, destinationObjectKey)
}

// CreateBucket mocks base method.
func (m *MockObjectStorage) CreateBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockObjectStorageMockRecorder) CreateBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockObjectStorage)(nil).CreateBucket), ctx, bucketName)
}

// DeleteBucket mocks base method.
func (m *MockObjectStorage) DeleteBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockObjectStorageMockRecorder) DeleteBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockObjectStorage)(nil).DeleteBucket), ctx, bucketName)
}

// DeleteObject mocks base method.
func (m *MockObjectStorage) DeleteObject(ctx context.Context, bucketName, objectKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStorageMockRecorder) DeleteObject(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStorage)(nil).DeleteObject), ctx, bucketName, objectKey)
}

// GetBucketMetadata mocks base method.
func (m *MockObjectStorage) GetBucketMetadata(ctx context.Context, bucketName string) (*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketMetadata", ctx, bucketName)
	ret0, _ := ret[0].(*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketMetadata indicates an expected call of GetBucketMetadata.
func (mr *MockObjectStorageMockRecorder) GetBucketMetadata(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetBucketMetadata), ctx, bucketName)
}

// GetObjectMetadata mocks base method.
func (m *MockObjectStorage) GetObjectMetadata(ctx context.Context, bucketName, objectKey string) (*objectstorage.ObjectMetadata, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadata", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetObjectMetadata indicates an expected call of GetObjectMetadata.
func (mr *MockObjectStorageMockRecorder) GetObjectMetadata(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetObjectMetadata), ctx, bucketName, objectKey)
}

// GetObjectMetadatas mocks base method.
func (m *MockObjectStorage) GetObjectMetadatas(ctx context.Context, bucketName, prefix, marker, delimiter string, limit int64) ([]*objectstorage.ObjectMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadatas", ctx, bucketName, prefix, marker, delimiter, limit)
	ret0, _ := ret[0].([]*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectMetadatas indicates an expected call of GetObjectMetadatas.
func (mr *MockObjectStorageMockRecorder) GetObjectMetadatas(ctx, bucketName, prefix, marker, delimiter, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).GetObjectMetadatas), ctx, bucketName, prefix, marker, delimiter, limit)
}

// GetOject mocks base method.
func (m *MockObjectStorage) GetOject(ctx context.Context, bucketName, objectKey string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOject", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOject indicates an expected call of GetOject.
func (mr *MockObjectStorageMockRecorder) GetOject(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOject", reflect.TypeOf((*MockObjectStorage)(nil).GetOject), ctx, bucketName, objectKey)
}

// GetSignURL mocks base method.
func (m *MockObjectStorage) GetSignURL(ctx context.Context, bucketName, objectKey string, method objectstorage.Method, expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignURL", ctx, bucketName, objectKey, method, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignURL indicates an expected call of GetSignURL.
func (mr *MockObjectStorageMockRecorder) GetSignURL(ctx, bucketName, objectKey, method, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignURL", reflect.TypeOf((*MockObjectStorage)(nil).GetSignURL), ctx, bucketName, objectKey, method, expire)
}

// IsBucketExist mocks base method.
func (m *MockObjectStorage) IsBucketExist(ctx context.Context, bucketName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBucketExist", ctx, bucketName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsBucketExist indicates an expected call of IsBucketExist.
func (mr *MockObjectStorageMockRecorder) IsBucketExist(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBucketExist", reflect.TypeOf((*MockObjectStorage)(nil).IsBucketExist), ctx, bucketName)
}

// IsObjectExist mocks base method.
func (m *MockObjectStorage) IsObjectExist(ctx context.Context, bucketName, objectKey string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsObjectExist", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsObjectExist indicates an expected call of IsObjectExist.
func (mr *MockObjectStorageMockRecorder) IsObjectExist(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsObjectExist", reflect.TypeOf((*MockObjectStorage)(nil).IsObjectExist), ctx, bucketName, objectKey)
}

// ListBucketMetadatas mocks base method.
func (m *MockObjectStorage) ListBucketMetadatas(ctx context.Context) ([]*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBucketMetadatas", ctx)
	ret0, _ := ret[0].([]*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketMetadatas indicates an expected call of ListBucketMetadatas.
func (mr *MockObjectStorageMockRecorder) ListBucketMetadatas(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).ListBucketMetadatas), ctx)
}

// PutObject mocks base method.
func (m *MockObjectStorage) PutObject(ctx context.Context, bucketName, objectKey, digest string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", ctx, bucketName, objectKey, digest, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject.
func (mr *MockObjectStorageMockRecorder) PutObject(ctx, bucketName, objectKey, digest, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockObjectStorage)(nil).PutObject), ctx, bucketName, objectKey, digest, reader)
}
