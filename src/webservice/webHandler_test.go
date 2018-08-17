package webservice

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestProcessRequest(t *testing.T){
	testreq, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessRequest)
	handler.ServeHTTP(reqRecorder,testreq)
	if status:=reqRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
