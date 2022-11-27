package getOnePatientFileService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getOnePatientFile"
	"sgp-info-svc/kit/constants"
	"strconv"
)

type Service struct {
	RepoDB getOnePatientFile.Repository
	logger log.Logger
}

func NewService(repoDB getOnePatientFile.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoDB, logger: logger}
}

func (s *Service) GetOnePatientFileService(ctx context.Context, id string) (getOnePatientFile.GetOnePatientFileResponse, error) {
	s.logger.Log("Start Endpoint GetInfoPatientFile", constants.UUID, ctx.Value(constants.UUID))
	idConverter, _ := strconv.Atoi(id)
	resp, err := s.RepoDB.GetOnePatientFileRepository(ctx, idConverter)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return getOnePatientFile.GetOnePatientFileResponse{}, err
	}

	respComorbidity, err := s.RepoDB.GetComorbidityPatient(ctx, resp.IdPatientFile)
	if err != nil {
		s.logger.Log("Error, not get response for comorbidity", "Error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOnePatientFile.GetOnePatientFileResponse{}, err
	}

	resp.Comorbidity = respComorbidity

	respSymptom, err := s.RepoDB.GetSymptomPatient(ctx, resp.IdPatientFile)
	if err != nil {
		s.logger.Log("Error, not get response for symptom", "Error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOnePatientFile.GetOnePatientFileResponse{}, err
	}

	resp.Symptom = respSymptom

	return resp, nil
}
