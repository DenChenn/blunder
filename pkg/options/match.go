package options

type MatchOption struct {
	FromGrpc *bool
}

func NewMatchOption() *MatchOption {
	return &MatchOption{}
}

func (o *MatchOption) SetFromGrpc(b bool) *MatchOption {
	o.FromGrpc = &b
	return o
}
