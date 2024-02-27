package api

import (
	"encoding/json"
	"net/http"
)

//coinBalance params

type CoinBalanceParams struct {
	Username string
}

// Response
type CoinBalanceResponse struct {
	Balance uint

	Code uint
}

// Error

type Error struct {
	Message string

	Code int
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// declaring package lvl types and since they are of the same type hence wrapped around together
var (
	ResponseErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalServerError = func(w http.ResponseWriter, err error) {
		writeError(w, "An error occured", http.StatusInternalServerError)
	}
)
