package getComorbidity

import "context"

type Repository interface {
	GetComorbidityRepository(ctx context.Context) ([]GetComorbidityResponse, error)
}

type Service interface {
	GetComorbidityService(ctx context.Context) ([]GetComorbidityResponse, error)
}

type GetComorbidityResponse struct {
	Id                     int    `json:"id"`
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
}
