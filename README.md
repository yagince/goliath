# Golidation

validation library for Golang

## Instalation

```
go get github.com/yagince/golidation
```

## Usage

```go
validation := NewValidation()
validation.Field("name").Required().MinLength(2).MaxLength(3)

var (
	result ValidationResult
	errors Errors
	error  Error
)

result = validation.Validate(map[string]interface{}{})
result.HasError() // => true

errors = result.Errors()
error, ok := errors["name"]
error.Message() // => "required"


result = validation.Validate(map[string]interface{}{"name": "1"})
errors = result.Errors()
error, ok := errors["name"]
error.Message() // => "must have 2 or more elements"
```
