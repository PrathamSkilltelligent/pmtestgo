package health

type HealthResponse struct {
	Message string `json:"message"`
}

func (res HealthResponse) ToJson() HealthResponse {
	return res
}

func NewHealthResponse(message string) HealthResponse {
	return HealthResponse{
		Message: message,
	}
}
