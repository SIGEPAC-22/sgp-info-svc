package getDepartmentService

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getDepartment"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getDepartment.Repository
	logger log.Logger
}

func NewService(repoBD getDepartment.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoBD, logger: logger}
}

func (s *Service) GetDepartmentService(ctx context.Context) ([]getDepartment.GetDepartmentResponse, error) {
	s.logger.Log("Start Endpoint GetDepartment", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetDepartmentRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getDepartment.GetDepartmentResponse{}, err
	}
	return resp, nil

}
