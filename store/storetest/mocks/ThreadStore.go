// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/tetrafolium/mattermost-server/model"
	mock "github.com/stretchr/testify/mock"
)

// ThreadStore is an autogenerated mock type for the ThreadStore type
type ThreadStore struct {
	mock.Mock
}

// CollectThreadsWithNewerReplies provides a mock function with given fields: userId, channelIds, timestamp
func (_m *ThreadStore) CollectThreadsWithNewerReplies(userId string, channelIds []string, timestamp int64) ([]string, error) {
	ret := _m.Called(userId, channelIds, timestamp)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, []string, int64) []string); ok {
		r0 = rf(userId, channelIds, timestamp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, int64) error); ok {
		r1 = rf(userId, channelIds, timestamp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMembershipIfNeeded provides a mock function with given fields: userId, postId, following, incrementMentions, updateFollowing
func (_m *ThreadStore) CreateMembershipIfNeeded(userId string, postId string, following bool, incrementMentions bool, updateFollowing bool) error {
	ret := _m.Called(userId, postId, following, incrementMentions, updateFollowing)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool) error); ok {
		r0 = rf(userId, postId, following, incrementMentions, updateFollowing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: postId
func (_m *ThreadStore) Delete(postId string) error {
	ret := _m.Called(postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMembershipForUser provides a mock function with given fields: userId, postId
func (_m *ThreadStore) DeleteMembershipForUser(userId string, postId string) error {
	ret := _m.Called(userId, postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ThreadStore) Get(id string) (*model.Thread, error) {
	ret := _m.Called(id)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(string) *model.Thread); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
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

// GetMembershipForUser provides a mock function with given fields: userId, postId
func (_m *ThreadStore) GetMembershipForUser(userId string, postId string) (*model.ThreadMembership, error) {
	ret := _m.Called(userId, postId)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string) *model.ThreadMembership); ok {
		r0 = rf(userId, postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipsForUser provides a mock function with given fields: userId, teamId
func (_m *ThreadStore) GetMembershipsForUser(userId string, teamId string) ([]*model.ThreadMembership, error) {
	ret := _m.Called(userId, teamId)

	var r0 []*model.ThreadMembership
	if rf, ok := ret.Get(0).(func(string, string) []*model.ThreadMembership); ok {
		r0 = rf(userId, teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, teamId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with given fields: threadId, since
func (_m *ThreadStore) GetPosts(threadId string, since int64) ([]*model.Post, error) {
	ret := _m.Called(threadId, since)

	var r0 []*model.Post
	if rf, ok := ret.Get(0).(func(string, int64) []*model.Post); ok {
		r0 = rf(threadId, since)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(threadId, since)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadsForUser provides a mock function with given fields: userId, teamId, opts
func (_m *ThreadStore) GetThreadsForUser(userId string, teamId string, opts model.GetUserThreadsOpts) (*model.Threads, error) {
	ret := _m.Called(userId, teamId, opts)

	var r0 *model.Threads
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) *model.Threads); ok {
		r0 = rf(userId, teamId, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Threads)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userId, teamId, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkAllAsRead provides a mock function with given fields: userId, teamId
func (_m *ThreadStore) MarkAllAsRead(userId string, teamId string) error {
	ret := _m.Called(userId, teamId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, teamId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAsRead provides a mock function with given fields: userId, threadId, timestamp
func (_m *ThreadStore) MarkAsRead(userId string, threadId string, timestamp int64) error {
	ret := _m.Called(userId, threadId, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userId, threadId, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: thread
func (_m *ThreadStore) Save(thread *model.Thread) (*model.Thread, error) {
	ret := _m.Called(thread)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(*model.Thread) *model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Thread) error); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMembership provides a mock function with given fields: membership
func (_m *ThreadStore) SaveMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveMultiple provides a mock function with given fields: thread
func (_m *ThreadStore) SaveMultiple(thread []*model.Thread) ([]*model.Thread, int, error) {
	ret := _m.Called(thread)

	var r0 []*model.Thread
	if rf, ok := ret.Get(0).(func([]*model.Thread) []*model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Thread)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func([]*model.Thread) int); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]*model.Thread) error); ok {
		r2 = rf(thread)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: thread
func (_m *ThreadStore) Update(thread *model.Thread) (*model.Thread, error) {
	ret := _m.Called(thread)

	var r0 *model.Thread
	if rf, ok := ret.Get(0).(func(*model.Thread) *model.Thread); ok {
		r0 = rf(thread)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Thread) error); ok {
		r1 = rf(thread)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMembership provides a mock function with given fields: membership
func (_m *ThreadStore) UpdateMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	var r0 *model.ThreadMembership
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUnreadsByChannel provides a mock function with given fields: userId, changedThreads, timestamp, updateViewedTimestamp
func (_m *ThreadStore) UpdateUnreadsByChannel(userId string, changedThreads []string, timestamp int64, updateViewedTimestamp bool) error {
	ret := _m.Called(userId, changedThreads, timestamp, updateViewedTimestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, int64, bool) error); ok {
		r0 = rf(userId, changedThreads, timestamp, updateViewedTimestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
