package api

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func sendSuccessResponse(w http.ResponseWriter, statusCode int, body interface{}) error {
	return sendResponse(w, statusCode, body, nil)
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, err error) error {
	return sendResponse(w, statusCode, nil, err)
}

func sendResponse(w http.ResponseWriter, statusCode int, body interface{}, err error) error {
	resp := httpResponse{
		Data:  body,
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
