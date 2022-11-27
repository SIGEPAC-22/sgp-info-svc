package getInfoPatientFileFileSvc

import (
	"context"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getInfoPatientFile"
	"sgp-info-svc/kit/constants"
)

type Service struct {
	RepoDB getInfoPatientFile.Repository
	logger log.Logger
}

func NewService(repoDB getInfoPatientFile.Repository, logger log.Logger) *Service {
	return &Service{RepoDB: repoDB, logger: logger}
}

func (s *Service) GetPatientFileService(ctx context.Context) ([]getInfoPatientFile.GetInfoPatientFileResponse, error) {
	s.logger.Log("Start Endpoint GetInfoPatientFile", constants.UUID, ctx.Value(constants.UUID))
	resp, err := s.RepoDB.GetPatientFileRepository(ctx)
	if err != nil {
		s.logger.Log("Error, Failed Repository of Database", "Error", err.Error(), constants.UUID)
		return []getInfoPatientFile.GetInfoPatientFileResponse{}, err
	}

	var rowList []getInfoPatientFile.GetInfoPatientFileResponse

	for _, respBD := range resp {
		var row getInfoPatientFile.GetInfoPatientFileResponse
		idPatientFile := respBD.IdPatientFile
		row.IdPatient = respBD.IdPatient
		row.IdPatientFile = respBD.IdPatientFile
		row.FullName = respBD.FullName
		row.DocumentType = respBD.DocumentType
		row.DocumentNumber = respBD.DocumentNumber
		row.CellphoneNumber = respBD.CellphoneNumber
		row.ResponsibleFamily = respBD.ResponsibleFamily
		row.ResponsibleFamilyPhoneNumber = respBD.ResponsibleFamilyPhoneNumber
		row.Sex = respBD.Sex
		row.StatePatient = respBD.StatePatient
		row.AdmissionDate = respBD.AdmissionDate
		row.HighDate = respBD.HighDate
		row.LowDate = respBD.LowDate

		respComorbidity, err := s.RepoDB.GetComorbidityPatient(ctx, idPatientFile)
		if err != nil {
			s.logger.Log("Error, not get response for comorbidity", "Error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			continue
		}

		row.Comorbidity = respComorbidity

		respSymptom, err := s.RepoDB.GetSymptomPatient(ctx, idPatientFile)
		if err != nil {
			s.logger.Log("Error, not get response for symptom", "Error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			continue
		}

		row.Symptom = respSymptom

		rowList = append(rowList, row)
	}

	resp = rowList

	return resp, nil
}
