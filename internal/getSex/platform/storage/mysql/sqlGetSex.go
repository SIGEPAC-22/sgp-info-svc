package mysql

type SqlGetSex struct {
	Id      int64  `db:"spt_id_sex"`
	NameSex string `db:"spt_gender_type"`
}
