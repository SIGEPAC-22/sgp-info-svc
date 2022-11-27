package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"gopkg.in/validator.v2"
	"sgp-info-svc/kit/constants"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func GetOneInfoPatientTransportMiddleware(logger log.Logger) Middleware {
	return func(endpoint endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(GetOneInfoPatientInternalRequest)
			if err := validator.Validate(&req); err != nil {
				logger.Log("Invalid request", "Error", err.Error(), "Request", req)
				return GetOneInfoPatientInternalResponse{
					Response: constants.ErrBadRequest.Error() + " - " + err.Error(),
					Err:      constants.ErrBadRequest,
				}, nil
			}
			defer logger.Log("Process Finished", "Request", req)

			return endpoint(ctx, request)
		}

	}
}
