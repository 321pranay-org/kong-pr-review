package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	body := "{\"commit_id\": \"fac510cdb65fcd54135bd891d54263818f88eb95\", \"event\": \"APPROVE\"}"

	req, err := http.NewRequest(http.MethodPost, "https://api.github.com/repos/321pranay-org/yaml/pulls/13/reviews", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
