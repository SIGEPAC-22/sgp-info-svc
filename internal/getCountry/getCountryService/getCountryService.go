package getCountryService

import (
	"context"
	"github.com/go-kit/log"
	GetCountry "sgp-info-svc/internal/getCountry"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB GetCountry.Repository
	logger log.Logger
}

func NewService(repoBD GetCountry.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetCountryService(ctx context.Context) ([]GetCountry.GetCountryResponse, error) {
	s.logger.Log("Start Endpoint GetCountry", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetCountryRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []GetCountry.GetCountryResponse{}, err
	}
	return resp, nil

}
