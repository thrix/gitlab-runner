// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package certificate

import mock "github.com/stretchr/testify/mock"
import tls "crypto/tls"

// MockGenerator is an autogenerated mock type for the Generator type
type MockGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields: host
func (_m *MockGenerator) Generate(host string) (tls.Certificate, []byte, error) {
	ret := _m.Called(host)

	var r0 tls.Certificate
	if rf, ok := ret.Get(0).(func(string) tls.Certificate); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Get(0).(tls.Certificate)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string) []byte); ok {
		r1 = rf(host)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(host)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
