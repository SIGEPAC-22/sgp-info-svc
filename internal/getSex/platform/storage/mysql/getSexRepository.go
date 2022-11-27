package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	GetSex "sgp-info-svc/internal/getSex"
	"sgp-info-svc/kit/constants"
)

type GetSexRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetSexRepo(db *sql.DB, log log.Logger) *GetSexRepo {
	return &GetSexRepo{db, log}
}

func (g *GetSexRepo) GetSexRepository(ctx context.Context) ([]GetSex.GetSexResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT spt_id_sex, spt_gender_type FROM sex_patient WHERE spt_id_state_data = ?;", id)
	if errDB != nil {
		g.log.Log("Error while trying to get information for sex", constants.UUID, ctx.Value(constants.UUID))
		return []GetSex.GetSexResponse{}, errDB
	}
	defer rows.Close()
	var resp []GetSex.GetSexResponse
	for rows.Next() {
		var respDB SqlGetSex
		if err := rows.Scan(&respDB.Id, &respDB.NameSex); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []GetSex.GetSexResponse{}, err
		}
		resp = append(resp, GetSex.GetSexResponse{
			Id:      int(respDB.Id),
			NameSex: respDB.NameSex,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
