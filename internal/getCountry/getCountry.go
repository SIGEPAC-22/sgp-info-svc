package GetCountry

import "context"

type Repository interface {
	GetCountryRepository(ctx context.Context) ([]GetCountryResponse, error)
}

type Service interface {
	GetCountryService(ctx context.Context) ([]GetCountryResponse, error)
}

type GetCountryResponse struct {
	Id          int    `json:"id"`
	NameCountry string `json:"nameCountry"`
}
