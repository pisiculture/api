package domain_test

import (
	"testing"

	"github.com/pisiculture/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewInstallation(t *testing.T) {
	installation := domain.NewInstallation()
	assert.NotNil(t, installation)
}

func TestInstallationValidate(t *testing.T) {

	installation := domain.NewInstallation()

	err := installation.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Param { Name } required", err.Error())

	installation.Name = "Installation test"

	err = installation.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "Param { Description } required", err.Error())
}
