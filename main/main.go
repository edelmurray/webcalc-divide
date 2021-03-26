package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//https://golangcode.com/get-a-url-parameter-from-a-request/
	output := map[string]interface{}{"calculation: ": "/", "answer: ": 0}
	x, ok := r.URL.Query()["x"]

	if !ok || len(x[0]) < 1 {
		log.Println("Error: values not appropriate")
		output["error"] = true
		http.Error(w, "No x parameter provided", http.StatusBadRequest)
		return
	}
	y, ok := r.URL.Query()["y"]

	if !ok || len(y[0]) < 1 {
		log.Println("Error: values not appropriate")
		output["error"] = true
		http.Error(w, "No y parameter provided", http.StatusBadRequest)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	xInt, err := strconv.Atoi(x[0])
	if err != nil {
		log.Println("Error: values not appropriate")
		fmt.Print("x error")
		output["error"] = true
		http.Error(w, "Illegal division by :"+x[0], http.StatusBadRequest)
		return
	}
	yInt, err := strconv.Atoi(y[0])
	if err != nil || (yInt <= 0) {
		log.Println("Error: values not appropriate")
		fmt.Print("y error")
		output["error"] = true
		http.Error(w, "Illegal division by :"+y[0], http.StatusBadRequest)
		return
	}

	output["answer"] = CalculateDivide(xInt, yInt)
	output["x-value"] = xInt
	output["y-value"] = yInt
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	returnAnswer, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	w.Write(returnAnswer)
}

func CalculateDivide(x int, y int) int {
	answer := 0
	answer = x / y
	return answer
}
