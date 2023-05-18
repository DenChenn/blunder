package blunder

type Map struct {
	detail map[error]Error
}

func NewMap(givenMap ...map[error]Error) *Map {
	if len(givenMap) > 0 {
		initMap := make(map[error]Error)
		for _, m := range givenMap {
			for k, v := range m {
				initMap[k] = v
			}
		}
		return &Map{
			detail: initMap,
		}
	}
	return &Map{
		detail: make(map[error]Error),
	}
}

func (m *Map) Add(target error, comp Error) *Map {
	m.detail[target] = comp
	return m
}
