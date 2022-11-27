package mysql

import "time"

type SqlGetDataHistorical struct {
	IdPatient      int        `db:"pat_id_patient"`
	IdPatientFile  int        `db:"pfl_id_patient_file"`
	FirstName      string     `db:"pat_first_name"`
	SecondName     string     `db:"pat_second_name"`
	FirstLastName  string     `db:"pat_first_last_name"`
	SecondLastName string     `db:"pat_second_last_name"`
	AdmissionDate  time.Time  `db:"pfl_admission_date"`
	HighDate       *time.Time `db:"pfl_high_date"`
	LowDate        *time.Time `db:"pfl_low_date"`
}
