// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	protobuf "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *Service) Create(ctx context.Context, input *protobuf.CreateCompanyRequest) (*protobuf.CreateCompanyResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *protobuf.CreateCompanyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protobuf.CreateCompanyRequest) *protobuf.CreateCompanyResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobuf.CreateCompanyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protobuf.CreateCompanyRequest) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Service) Delete(ctx context.Context, input *protobuf.DeleteCompanyRequest) (*protobuf.DeleteCompanyResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *protobuf.DeleteCompanyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protobuf.DeleteCompanyRequest) *protobuf.DeleteCompanyResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobuf.DeleteCompanyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protobuf.DeleteCompanyRequest) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, input
func (_m *Service) List(ctx context.Context, input *protobuf.ListCompanyRequest) (*protobuf.ListCompanyResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *protobuf.ListCompanyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protobuf.ListCompanyRequest) *protobuf.ListCompanyResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobuf.ListCompanyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protobuf.ListCompanyRequest) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: ctx, input
func (_m *Service) Read(ctx context.Context, input *protobuf.ReadCompanyRequest) (*protobuf.ReadCompanyResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *protobuf.ReadCompanyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protobuf.ReadCompanyRequest) *protobuf.ReadCompanyResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobuf.ReadCompanyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protobuf.ReadCompanyRequest) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, input
func (_m *Service) Update(ctx context.Context, input *protobuf.UpdateCompanyRequest) (*protobuf.UpdateCompanyResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *protobuf.UpdateCompanyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protobuf.UpdateCompanyRequest) *protobuf.UpdateCompanyResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protobuf.UpdateCompanyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protobuf.UpdateCompanyRequest) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
