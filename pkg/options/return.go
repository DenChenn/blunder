package options

type ReturnOption struct {
	ForGrpc      *bool
	ForGraphQL   *bool
	Log          *bool
	DiscordAlert *bool
}

func NewReturnOption() *ReturnOption {
	return &ReturnOption{}
}

func (o *ReturnOption) SetForGrpc(b bool) *ReturnOption {
	o.ForGrpc = &b
	return o
}

func (o *ReturnOption) SetForGraphQL(b bool) *ReturnOption {
	o.ForGraphQL = &b
	return o
}

func (o *ReturnOption) SetLog(b bool) *ReturnOption {
	o.Log = &b
	return o
}

func (o *ReturnOption) SetDiscordAlert(b bool) *ReturnOption {
	o.DiscordAlert = &b
	return o
}
