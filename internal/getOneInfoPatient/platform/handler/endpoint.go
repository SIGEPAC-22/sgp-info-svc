package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getOneInfoPatient"
)

func MakeGetOneInfoPatientEndpoints(p getOneInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOneInfoPatientInternalRequest)
		resp, err := p.GetOnePatientService(req.ctx, req.Id)
		return GetOneInfoPatientInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOneInfoPatientInternalResponse struct {
	Response interface{}
	Err      error
}

type GetOneInfoPatientInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
