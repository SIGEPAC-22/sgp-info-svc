package getTypeDocumentService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/GetTypeDocument"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB GetTypeDocument.Repository
	logger log.Logger
}

func NewService(repoBD GetTypeDocument.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetTypeDocumentService(ctx context.Context) ([]GetTypeDocument.GetTypeDocumentResponse, error) {
	s.logger.Log("Start Endpoint GetTypeDocument", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetTypeDocumentRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []GetTypeDocument.GetTypeDocumentResponse{}, err
	}
	return resp, nil

}
