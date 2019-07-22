package gonextcloud

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	dir     string
	wd      WebDav
	folders  = []string{
		"folder1",
		"folder1/sub1",
		"folder1/sub1/ssub1",
		"folder1/sub2",
		"folder1/sub2/ssub1",
		"folder1/sub2/ssub2",
		"folder2",
		"folder2/sub2",
		"folder2/sub3",
		"folder2/sub3/ssub1",
		"folder2/sub4",
	}
	wtests = []struct {
		name string
		test test
	}{
		{
			name: "CreateFolder",
			test: testCreateFolder,
		},
		{
			name: "TestCreateSubFolders",
			test: testCreateSubFolders,
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

func testCreateFolder(t *testing.T) {
	err := wd.Mkdir(dir, 0777)
	require.NoError(t, err)
}

func testStat(t *testing.T) {
	i, err := wd.Stat(dir)
	require.NoError(t, err)
	assert.Equal(t, dir, i.Name())
	assert.True(t, i.IsDir())
}

func testCreateSubFolders(t *testing.T) {
	sort.Strings(folders)
	d := strings.TrimRight(dir, "/")
	var ds []string
	ds = append(ds, d)
	for _, f := range folders {
		p := d + "/" + f
		err := wd.MkdirAll(p, 0777)
		assert.NoError(t, err)
		ds = append(ds, p)
	}
	folders = ds
}

func testWalk(t *testing.T) {
	err := wd.Walk(dir, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err)
		assert.Equal(t, filepath.Base(path), info.Name())
		assert.Contains(t, folders, path)
		return nil
	})
	assert.NoError(t, err)
}

func testDelete(t *testing.T) {
	err := wd.Remove(dir)
	require.NoError(t, err)
}
