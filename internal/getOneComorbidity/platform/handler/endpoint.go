package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getOneComorbidity"
)

func MakeGetOneComorbidityEndpoint(c getOneComorbidity.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetOneComorbidityInternalRequest)
		resp, err := c.GetOneComorbiditySvc(req.ctx, req.Id)
		return GetOneComorbidityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetOneComorbidityInternalRequest struct {
	Id  string `json:"id" example:"1" validate:"nonzero"`
	ctx context.Context
}

type GetOneComorbidityInternalResponse struct {
	Response interface{}
	Err      error
}
