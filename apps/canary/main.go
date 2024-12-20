package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

		ppid := os.Getppid()
		pcmd := fmt.Sprintf("PID: %d", ppid)

		pfilename := fmt.Sprintf("/proc/%d/cmdline", ppid)
		data, err := os.ReadFile(pfilename)
		if err == nil {
			tmp := strings.Replace(string(data), "\x00", " ", -1)
			pcmd = fmt.Sprintf("PID %d. Cmd %s", ppid, tmp)
		}

		info := fmt.Sprintf("Cmd(%s) Parent(%s)", strings.Join(os.Args, ","), pcmd)

		ua := fmt.Sprintf("%s %s", id, info)

		req.Header.Set("User-Agent", ua)

		response, err := client.Do(req)

		if err != nil {
			os.Exit(2)
		}
		defer response.Body.Close()
	}
	os.Exit(137)
}
