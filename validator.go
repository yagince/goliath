package goliath

type Validator interface {
	IsSatisfied(v interface{}) bool
	Message() string
}
