package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getDataHistorical"
	"sgp-info-svc/kit/constants"
	"time"
)

type GetDataHistoricalRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetDataHistoricalRepo(db *sql.DB, log log.Logger) *GetDataHistoricalRepo {
	return &GetDataHistoricalRepo{db, log}
}

func (g *GetDataHistoricalRepo) GetDataHistoricalRepository(ctx context.Context) ([]getDataHistorical.GetDataHistoricalResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	status := config.GetBool("app-properties.getDataHistorical")

	rows, errDB := g.db.QueryContext(ctx, "SELECT pat_id_patient, pfl_id_patient_file, pat_first_name, pat_second_name, pat_first_last_name, pat_second_last_name, pfl_admission_date, pfl_high_date, pfl_low_date FROM pat_patient\nINNER JOIN pfl_patient_file AS pfl\nON pfl_id_patient = pat_id_patient WHERE pft_status_bot = ?;", status)
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []getDataHistorical.GetDataHistoricalResponse{}, errDB
	}
	defer rows.Close()
	var resp []getDataHistorical.GetDataHistoricalResponse
	for rows.Next() {
		var respDB SqlGetDataHistorical
		if err := rows.Scan(&respDB.IdPatient, &respDB.IdPatientFile, &respDB.FirstName, &respDB.SecondName, &respDB.FirstLastName, &respDB.SecondLastName, &respDB.AdmissionDate, &respDB.HighDate, &respDB.LowDate); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getDataHistorical.GetDataHistoricalResponse{}, err
		}
		resp = append(resp, getDataHistorical.GetDataHistoricalResponse{
			IdPatient:      respDB.IdPatient,
			IdPatientFile:  respDB.IdPatientFile,
			FirstName:      respDB.FirstName,
			SecondName:     respDB.SecondName,
			FirstLastName:  respDB.FirstLastName,
			SecondLastName: respDB.SecondLastName,
			AdmissionDate:  respDB.AdmissionDate.Format(config.GetString("app-PropertiesUpdateBot.date")),
			HighDate:       transformerPointer(respDB.HighDate),
			LowDate:        transformerPointer(respDB.LowDate),
		})
	}

	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
func transformerPointer(date *time.Time) string {

	if date != nil {
		var dateConverter string

		config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
		dateConverter = date.Format(config.GetString("app-PropertiesUpdateBot.date"))
		return dateConverter
	} else {
		return "0000-00-00"
	}
}

func (g *GetDataHistoricalRepo) UpdateIdStatusBot(ctx context.Context, idPatient, idPatientFile int) error {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	false := config.GetBool("app-PropertiesUpdateBot.false")
	true := config.GetBool("app-PropertiesUpdateBot.true")

	sql, err := g.db.ExecContext(ctx, "UPDATE pfl_patient_file SET pft_status_bot = ? WHERE pfl_id_patient_file = ? AND pfl_id_patient = ? AND pft_status_bot = ?;", false, idPatientFile, idPatient, true)
	g.log.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		g.log.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			g.log.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return err
		}
	}
	return nil
}
