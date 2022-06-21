package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Data struct {
	Origin string
	User   string
	Active bool
}

func main() {
	content, err := ioutil.ReadFile("./arquivo.json")
	if err != nil {
		log.Fatal("Error when opering file: ", err)
	}

	var payload Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during unmarshal(): ", err)
	}

	log.Printf("origin: %s\n", payload.Origin)
	log.Printf("user: %s\n", payload.User)
	log.Printf("status: %t\n", payload.Active)

}
