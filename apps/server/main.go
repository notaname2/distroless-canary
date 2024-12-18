package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func appHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request")
	appName := r.PathValue("APP")
	if appName == "" {
		log.Printf("No app name")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Printf("App name is %s", appName)

	params := r.URL.Query()["ARG"]

	binName := fmt.Sprintf("/app/bin/%s", appName)

	log.Printf("setting to run command %s with params %v", binName, params)
	cmd := exec.Command(binName, params...)

	cmd.Stdout = w

	err := cmd.Run()
	if err != nil {
		log.Printf("Could not run %s: %s", appName, err.Error())
	}
}

func ackHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Doing ACK")
	_, err := io.WriteString(w, "ack")

	if err != nil {
		log.Printf("error sending ack: %s", err.Error())
	}

	log.Printf("ACK done")
}

func main() {
	log.Printf("Setting up handler")
	http.HandleFunc("GET /tool/{APP}", appHandler)
	http.HandleFunc("GET /ack", ackHandler)

	log.Printf("Staring listener")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
