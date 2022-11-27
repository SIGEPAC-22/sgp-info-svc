package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"sgp-info-svc/kit/constants"
)

func NewHttpGetInfoPatientHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle(path,
		httptransport.NewServer(endpoints,
			DecodeRequestGetInfoPatient,
			EncodeRequestGetInfoPatient,
		)).Methods(http.MethodGet)
	return r
}

func DecodeRequestGetInfoPatient(ctx context.Context, r *http.Request) (interface{}, error) {
	processID, _ := uuid.NewUUID()
	ctx = context.WithValue(ctx, constants.UUID, processID.String())
	return GetInfoPatientInternalRequest{
		ctx: ctx,
	}, nil
}

func EncodeRequestGetInfoPatient(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset = utf-8")
	resp, _ := response.(GetInfoPatientInternalResponse)
	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset = utf-8")
		switch resp.Err {
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return json.NewEncoder(w).Encode(resp.Err.Error())
	}
	return json.NewEncoder(w).Encode(resp.Response)
}