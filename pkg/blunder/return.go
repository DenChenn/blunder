package blunder

import (
	"github.com/DenChenn/blunder/pkg/options"
	"github.com/gin-gonic/gin"
)

func (mr *MatchResult) Return(opt *options.ReturnOption) error {
	// TODO: logging, alerting, etc.
	return mr.Result
}

func (mr *MatchResult) ReturnForGin(ginContext *gin.Context, opt *options.ReturnOption) {
	ginContext.JSON(
		mr.Result.GetHttpStatusCode(),
		mr.Result,
	)
}

func (mr *MatchResult) ReturnForGqlgen(opt *options.ReturnOption) error {
	panic("implement me")
}

func (mr *MatchResult) ReturnForGrpc(opt *options.ReturnOption) error {
	panic("implement me")
}
