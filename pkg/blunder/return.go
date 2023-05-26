package blunder

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Return is a function to return the error
func (mr *MatchResult) Return() error {
	return mr.Result
}

// ReturnForGin is a function to return the error for gin framework
func (mr *MatchResult) ReturnForGin(ginContext *gin.Context) {
	ginContext.JSON(
		mr.Result.GetHttpStatusCode(),
		mr.Result,
	)
}

// ReturnForGrpc is a function to return the error for grpc
func (mr *MatchResult) ReturnForGrpc() error {
	st := status.New(codes.Code(mr.Result.GetGrpcStatusCode()), mr.Result.GetId())
	return st.Err()
}
