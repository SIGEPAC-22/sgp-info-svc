package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getCountry"
	"sgp-info-svc/kit/constants"
)

type GetCountryRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetCountryRepo(db *sql.DB, log log.Logger) *GetCountryRepo {
	return &GetCountryRepo{db, log}
}

func (g *GetCountryRepo) GetCountryRepository(ctx context.Context) ([]GetCountry.GetCountryResponse, error) {

	rows, errDB := g.db.QueryContext(ctx, "SELECT cty_id_country, cty_country_name FROM cty_country;")
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []GetCountry.GetCountryResponse{}, errDB
	}
	defer rows.Close()
	var resp []GetCountry.GetCountryResponse
	for rows.Next() {
		var respDB SqlGetCountry
		if err := rows.Scan(&respDB.Id, &respDB.NameCountry); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []GetCountry.GetCountryResponse{}, err
		}
		resp = append(resp, GetCountry.GetCountryResponse{
			Id:          int(respDB.Id),
			NameCountry: respDB.NameCountry,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
