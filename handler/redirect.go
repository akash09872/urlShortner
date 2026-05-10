package handler

import (
	"encoding/json"
	"net/http"
	"urlshortner/model"
	"urlshortner/storage"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // remove "/"
	url := storage.GetURL(code)
	if url == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Response{
			Message: "NOT Found",
		})
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
