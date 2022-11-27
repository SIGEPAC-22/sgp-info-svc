package getComorbidityService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getComorbidity"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getComorbidity.Repository
	logger log.Logger
}

func NewService(repoBD getComorbidity.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetComorbidityService(ctx context.Context) ([]getComorbidity.GetComorbidityResponse, error) {
	s.logger.Log("Start Endpoint GetComorbidity", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetComorbidityRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getComorbidity.GetComorbidityResponse{}, err
	}
	return resp, nil

}
