package goliath

type Required struct{}

func (validator Required) IsSatisfied(value interface{}) bool {
	if value == nil {
		return false
	}

	if str, ok := value.(string); ok {
		return len(str) > 0
	}

	return true
}

func (validator Required) Message() string {
	return "required"
}
