package gonextcloud

import (
	"errors"
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"strconv"
)

//NotificationsList returns all the notifications
func (c *Client) NotificationsList() ([]types.Notification, error) {
	if err := c.notificationsAvailable(); err != nil {
		return nil, err
	}
	res, err := c.baseRequest(http.MethodGet, routes.notifications, nil)
	if err != nil {
		return nil, err
	}
	var r types.NotificationsListResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//Notifications returns the notification corresponding to the id
func (c *Client) Notifications(id int) (types.Notification, error) {
	if err := c.notificationsAvailable(); err != nil {
		return types.Notification{}, err
	}
	res, err := c.baseRequest(http.MethodGet, routes.notifications, nil, strconv.Itoa(id))
	if err != nil {
		return types.Notification{}, err
	}
	var r types.NotificationResponse
	res.JSON(&r)
	return r.Ocs.Data, nil
}

//NotificationsDelete deletes the notification corresponding to the id
func (c *Client) NotificationsDelete(id int) error {
	if err := c.notificationsAvailable(); err != nil {
		return err
	}
	_, err := c.baseRequest(http.MethodDelete, routes.notifications, nil, strconv.Itoa(id))
	return err
}

//NotificationsDeleteAll deletes all notifications
func (c *Client) NotificationsDeleteAll() error {
	if err := c.notificationsAvailable(); err != nil {
		return err
	}
	_, err := c.baseRequest(http.MethodDelete, routes.notifications, nil)
	return err
}

//NotificationsCreate creates a notification (if the user is an admin)
func (c *Client) NotificationsCreate(userID, title, message string) error {
	if err := c.adminNotificationsAvailable(); err != nil {
		return err
	}
	ro := &req.RequestOptions{
		Data: map[string]string{
			"shortMessage": title,
			"longMessage":  message,
		},
	}
	_, err := c.baseRequest(http.MethodPost, routes.adminNotifications, ro, userID)
	return err
}

func (c *Client) adminNotificationsAvailable() error {
	if len(c.capabilities.Notifications.AdminNotifications) == 0 {
		return errors.New("'admin notifications' not available on this instance")
	}
	return nil
}
func (c *Client) notificationsAvailable() error {
	if len(c.capabilities.Notifications.OcsEndpoints) == 0 {
		return errors.New("notifications not available on this instance")
	}
	return nil
}
