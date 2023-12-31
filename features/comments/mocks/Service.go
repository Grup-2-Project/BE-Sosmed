// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	comments "BE-Sosmed/features/comments"

	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateComment provides a mock function with given fields: token, newComment
func (_m *Service) CreateComment(token *jwt.Token, newComment comments.Comment) (comments.Comment, error) {
	ret := _m.Called(token, newComment)

	var r0 comments.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(*jwt.Token, comments.Comment) (comments.Comment, error)); ok {
		return rf(token, newComment)
	}
	if rf, ok := ret.Get(0).(func(*jwt.Token, comments.Comment) comments.Comment); ok {
		r0 = rf(token, newComment)
	} else {
		r0 = ret.Get(0).(comments.Comment)
	}

	if rf, ok := ret.Get(1).(func(*jwt.Token, comments.Comment) error); ok {
		r1 = rf(token, newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteComment provides a mock function with given fields: token, CommentID
func (_m *Service) DeleteComment(token *jwt.Token, CommentID uint) error {
	ret := _m.Called(token, CommentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*jwt.Token, uint) error); ok {
		r0 = rf(token, CommentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutComment provides a mock function with given fields: token, updatedComment
func (_m *Service) PutComment(token *jwt.Token, updatedComment comments.Comment) (comments.Comment, error) {
	ret := _m.Called(token, updatedComment)

	var r0 comments.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(*jwt.Token, comments.Comment) (comments.Comment, error)); ok {
		return rf(token, updatedComment)
	}
	if rf, ok := ret.Get(0).(func(*jwt.Token, comments.Comment) comments.Comment); ok {
		r0 = rf(token, updatedComment)
	} else {
		r0 = ret.Get(0).(comments.Comment)
	}

	if rf, ok := ret.Get(1).(func(*jwt.Token, comments.Comment) error); ok {
		r1 = rf(token, updatedComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
