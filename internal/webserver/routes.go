package webserver

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func static(w http.ResponseWriter, req *http.Request) {
	file, err := fileSystem.Open(req.URL.Path[1:])
	if err == nil {
		file.Close()
		fileServer.ServeHTTP(w, req)
		return
	}

	indexFile, err := fileSystem.Open("index.html")
	if err != nil {
		http.Error(w, "index.html not found", http.StatusNotFound)
		return
	}
	defer indexFile.Close()

	stat, err := indexFile.Stat()
	http.ServeContent(w, req, "index.html", stat.ModTime(), indexFile.(io.ReadSeeker))
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

	dst, err := os.Create(filepath.Join("data", "samples", handler.Filename))
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

	err := os.Remove(filepath.Join("data", "samples", fileName))
	if err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}
}
