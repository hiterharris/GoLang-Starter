package main

import (
	"fmt";
	"net/http";
	"io/ioutil"
	"log"
	"encoding/json"
)


func main() {
	URL := "https://www.boredapi.com/api/activity/"

	// Make API request
	resp, err := http.Get(URL)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Fatalln(err) }

	// Parse response
	resBytes := []byte(body)
	var jsonRes map[string]interface{}
	_ = json.Unmarshal(resBytes, &jsonRes)
	activity := jsonRes["activity"].(string)
	fmt.Println("activity: ", activity)
}

