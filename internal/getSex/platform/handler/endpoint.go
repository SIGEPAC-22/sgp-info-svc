package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/GetSex"
)

func MakeGetSexEndpoints(c GetSex.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetSexInternalRequest)
		resp, err := c.GetSexService(req.ctx)
		return GetSexInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetSexInternalResponse struct {
	Response interface{}
	Err      error
}

type GetSexInternalRequest struct {
	ctx context.Context
}
