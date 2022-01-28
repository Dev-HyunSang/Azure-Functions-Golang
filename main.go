package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type message struct {
	Body string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message{
		Body: "Hello World!",
	})
}

func main() {
	http.HandleFunc("/api/hello-world", HelloWorld)

	FunctionPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + FunctionPort)
	}
	log.Println("Azure Function Running Server...!")
	log.Fatal(http.ListenAndServe(":"+FunctionPort, nil))
}
