package client

import (
	"fmt"
	"log"
	"net/http"
)

func InitClient() {
	fileServer := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting backend at port 8082...\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
