package usecase

import (
	health "gefferson.com.br/geffws/api/usecase/health"
)

type (
	UseCases interface {
		NewPingGetUseCase() *health.PingResponse
	}
	useCases struct {
	}
)

func NewUseCases() UseCases {
	return &useCases{}
}

func (u *useCases) NewPingGetUseCase() *health.PingResponse {
	return health.GetPingResponse()
}
