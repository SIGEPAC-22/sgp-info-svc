package mysql

import "time"

type SqlGetInfoPatientFile struct {
	IdPatient                    int        `json:"idPatient"`
	IdPatientFile                int        `json:"idPatientFile"`
	FullName                     *string    `json:"fullName"`
	DocumentType                 *string    `json:"documentType"`
	DocumentNumber               *string    `json:"documentNumber"`
	CellphoneNumber              *string    `json:"cellphoneNumber"`
	ResponsibleFamily            *string    `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber *string    `json:"responsibleFamilyPhoneNumber"`
	Sex                          *string    `json:"sex"`
	Pregnant                     bool       `json:"pregnant"`
	StatePatient                 *string    `json:"statePatient"`
	AdmissionDate                time.Time  `json:"admissionDate"`
	HighDate                     *time.Time `json:"highDate"`
	LowDate                      *time.Time `json:"lowDate"`
}

type SqlGetComorbidityPatient struct {
	IdPatientFile   int    `db:"fhs_id_patient_file"`
	NameComorbidity string `db:"cby_name_comorbidity"`
}

type SqlGetSymptomPatient struct {
	IdPatientFile int    `db:"fhs_id_patient_file"`
	NameSymptom   string `db:"stm_name_symptons"`
}
