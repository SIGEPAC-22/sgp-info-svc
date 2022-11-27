package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/getInfoPatient"
	"sgp-info-svc/kit/constants"
)

type GetInfoPatientRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetInfoPatientRepo(db *sql.DB, log log.Logger) *GetInfoPatientRepo {
	return &GetInfoPatientRepo{db: db, log: log}
}

func (g *GetInfoPatientRepo) GetPatientRepository(ctx context.Context) ([]getInfoPatient.GetInfoPatientResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)

	rows, errDB := g.db.QueryContext(ctx, "SELECT pat.pat_id_patient, pat.pat_first_name, pat.pat_second_name, pat.pat_first_last_name, pat.pat_second_last_name, pat.pat_date_birth, doctype.dct_document_type_name,\npat.pat_document_number,pat.pat_cellphone_number, pat.pat_phone_number, pat.pat_reponsible_family, pat.pat_responsible_family_phone_number,\nsexp.spt_gender_type, dep.dpt_name_dapartment, con.cty_country_name\nFROM pat_patient AS pat\nINNER JOIN sex_patient AS sexp\nON pat.pat_id_patient_sex = sexp.spt_id_sex\nINNER JOIN dct_document_type AS doctype\nON pat.pat_id_document_type = doctype.dct_id_document_type\nINNER JOIN dpt_department AS dep\nON pat.pat_id_department = dep.dpt_id_department\nINNER JOIN cty_country AS con\nON con.cty_id_country = pat.pat_id_country;")
	if errDB != nil {
		g.log.Log("Error while trying to get information for patient", constants.UUID, ctx.Value(constants.UUID))
		return []getInfoPatient.GetInfoPatientResponse{}, errDB
	}
	defer rows.Close()
	var resp []getInfoPatient.GetInfoPatientResponse
	for rows.Next() {
		var respDB SqlGetInfoPatient
		if err := rows.Scan(&respDB.Id, &respDB.firstName, &respDB.secondName, &respDB.lastName, &respDB.motherLastName, &respDB.dateBirth, &respDB.documentType, &respDB.documentNumber, &respDB.cellphoneNumber, &respDB.phoneNumber, &respDB.responsibleFamily, &respDB.responsibleFamilyPhoneNumber, &respDB.patientSex, &respDB.department, &respDB.country); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getInfoPatient.GetInfoPatientResponse{}, err
		}
		resp = append(resp, getInfoPatient.GetInfoPatientResponse{
			Id:                           respDB.Id,
			FirstName:                    respDB.firstName,
			SecondName:                   respDB.secondName,
			LastName:                     respDB.lastName,
			MotherLastName:               respDB.motherLastName,
			PatientSex:                   respDB.patientSex,
			DateBirth:                    respDB.dateBirth.Format(config.GetString("app-properties.getInfoPatient.dateBirth-Format")),
			DocumentType:                 respDB.documentType,
			DocumentNumber:               respDB.documentNumber,
			CellphoneNumber:              respDB.cellphoneNumber,
			PhoneNumber:                  respDB.phoneNumber,
			ResponsibleFamily:            respDB.responsibleFamily,
			ResponsibleFamilyPhoneNumber: respDB.responsibleFamilyPhoneNumber,
			Department:                   respDB.department,
			Country:                      respDB.country,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
