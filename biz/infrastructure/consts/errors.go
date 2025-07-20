package consts

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Errno struct {
	err  error
	code codes.Code
}

// GRPCStatus 实现 GRPCStatus 方法
func (en *Errno) GRPCStatus() *status.Status {
	return status.New(en.code, en.err.Error())
}

// 实现 Error 方法
func (en *Errno) Error() string {
	return en.err.Error()
}

// NewErrno 创建自定义错误
func NewErrno(code codes.Code, err error) *Errno {
	return &Errno{
		err:  err,
		code: code,
	}
}

// ErrInvalidParams 调用时错误
var (
	ErrInvalidParams = NewErrno(codes.InvalidArgument, errors.New("参数错误"))
	ErrCall          = NewErrno(codes.Unknown, errors.New("调用接口失败，请重试"))
	ErrAuth          = NewErrno(codes.PermissionDenied, errors.New("无权限调用"))
)

// 数据库相关错误
var (
	ErrNotFound        = NewErrno(codes.NotFound, errors.New("not found"))
	ErrInvalidObjectId = NewErrno(codes.InvalidArgument, errors.New("无效的id "))
)
