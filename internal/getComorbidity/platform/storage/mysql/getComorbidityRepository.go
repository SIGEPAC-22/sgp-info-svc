package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getComorbidity"
	"sgp-info-svc/kit/constants"
)

type GetComorbidityRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetComorbidityRepo(db *sql.DB, log log.Logger) *GetComorbidityRepo {
	return &GetComorbidityRepo{db, log}
}

func (g *GetComorbidityRepo) GetComorbidityRepository(ctx context.Context) ([]getComorbidity.GetComorbidityResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT cby_id_comorbidity, cby_name_comorbidity, cby_comorbidity_description FROM cby_comorbidity where cby_id_state_data = ?;", id)
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []getComorbidity.GetComorbidityResponse{}, errDB
	}
	defer rows.Close()
	var resp []getComorbidity.GetComorbidityResponse
	for rows.Next() {
		var respDB SqlGetComorbidity
		if err := rows.Scan(&respDB.Id, &respDB.NameComorbidity, &respDB.DescriptionComorbidity); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getComorbidity.GetComorbidityResponse{}, err
		}
		resp = append(resp, getComorbidity.GetComorbidityResponse{
			Id:                     int(respDB.Id),
			NameComorbidity:        respDB.NameComorbidity,
			DescriptionComorbidity: respDB.DescriptionComorbidity,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
