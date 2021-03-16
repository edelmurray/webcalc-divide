package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//https://golangcode.com/get-a-url-parameter-from-a-request/
	x, ok := r.URL.Query()["x"]
	y, ok := r.URL.Query()["y"]

	if !ok || len(x[0]) < 1 || len(y[0]) < 1 {
		log.Println("Error: values not appropriate")
		http.Error(w, "values not appropriate", http.StatusBadRequest)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	xInt, err := strconv.Atoi(x[0])
	if err != nil {
		log.Println("Error: values not appropriate")
		fmt.Print("x error")
		http.Error(w, "cannot divide :"+x[0], http.StatusBadRequest)
		return
	}
	yInt, err := strconv.Atoi(y[0])
	if err != nil || (yInt <= 0) {
		log.Println("Error: values not appropriate")
		fmt.Print("y error")
		http.Error(w, "cannot divide :"+y[0], http.StatusBadRequest)
		return
	}
	output := map[string]interface{}{"x-value": 0, "y-value": 0, "answer": 0}
	answer := CalculateDivide(xInt, yInt)
	output["answer"] = answer
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
