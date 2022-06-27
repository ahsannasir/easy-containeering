package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var buildID string

// Test_Publish: Tests the complete Build and Push workflow
// uses a helper publishBuild to create a decorated HTTP request object
// that is then used to call the "publisher" handler.
// Test passes if build is successfully initiated and returns a bulid_id
func Test_Publish(t *testing.T) {
	request, w, err := publishBuild()
	publisher(w, request)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	data := make(map[string]string)
	_ = json.NewDecoder(res.Body).Decode(&data)
	if data != nil && data["build_id"] != "" {
		buildID = data["build_id"]
	} else {
		t.Fatal("Failed: Unexpected response received!")
	}
}

// Test_Get_Logs: Tests wether a particular builds logs are successfully fetched
// uses a test build_id to fetch logs and test passes if there are no failures during fetch
func Test_Get_Logs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/logs?build_id=03fbbda8-4454-499c-8a30-d4322cf688eb", nil)
	w := httptest.NewRecorder()
	getlogger(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	fmt.Println(string(data))
	if res.StatusCode == http.StatusInternalServerError {
		t.Errorf("expected logs got %v", string(data))
	}
}

// Test_Get_Status: Tests wether a build's status can be retrieved successfully
// uses the helper publishBuild to get a decorated request object for builds
// publishes a build and fetches its status
// Test passes if expected status is received
func Test_Get_Status(t *testing.T) {

	request, w, err := publishBuild()
	publisher(w, request)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Fatal("Failed: ", err)
	}
	data := make(map[string]string)
	_ = json.NewDecoder(res.Body).Decode(&data)
	if data != nil && data["build_id"] != "" {
		buildID := data["build_id"]

		req := httptest.NewRequest(http.MethodGet, "/api/status?build_id="+buildID, nil)
		w := httptest.NewRecorder()
		getstatus(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		if string(data) == "" || string(data) == "An Error Occurred" {
			t.Errorf("expected logs got %v", string(data))
		}
		fmt.Println(string(data))
	}
}

// publishBuild: A helper function that allows to build a test http request object
// to send a multipart file and other data fields to /api/publish API endpoint
// to create a build for testing.
func publishBuild() (*http.Request, *httptest.ResponseRecorder, error) {
	url := "/api/publish"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open("./resources/Dockerfile")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("Dockerfile", "Dockerfile")
	_, err = io.Copy(part1, file)
	if err != nil {
		return nil, nil, err
	}

	_ = writer.WriteField("name", "pyhoncontained")
	_ = writer.WriteField("repository", "ahsannasir")

	err = writer.Close()
	if err != nil {
		return nil, nil, err
	}
	// We read from the pipe which receives data
	// from the multipart writer, which, in turn,
	// receives data from png.Encode().
	// We have 3 chained writers!

	request := httptest.NewRequest(method, url, payload)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	return request, w, nil
}
