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
	ac, err := c.AppsConfig().Get()
	assert.NoError(t, err)
	assert.NotEmpty(t, ac)
}

func TestAppsConfigList(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	a, err := c.AppsConfig().List()
	assert.NoError(t, err)
	assert.Contains(t, a, "files")
}

func TestAppsConfigKeys(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	ks, err := c.AppsConfig().Keys("activity")
	assert.NoError(t, err)
	assert.Contains(t, ks, "enabled")
}

func TestAppsConfigValue(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	k, err := c.AppsConfig().Value("files", "enabled")
	assert.NoError(t, err)
	assert.Equal(t, "yes", k)
}

func TestAppConfigDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	d, err := c.AppsConfig().Details("activity")
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
}
