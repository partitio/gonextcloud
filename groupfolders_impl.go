package gonextcloud

import (
	"fmt"
	"net/http"
	"strconv"

	req "github.com/levigross/grequests"
)

//groupFolders contains all groups Folders available actions
type groupFolders struct {
	c *client
}

//List returns the groups folders
func (g *groupFolders) List() (map[int]GroupFolder, error) {
	res, err := g.c.baseOcsRequest(http.MethodGet, routes.groupfolders, nil)
	if err != nil {
		return nil, err
	}
	var r groupFoldersListResponse
	res.JSON(&r)
	gfs := formatBadIDAndGroups(r.Ocs.Data)
	return gfs, nil
}

//Get returns the group folder details
func (g *groupFolders) Get(id int) (GroupFolder, error) {
	res, err := g.c.baseOcsRequest(http.MethodGet, routes.groupfolders, nil, strconv.Itoa(id))
	if err != nil {
		return GroupFolder{}, err
	}
	var r groupFoldersResponse
	res.JSON(&r)
	if r.Ocs.Data.ID == 0 {
		return GroupFolder{}, fmt.Errorf("%d is not a valid groupfolder's id", id)
	}
	return r.Ocs.Data.FormatGroupFolder(), nil
}

//Create creates a group folder
func (g *groupFolders) Create(name string) (id int, err error) {
	// TODO: Validate Folder name
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	res, err := g.c.baseOcsRequest(http.MethodPost, routes.groupfolders, ro)
	if err != nil {
		return 0, err
	}
	var r groupFoldersCreateResponse
	res.JSON(&r)
	id, _ = strconv.Atoi(r.Ocs.Data.ID)
	return id, nil
}

//Rename renames the group folder
func (g *groupFolders) Rename(groupID int, name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	// groupFolders's response does not give any clues about success or failure
	_, err := g.c.baseOcsRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(groupID), "mountpoint")
	if err != nil {
		return err
	}
	return nil
}

//TODO func (c *client) GroupFoldersDelete(id int) error {

//AddGroup adds group to folder
func (g *groupFolders) AddGroup(folderID int, groupName string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"group": groupName,
		},
	}
	// groupFolders's response does not give any clues about success or failure
	_, err := g.c.baseOcsRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups")
	if err != nil {
		return err
	}
	return nil
}

//RemoveGroup remove a group from the group folder
func (g *groupFolders) RemoveGroup(folderID int, groupName string) error {
	// groupFolders's response does not give any clues about success or failure
	_, err := g.c.baseOcsRequest(http.MethodDelete, routes.groupfolders, nil, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

//SetGroupPermissions set groups permissions
func (g *groupFolders) SetGroupPermissions(folderID int, groupName string, permission SharePermission) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"permissions": strconv.Itoa(int(permission)),
		},
	}
	// groupFolders's response does not give any clues about success or failure
	_, err := g.c.baseOcsRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

//SetQuota set quota on the group folder. quota in bytes, use -3 for unlimited
func (g *groupFolders) SetQuota(folderID int, quota int) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"quota": strconv.Itoa(int(quota)),
		},
	}
	// groupFolders's response does not give any clues about success or failure
	_, err := g.c.baseOcsRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "quota")
	if err != nil {
		return err
	}
	return nil
}

func formatBadIDAndGroups(g map[string]groupFolderBadFormatIDAndGroups) map[int]GroupFolder {
	var gfs = map[int]GroupFolder{}
	for k := range g {
		i, _ := strconv.Atoi(k)
		d := g[k]
		gfs[i] = d.FormatGroupFolder()
	}
	return gfs
}
