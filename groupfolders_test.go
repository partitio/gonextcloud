package gonextcloud

import (
	"github.com/stretchr/testify/assert"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"testing"
)

var (
	groupFoldersTests = []struct {
		string
		test
	}{
		{
			"TestGroupFoldersCreate",
			func(t *testing.T) {
				// Recreate client
				var err error
				groupID, err = c.GroupFolders().Create("API")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersList",
			func(t *testing.T) {
				gfs, err := c.GroupFolders().List()
				assert.NoError(t, err)
				assert.NotNil(t, gfs[groupID])
			},
		},
		{
			"TestGroupFolders",
			func(t *testing.T) {
				gf, err := c.GroupFolders().Get(groupID)
				assert.NoError(t, err)
				assert.NotNil(t, gf)
			},
		},
		{
			"TestGroupFolderRename",
			func(t *testing.T) {
				err := c.GroupFolders().Rename(groupID, "API_Renamed")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersAddGroup",
			func(t *testing.T) {
				err := c.GroupFolders().AddGroup(groupID, "admin")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersSetGroupPermissions",
			func(t *testing.T) {
				err := c.GroupFolders().SetGroupPermissions(groupID, "admin", types.ReadPermission)
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersSetQuota",
			func(t *testing.T) {
				err := c.GroupFolders().SetQuota(groupID, 100)
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFolderRemoveGroup",
			func(t *testing.T) {
				err := c.GroupFolders().RemoveGroup(groupID, "admin")
				assert.NoError(t, err)
			},
		},
	}
)

func TestGroupFolders(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	for _, tt := range groupFoldersTests {
		t.Run(tt.string, tt.test)
	}
}
