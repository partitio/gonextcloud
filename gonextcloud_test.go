package gonextcloud

import (
	"github.com/partitio/gonextcloud/client"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

var config = Config{}
var c *client.Client

type Config struct {
	URL              string   `yaml:"url"`
	Login            string   `yaml:"login"`
	Password         string   `yaml:"password"`
	AppName          string   `yaml:"app-name"`
	GroupsToCreate   []string `yaml:"groups-to-create"`
	NotExistingUser  string   `yaml:"not-existing-user"`
	NotExistingGroup string   `yaml:"not-existing-group"`
}

func LoadConfig() error {
	f, err := os.Open("./config.yml")
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		return err
	}
	return nil
}

func TestTruth(t *testing.T) {
	assert.Equal(t, true, true, "seriously ??!")
}

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	assert.Nil(t, err)
}

func TestClient(t *testing.T) {
	var err error
	c, err = client.NewClient(config.URL)
	assert.Nil(t, err, "aie")
}

func TestLoginFail(t *testing.T) {
	err := c.Login("", "")
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	err := c.Login(config.Login, config.Password)
	assert.Nil(t, err)
}

func TestUserList(t *testing.T) {
	us, err := c.UserList()
	assert.Nil(t, err)

	assert.Contains(t, us, config.Login)
}

func TestExistingUser(t *testing.T) {
	u, err := c.User(config.Login)
	assert.Nil(t, err)
	assert.NotNil(t, u)
}

func TestEmptyUser(t *testing.T) {
	u, err := c.User("")
	assert.NotNil(t, err)
	assert.Empty(t, u)
}

func TestNonExistingUser(t *testing.T) {
	_, err := c.User(config.NotExistingUser)
	assert.NotNil(t, err)
}

func TestUserSearch(t *testing.T) {
	us, err := c.UserSearch(config.Login)
	assert.Nil(t, err)
	assert.Contains(t, us, config.Login)
}

func TestUserCreate(t *testing.T) {
	err := c.UserCreate(config.NotExistingUser, "somecomplicatedpassword")
	assert.Nil(t, err)
}

func TestUserCreateExisting(t *testing.T) {
	err := c.UserCreate(config.NotExistingUser, "somecomplicatedpassword")
	assert.NotNil(t, err)
}

func TestGroupList(t *testing.T) {
	gs, err := c.GroupList()
	assert.Nil(t, err)
	assert.Contains(t, gs, "admin")
}

func TestGroupCreate(t *testing.T) {
	err := c.GroupCreate(config.NotExistingGroup)
	assert.Nil(t, err)
}

func TestUserUpdateEmail(t *testing.T) {
	email := "my@mail.com"
	err := c.UserUpdateEmail(config.NotExistingUser, email)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, email, u.Email)
}

func TestUserUpdateDisplayName(t *testing.T) {
	displayName := "Display Name"
	err := c.UserUpdateDisplayName(config.NotExistingUser, displayName)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, displayName, u.Displayname)
}

func TestUserUpdatePhone(t *testing.T) {
	phone := "+33 42 42 42 42"
	err := c.UserUpdatePhone(config.NotExistingUser, phone)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, phone, u.Phone)
}

func TestUserUpdateAddress(t *testing.T) {
	address := "Main Street, Galifrey"
	err := c.UserUpdateAddress(config.NotExistingUser, address)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, address, u.Address)
}

func TestUserUpdateWebSite(t *testing.T) {
	website := "www.doctor.who"
	err := c.UserUpdateWebSite(config.NotExistingUser, website)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, website, u.Website)
}

func TestUserUpdateTwitter(t *testing.T) {
	twitter := "@doctorwho"
	err := c.UserUpdateTwitter(config.NotExistingUser, twitter)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Equal(t, twitter, u.Twitter)
}

func TestUserUpdateQuota(t *testing.T) {
	quota := 1024 * 1024 * 1024
	err := c.UserUpdateQuota(config.NotExistingUser, quota)
	assert.Nil(t, err)
	// TODO : Find better verification : A never connected User does not have quota available
	//u, err := c.User(config.NotExistingUser)
	//assert.Nil(t, err)
	//assert.Equal(t, quota, u.Quota.Quota)
}

func TestUserUpdatePassword(t *testing.T) {
	password := "newcomplexpassword"
	err := c.UserUpdatePassword(config.NotExistingUser, password)
	assert.Nil(t, err)
}

func TestUserGroupAdd(t *testing.T) {
	err := c.UserGroupAdd(config.NotExistingUser, config.NotExistingGroup)
	assert.Nil(t, err)
	gs, err := c.UserGroupList(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Contains(t, gs, config.NotExistingGroup)
}

func TestUserGroupSubAdminList(t *testing.T) {
	gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	assert.NotNil(t, err)
	assert.Empty(t, gs)
}

func TestUserGroupPromote(t *testing.T) {
	err := c.UserGroupPromote(config.NotExistingUser, config.NotExistingGroup)
	assert.Nil(t, err)
	gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	assert.Nil(t, err)
	assert.Contains(t, gs, config.NotExistingGroup)
}

func TestUserGroupDemote(t *testing.T) {
	err := c.UserGroupDemote(config.NotExistingUser, config.NotExistingGroup)
	assert.Nil(t, err)
	//gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	//assert.Nil(t, err)
	//assert.Empty(t, gs)
}

func TestUserDisable(t *testing.T) {
	err := c.UserDisable(config.NotExistingUser)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.False(t, u.Enabled)
}

func TestUserEnable(t *testing.T) {
	err := c.UserEnable(config.NotExistingUser)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.Nil(t, err)
	assert.True(t, u.Enabled)
}

func TestGroupDelete(t *testing.T) {
	err := c.GroupDelete(config.NotExistingGroup)
	assert.Nil(t, err)
}

func TestUserDelete(t *testing.T) {
	err := c.UserDelete(config.NotExistingUser)
	assert.Nil(t, err)
}

func TestLogout(t *testing.T) {
	err := c.Logout()
	assert.Nil(t, err)
}
