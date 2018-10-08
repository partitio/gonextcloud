package gonextcloud

import (
	"github.com/partitio/gonextcloud/types"
	"github.com/partitio/swarmmanager/libnextcloudpartitio/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"
)

var config = Config{}
var c *Client

const password = "somecomplicatedpassword"

type Config struct {
	URL              string `yaml:"url"`
	Login            string `yaml:"login"`
	Password         string `yaml:"password"`
	AppName          string `yaml:"app-name"`
	ShareFolder      string `yaml:"share-folder"`
	NotExistingUser  string `yaml:"not-existing-user"`
	NotExistingGroup string `yaml:"not-existing-group"`
}

// LoadConfig loads the test configuration
func LoadConfig() error {
	f, err := os.Open("./config.yml")
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, &config)
}

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	assert.NoError(t, err)
}

func TestClient(t *testing.T) {
	var err error
	c, err = NewClient(config.URL)
	assert.NoError(t, err, "aie")
}

func TestLoginFail(t *testing.T) {
	err := c.Login("", "")
	assert.Error(t, err)
}

func TestLogin(t *testing.T) {
	err := c.Login(config.Login, config.Password)
	assert.NoError(t, err)
}

func TestUserList(t *testing.T) {
	us, err := c.UserList()
	assert.NoError(t, err)

	assert.Contains(t, us, config.Login)
}

func TestExistingUser(t *testing.T) {
	u, err := c.User(config.Login)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestEmptyUser(t *testing.T) {
	u, err := c.User("")
	assert.Error(t, err)
	assert.Empty(t, u)
}

func TestNonExistingUser(t *testing.T) {
	_, err := c.User(config.NotExistingUser)
	assert.Error(t, err)
}

func TestUserSearch(t *testing.T) {
	us, err := c.UserSearch(config.Login)
	assert.NoError(t, err)
	assert.Contains(t, us, config.Login)
}

func TestUserCreate(t *testing.T) {
	err := c.UserCreate(config.NotExistingUser, password, nil)
	assert.NoError(t, err)
}

//func TestUserCreateFull(t *testing.T) {
//	if err := initClient(); err != nil {
//		return
//	}
//	username := fmt.Sprintf("%s-2", config.NotExistingUser)
//	user := &types.User{
//		ID:          username,
//		Displayname: strings.ToUpper(username),
//		Email:       "some@address.com",
//		Address:     "Main Street, City",
//		Twitter:     "@me",
//		Phone:       "42 42 242 424",
//		Website:     "my.site.com",
//	}
//	err := c.UserCreate(username, password, user)
//	assert.Nil(t, err)
//	u, err := c.User(username)
//	assert.NoError(t, err)
//	o := structs.Map(user)
//	r := structs.Map(u)
//	for k := range o {
//		if ignoredUserField(k) {
//			continue
//		}
//		assert.Equal(t, o[k], r[k])
//	}
//	// Clean up
//	err = c.UserDelete(u.ID)
//	assert.NoError(t, err)
//}

//func TestUserUpdate(t *testing.T) {
//	if err := initClient(); err != nil {
//		return
//	}
//	username := fmt.Sprintf("%s-2", config.NotExistingUser)
//	err := c.UserCreate(username, password, nil)
//	assert.Nil(t, err)
//	user := &types.User{
//		ID:          username,
//		Displayname: strings.ToUpper(username),
//		Email:       "some@address.com",
//		Address:     "Main Street, City",
//		Twitter:     "@me",
//		Phone:       "42 42 242 424",
//		Website:     "my.site.com",
//	}
//	err = c.UserUpdate(user)
//	assert.Nil(t, err)
//	u, err := c.User(username)
//	assert.Nil(t, err)
//	o := structs.Map(user)
//	r := structs.Map(u)
//	for k := range o {
//		if ignoredUserField(k) {
//			continue
//		}
//		assert.Equal(t, o[k], r[k])
//	}
//	// Clean up
//	err = c.UserDelete(u.ID)
//	assert.NoError(t, err)
//}

func TestUserCreateExisting(t *testing.T) {
	err := c.UserCreate(config.NotExistingUser, password, nil)
	assert.Error(t, err)
}

func TestGroupList(t *testing.T) {
	gs, err := c.GroupList()
	assert.NoError(t, err)
	assert.Contains(t, gs, "admin")
}

func TestGroupCreate(t *testing.T) {
	err := c.GroupCreate(config.NotExistingGroup)
	assert.NoError(t, err)
}

func TestUserUpdateEmail(t *testing.T) {
	email := "my@mail.com"
	err := c.UserUpdateEmail(config.NotExistingUser, email)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, email, u.Email)
}

func TestUserUpdateDisplayName(t *testing.T) {
	displayName := "Display Name"
	err := c.UserUpdateDisplayName(config.NotExistingUser, displayName)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, displayName, u.Displayname)
}

func TestUserUpdatePhone(t *testing.T) {
	phone := "+33 42 42 42 42"
	err := c.UserUpdatePhone(config.NotExistingUser, phone)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, phone, u.Phone)
}

func TestUserUpdateAddress(t *testing.T) {
	address := "Main Street, Galifrey"
	err := c.UserUpdateAddress(config.NotExistingUser, address)
	assert.NoError(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, address, u.Address)
}

func TestUserUpdateWebSite(t *testing.T) {
	website := "www.doctor.who"
	err := c.UserUpdateWebSite(config.NotExistingUser, website)
	assert.NoError(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, website, u.Website)
}

func TestUserUpdateTwitter(t *testing.T) {
	twitter := "@doctorwho"
	err := c.UserUpdateTwitter(config.NotExistingUser, twitter)
	assert.NoError(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Equal(t, twitter, u.Twitter)
}

func TestUserUpdateQuota(t *testing.T) {
	quota := 1024 * 1024 * 1024
	err := c.UserUpdateQuota(config.NotExistingUser, quota)
	assert.NoError(t, err)
	// TODO : Find better verification : A never connected User does not have quota available
	//u, err := c.User(config.NotExistingUser)
	//assert.Nil(t, err)
	//assert.Equal(t, quota, u.Quota.Quota)
}

func TestUserUpdatePassword(t *testing.T) {
	password := "newcomplexpassword"
	err := c.UserUpdatePassword(config.NotExistingUser, password)
	assert.NoError(t, err)
}

func TestUserGroupAdd(t *testing.T) {
	err := c.UserGroupAdd(config.NotExistingUser, config.NotExistingGroup)
	assert.Nil(t, err)
	gs, err := c.UserGroupList(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Contains(t, gs, config.NotExistingGroup)
}

func TestUserGroupSubAdminList(t *testing.T) {
	gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Empty(t, gs)
}

func TestUserGroupPromote(t *testing.T) {
	err := c.UserGroupPromote(config.NotExistingUser, config.NotExistingGroup)
	assert.Nil(t, err)
	gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	assert.NoError(t, err)
	assert.Contains(t, gs, config.NotExistingGroup)
}

func TestUserGroupDemote(t *testing.T) {
	err := c.UserGroupDemote(config.NotExistingUser, config.NotExistingGroup)
	assert.NoError(t, err)
	//gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
	//assert.Nil(t, err)
	//assert.Empty(t, gs)
}

func TestUserDisable(t *testing.T) {
	err := c.UserDisable(config.NotExistingUser)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.False(t, u.Enabled)
}

func TestUserEnable(t *testing.T) {
	err := c.UserEnable(config.NotExistingUser)
	assert.Nil(t, err)
	u, err := c.User(config.NotExistingUser)
	assert.NoError(t, err)
	assert.True(t, u.Enabled)
}

func TestGroupDelete(t *testing.T) {
	err := c.GroupDelete(config.NotExistingGroup)
	assert.Nil(t, err)
}

func TestUserDelete(t *testing.T) {
	err := c.UserDelete(config.NotExistingUser)
	assert.NoError(t, err)
}

func TestInvalidBaseRequest(t *testing.T) {
	c.baseURL = &url.URL{}
	_, err := c.baseRequest(http.MethodGet, routes.capabilities, nil, "admin", "invalid")
	c = nil
	assert.Error(t, err)
}

func TestShareList(t *testing.T) {
	if err := initClient(); err != nil {
		return
	}
	s, err := c.SharesList()
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func TestLogout(t *testing.T) {
	err := c.Logout()
	assert.NoError(t, err)
	assert.Nil(t, c.session.HTTPClient.Jar)
}

func TestLoggedIn(t *testing.T) {
	c := &Client{}
	c.capabilities = &types.Capabilities{}
	assert.False(t, c.loggedIn())
}

func TestLoginInvalidURL(t *testing.T) {
	c, _ = NewClient("")
	err := c.Login("", "")
	assert.Error(t, err)
}

func TestBaseRequest(t *testing.T) {
	c, _ = NewClient("")
	_, err := c.baseRequest(http.MethodGet, routes.capabilities, nil, "admin", "invalid")
	assert.Error(t, err)
}

var groupID = 37

func TestGroupFoldersCreate(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	var err error
	groupID, err = c.GroupFoldersCreate("API")
	assert.NoError(t, err)
}

func TestGroupFoldersList(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	gfs, err := c.GroupFoldersList()
	assert.NoError(t, err)
	utils.PrettyPrint(gfs)
	assert.NotNil(t, gfs[groupID])
}

func TestGroupFolders(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	gf, err := c.GroupFolders(groupID)
	assert.NoError(t, err)
	utils.PrettyPrint(gf)
	assert.NotNil(t, gf)
}

func TestGroupFolderRename(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	err := c.GroupFoldersRename(groupID, "API_Renamed")
	assert.NoError(t, err)
}

func TestGroupFoldersAddGroup(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	err := c.GroupFoldersAddGroup(groupID, "admin")
	assert.NoError(t, err)
}

func TestGroupFoldersSetGroupPermissions(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	err := c.GroupFoldersSetGroupPermissions(groupID, "admin", types.ReadPermission)
	assert.NoError(t, err)
}

func TestGroupFoldersSetQuota(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	err := c.GroupFoldersSetQuota(groupID, 100)
	assert.NoError(t, err)
}

func TestGroupFolderRemoveGroup(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		return
	}
	err := c.GroupFoldersRemoveGroup(groupID, "admin")
	assert.NoError(t, err)
}

func initClient() error {
	if c == nil {
		if err := LoadConfig(); err != nil {
			return err
		}
		var err error
		c, err = NewClient(config.URL)
		if err != nil {
			return err
		}
		if err = c.Login(config.Login, config.Password); err != nil {
			return err
		}
	}
	return nil
}
