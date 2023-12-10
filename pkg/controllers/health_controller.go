package controllers

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"health": "ok",
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to marshal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
