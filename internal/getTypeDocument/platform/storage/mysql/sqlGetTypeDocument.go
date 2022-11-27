package mysql

type SqlGetTypeDocument struct {
	Id               int64  `db:"dct_id_document_type"`
	NameTypeDocument string `db:"dct_document_type_name"`
}
