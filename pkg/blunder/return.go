package blunder

import "github.com/DenChenn/blunder/pkg/options"

func (mr *MatchResult) Return(opt *options.ReturnOption) error {
	// TODO: logging, alerting, etc.
	if mr.IsMatched {
		return mr.Result
	}
	return nil
}
