package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getStatePatient"
)

func MakeGetStatePatientEndpoints(c getStatePatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetStatePatientInternalRequest)
		resp, err := c.GetStatePatientService(req.ctx)
		return GetStatePatientInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetStatePatientInternalResponse struct {
	Response interface{}
	Err      error
}

type GetStatePatientInternalRequest struct {
	ctx context.Context
}
