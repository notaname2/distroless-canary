package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	token, set := os.LookupEnv("TOKEN")

	if set {
		url := fmt.Sprintf("https://canarytokens.com/%s", token)

		id, set := os.LookupEnv("IDENT")
		if !set {
			id = "SOMEDEFAULT"
		}

		client := &http.Client{
			Transport: &http.Transport{},
		}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			os.Exit(1)
		}

		req.Header.Set("User-Agent", id)

		response, err := client.Do(req)

		if err != nil {
			os.Exit(2)
		}
		defer response.Body.Close()
	}
	os.Exit(1337)
}
