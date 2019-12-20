package gonextcloud

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	req "github.com/levigross/grequests"
)

//shares contains all shares available actions
type shares struct {
	c *client
}

//List list all shares of the logged in user
func (s *shares) List() ([]Share, error) {
	res, err := s.c.baseOcsRequest(http.MethodGet, routes.shares, nil)
	if err != nil {
		return nil, err
	}
	var r sharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//GetFromPath return shares from a specific file or folder
func (s *shares) GetFromPath(path string, reshares bool, subfiles bool) ([]Share, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{
			"path":     path,
			"reshares": strconv.FormatBool(reshares),
			"subfiles": strconv.FormatBool(subfiles),
		},
	}
	res, err := s.c.baseOcsRequest(http.MethodGet, routes.shares, ro)
	if err != nil {
		return nil, err
	}
	var r sharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Get information about a known Share
func (s *shares) Get(shareID string) (Share, error) {
	res, err := s.c.baseOcsRequest(http.MethodGet, routes.shares, nil, shareID)
	if err != nil {
		return Share{}, err
	}
	var r sharesListResponse
	res.JSON(&r)
	return r.Ocs.Data[0], nil
}

//Create create a share
func (s *shares) Create(
	path string,
	shareType ShareType,
	permission SharePermission,
	shareWith string,
	publicUpload bool,
	password string,
) (Share, error) {

	if (shareType == UserShare || shareType == GroupShare) && shareWith == "" {
		return Share{}, fmt.Errorf("shareWith cannot be empty for ShareType UserShare or GroupShare")
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
	res, err := s.c.baseOcsRequest(http.MethodPost, routes.shares, ro)
	if err != nil {
		return Share{}, err
	}
	var r sharesResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Delete Remove the given share.
func (s *shares) Delete(shareID int) error {
	_, err := s.c.baseOcsRequest(http.MethodDelete, routes.shares, nil, strconv.Itoa(shareID))
	return err
}

// Update update share details
// expireDate expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (s *shares) Update(shareUpdate ShareUpdate) error {
	errs := make(chan *UpdateError)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		if err := s.UpdatePassword(shareUpdate.ShareID, shareUpdate.Password); err != nil {
			errs <- &UpdateError{
				Field: "password",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdateExpireDate(shareUpdate.ShareID, shareUpdate.ExpireDate); err != nil {
			errs <- &UpdateError{
				Field: "expireDate",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdatePermissions(shareUpdate.ShareID, shareUpdate.Permissions); err != nil {
			errs <- &UpdateError{
				Field: "permissions",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := s.UpdatePublicUpload(shareUpdate.ShareID, shareUpdate.PublicUpload); err != nil {
			errs <- &UpdateError{
				Field: "publicUpload",
				Error: err,
			}
		}
	}()
	go func() {
		wg.Wait()
		close(errs)
	}()
	if err := newUpdateError(errs); err != nil {
		return err
	}
	return nil
}

//UpdateExpireDate updates the share's expire date
// expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (s *shares) UpdateExpireDate(shareID int, expireDate string) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "expireDate", expireDate)
}

//UpdatePublicUpload enable or disable public upload
func (s *shares) UpdatePublicUpload(shareID int, public bool) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "publicUpload", strconv.FormatBool(public))
}

//UpdatePassword updates share password
func (s *shares) UpdatePassword(shareID int, password string) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "password", password)
}

//UpdatePermissions update permissions
func (s *shares) UpdatePermissions(shareID int, permissions SharePermission) error {
	return s.baseShareUpdate(strconv.Itoa(shareID), "permissions", strconv.Itoa(int(permissions)))
}

func (s *shares) baseShareUpdate(shareID string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{key: value},
	}
	_, err := s.c.baseOcsRequest(http.MethodPut, routes.shares, ro, shareID)
	return err
}
