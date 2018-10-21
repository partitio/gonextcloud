package gonextcloud

import (
	"fmt"
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"strconv"
)

//GroupFoldersI available methods
type GroupFoldersI interface {
	List() (map[int]types.GroupFolder, error)
	Get(id int) (types.GroupFolder, error)
	Create(name string) (id int, err error)
	Rename(groupID int, name string) error
	AddGroup(folderID int, groupName string) error
	RemoveGroup(folderID int, groupName string) error
	SetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error
	SetQuota(folderID int, quota int) error
}

//GroupFolders contains all Groups Folders available actions
type GroupFolders struct {
	c *Client
}

//List returns the groups folders
func (g *GroupFolders) List() (map[int]types.GroupFolder, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groupfolders, nil)
	if err != nil {
		return nil, err
	}
	var r types.GroupFoldersListResponse
	res.JSON(&r)
	gfs := formatBadIDAndGroups(r.Ocs.Data)
	return gfs, nil
}

//Get returns the group folder details
func (g *GroupFolders) Get(id int) (types.GroupFolder, error) {
	res, err := g.c.baseRequest(http.MethodGet, routes.groupfolders, nil, strconv.Itoa(id))
	if err != nil {
		return types.GroupFolder{}, err
	}
	var r types.GroupFoldersResponse
	res.JSON(&r)
	if r.Ocs.Data.ID == 0 {
		return types.GroupFolder{}, fmt.Errorf("%d is not a valid groupfolder's id", id)
	}
	return r.Ocs.Data.FormatGroupFolder(), nil
}

//Create creates a group folder
func (g *GroupFolders) Create(name string) (id int, err error) {
	// TODO: Validate Folder name
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	res, err := g.c.baseRequest(http.MethodPost, routes.groupfolders, ro)
	if err != nil {
		return 0, err
	}
	var r types.GroupFoldersCreateResponse
	res.JSON(&r)
	id, _ = strconv.Atoi(r.Ocs.Data.ID)
	return id, nil
}

//Rename renames the group folder
func (g *GroupFolders) Rename(groupID int, name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := g.c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(groupID), "mountpoint")
	if err != nil {
		return err
	}
	return nil
}

//TODO func (c *Client) GroupFoldersDelete(id int) error {

//AddGroup adds group to folder
func (g *GroupFolders) AddGroup(folderID int, groupName string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"group": groupName,
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := g.c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups")
	if err != nil {
		return err
	}
	return nil
}

//RemoveGroup remove a group from the group folder
func (g *GroupFolders) RemoveGroup(folderID int, groupName string) error {
	// GroupFolders's response does not give any clues about success or failure
	_, err := g.c.baseRequest(http.MethodDelete, routes.groupfolders, nil, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

//SetGroupPermissions set groups permissions
func (g *GroupFolders) SetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"permissions": strconv.Itoa(int(permission)),
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := g.c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

//SetQuota set quota on the group folder. quota in bytes, use -3 for unlimited
func (g *GroupFolders) SetQuota(folderID int, quota int) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"quota": strconv.Itoa(int(quota)),
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := g.c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "quota")
	if err != nil {
		return err
	}
	return nil
}

func formatBadIDAndGroups(g map[string]types.GroupFolderBadFormatIDAndGroups) map[int]types.GroupFolder {
	var gfs = map[int]types.GroupFolder{}
	for k := range g {
		i, _ := strconv.Atoi(k)
		d := g[k]
		gfs[i] = d.FormatGroupFolder()
	}
	return gfs
}
