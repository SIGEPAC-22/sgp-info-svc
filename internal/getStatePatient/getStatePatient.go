package getStatePatient

import "context"

type Repository interface {
	GetStatePatientRepository(ctx context.Context) ([]GetStatePatientResponse, error)
}

type Service interface {
	GetStatePatientService(ctx context.Context) ([]GetStatePatientResponse, error)
}

type GetStatePatientResponse struct {
	Id               int    `json:"id"`
	NameStatePatient string `json:"nameStatePatient"`
}
