package blunder

import (
	"errors"
	"github.com/DenChenn/blunder/pkg/options"
	"google.golang.org/grpc/status"
)

type MatchResult struct {
	IsMatched bool
	Result    Error
}

func (mr *MatchResult) GetIsMatched() bool {
	return mr.IsMatched
}

func MatchOne(target error, comp error, output Error, opt *options.MatchOption) *MatchResult {
	matched := &MatchResult{
		IsMatched: true,
		Result:    output,
	}

	// if option is set to FromGrpc, then check if the target error is grpc standard status
	if opt != nil && opt.FromGrpc != nil && *opt.FromGrpc {
		st, ok := status.FromError(target)
		e, isError := comp.(Error)
		if ok && isError {
			if e.GetId() == st.Message() {
				return matched
			}
		}
	}

	if errors.Is(target, comp) {
		return matched
	}

	return &MatchResult{
		IsMatched: false,
		Result:    ErrUndefined.WithCustomMessage(target.Error()),
	}
}

func MatchMany(target error, comp *Map, opt *options.MatchOption) *MatchResult {
	isFromGrpc := opt != nil && opt.FromGrpc != nil && *opt.FromGrpc

	for k, v := range comp.detail {
		matched := &MatchResult{
			IsMatched: true,
			Result:    v,
		}

		if isFromGrpc {
			st, ok := status.FromError(target)
			e, isError := k.(Error)
			if ok && isError {
				if e.GetId() == st.Message() {
					return matched
				}
			}
		}

		if errors.Is(target, k) {
			return matched
		}
	}
	return &MatchResult{
		IsMatched: false,
		Result:    ErrUndefined.WithCustomMessage(target.Error()),
	}
}
