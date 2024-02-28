package api

import (
	"encoding/json"
	"net/http"

	"dantejs.com/go/http-server/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id != "" {
		json.NewEncoder(w).Encode(data.GetOne(id))
		return
	}
	json.NewEncoder(w).Encode(data.GetAll())
}
