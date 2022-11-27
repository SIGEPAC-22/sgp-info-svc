package getOneComorbidity

import "context"

type Service interface {
	GetOneComorbiditySvc(ctx context.Context, id string) (GetOneComorbidityResponse, error)
}

type Repository interface {
	GetOneComorbidity(ctx context.Context, id string) (GetOneComorbidityResponse, error)
}

type GetOneComorbidityRequest struct {
	Id string `json:"id"`
}

type GetOneComorbidityResponse struct {
	Id                     int    `json:"id"`
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
}
