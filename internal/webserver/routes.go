package webserver

import (
	"errors"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "web/index.html")
}

func static(w http.ResponseWriter, req *http.Request) {
	path := "web" + req.URL.Path
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, req, path)
}
