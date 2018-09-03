package gonextcloud

import (
	"fmt"
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/types"
	"net/http"
	"strconv"
	"sync"
)

func (c *Client) SharesList() ([]types.Share, error) {
	res, err := c.baseRequest(http.MethodGet, routes.shares, nil)
	if err != nil {
		return nil, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (c *Client) Shares(path string, reshares bool, subfiles bool) ([]types.Share, error) {
	ro := &req.RequestOptions{
		Params: map[string]string{
			"path":     path,
			"reshares": strconv.FormatBool(reshares),
			"subfiles": strconv.FormatBool(subfiles),
		},
	}
	res, err := c.baseRequest(http.MethodGet, routes.shares, ro)
	if err != nil {
		return nil, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (c *Client) Share(shareID string) (types.Share, error) {
	res, err := c.baseRequest(http.MethodGet, routes.shares, nil, shareID)
	if err != nil {
		return types.Share{}, err
	}
	var r types.SharesListResponse
	res.JSON(&r)
	return r.Ocs.Data[0], nil
}

func (c *Client) ShareCreate(
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
	res, err := c.baseRequest(http.MethodPost, routes.shares, ro)
	if err != nil {
		return types.Share{}, err
	}
	var r types.SharesResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

func (c *Client) ShareDelete(shareID int) error {
	_, err := c.baseRequest(http.MethodDelete, routes.shares, nil, strconv.Itoa(shareID))
	return err
}

// expireDate expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (c *Client) ShareUpdate(shareUpdate types.ShareUpdate) error {
	errs := make(chan types.UpdateError)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		if err := c.ShareUpdatePassword(shareUpdate.ShareID, shareUpdate.Password); err != nil {
			errs <- types.UpdateError{
				Field: "password",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := c.ShareUpdateExpireDate(shareUpdate.ShareID, shareUpdate.ExpireDate); err != nil {
			errs <- types.UpdateError{
				Field: "expireDate",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := c.ShareUpdatePermissions(shareUpdate.ShareID, shareUpdate.Permissions); err != nil {
			errs <- types.UpdateError{
				Field: "permissions",
				Error: err,
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := c.ShareUpdatePublicUpload(shareUpdate.ShareID, shareUpdate.PublicUpload); err != nil {
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

// expireDate expects a well formatted date string, e.g. ‘YYYY-MM-DD’
func (c *Client) ShareUpdateExpireDate(shareID int, expireDate string) error {
	return c.baseShareUpdate(strconv.Itoa(shareID), "expireDate", expireDate)
}

func (c *Client) ShareUpdatePublicUpload(shareID int, public bool) error {
	return c.baseShareUpdate(strconv.Itoa(shareID), "publicUpload", strconv.FormatBool(public))
}

func (c *Client) ShareUpdatePassword(shareID int, password string) error {
	return c.baseShareUpdate(strconv.Itoa(shareID), "password", password)
}

func (c *Client) ShareUpdatePermissions(shareID int, permissions types.SharePermission) error {
	return c.baseShareUpdate(strconv.Itoa(shareID), "permissions", strconv.Itoa(int(permissions)))
}

func (c *Client) baseShareUpdate(shareID string, key string, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{key: value},
	}
	_, err := c.baseRequest(http.MethodPut, routes.shares, ro, shareID)
	return err
}
