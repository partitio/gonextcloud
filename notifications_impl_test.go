package gonextcloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	notificationID int
	createdID      int
	title          = "Short Message"
	message        = "Longer notification message"
	tests          = []struct {
		string
		test
	}{
		{
			"notificationCreate",
			func(t *testing.T) {
				err := c.Notifications().Create(config.Login, title, message)
				assert.NoError(t, err)
			},
		}, {
			"notificationDelete",
			func(t *testing.T) {
				// Get created Notification ID
				ns, err := c.Notifications().List()
				if err != nil {
					t.SkipNow()
				}
				for _, n := range ns {
					if n.Subject == title {
						createdID = n.NotificationID
						break
					}
				}
				if createdID == 0 {
					t.SkipNow()
				}
				err = c.Notifications().Delete(createdID)
				assert.NoError(t, err)
			},
		},
	}
)

func TestNotificationsList(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	if err := c.Notifications().Available(); err != nil {
		t.SkipNow()
	}
	ns, err := c.Notifications().List()
	assert.NoError(t, err)
	if len(ns) > 0 {
		notificationID = ns[0].NotificationID
	}
}

func TestNotifications(t *testing.T) {
	if notificationID == 0 {
		t.SkipNow()
	}
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	if err := c.Notifications().Available(); err != nil {
		t.SkipNow()
	}
	n, err := c.Notifications().Get(notificationID)
	assert.NoError(t, err)
	assert.NotEmpty(t, n)
}

// Disable due to very long response time
//func TestNotificationsAdmin(t *testing.T) {
//	c = nil
//	if err := initClient(); err != nil {
//		t.Fatal(err)
//	}
//	if err := c.AdminAvailable(); err != nil {
//		t.SkipNow()
//	}
//	for _, test := range tests {
//		t.Run(test.string, test.test)
//	}
//}
