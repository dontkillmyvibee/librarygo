package validation

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *FieldError) Error() string {
	return e.Field + ": " + e.Message
}
