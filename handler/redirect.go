package handler

import (
	"net/http"
	"urlshortner/storage"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // remove "/"
	url := storage.GetURL(code)

	http.Redirect(w, r, url, http.StatusFound)
}
