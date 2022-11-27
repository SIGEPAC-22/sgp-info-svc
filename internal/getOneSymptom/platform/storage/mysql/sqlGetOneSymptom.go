package mysql

type SqlOneGetSymptom struct {
	Id                 int64  `db:"stm_id_sympton"`
	NameSymptom        string `db:"stm_name_symptons"`
	DescriptionSymptom string `db:"stm_sympton_description"`
}
