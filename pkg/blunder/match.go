package blunder

import (
	"errors"
	"google.golang.org/grpc/status"
)

type MatchResult struct {
	IsMatched bool
	Result    Error
}

func (mr *MatchResult) GetIsMatched() bool {
	return mr.IsMatched
}

func (mr *MatchResult) Err() Error {
	return mr.Result
}

// Match is a function to match the error with the given condition
func Match(happened error, is error, shouldReturn Error) *MatchResult {
	matched := &MatchResult{
		IsMatched: true,
		Result:    shouldReturn,
	}

	// check if the error is from grpc standard status
	st, ok := status.FromError(happened)
	e, isError := is.(Error)
	if ok && isError {
		if e.GetId() == st.Message() {
			return matched
		}
	}

	if errors.Is(happened, is) {
		return matched
	}

	return &MatchResult{
		IsMatched: false,
		Result:    ErrUndefined.WithCustomMessage(happened.Error()),
	}
}

// MatchCondition is a function to match the error with the given condition
func MatchCondition(happened error, match *Condition) *MatchResult {
	for k, v := range match.detail {
		matched := &MatchResult{
			IsMatched: true,
			Result:    v,
		}

		// check if the error is from grpc standard status
		st, ok := status.FromError(happened)
		e, isError := k.(Error)
		if ok && isError {
			if e.GetId() == st.Message() {
				return matched
			}
		}

		if errors.Is(happened, k) {
			return matched
		}
	}
	return &MatchResult{
		IsMatched: false,
		Result:    ErrUndefined.WithCustomMessage(happened.Error()),
	}
}
