package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getOnePatientFile"
)

func MakeGetOnePatientFileEndpoint(c getOnePatientFile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOnePatientFileInternalRequest)
		resp, err := c.GetOnePatientFileService(req.ctx, req.Id)
		return GetOnePatientFileInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOnePatientFileInternalRequest struct {
	Id  string `json:"id" example:"1" validate:"nonzero"`
	ctx context.Context
}

type GetOnePatientFileInternalResponse struct {
	Response interface{}
	Err      error
}
