package getStatePatientService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getStatePatient"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getStatePatient.Repository
	logger log.Logger
}

func NewService(repoBD getStatePatient.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetStatePatientService(ctx context.Context) ([]getStatePatient.GetStatePatientResponse, error) {
	s.logger.Log("Start Endpoint GetStatePatient", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetStatePatientRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getStatePatient.GetStatePatientResponse{}, err
	}
	return resp, nil

}
