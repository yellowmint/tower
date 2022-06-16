package validate

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func Init() {
	validate = validator.New()
}

func Validate(s interface{}) error {
	return validate.Struct(s)
}
