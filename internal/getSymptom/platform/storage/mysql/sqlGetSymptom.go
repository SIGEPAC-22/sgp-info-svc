package mysql

type SqlGetSymptom struct {
	id                 int    `db:"stm_id_sympton"`
	nameSymptom        string `db:"stm_name_symptons"`
	descriptionSymptom string `db:"stm_sympton_description"`
}
