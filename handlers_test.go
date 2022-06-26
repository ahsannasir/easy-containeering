package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get_Status(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/logs?build_id=9c19f38f-6114-4c8c-8506-6334d7174b95", nil)
	w := httptest.NewRecorder()
	getlogger(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) == "" {
		t.Errorf("expected ABC got %v", string(data))
	}
}
