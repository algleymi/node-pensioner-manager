package registry

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createServerWithFixture(fixture string, response int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(response)
		bytes, _ := ioutil.ReadFile(fixture)
		w.Write(bytes)
	}))
}

func Test_GetMetadataReturnsAge(t *testing.T) {
	server := createServerWithFixture("./fixtures/2016_package.json", http.StatusOK)
	defer server.Close()

	age, _ := GetPackageAge(server.URL)

	assert.Greater(t, age, 2)
}

func Test_GetMetadataFromPackageThatDoesntExist(t *testing.T) {
	server := createServerWithFixture("./fixtures/no_response.json", http.StatusNotFound)
	defer server.Close()

	_, err := GetPackageAge(server.URL)

	assert.ErrorIs(t, err, ErrorPackageNotFound)
}

func Test_GetMetadataFromPackageWithStrangeNameHasAnError(t *testing.T) {
	server := createServerWithFixture("./fixtures/no_response.json", http.StatusBadRequest)
	defer server.Close()

	_, err := GetPackageAge("")

	assert.ErrorIs(t, err, ErrorFetchingMetadata)
}
