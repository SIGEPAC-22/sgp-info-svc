package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-info-svc/internal/getOneComorbidity"
	"sgp-info-svc/kit/constants"
)

type getOneComorbiditySvc struct {
	repoDB getOneComorbidity.Repository
	logger kitlog.Logger
}

func NewGetOneComorbiditySvc(repoDB getOneComorbidity.Repository, logger kitlog.Logger) *getOneComorbiditySvc {
	return &getOneComorbiditySvc{repoDB: repoDB, logger: logger}
}
func (g *getOneComorbiditySvc) GetOneComorbiditySvc(ctx context.Context, id string) (getOneComorbidity.GetOneComorbidityResponse, error) {
	g.logger.Log("Starting getOneComorbidity process", constants.UUID, ctx.Value(constants.UUID))

	resp, err := g.repoDB.GetOneComorbidity(ctx, id)
	if err != nil {
		g.logger.Log("Error getting data from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return getOneComorbidity.GetOneComorbidityResponse{}, err
	}
	return resp, nil
}
