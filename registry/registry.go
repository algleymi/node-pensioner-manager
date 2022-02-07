package registry

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RegistryMetadata struct {
	DbName string `json:"db_name"`
}

func GetRegistryMetadata() string {
	response, err := http.Get("https://registry.npmjs.org")	
	fatalIfError(err)
	defer response.Body.Close()

	var metadata RegistryMetadata
	body, err := ioutil.ReadAll(response.Body)
	fatalIfError(err)

	json.Unmarshal(body, &metadata)

	return metadata.DbName
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
