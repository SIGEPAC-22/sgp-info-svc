package mysql

import "time"

type SqlGetInfoPatient struct {
	Id                           int       `json:"id"`
	firstName                    string    `db:"pat_first_name"`
	secondName                   string    `db:"pat_second_name"`
	lastName                     string    `db:"pat_last_name"`
	motherLastName               string    `db:"pat_mothers_last_name"`
	patientSex                   string    `db:"spt_gender_type"`
	dateBirth                    time.Time `db:"pat_date_birth"`
	documentType                 string    `db:"dct_document_type_name"`
	documentNumber               string    `db:"pat_document_number"`
	cellphoneNumber              string    `db:"pat_cellphone_number"`
	phoneNumber                  string    `db:"pat_phone_number"`
	responsibleFamily            string    `db:"pat_reponsible_family"`
	responsibleFamilyPhoneNumber string    `db:"pat_responsible_family_phone_number"`
	department                   string    `db:"dpt_name_dapartment"`
	country                      string    `db:"cty_country_name"`
}
