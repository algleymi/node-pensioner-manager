package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainPanics(t *testing.T) {
	assert.Panics(t, func() {
		main()
	})
}
