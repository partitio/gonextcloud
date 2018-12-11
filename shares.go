package gonextcloud

import (
	"fmt"
	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"strconv"
	"sync"
)

//Shares contains all Shares available actions
type Shares struct {
	c *Client
}

//List list all shares of the logged in user
func (s *Shares) List() ([]types.Share, error) {
	res, err := s.c.baseRequest(http.MethodGet, routes.shares, nil)
	if err != nil {
		return nil, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//GetFromPath return shares from a specific file or folder
func (s *Shares) GetFromPath(path string, reshares bool, subfiles bool) ([]types.Share, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{
			"path":     path,
			"reshares": strconv.FormatBool(reshares),
			"subfiles": strconv.FormatBool(subfiles),
		},
	}
	res, err := s.c.baseRequest(http.MethodGet, routes.shares, ro)
	if err != nil {
		return nil, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Get information about a known Share
func (s *Shares) Get(shareID string) (types.Share, error) {
	res, err := s.c.baseRequest(http.MethodGet, routes.shares, nil, shareID)
	if err != nil {
		return types.Share{}, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data[0], nil
}

//Create create a share
func (s *Shares) Create(
	path string,
	shareType types.ShareType,
	permission types.SharePermission,
	shareWith string,
	publicUpload bool,
	password string,
) (types.Share, error) {

	if (shareType == types.UserShare || shareType == types.GroupShare) && shareWith == "" {
		return types.Share{}, fmt.Errorf("shareWith cannot be empty for ShareType UserShare or GroupShare")
	}
	ro := &req.RequestOptions{
		Data: map[string]string{
			"path":         path,
			"shareType":    strconv.Itoa(int(shareType)),
			"shareWith":    shareWith,
			"publicUpload": strconv.FormatBool(publicUpload),
			"password":     password,
			"permissions":  strconv.Itoa(int(permission)),
		},
	}
	res, err := s.c.baseRequest(http.MethodPost, routes.shares, ro)
	if err != nil {
		return types.Share{}, err
	}
	var r types.SharesResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Delete Remove the given share.
func (s *Shares) Delete(shareID int) error {
	_, err := s.c.baseRequest(http.MethodDelete, routes.shares, nil, strconv.Itoa(shareID))
	return err
}

// Update update share details
// expireDate expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (s *Shares) Update(shareUpdate types.ShareUpdate) error {
	errs := make(chan types.UpdateError)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		if err := s.UpdatePassword(shareUpdate.ShareID, shareUpdate.Password); err != nil {
			errs <- types.UpdateError{
				Field: "password",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdateExpireDate(shareUpdate.ShareID, shareUpdate.ExpireDate); err != nil {
			errs <- types.UpdateError{
				Field: "expireDate",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdatePermissions(shareUpdate.ShareID, shareUpdate.Permissions); err != nil {
			errs <- types.UpdateError{
				Field: "permissions",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdatePublicUpload(shareUpdate.ShareID, shareUpdate.PublicUpload); err != nil {
			errs <- types.UpdateError{
				Field: "publicUpload",
				Error: err,
			}
		}
	}()
	go func() {
		wg.Wait()
		close(errs)
	}()
	return types.NewUpdateError(errs)
}

//UpdateExpireDate updates the share's expire date
// expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (s *Shares) UpdateExpireDate(shareID int, expireDate string) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "expireDate", expireDate)
}

//UpdatePublicUpload enable or disable public upload
func (s *Shares) UpdatePublicUpload(shareID int, public bool) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "publicUpload", strconv.FormatBool(public))
}

//UpdatePassword updates share password
func (s *Shares) UpdatePassword(shareID int, password string) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "password", password)
}

//UpdatePermissions update permissions
func (s *Shares) UpdatePermissions(shareID int, permissions types.SharePermission) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "permissions", strconv.Itoa(int(permissions)))
}

func (s *Shares) baseShareUpdate(shareID string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{key: value},
	}
	_, err := s.c.baseRequest(http.MethodPut, routes.shares, ro, shareID)
	return err
}
