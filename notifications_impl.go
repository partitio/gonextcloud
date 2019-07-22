package gonextcloud

import (
	"errors"
	"net/http"
	"strconv"

	req "github.com/levigross/grequests"
)

//notifications contains all notifications available actions
type notifications struct {
	c *client
}

//List returns all the notifications
func (n *notifications) List() ([]Notification, error) {
	if err := n.Available(); err != nil {
		return nil, err
	}
	res, err := n.c.baseRequest(http.MethodGet, routes.notifications, nil)
	if err != nil {
		return nil, err
	}
	var r notificationsListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Get returns the notification corresponding to the id
func (n *notifications) Get(id int) (Notification, error) {
	if err := n.Available(); err != nil {
		return Notification{}, err
	}
	res, err := n.c.baseRequest(http.MethodGet, routes.notifications, nil, strconv.Itoa(id))
	if err != nil {
		return Notification{}, err
	}
	var r notificationResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Delete deletes the notification corresponding to the id
func (n *notifications) Delete(id int) error {
	if err := n.Available(); err != nil {
		return err
	}
	_, err := n.c.baseRequest(http.MethodDelete, routes.notifications, nil, strconv.Itoa(id))
	return err
}

//DeleteAll deletes all notifications
func (n *notifications) DeleteAll() error {
	if err := n.Available(); err != nil {
		return err
	}
	_, err := n.c.baseRequest(http.MethodDelete, routes.notifications, nil)
	return err
}

//Create creates a notification (if the user is an admin)
func (n *notifications) Create(userID, title, message string) error {
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
func (n *notifications) AdminAvailable() error {
	if len(n.c.capabilities.Notifications.AdminNotifications) == 0 {
		return errors.New("'admin notifications' not available on this instance")
	}
	return nil
}

//Available returns an error if the notifications app is not installed
func (n *notifications) Available() error {
	if len(n.c.capabilities.Notifications.OcsEndpoints) == 0 {
		return errors.New("notifications not available on this instance")
	}
	return nil
}
