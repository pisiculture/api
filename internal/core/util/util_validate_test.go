package util_test

import (
	"testing"

	"github.com/pisiculture/internal/core/util"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {

	t.Run("Email_with_@", func(t *testing.T) {
		err := util.ValidateEmail("DjontanWillenz.com.br")
		assert.NotNil(t, err, "Email not with @")
	})

	t.Run("Email_does_not_end_with_.com", func(t *testing.T) {
		err := util.ValidateEmail("Djontan@teste.br")
		assert.NotNil(t, err, "Email does not end with .com")
	})
}
