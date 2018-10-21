package gonextcloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientAppsList(t *testing.T) {
	if err := initClient(); err != nil {
		t.FailNow()
	}
	l, err := c.Apps.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, l)
}
