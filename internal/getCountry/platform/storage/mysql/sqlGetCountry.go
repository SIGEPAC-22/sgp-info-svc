package mysql

type SqlGetCountry struct {
	Id          int64  `db:"cty_id_country"`
	NameCountry string `db:"cty_country_name"`
}
