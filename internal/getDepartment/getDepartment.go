package getDepartment

import "context"

type Repository interface {
	GetDepartmentRepository(ctx context.Context) ([]GetDepartmentResponse, error)
}

type Service interface {
	GetDepartmentService(ctx context.Context) ([]GetDepartmentResponse, error)
}

type GetDepartmentResponse struct {
	Id             int    `json:"id"`
	NameDepartment string `json:"nameDepartment"`
}
