package mysql

type SqlGetDepartment struct {
	Id             int64  `db:"dct_id_document_type"`
	NameDepartment string `db:"dct_document_type_name"`
}
