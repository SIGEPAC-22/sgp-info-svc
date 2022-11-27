package getInfoPatientService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getInfoPatient"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getInfoPatient.Repository
	logger log.Logger
}

func NewService(repoDB getInfoPatient.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoDB, logger: logger}
}

func (s *Service) GetPatientService(ctx context.Context) ([]getInfoPatient.GetInfoPatientResponse, error) {
	s.logger.Log("Start Endpoint GetInfoPatient", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetPatientRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getInfoPatient.GetInfoPatientResponse{}, err
	}
	return resp, nil
}
