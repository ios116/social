package webserver

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, js interface{}, status int) {
	b, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}
