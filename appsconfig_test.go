package gonextcloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppsConfig(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	ac, err := c.AppsConfig()
	assert.NoError(t, err)
	assert.NotEmpty(t, ac)
}

func TestAppsConfigList(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	a, err := c.AppsConfigList()
	assert.NoError(t, err)
	assert.Contains(t, a, "files")
}

func TestAppsConfigKeys(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	ks, err := c.AppsConfigKeys("activity")
	assert.NoError(t, err)
	assert.Contains(t, ks, "enabled")
}

func TestAppsConfigValue(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	k, err := c.AppsConfigValue("files", "enabled")
	assert.NoError(t, err)
	assert.Equal(t, "yes", k)
}

func TestAppConfigDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	d, err := c.AppsConfigDetails("activity")
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
}
