package handler

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	bb, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("JSON Marshalling error: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(bb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while trying to write to the ResponseWriter: " + err.Error()))
		return
	}
}
