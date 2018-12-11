package gonextcloud

import (
	"errors"
	req "github.com/levigross/grequests"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"strconv"
)

//Notifications contains all Notifications available actions
type Notifications struct {
	c *Client
}

//List returns all the notifications
func (n *Notifications) List() ([]types.Notification, error) {
	if err := n.Available(); err != nil {
		return nil, err
	}
	res, err := n.c.baseRequest(http.MethodGet, routes.notifications, nil)
	if err != nil {
		return nil, err
	}
	var r types.NotificationsListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Get returns the notification corresponding to the id
func (n *Notifications) Get(id int) (types.Notification, error) {
	if err := n.Available(); err != nil {
		return types.Notification{}, err
	}
	res, err := n.c.baseRequest(http.MethodGet, routes.notifications, nil, strconv.Itoa(id))
	if err != nil {
		return types.Notification{}, err
	}
	var r types.NotificationResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Delete deletes the notification corresponding to the id
func (n *Notifications) Delete(id int) error {
	if err := n.Available(); err != nil {
		return err
	}
	_, err := n.c.baseRequest(http.MethodDelete, routes.notifications, nil, strconv.Itoa(id))
	return err
}

//DeleteAll deletes all notifications
func (n *Notifications) DeleteAll() error {
	if err := n.Available(); err != nil {
		return err
	}
	_, err := n.c.baseRequest(http.MethodDelete, routes.notifications, nil)
	return err
}

//Create creates a notification (if the user is an admin)
func (n *Notifications) Create(userID, title, message string) error {
	if err := n.AdminAvailable(); err != nil {
		return err
	}
	ro := &req.RequestOptions{
		Data: map[string]string{
			"shortMessage": title,
			"longMessage":  message,
		},
	}
	_, err := n.c.baseRequest(http.MethodPost, routes.adminNotifications, ro, userID)
	return err
}

//AdminAvailable returns an error if the admin-notifications app is not installed
func (n *Notifications) AdminAvailable() error {
	if len(n.c.capabilities.Notifications.AdminNotifications) == 0 {
		return errors.New("'admin notifications' not available on this instance")
	}
	return nil
}

//Available returns an error if the notifications app is not installed
func (n *Notifications) Available() error {
	if len(n.c.capabilities.Notifications.OcsEndpoints) == 0 {
		return errors.New("notifications not available on this instance")
	}
	return nil
}
