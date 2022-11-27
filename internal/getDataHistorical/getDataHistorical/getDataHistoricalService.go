package getDataHistorical

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getDataHistorical"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getDataHistorical.Repository
	logger log.Logger
}

func NewService(repoBD getDataHistorical.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetDataHistoricalService(ctx context.Context) ([]getDataHistorical.GetDataHistoricalResponse, error) {
	s.logger.Log("Start Endpoint GetDataHistorical", constants.UUID, ctx.Value(constants.UUID))

	resp, err := s.RepoDB.GetDataHistoricalRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getDataHistorical.GetDataHistoricalResponse{}, err
	}

	for _, respUpdate := range resp {
		IdPatient := respUpdate.IdPatient
		IdPatientFile := respUpdate.IdPatientFile

		errUpdate := s.RepoDB.UpdateIdStatusBot(ctx, IdPatient, IdPatientFile)
		if errUpdate != nil {
			s.logger.Log("Error no execute update", constants.UUID, ctx.Value(constants.UUID))
			continue
		}
	}

	return resp, nil

}
