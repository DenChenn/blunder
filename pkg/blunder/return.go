package blunder

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mr *MatchResult) Return() error {
	return mr.Result
}

func (mr *MatchResult) ReturnForGin(ginContext *gin.Context) {
	ginContext.JSON(
		mr.Result.GetHttpStatusCode(),
		mr.Result,
	)
}

func (mr *MatchResult) ReturnForGrpc() error {
	st := status.New(codes.Code(mr.Result.GetGrpcStatusCode()), mr.Result.GetId())
	return st.Err()
}
