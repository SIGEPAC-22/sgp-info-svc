package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getOnePatientFile"
	"sgp-info-svc/kit/constants"
	"time"
)

type GetOnePatientFileRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetOnePatientFileRepo(db *sql.DB, log log.Logger) *GetOnePatientFileRepo {
	return &GetOnePatientFileRepo{db: db, log: log}
}

func (g *GetOnePatientFileRepo) GetOnePatientFileRepository(ctx context.Context, id int) (getOnePatientFile.GetOnePatientFileResponse, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	rows := g.db.QueryRowContext(ctx, "SELECT pat.pat_id_patient, pfl.pfl_id_patient_file, CONCAT_WS(' ',pat.pat_first_name, pat.pat_second_name, pat.pat_first_last_name, pat.pat_second_last_name) AS Fullname, \ndct.dct_document_type_name,pat.pat_document_number, pat.pat_cellphone_number, pat.pat_reponsible_family, pat.pat_responsible_family_phone_number,\nsex.spt_gender_type, spt.spt_name_state_patient, pfl.pfl_admission_date, pfl.pfl_high_date, pfl.pfl_low_date\nFROM pat_patient AS pat\nINNER JOIN pfl_patient_file AS pfl ON pfl.pfl_id_patient = pat.pat_id_patient\nINNER JOIN sex_patient AS sex ON pat.pat_id_patient_sex = sex.spt_id_sex\nINNER JOIN spt_state_patient AS spt ON pfl.pfl_id_state_patient = spt.spt_id_state_patient\nINNER JOIN dct_document_type AS dct ON pat.pat_id_document_type = dct.dct_id_document_type WHERE pfl.pfl_id_patient_file = ? AND pat.pat_id_patient = ?;", id, id)
	g.log.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))
	var respDB SqlGetOnePatientFile
	if err := rows.Scan(&respDB.IdPatient, &respDB.IdPatientFile, &respDB.FullName, &respDB.DocumentType, &respDB.DocumentNumber, &respDB.CellphoneNumber, &respDB.ResponsibleFamily, &respDB.ResponsibleFamilyPhoneNumber, &respDB.Sex, &respDB.StatePatient, &respDB.AdmissionDate, &respDB.HighDate, &respDB.LowDate); err != nil {
		g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOnePatientFile.GetOnePatientFileResponse{}, err
	}

	resp := getOnePatientFile.GetOnePatientFileResponse{
		IdPatient:                    respDB.IdPatient,
		IdPatientFile:                respDB.IdPatientFile,
		FullName:                     respDB.FullName,
		DocumentType:                 respDB.DocumentType,
		DocumentNumber:               respDB.DocumentNumber,
		CellphoneNumber:              respDB.CellphoneNumber,
		ResponsibleFamily:            respDB.ResponsibleFamily,
		ResponsibleFamilyPhoneNumber: respDB.ResponsibleFamilyPhoneNumber,
		Sex:                          respDB.Sex,
		StatePatient:                 respDB.StatePatient,
		AdmissionDate:                respDB.AdmissionDate.Format(config.GetString("app-properties.getInfoPatient.dateBirth-Format")),
		HighDate:                     transformerPointer(respDB.HighDate),
		LowDate:                      transformerPointer(respDB.LowDate),
	}
	return resp, nil
}

func (g *GetOnePatientFileRepo) GetComorbidityPatient(ctx context.Context, idPatientFile int) ([]string, error) {

	var resp []string
	rows, errDB := g.db.QueryContext(ctx, "SELECT fhc_id_patient_file, cby.cby_name_comorbidity FROM cby_comorbidity as cby \nINNER JOIN fhc_file_has_cormobility AS fhc\nON cby.cby_id_comorbidity = fhc.fhc_id_conmorbilities\nINNER JOIN pfl_patient_file AS pfl\nON pfl.pfl_id_patient_file =  fhc.fhc_id_patient_file\nINNER JOIN pat_patient as pat\nON pat.pat_id_patient = pfl.pfl_id_patient\nWHERE fhc.fhc_id_patient_file = ?;", idPatientFile)
	if errDB != nil {
		g.log.Log("Error while trying to get information for Comorbidity-patientFile", constants.UUID, ctx.Value(constants.UUID))
		return nil, errDB
	}
	defer rows.Close()
	for rows.Next() {
		var respDB SqlGetComorbidityPatient
		if err := rows.Scan(&respDB.IdPatientFile, &respDB.NameComorbidity); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return nil, err
		}

		resp = append(resp, respDB.NameComorbidity)

	}

	if resp == nil {
		resp = append(resp, "not available")
	}

	return resp, nil
}

func (g *GetOnePatientFileRepo) GetSymptomPatient(ctx context.Context, idPatientFile int) ([]string, error) {
	var resp []string
	rows, errDB := g.db.QueryContext(ctx, "SELECT fhs_id_patient_file, stm.stm_name_symptons FROM stm_symptom as stm \nINNER JOIN fhs_file_has_sympton AS fhs\nON stm.stm_id_sympton = fhs.fhs_id_symptom\nINNER JOIN pfl_patient_file AS pfl\nON pfl.pfl_id_patient_file =  fhs.fhs_id_patient_file\nINNER JOIN pat_patient as pat\nON pat.pat_id_patient = pfl.pfl_id_patient\nWHERE fhs.fhs_id_patient_file = ?;", idPatientFile)
	if errDB != nil {
		g.log.Log("Error while trying to get information for Symptom-PatientFile", constants.UUID, ctx.Value(constants.UUID))
		return nil, errDB
	}

	defer rows.Close()
	for rows.Next() {
		var respDB SqlGetSymptomPatient
		if err := rows.Scan(&respDB.IdPatientFile, &respDB.NameSymptom); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return nil, err
		}

		resp = append(resp, respDB.NameSymptom)

	}
	if resp == nil {
		resp = append(resp, "not available")
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
