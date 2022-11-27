package getSexService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/GetSex"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB GetSex.Repository
	logger log.Logger
}

func NewService(repoBD GetSex.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetSexService(ctx context.Context) ([]GetSex.GetSexResponse, error) {
	s.logger.Log("Start Endpoint GetSex", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetSexRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []GetSex.GetSexResponse{}, err
	}
	return resp, nil

}
