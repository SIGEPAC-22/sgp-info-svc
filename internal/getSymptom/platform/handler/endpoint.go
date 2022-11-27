package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getSymptom"
)

func MakeSymtpomEndpoints(c getSymptom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetSymptomInternalRequest)
		resp, err := c.GetSymptomService(req.ctx)
		return GetSymptomInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetSymptomInternalResponse struct {
	Response interface{}
	Err      error
}

type GetSymptomInternalRequest struct {
	ctx context.Context
}
