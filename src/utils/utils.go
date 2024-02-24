package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string
}

type SuccessResponse struct {
	Message string
}

func JSONError(w http.ResponseWriter, errMsg string, errCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(errorResponse{Message: errMsg})
}

func JSONResponse(w http.ResponseWriter, resp interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
