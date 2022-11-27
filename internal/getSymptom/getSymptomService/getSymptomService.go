package getSymptomService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getSymptom"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getSymptom.Repository
	logger log.Logger
}

func NewService(repoBD getSymptom.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetSymptomService(ctx context.Context) ([]getSymptom.GetSymptomResponse, error) {
	s.logger.Log("Start Endpoint GetSymptom", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetSymptomRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getSymptom.GetSymptomResponse{}, err
	}
	return resp, nil
}
