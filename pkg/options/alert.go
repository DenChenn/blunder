package options

type Alert struct {
	DiscordAlert bool
	LineAlert    bool
}

func NewAlert() *Alert {
	return &Alert{}
}

func (o *Alert) SetDiscordAlert(b bool) *Alert {
	o.DiscordAlert = b
	return o
}

func (o *Alert) SetLineAlert(b bool) *Alert {
	o.LineAlert = b
	return o
}
