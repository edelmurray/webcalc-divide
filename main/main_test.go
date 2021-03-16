package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

//TestHandler for main.go
func TestHandler(t *testing.T) {
	//https://www.youtube.com/watch?v=hVFEV-ieeew&ab_channel=justforfunc%3AProgramminginGo
	req, err := http.NewRequest("GET", "localhost:8080/?x=10&y=2", nil)
	if err != nil {
		t.Fatalf("could not create request: %x", err)
	}

	rec := httptest.NewRecorder()

	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected stauts ok; got %x ", err)
	}
	//expected := map[string]interface{}{"x-value": 10, "y-value": 2, "answer": 5}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("get:\n", keepLines(string(body), 1))
	returnedAnswer := keepLines(string(body), 1)
	if err != nil {
		t.Fatalf("expected an int got a string %s", returnedAnswer)
	}
}

func TestHandlerIncorrectParamsEmpty(t *testing.T) {
	//https://www.youtube.com/watch?v=hVFEV-ieeew&ab_channel=justforfunc%3AProgramminginGo
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("could not create request: %x", err)
	}

	rec := httptest.NewRecorder()

	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status bad request; got %x ", err)
	}
}

func TestHandlerIncorrectParamsXIncorrect(t *testing.T) {
	//https://www.youtube.com/watch?v=hVFEV-ieeew&ab_channel=justforfunc%3AProgramminginGo
	req, err := http.NewRequest("GET", "localhost:8080/?x=h&y=7", nil)
	if err != nil {
		t.Fatalf("could not create request: %x", err)
	}

	rec := httptest.NewRecorder()

	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status bad request; got %x ", err)
	}
}

func TestHandlerIncorrectParamsYIncorrect(t *testing.T) {
	//https://www.youtube.com/watch?v=hVFEV-ieeew&ab_channel=justforfunc%3AProgramminginGo
	req, err := http.NewRequest("GET", "localhost:8080/?x=3&y=0", nil)
	if err != nil {
		t.Fatalf("could not create request: %x", err)
	}

	rec := httptest.NewRecorder()

	Handler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status bad request; got %x ", err)
	}
}
