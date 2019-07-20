package gonextcloud

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
)

var (
	dir string
	wd types.WebDav

	wtests = []struct {
		name string
		test test
	}{
		{
			name: "CreateFolder",
			test: testCreateFolder,
		},
		{
			name: "TestStat",
			test: testStat,
		},
		{
			name: "TestWalk",
			test: testWalk,
		},
		{
			name: "TestDelete",
			test: testDelete,
		},
	}
)

func TestWebDav(t *testing.T) {
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	if config.NotExistingFolder == "" {
		config.NotExistingFolder = "/not-existing-folder"
	}
	dir = config.NotExistingFolder
	wd = c.WebDav()
	for _, tt := range wtests {
		t.Run(tt.name, tt.test)
	}
}

func testCreateFolder(t *testing.T){
	err := wd.Mkdir(dir, 0777)
	require.NoError(t, err)
}

func testStat(t *testing.T) {
	i, err := wd.Stat(dir)
	require.NoError(t, err)
	// TODO : there is a problem with fileinfo's Name for directories: find a fix
	// assert.Equal(t, dir, i.Name())
	assert.True(t, i.IsDir())
}

func testWalk(t *testing.T) {
	found := false
	err := wd.Walk("/", func(path string, info os.FileInfo, err error) error {
		path = strings.Trim(path, "/")
		assert.NoError(t, err)
		if path == dir {
			found = true
		}
		p := strings.Split(path, "/")
		assert.Equal(t, p[len(p)-1], info.Name())
		return nil
	})
	assert.NoError(t, err)
	assert.True(t, found)
}

func testDelete(t *testing.T) {
	err := wd.Remove(dir)
	require.NoError(t, err)
}

