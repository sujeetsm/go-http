package webservice

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestProcessRequestPost(t *testing.T){
	testreq, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessRequest)
	handler.ServeHTTP(reqRecorder,testreq)
	if status:=reqRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

}

func TestProcessRequestPut(t *testing.T){
	testreq, err := http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessRequest)
	handler.ServeHTTP(reqRecorder,testreq)
	if status:=reqRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

}
func TestProcessRequestGet(t *testing.T){
	testreq, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProcessRequest)
	handler.ServeHTTP(reqRecorder,testreq)
	if status:=reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

}

