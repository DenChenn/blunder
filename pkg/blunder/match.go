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

func (mr *MatchResult) GetIsMatched(opt *options.ReturnOption) bool {
	return mr.IsMatched
}

func MatchOne(target error, comp error, output Error, opt *options.MatchOption) *MatchResult {
	if opt != nil && opt.FromGrpc != nil && *opt.FromGrpc {
		st, ok := status.FromError(target)
		if !ok {
			return &MatchResult{
				IsMatched: false,
				Result:    PackageError("target is not a grpc standard status error"),
			}
		}

		e, isError := comp.(Error)
		if !isError {
			return &MatchResult{
				IsMatched: false,
				Result:    PackageError("Given comparative error is not a blunder error, specifying FromGrpc is not allowed"),
			}
		}

		if e.GetId() == st.Message() {
			return &MatchResult{
				IsMatched: true,
				Result:    output,
			}
		}
		return &MatchResult{
			IsMatched: false,
		}
	}

	if errors.Is(target, comp) {
		return &MatchResult{
			IsMatched: true,
			Result:    output,
		}
	}
	return &MatchResult{
		IsMatched: false,
	}
}

func MatchMany(target error, comp *Map, opt *options.MatchOption) *MatchResult {
	for k, v := range comp.detail {
		if opt != nil && opt.FromGrpc != nil && *opt.FromGrpc {
			st, ok := status.FromError(target)
			if !ok {
				return &MatchResult{
					IsMatched: false,
					Result:    PackageError("target is not a grpc standard status error"),
				}
			}

			e, ok := k.(Error)
			if !ok {
				return &MatchResult{
					IsMatched: false,
					Result:    PackageError("Some of the given comparative error is not a blunder error, specifying FromGrpc is not allowed"),
				}
			}

			if e.GetErrorCode() == st.Message() {
				return &MatchResult{
					IsMatched: true,
					Result:    v,
				}
			}
		}

		if errors.Is(target, k) {
			return &MatchResult{
				IsMatched: true,
				Result:    v,
			}
		}
	}
	return &MatchResult{
		IsMatched: false,
	}
}
