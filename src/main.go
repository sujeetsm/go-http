package main

import (
	"log"
	"net/http"
	"strconv"
	"webservice"
)

const PORT = 8080




func main() {

	log.Println("Attempting to start HTTP Server.")

	http.HandleFunc("/", webservice.ProcessRequest)

	var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}
}
