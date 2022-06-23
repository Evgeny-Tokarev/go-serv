package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Form struct {
	Name    string
	Address string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, PATCH, DELETE, GET")
	var f Form
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	json1 := make(map[string]string)
	json1["name"] = f.Name
	json1["address"] = f.Address
	jsonStr, err := json.Marshal(json1)
	if err != nil {

		log.Fatal(err)
	}
	w.Write(jsonStr)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func InitServer() {
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port 9000\n")
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
	}()
}
