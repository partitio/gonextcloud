package client

import "github.com/partitio/gonextcloud/client/types"

type BaseClient interface {
	NewClient(hostname string) (*Client, error)

	Login(username string, password string) error
	Logout() error

	User(name string) (*types.User, error)
	UserSearch(search string) ([]string, error)
	UserList() ([]string, error)
	UserCreate(username string, password string) error
	UserDelete(name string) error
	UserDisable(name string) error
	UserEnable(name string) error
	UserGroupAdd(name string, group string) error
	UserGroupDemote(name string, group string) error
	UserGroupList(name string) ([]string, error)
	UserGroupPromote(name string, group string) error
	UserGroupRemove(name string, group string) error
	UserGroupSubAdminList(name string) ([]string, error)
	UserSendWelcomeEmail(name string) error

	UserUpdateAddress(name string, address string) error
	UserUpdateDisplayName(name string, displayName string) error
	UserUpdateEmail(name string, email string) error
	UserUpdatePassword(name string, password string) error
	UserUpdatePhone(name string, phone string) error
	UserUpdateQuota(name string, quota string) error
	UserUpdateTwitter(name string, twitter string) error
	UserUpdateWebSite(name string, website string) error

	GroupSearch(search string) ([]string, error)
	GroupList() ([]string, error)
	GroupUsers(name string) ([]string, error)
	GroupCreate(name string) error
	GroupDelete(name string) error
	GroupSubAdminList(name string) ([]string, error)
}
