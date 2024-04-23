// Code generated by mockery v2.42.2. DO NOT EDIT.

package v1

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestResolver is an autogenerated mock type for the RequestResolver type
type RequestResolver struct {
	mock.Mock
}

// Execute provides a mock function with given fields: req
func (_m *RequestResolver) Execute(req *http.Request) {
	_m.Called(req)
}

// NewRequestResolver creates a new instance of RequestResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestResolver(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestResolver {
	mock := &RequestResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
