// based on https://blog.alexellis.io/golang-json-api-client
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://api.open-notify.org/astros.json"
	ppl, err := fetchAstros(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("there are %d astronauts currently in space\n", ppl.Number)
}

func fetchAstros(url string) (*People, error) {
	// create an HTTP request
	cli := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	// execute the request
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// check the response code
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got response %d", res.StatusCode)
	}

	// unmarshal the response
	var people People
	err = json.NewDecoder(res.Body).Decode(&people)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling json: %w", err)
	}
	return &people, nil
}
