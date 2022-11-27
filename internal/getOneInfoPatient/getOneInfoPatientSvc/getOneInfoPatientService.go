package getOneInfoPatientSvc

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getOneInfoPatient"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getOneInfoPatient.Repository
	logger log.Logger
}

func NewService(repoDB getOneInfoPatient.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoDB, logger: logger}
}

func (s *Service) GetOnePatientService(ctx context.Context, id string) (getOneInfoPatient.GetOneInfoPatientResponse, error) {
	s.logger.Log("Start Endpoint GetInfoPatient", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetOnePatientRepository(ctx, id)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return getOneInfoPatient.GetOneInfoPatientResponse{}, err
	}
	return resp, nil
}
