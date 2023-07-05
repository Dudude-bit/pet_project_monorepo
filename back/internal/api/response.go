package api

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func sendResponse(w http.ResponseWriter, statusCode int, body interface{}) error {
	resp := httpResponse{
		Data: body,
	}

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	return nil
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, err error) error {
	resp := httpResponse{
		Error: err,
	}

	w.WriteHeader(statusCode)

	errEncode := json.NewEncoder(w).Encode(resp)
	if errEncode != nil {
		return errEncode
	}

	w.Header().Set("Content-Type", "application/json")

	return nil
}
