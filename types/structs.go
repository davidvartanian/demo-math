package types

import "gopkg.in/go-playground/validator.v9"

type (
	Payload struct {
		A float64 `json:"a" validate:"required"`
		B float64 `json:"b" validate:"required"`
	}

	Result struct {
		Result float64 `json:"result,omitempty"`
		Error  string  `json:"error,omitempty"`
	}

	PayloadValidator struct {
		Validator *validator.Validate
	}
)
