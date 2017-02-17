package controllers

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type ValidationError struct {
	Field string
	Kind  reflect.Kind
	Type  reflect.Type
	Value interface{}
}

func ValidationErrors(errs error) []ValidationError {
	var resErrors []ValidationError

	for _, err := range errs.(validator.ValidationErrors) {
		resErrors = append(resErrors, ValidationError{
			Field: err.Field(),
			Kind:  err.Kind(),
			Type:  err.Type(),
			Value: err.Value(),
		})
	}

	return resErrors
}
