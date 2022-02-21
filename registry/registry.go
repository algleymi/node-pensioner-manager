package registry

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func GetPackageAge(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, ErrorFetchingMetadata
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return 0, ErrorPackageNotFound
	}

	var metadata PackageMetadata
	body, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &metadata)

	parsedTime, _ := time.Parse(time.RFC3339, metadata.Time[metadata.DistTags.Latest])

	return int(time.Since(parsedTime).Hours() / 24 / 365), nil
}

var ErrorFetchingMetadata = errors.New("error fetching metadata")
var ErrorPackageNotFound = errors.New("package not found")
