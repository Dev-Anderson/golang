package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./arquivo.json")
	if err != nil {
		log.Fatal("Error when opering file: ", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("origin: %s\n", payload["origin"])
	log.Printf("user: %s\n", payload["user"])
	log.Printf("status: %s\n", payload["active"])
}
