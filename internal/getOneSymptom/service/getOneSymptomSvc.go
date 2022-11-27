package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-info-svc/internal/getOneSymptom"
	"sgp-info-svc/kit/constants"
)

type getOneSymptomSvc struct {
	repoDB getOneSymptom.Repository
	logger kitlog.Logger
}

func NewGetOneSymptomSvc(repoDB getOneSymptom.Repository, logger kitlog.Logger) *getOneSymptomSvc {
	return &getOneSymptomSvc{repoDB: repoDB, logger: logger}
}
func (g *getOneSymptomSvc) GetOneSymptomSvc(ctx context.Context, id string) (getOneSymptom.GetOneSymptomResponse, error) {
	g.logger.Log("Starting getOneSymptom process", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetOneSymptom(ctx, id)
	if err != nil {
		g.logger.Log("Error getting data from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOneSymptom.GetOneSymptomResponse{}, err
	}
	return resp, nil
}
