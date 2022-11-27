package mysql

type SqlOneGetComorbidity struct {
	Id                     int64  `db:"cby_id_comorbidity"`
	NameComorbidity        string `db:"cby_name_comorbidity"`
	DescriptionComorbidity string `db:"cby_comorbidity_description"`
}
