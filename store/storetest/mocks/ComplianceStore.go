// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/tetrafolium/mattermost-server/v5/model"
)

// ComplianceStore is an autogenerated mock type for the ComplianceStore type
type ComplianceStore struct {
	mock.Mock
}

// ComplianceExport provides a mock function with given fields: compliance
func (_m *ComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, error) {
	ret := _m.Called(compliance)

	var r0 []*model.CompliancePost
	if rf, ok := ret.Get(0).(func(*model.Compliance) []*model.CompliancePost); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.CompliancePost)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Compliance) error); ok {
		r1 = rf(compliance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *ComplianceStore) Get(id string) (*model.Compliance, error) {
	ret := _m.Called(id)

	var r0 *model.Compliance
	if rf, ok := ret.Get(0).(func(string) *model.Compliance); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: offset, limit
func (_m *ComplianceStore) GetAll(offset int, limit int) (model.Compliances, error) {
	ret := _m.Called(offset, limit)

	var r0 model.Compliances
	if rf, ok := ret.Get(0).(func(int, int) model.Compliances); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Compliances)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageExport provides a mock function with given fields: after, limit
func (_m *ComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, error) {
	ret := _m.Called(after, limit)

	var r0 []*model.MessageExport
	if rf, ok := ret.Get(0).(func(int64, int) []*model.MessageExport); ok {
		r0 = rf(after, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.MessageExport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(after, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: compliance
func (_m *ComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, error) {
	ret := _m.Called(compliance)

	var r0 *model.Compliance
	if rf, ok := ret.Get(0).(func(*model.Compliance) *model.Compliance); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Compliance) error); ok {
		r1 = rf(compliance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: compliance
func (_m *ComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, error) {
	ret := _m.Called(compliance)

	var r0 *model.Compliance
	if rf, ok := ret.Get(0).(func(*model.Compliance) *model.Compliance); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Compliance) error); ok {
		r1 = rf(compliance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
