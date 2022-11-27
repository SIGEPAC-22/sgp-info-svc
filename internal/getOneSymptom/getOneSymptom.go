package getOneSymptom

import "context"

type Service interface {
	GetOneSymptomSvc(ctx context.Context, id string) (GetOneSymptomResponse, error)
}

type Repository interface {
	GetOneSymptom(ctx context.Context, id string) (GetOneSymptomResponse, error)
}

type GetOneSymptomRequest struct {
	Id string `json:"id"`
}

type GetOneSymptomResponse struct {
	Id                 int    `json:"id"`
	NameSymptom        string `json:"nameSymptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
}
