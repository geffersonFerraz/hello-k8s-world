package usecase

type PingResponse struct {
	Message string `json:"message"`
}

func GetPingResponse() *PingResponse {
	return &PingResponse{Message: "pong"}
}
