package domain

type ServiceResponse struct {
	Error      error
	StatusCode int
	Data       map[string]interface{}
}
