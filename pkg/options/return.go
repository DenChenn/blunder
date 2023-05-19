package options

type ReturnOption struct {
	Log          *bool
	DiscordAlert *bool
	LineAlert    *bool
}

func NewReturnOption() *ReturnOption {
	return &ReturnOption{}
}

func (o *ReturnOption) SetLog(b bool) *ReturnOption {
	o.Log = &b
	return o
}

func (o *ReturnOption) SetDiscordAlert(b bool) *ReturnOption {
	o.DiscordAlert = &b
	return o
}

func (o *ReturnOption) SetLineAlert(b bool) *ReturnOption {
	o.LineAlert = &b
	return o
}
