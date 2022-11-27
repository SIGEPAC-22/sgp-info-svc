package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getInfoPatientFile"
)

func MakeInfoPatientFileEndpoints(p getInfoPatientFile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetInfoPatientFileInternalRequest)
		resp, err := p.GetPatientFileService(req.ctx)
		return GetInfoPatientFileInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetInfoPatientFileInternalResponse struct {
	Response interface{}
	Err      error
}

type GetInfoPatientFileInternalRequest struct {
	ctx context.Context
}
