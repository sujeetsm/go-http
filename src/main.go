package main

import (
	"log"
	"net/http"
	"webservice"
	"os"
)



type Config struct {
	Port   string
}



func main() {
	log.Println("Starting HTTP Server.")
	config := Config{}
	val, ok := os.LookupEnv("HTTP_PORT")
		if !ok {
			log.Println("env var HTTP_PORT not set.. defaulting to 8080\n")
			//set to default value
			config.Port = "8080"
		} else {
			log.Println("Http port = ", val)
			config.Port = val
		}
	http.HandleFunc("/", webservice.ProcessRequest)

	var err = http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err)
	}

}
