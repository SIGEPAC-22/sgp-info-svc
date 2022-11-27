package getInfoPatient

import "context"

type Repository interface {
	GetPatientRepository(ctx context.Context) ([]GetInfoPatientResponse, error)
}

type Service interface {
	GetPatientService(ctx context.Context) ([]GetInfoPatientResponse, error)
}

type GetInfoPatientResponse struct {
	Id                           int    `json:"id"`
	FirstName                    string `json:"firstName"`
	SecondName                   string `json:"secondName"`
	LastName                     string `json:"lastName"`
	MotherLastName               string `json:"motherLastName"`
	PatientSex                   string `json:"patientSex"`
	DateBirth                    string `json:"dateBirth"`
	DocumentType                 string `json:"documentType"`
	DocumentNumber               string `json:"documentNumber"`
	CellphoneNumber              string `json:"cellphoneNumber"`
	PhoneNumber                  string `json:"phoneNumber"`
	ResponsibleFamily            string `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber string `json:"responsibleFamilyPhoneNumber"`
	Department                   string `json:"department"`
	Country                      string `json:"country"`
}
