package webserver

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func sample(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("data", handler.Filename))
	if err != nil {
		http.Error(w, "Unable to create local file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteSample(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileName := req.PathValue("name")

	err := os.Remove(filepath.Join("data", fileName))
	if err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}
}
