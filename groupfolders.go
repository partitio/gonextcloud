package gonextcloud

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/http"
	"strconv"
)

func (c *Client) GroupFoldersList() (map[int]types.GroupFolder, error) {
	res, err := c.baseRequest(http.MethodGet, routes.groupfolders, nil)
	if err != nil {
		return nil, err
	}
	var r types.GroupFoldersListResponse
	res.JSON(&r)
	gfs := formatBadIDAndGroups(r.Ocs.Data)
	return gfs, nil
}

func (c *Client) GroupFolders(id int) (types.GroupFolder, error) {
	res, err := c.baseRequest(http.MethodGet, routes.groupfolders, nil, strconv.Itoa(id))
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

func (c *Client) GroupFoldersCreate(name string) (id int, err error) {
	// TODO: Validate Folder name
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	res, err := c.baseRequest(http.MethodPost, routes.groupfolders, ro)
	if err != nil {
		return 0, err
	}
	var r types.GroupFoldersCreateResponse
	res.JSON(&r)
	id, _ = strconv.Atoi(r.Ocs.Data.ID)
	return id, nil
}

func (c *Client) GroupFoldersRename(groupID int, name string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"mountpoint": name,
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(groupID), "mountpoint")
	if err != nil {
		return err
	}
	return nil
}

//TODO func (c *Client) GroupFoldersDelete(id int) error {
//	// GroupFolders's response does not give any clues about success or failure
//	return nil
//}

func (c *Client) GroupFoldersAddGroup(folderID int, groupName string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"group": groupName,
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups")
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GroupFoldersRemoveGroup(folderID int, groupName string) error {
	// GroupFolders's response does not give any clues about success or failure
	_, err := c.baseRequest(http.MethodDelete, routes.groupfolders, nil, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GroupFoldersSetGroupPermissions(folderID int, groupName string, permission types.SharePermission) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"permissions": strconv.Itoa(int(permission)),
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "groups", groupName)
	if err != nil {
		return err
	}
	return nil
}

//GroupFoldersSetQuota set quota on the group folder. quota in bytes, use -3 for unlimited
func (c *Client) GroupFoldersSetQuota(folderID int, quota int) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"quota": strconv.Itoa(int(quota)),
		},
	}
	// GroupFolders's response does not give any clues about success or failure
	_, err := c.baseRequest(http.MethodPost, routes.groupfolders, ro, strconv.Itoa(folderID), "quota")
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
