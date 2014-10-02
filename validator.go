package goliath

type Validator interface {
	IsSatisfied(v interface{}) bool
	Message() string
}

func ChoiceMessage(m ...string) string {
	empty := ""
	for _, val := range m {
		if val != empty {
			return val
		}
	}
	return empty
}
