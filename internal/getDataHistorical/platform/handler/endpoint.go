package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getDataHistorical"
)

func MakeGetDataHistoricalEndpoints(c getDataHistorical.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetDataHistoricalInternalRequest)
		resp, err := c.GetDataHistoricalService(req.ctx)
		return GetDataHistoricalInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetDataHistoricalInternalResponse struct {
	Response interface{}
	Err      error
}

type GetDataHistoricalInternalRequest struct {
	ctx context.Context
}
