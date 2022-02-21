package registry

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type PackageMetadata struct {
	Time     map[string]string `json:"time"`
	DistTags DistTags          `json:"dist-tags"`
}

type DistTags struct {
	Latest string `json:"latest"`
}

type PackageTimeMetadata struct {
	Modified string `json:"modified"`
}

func GetPackageAge(packageName string) int {
	response, err := http.Get("https://registry.npmjs.org/" + packageName)
	fatalIfError(err)
	defer response.Body.Close()

	var metadata PackageMetadata
	body, err := ioutil.ReadAll(response.Body)
	fatalIfError(err)

	json.Unmarshal(body, &metadata)

	parsedTime, err := time.Parse(time.RFC3339, metadata.Time[metadata.DistTags.Latest])
	fatalIfError(err)

	return int(time.Since(parsedTime).Hours() / 24 / 365)
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
