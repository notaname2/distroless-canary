package main

import (
	"encoding/json"
	"os"
)

type output struct {
	Message []string `json:"message"`
}

func main() {

	data := output{}
	data.Message = os.Args[1:]
	b, err := json.Marshal(data)

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stdout.WriteString("{}")
		return
	}

	os.Stdout.Write(b)
}
