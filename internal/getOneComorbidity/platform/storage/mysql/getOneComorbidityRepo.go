package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getOneComorbidity"
	"sgp-info-svc/kit/constants"
)

type getOneComorbidityRepo struct {
	db  *sql.DB
	log kitlog.Logger
}

func NewGetOneComorbidityRepo(db *sql.DB, log kitlog.Logger) *getOneComorbidityRepo {
	return &getOneComorbidityRepo{db: db, log: log}
}

func (g *getOneComorbidityRepo) GetOneComorbidity(ctx context.Context, id string) (getOneComorbidity.GetOneComorbidityResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	statusActive := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows := g.db.QueryRowContext(ctx, "SELECT cby_id_comorbidity, cby_name_comorbidity, cby_comorbidity_description FROM cby_comorbidity where cby_id_comorbidity = ? AND cby_id_state_data = ?", id, statusActive)
	g.log.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))

	var respDB SqlOneGetComorbidity

	if err := rows.Scan(&respDB.Id, &respDB.NameComorbidity, &respDB.DescriptionComorbidity); err != nil {
		g.log.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOneComorbidity.GetOneComorbidityResponse{}, errors.New("Data not found")
	}

	resp := getOneComorbidity.GetOneComorbidityResponse{
		Id:                     int(respDB.Id),
		NameComorbidity:        respDB.NameComorbidity,
		DescriptionComorbidity: respDB.DescriptionComorbidity,
	}

	return resp, nil
}
