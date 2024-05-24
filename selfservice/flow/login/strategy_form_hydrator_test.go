package login

import (
	"github.com/ory/kratos/identity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithIdentityHint(t *testing.T) {
	expected := new(identity.Identity)
	opts := NewFormHydratorOptions([]FormHydratorModifier{WithIdentityHint(expected)})
	assert.Equal(t, expected, opts.IdentityHint)
}
