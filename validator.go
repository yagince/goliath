package goliath

type Validator interface {
	IsSatisfied(v interface{}) bool
	Message() string
}

type CustomMessageValidator struct {
	Validator
	CustomMessage string
}

func (validator CustomMessageValidator) Message() string {
	return validator.CustomMessage
}
