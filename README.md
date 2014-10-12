# Goliath

validation library for Golang

## Instalation

```
go get github.com/yagince/goliath
```

## Usage

```go
validation := NewValidation()
validation.Field("name").Required().MinLength(2).MaxLength(3)

var (
	result ValidationResult
	errors ValidtaionErrors
	error  ValidtaionError
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

### Required

```go
validation := NewValidation()
validation.Field("name").Required()
```

### Min and Max

```go
validation := NewValidation()
validation.Field("age").Min(0).Max(999)
validation.Field("rate").Min(0.5).Max(0.9)
```

### MinLength and MaxLength

```go
validation := NewValidation()
validation.Field("password").MinLength(6).MaxLength(8)
```

### Pattern

```go
validation := NewValidation()
validation.Field("num").Pattern("^[0-9]+$")
```

### Each ( for Array Value )

json
```json
{
  "ids":[1,2,3,4,5]
}
```
validation
```go
validation := NewValidation()
validation.Field("ids").Each().Min(0).Max(6)
```
