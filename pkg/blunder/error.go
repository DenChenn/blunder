package blunder

import (
	"fmt"
)

type Error interface {
	Error() string
	GetId() string
	GetHttpStatusCode() int
	GetGrpcStatusCode() int
	GetErrorCode() string
	GetMessage() string
}

type Unknown struct {
	Id             string
	HttpStatusCode int
	GrpcStatusCode int
	ErrorCode      string
	Message        string
}

func NewUnknown(httpStatusCode int, grpcStatusCode int, errorCode string, message string) Error {
	return &Unknown{
		Id:             "Unknown",
		HttpStatusCode: httpStatusCode,
		GrpcStatusCode: grpcStatusCode,
		ErrorCode:      errorCode,
		Message:        message,
	}
}

func (e *Unknown) Error() string {
	return e.Message
}

func (e *Unknown) GetId() string {
	return e.Id
}

func (e *Unknown) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e *Unknown) GetGrpcStatusCode() int {
	return e.GrpcStatusCode
}

func (e *Unknown) GetErrorCode() string {
	return e.ErrorCode
}

func (e *Unknown) GetMessage() string {
	return e.Message
}

func PackageError(message string) Error {
	return NewUnknown(
		500,
		13,
		"Unknown",
		fmt.Sprintf("[blunder error]: %v", message),
	)
}
