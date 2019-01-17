package common

import "gopkg.in/go-playground/validator.v9"

var validate = validator.New()

func Validate() *validator.Validate{
	return validate
}