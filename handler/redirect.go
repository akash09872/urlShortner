package handler

import (
	"fmt"
	"net/http"
	"urlshortner/storage"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // remove "/"

	url, ok := storage.Store[code]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found")
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
