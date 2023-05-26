package blunder

// Condition is a struct to define a map of error and Error, which is used to match the error
type Condition struct {
	detail map[error]Error
}

// NewCondition is a function to create a new condition
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

// OneToOne is a function to add a single condition to the condition map, just like OneToOne function in mathematics
func (m *Condition) OneToOne(is error, shouldReturn Error) *Condition {
	m.detail[is] = shouldReturn
	return m
}

// ManyToOne is a function to add multiple conditions to the condition map, just like ManyToOne function in mathematics
func (m *Condition) ManyToOne(isThese []error, shouldReturn Error) *Condition {
	for _, is := range isThese {
		m.detail[is] = shouldReturn
	}
	return m
}
