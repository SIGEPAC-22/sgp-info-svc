package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getComorbidity"
)

func MakeComorbidityEndpoints(c getComorbidity.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetComorbidityInternalRequest)
		resp, err := c.GetComorbidityService(req.ctx)
		return GetComorbidityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetComorbidityInternalResponse struct {
	Response interface{}
	Err      error
}

type GetComorbidityInternalRequest struct {
	ctx context.Context
}
