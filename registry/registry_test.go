package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetVleesbroodsUnbgPackageMetadataReturnsAge(t *testing.T) {
	response := GetPackageAge("@vleesbrood/unbg")

	assert.Greater(t, response, 2)
}

func Test_GetInflightMetadataReturnsAge(t *testing.T) {
	response := GetPackageAge("inflight")

	assert.Greater(t, response, 2)
}
