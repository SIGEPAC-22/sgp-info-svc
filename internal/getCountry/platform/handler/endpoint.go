package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getCountry"
)

func MakeGetCountryEndpoints(c GetCountry.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetCountryInternalRequest)
		resp, err := c.GetCountryService(req.ctx)
		return GetCountryInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetCountryInternalResponse struct {
	Response interface{}
	Err      error
}

type GetCountryInternalRequest struct {
	ctx context.Context
}
