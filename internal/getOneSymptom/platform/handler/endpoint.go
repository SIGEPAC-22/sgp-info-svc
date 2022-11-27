package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getOneSymptom"
)

func MakeGetOneSymptomEndpoint(c getOneSymptom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOneSymptomInternalRequest)
		resp, err := c.GetOneSymptomSvc(req.ctx, req.Id)
		return GetOneSymptomInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOneSymptomInternalRequest struct {
	Id  string `json:"id" example:"1" validate:"nonzero"`
	ctx context.Context
}

type GetOneSymptomInternalResponse struct {
	Response interface{}
	Err      error
}
