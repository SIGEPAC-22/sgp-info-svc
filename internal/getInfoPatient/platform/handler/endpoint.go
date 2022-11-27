package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getInfoPatient"
)

func MakeInfoPatientEndpoints(p getInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInfoPatientInternalRequest)
		resp, err := p.GetPatientService(req.ctx)
		return GetInfoPatientInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetInfoPatientInternalResponse struct {
	Response interface{}
	Err      error
}

type GetInfoPatientInternalRequest struct {
	ctx context.Context
}
