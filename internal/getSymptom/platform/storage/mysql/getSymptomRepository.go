package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getSymptom"
	"sgp-info-svc/kit/constants"
)

type GetSymptomRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetSymptomRepo(db *sql.DB, log log.Logger) *GetSymptomRepo {
	return &GetSymptomRepo{db, log}
}

func (g *GetSymptomRepo) GetSymptomRepository(ctx context.Context) ([]getSymptom.GetSymptomResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	statusActive := config.GetInt("app-properties.getComorbidity.idStatusActive")

	rows, errDB := g.db.QueryContext(ctx, "SELECT stm_id_sympton, stm_name_symptons, stm_sympton_description FROM sgp_info_svc.stm_symptom WHERE stm_id_state_data = ?;", statusActive)
	if errDB != nil {
		g.log.Log("Error while trying to get information for symptoms", constants.UUID, ctx.Value(constants.UUID))
		return []getSymptom.GetSymptomResponse{}, errDB
	}
	defer rows.Close()
	var resp []getSymptom.GetSymptomResponse
	for rows.Next() {
		var respDB SqlGetSymptom
		if err := rows.Scan(&respDB.id, &respDB.nameSymptom, &respDB.descriptionSymptom); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getSymptom.GetSymptomResponse{}, err
		}
		resp = append(resp, getSymptom.GetSymptomResponse{
			Id:                 respDB.id,
			NameSymptom:        respDB.nameSymptom,
			DescriptionSymptom: respDB.descriptionSymptom,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
