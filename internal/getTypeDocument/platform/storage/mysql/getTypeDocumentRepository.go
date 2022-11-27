package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-info-svc/internal/GetTypeDocument"
	"sgp-info-svc/kit/constants"
)

type GetTypeDocumentRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetTypeDocumentRepo(db *sql.DB, log log.Logger) *GetTypeDocumentRepo {
	return &GetTypeDocumentRepo{db, log}
}

func (g *GetTypeDocumentRepo) GetTypeDocumentRepository(ctx context.Context) ([]GetTypeDocument.GetTypeDocumentResponse, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	id := config.GetInt("app-properties.getComorbidity.idStatusActive")
	rows, errDB := g.db.QueryContext(ctx, "SELECT dct_id_document_type, dct_document_type_name FROM dct_document_type WHERE dct_id_state_data = ?;", id)
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []GetTypeDocument.GetTypeDocumentResponse{}, errDB
	}
	defer rows.Close()
	var resp []GetTypeDocument.GetTypeDocumentResponse
	for rows.Next() {
		var respDB SqlGetTypeDocument
		if err := rows.Scan(&respDB.Id, &respDB.NameTypeDocument); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []GetTypeDocument.GetTypeDocumentResponse{}, err
		}
		resp = append(resp, GetTypeDocument.GetTypeDocumentResponse{
			Id:               int(respDB.Id),
			NameTypeDocument: respDB.NameTypeDocument,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
