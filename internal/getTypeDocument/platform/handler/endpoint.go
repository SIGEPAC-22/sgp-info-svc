package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	GetTypeDocument "sgp-info-svc/internal/getTypeDocument"
)

func MakeGetTypeDocumentEndpoints(c GetTypeDocument.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetTypeDocumentInternalRequest)
		resp, err := c.GetTypeDocumentService(req.ctx)
		return GetTypeDocumentInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetTypeDocumentInternalResponse struct {
	Response interface{}
	Err      error
}

type GetTypeDocumentInternalRequest struct {
	ctx context.Context
}
