package blunder

type Condition struct {
	detail map[error]Error
}

func NewCondition(givenCondition ...map[error]Error) *Condition {
	if len(givenCondition) > 0 {
		initCondition := make(map[error]Error)
		for _, m := range givenCondition {
			for k, v := range m {
				initCondition[k] = v
			}
		}
		return &Condition{
			detail: initCondition,
		}
	}
	return &Condition{
		detail: make(map[error]Error),
	}
}

func (m *Condition) Add(is error, shouldReturn Error) *Condition {
	m.detail[is] = shouldReturn
	return m
}

func (m *Condition) AddMany(isThese []error, shouldReturn Error) *Condition {
	for _, is := range isThese {
		m.detail[is] = shouldReturn
	}
	return m
}
