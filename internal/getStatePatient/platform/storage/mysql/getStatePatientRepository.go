package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getStatePatient"
	"sgp-info-svc/kit/constants"
)

type GetStatePatientRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetStatePatientRepo(db *sql.DB, log log.Logger) *GetStatePatientRepo {
	return &GetStatePatientRepo{db, log}
}

func (g *GetStatePatientRepo) GetStatePatientRepository(ctx context.Context) ([]getStatePatient.GetStatePatientResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")

	rows, errDB := g.db.QueryContext(ctx, "SELECT spt_id_state_patient, spt_name_state_patient FROM spt_state_patient WHERE spt_id_state_data = ?;", id)
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []getStatePatient.GetStatePatientResponse{}, errDB
	}
	defer rows.Close()
	var resp []getStatePatient.GetStatePatientResponse
	for rows.Next() {
		var respDB SqlGetStatePatient
		if err := rows.Scan(&respDB.Id, &respDB.NameStatePatient); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getStatePatient.GetStatePatientResponse{}, err
		}
		resp = append(resp, getStatePatient.GetStatePatientResponse{
			Id:               int(respDB.Id),
			NameStatePatient: respDB.NameStatePatient,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
