package main

import (
	"context"
	"encoding/json"
	auth "ml-cicd/auth"
	artifacts "ml-cicd/src/artifacts"
	builder "ml-cicd/src/builder"
	utils "ml-cicd/src/utilities"
	"net/http"

	"github.com/google/uuid"
)

// publisher: An API handler that allows consumers to publish their dockerfiles
// from where a docker image is built and pushed onto a docker registry
func publisher(w http.ResponseWriter, r *http.Request) {

	// generate a guid to identify builds
	buildID := uuid.New().String()

	// this ensures publisher only works if user is authenticated
	auth.Verify(w, r)
	if r.Method == "POST" {

		// fetch form fields
		repository := r.FormValue("repository")
		imagename := r.FormValue("name")
		file, handler, err := r.FormFile("Dockerfile")
		defer file.Close()

		// return bad request if invalid inputs are passed
		if err != nil || file == nil || repository == "" || imagename == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		// finally generate artifacts that are required to start the build process
		err = artifacts.GenArtifacts(file, handler.Filename, buildID)

		// get a client to docker daemon
		cli, err := utils.GetDockerClient(context.Background())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		// Finally initiate the build job. The build job executes as a go-routine
		// that unblocks the user from waiting for this API endpoint to return
		// and builds & pushes the image to registry concurrently.
		// To fetch logs or status of this concurrent build job, user can use /api/logs or /api/status endpoints.
		go builder.Build(cli, buildID, repository, imagename)

	}
	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"build_id": buildID})
}

// getLogger: An API handler that allows fetching the logs of a particular build process
// logs can be fetched during the build run or after the build has succeeded
func getlogger(w http.ResponseWriter, r *http.Request) {
	keyVal := r.URL.Query()["build_id"][0]
	if keyVal == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	logs, err := artifacts.FetchLog(keyVal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(logs))
}

// getstatus: An API handler that fetches the status of concurrent build.
// Build Status can be "running", "failed" or "success" depending upon the build progress.
func getstatus(w http.ResponseWriter, r *http.Request) {
	keyVal := r.URL.Query()["build_id"][0]
	if keyVal == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"build_status": utils.GetBuildStatus(keyVal)})
}
