package types

// Validate Payload for required values
// Known issue: Validator doesn't recognise 0 as value and fails
func (cv *PayloadValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
