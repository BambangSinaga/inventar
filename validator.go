package inventar

import _validator "gopkg.in/go-playground/validator.v9"

// Validate is struct for validate
type Validate struct {
	validator *_validator.Validate
}

// ValidateStruct is method implementation for validating struct
func (v Validate) ValidateStruct(data interface{}) error {
	return v.validator.Struct(data)
}

// NewValidator is function to init validator
func NewValidator() Validate {
	val := _validator.New()
	return Validate{
		validator: val,
	}
}
