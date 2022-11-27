package getSymptom

import "context"

type Repository interface {
	GetSymptomRepository(ctx context.Context) ([]GetSymptomResponse, error)
}

type Service interface {
	GetSymptomService(ctx context.Context) ([]GetSymptomResponse, error)
}

type GetSymptomResponse struct {
	Id                 int    `json:"id"`
	NameSymptom        string `json:"name_symptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
}
