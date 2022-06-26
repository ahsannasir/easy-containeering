package main

import (
	"encoding/json"
	artifacts "ml-cicd/src/artifacts"
	builder "ml-cicd/src/builder"
	utils "ml-cicd/src/utilities"
	"net/http"

	"github.com/google/uuid"
)

func publisher(w http.ResponseWriter, r *http.Request) {
	buildID := uuid.New().String()
	if r.Method == "POST" {

		repository := r.FormValue("repository")
		imagename := r.FormValue("name")
		// find a file submitted as a form
		file, handler, err := r.FormFile("Dockerfile")
		defer file.Close()

		err := artifacts.GenArtifacts(file, handler, buildID)
		if err != nil {
			w.Write([]byte("An Error Occurred!"))
		}
		go builder.Build(buildID, repository, imagename)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"build_id": buildID})
}

func getlogger(w http.ResponseWriter, r *http.Request) {
	keyVal := r.URL.Query()["build_id"][0]
	w.Write([]byte(artifacts.FetchLog(keyVal)))
}

func getstatus(w http.ResponseWriter, r *http.Request) {
	keyVal := r.URL.Query()["build_id"][0]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"build_status": utils.GetBuildStatus(keyVal)})
}
