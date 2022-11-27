package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-info-svc/internal/getDepartment"
)

func MakeGetDepartmentEndpoints(c getDepartment.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetDepartmentInternalRequest)
		resp, err := c.GetDepartmentService(req.ctx)
		return GetDepartmentInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetDepartmentInternalResponse struct {
	Response interface{}
	Err      error
}

type GetDepartmentInternalRequest struct {
	ctx context.Context
}
