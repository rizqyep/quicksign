package domain

type ServiceResponse struct {
	Error                 string
	StatusCode            int
	Data                  map[string]interface{}
	CustomResponseMessage string
}
