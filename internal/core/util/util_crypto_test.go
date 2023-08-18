package util_test

import (
	"testing"

	"github.com/pisiculture/internal/core/util"
	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	val, err := util.Crypto("Teste")
	assert.Nil(t, err, "Error generating encryption")
	assert.NotNil(t, val)
}
