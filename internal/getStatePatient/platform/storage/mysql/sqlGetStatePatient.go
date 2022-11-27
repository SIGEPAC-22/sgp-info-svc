package mysql

type SqlGetStatePatient struct {
	Id               int64  `db:"spt_id_state_patient"`
	NameStatePatient string `db:"spt_name_state_patient"`
}
