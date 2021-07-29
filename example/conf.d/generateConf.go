package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// initialisation of configuration
var AppConf interface{} = initConfig()

// write configuration in app.json
func main() {

	bytes, err := json.Marshal(AppConf)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = ioutil.WriteFile("app.json", bytes, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
