package goliath

type JsonObject struct{}

func (validator JsonObject) IsSatisfied(value interface{}) (ok bool) {
	if value == nil {
		return true
	}

	_, ok = value.(map[string]interface{})
	return
}

func (validator JsonObject) Message() string {
	return "must be json object and type of key is string"
}

func (validator JsonObject) cast(value interface{}) map[string]interface{} {
	return value.(map[string]interface{})
}
