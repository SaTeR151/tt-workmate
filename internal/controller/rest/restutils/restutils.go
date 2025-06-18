package restutils

import (
	"encoding/json"
	"net/http"
)

func Error(res http.ResponseWriter, err string, code int) {
	res.WriteHeader(code)
	json.NewEncoder(res).Encode(HTTPError{Code: code, Error: err})
	return
}

type HTTPError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
