package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomError struct {
	Errors map[string]string
}

func NewValidatorError(err error) CustomError {
	res := CustomError{}
	res.Errors = make(map[string]string)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// handles other err type
		res.Errors["Object"] = fmt.Sprintf(" %s", err.Error())
	}

	for _, v := range errs {
		field := v.Field()
		tag := v.Tag()
		res.Errors[field] = fmt.Sprintf("The field is %s", tag)
	}

	return res
}
