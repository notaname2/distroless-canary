package main

import (
	"encoding/json"
	"os"
	"time"
)

type output struct {
	CurrentTime string `json:"currentTime"`
}

func main() {

	now := time.Now()

	data := output{CurrentTime: now.String()}
	b, err := json.Marshal(data)

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stdout.WriteString("{}")
		return
	}

	os.Stdout.Write(b)
}
