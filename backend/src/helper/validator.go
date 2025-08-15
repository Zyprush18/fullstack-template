package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
)
	
func Validation(req interface{}) error {
	validate := validator.New()

	if err:= validate.Struct(&req);err != nil {
		var validateErr validator.ValidationErrors
		if errors.As(err, &validateErr) {
			for _, v := range validateErr {
				return  errors.New(v.Field() + v.Tag())
			}
		}
	}

	return nil
}