package main

import (
	"net/http"
	"os/exec"
	"log"
)

func appHandler(w http.ResponseWriter, r *http.Request) {
	appName := r.PathValue("APP")
	if appName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	params := r.URL.Query()["ARG"]

	cmd := exec.Command(appName, params...)

	cmd.Stdout = w

	err := cmd.Run()
	if err != nil {
		log.Printf("Could not run %s: $s", appName, err.Error())
	}
}

func main() {
	http.HandleFunc("GET /tool/{APP}", appHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
