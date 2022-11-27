package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getOneSymptom"
	"sgp-info-svc/kit/constants"
)

type getOneSymptomRepo struct {
	db  *sql.DB
	log kitlog.Logger
}

func NewGetOneSymptomRepo(db *sql.DB, log kitlog.Logger) *getOneSymptomRepo {
	return &getOneSymptomRepo{db: db, log: log}
}

func (g *getOneSymptomRepo) GetOneSymptom(ctx context.Context, id string) (getOneSymptom.GetOneSymptomResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	statusActive := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows := g.db.QueryRowContext(ctx, "SELECT stm_id_sympton, stm_name_symptons, stm_sympton_description FROM stm_symptom where stm_id_sympton = ? AND stm_id_state_data = ?", id, statusActive)
	g.log.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))

	var respDB SqlOneGetSymptom

	if err := rows.Scan(&respDB.Id, &respDB.NameSymptom, &respDB.DescriptionSymptom); err != nil {
		g.log.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOneSymptom.GetOneSymptomResponse{}, errors.New("Data not found")
	}

	resp := getOneSymptom.GetOneSymptomResponse{
		Id:                 int(respDB.Id),
		NameSymptom:        respDB.NameSymptom,
		DescriptionSymptom: respDB.DescriptionSymptom,
	}

	return resp, nil
}
