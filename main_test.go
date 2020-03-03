package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestHomeLink(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	homeLink(recorder, req)
	validateResponse(recorder, err, 200, "Could not connect")
}

func TestGetAllEvents(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:123/", nil)
	recorder := httptest.NewRecorder()
	getAllEvents(recorder, req)
	validateResponse(recorder, err, 400, "Expected 400")
}

func TestDeleteEvent(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:123/", nil)
	recorder := httptest.NewRecorder()
	deleteEvent(recorder, req)
	validateResponse(recorder, err, 200, "Expected 400")
}

func validateResponse(recorder *httptest.ResponseRecorder, err error, expectedCode int, message string) {
	trace()
	if recorder.Code != expectedCode {
		if err != nil {
			log.Fatal(message+" ", err)
		} else {
			log.Fatal(message)
		}
	}
}

func trace() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}
