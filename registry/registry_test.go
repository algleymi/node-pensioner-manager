package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRegistryReturnsBody(t *testing.T) {
	response := GetRegistryMetadata()

	assert.Equal(t, response, "registry")
}
