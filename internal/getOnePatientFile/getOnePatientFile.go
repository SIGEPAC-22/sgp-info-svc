package getOnePatientFile

import "context"

type Repository interface {
	GetOnePatientFileRepository(ctx context.Context, id int) (GetOnePatientFileResponse, error)
	GetComorbidityPatient(ctx context.Context, idPatientFile int) ([]string, error)
	GetSymptomPatient(ctx context.Context, idPatientFile int) ([]string, error)
}

type Service interface {
	GetOnePatientFileService(ctx context.Context, id string) (GetOnePatientFileResponse, error)
}

type GetOnePatientFileResponse struct {
	IdPatient                    int      `json:"idPatient"`
	IdPatientFile                int      `json:"idPatientFile"`
	FullName                     *string  `json:"fullName"`
	DocumentType                 *string  `json:"documentType"`
	DocumentNumber               *string  `json:"documentNumber"`
	CellphoneNumber              *string  `json:"cellphoneNumber"`
	ResponsibleFamily            *string  `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber *string  `json:"responsibleFamilyPhoneNumber"`
	Sex                          *string  `json:"sex"`
	Pregnant                     string   `json:"pregnant"`
	StatePatient                 *string  `json:"statePatient"`
	AdmissionDate                string   `json:"admissionDate"`
	HighDate                     string   `json:"highDate"`
	LowDate                      string   `json:"lowDate"`
	Comorbidity                  []string `json:"comorbidity"`
	Symptom                      []string `json:"symptom"`
}
