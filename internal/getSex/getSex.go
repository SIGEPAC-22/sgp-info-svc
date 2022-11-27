package GetSex

import "context"

type Repository interface {
	GetSexRepository(ctx context.Context) ([]GetSexResponse, error)
}

type Service interface {
	GetSexService(ctx context.Context) ([]GetSexResponse, error)
}

type GetSexResponse struct {
	Id      int    `json:"id"`
	NameSex string `json:"nameSex"`
}
