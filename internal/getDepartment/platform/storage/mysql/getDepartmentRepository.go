package mysql

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"sgp-info-svc/internal/getDepartment"
	"sgp-info-svc/kit/constants"
)

type GetDepartmentRepo struct {
	db  *sql.DB
	log log.Logger
}

func NewGetDepartmentRepo(db *sql.DB, log log.Logger) *GetDepartmentRepo {
	return &GetDepartmentRepo{db, log}
}

func (g *GetDepartmentRepo) GetDepartmentRepository(ctx context.Context) ([]getDepartment.GetDepartmentResponse, error) {

	rows, errDB := g.db.QueryContext(ctx, "SELECT dpt_id_department, dpt_name_dapartment FROM dpt_department;")
	if errDB != nil {
		g.log.Log("Error while trying to get information for conmorbidity", constants.UUID, ctx.Value(constants.UUID))
		return []getDepartment.GetDepartmentResponse{}, errDB
	}
	defer rows.Close()
	var resp []getDepartment.GetDepartmentResponse
	for rows.Next() {
		var respDB SqlGetDepartment
		if err := rows.Scan(&respDB.Id, &respDB.NameDepartment); err != nil {
			g.log.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return []getDepartment.GetDepartmentResponse{}, err
		}
		resp = append(resp, getDepartment.GetDepartmentResponse{
			Id:             int(respDB.Id),
			NameDepartment: respDB.NameDepartment,
		})
	}
	if len(resp) == 0 {
		g.log.Log("Data Not Found", constants.UUID, ctx.Value(constants.UUID))
	}
	return resp, nil
}
