package errbank

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrUnimplemetedEntity = errors.New("unimplemented entity")
)

type (
	FieldError struct {
		Field string `json:"field"`
		Error string `json:"error"`
	}

	Error string

	ValidationError        []FieldError
	ValidationAdvanceError struct {
		StatusCode int
		Err        error
	}
	AdvanceQueryValidationError []FieldError

	TranslationError struct {
		StatusCode int
		Errors     ValidationError
	}
)

func (e ValidationAdvanceError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e Error) Error() string {
	return string(e)
}

func (fe FieldError) AddPrefix(prefix string) FieldError {
	if prefix == "" {
		return fe
	}

	fe.Field = fmt.Sprintf("%s.%s", prefix, fe.Field)
	return fe
}

func (te TranslationError) Error() string {
	if te.Errors == nil || len(te.Errors) <= 0 {
		return "'validation error, err': JSON(nil)"
	}

	_byte, _ := json.Marshal(te.Errors)
	return fmt.Sprintf("'validation error, err': JSON(%s)", string(_byte))
}

func (te TranslationError) WithPrefix(prefix string) ValidationError {
	_ve := make(ValidationError, len(te.Errors))
	for _, field := range te.Errors {
		_ve = append(_ve, field.AddPrefix(prefix))
	}

	return _ve
}

func (ve ValidationError) Error() string {
	if ve == nil || len(ve) <= 0 {
		return "'validation error, err': JSON(nil)"
	}

	_byte, _ := json.Marshal(ve)
	return fmt.Sprintf("'validation error, err': JSON(%s)", string(_byte))
}

func (ve ValidationError) WithPrefix(prefix string) ValidationError {
	_ve := make(ValidationError, len(ve))
	for _, field := range ve {
		_ve = append(_ve, field.AddPrefix(prefix))
	}

	return _ve
}

func (ve AdvanceQueryValidationError) Error() string {
	if ve == nil || len(ve) <= 0 {
		return "'advance query validation error, err': JSON(nil)"
	}

	_byte, _ := json.Marshal(ve)
	return fmt.Sprintf("'advance query validation error, err': JSON(%s)", string(_byte))
}
