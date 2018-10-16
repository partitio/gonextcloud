package gonextcloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Config struct {
	URL              string `yaml:"url"`
	Login            string `yaml:"login"`
	Password         string `yaml:"password"`
	AppName          string `yaml:"app-name"`
	ShareFolder      string `yaml:"share-folder"`
	NotExistingUser  string `yaml:"not-existing-user"`
	NotExistingGroup string `yaml:"not-existing-group"`
	Email            string `yaml:"email"`
}

const password = "somecomplicatedpassword"

type test = func(t *testing.T)

var (
	config             = Config{}
	c                  *Client
	groupID            = 37
	provisionningTests = []struct {
		string
		test
	}{
		{
			"TestLoadConfig",
			func(t *testing.T) {
				err := LoadConfig()
				assert.NoError(t, err)
			},
		},
		{
			"create client",
			func(t *testing.T) {
				var err error
				c, err = NewClient(config.URL)
				assert.NoError(t, err, "aie")
			},
		},

		{
			"login failed",
			func(t *testing.T) {
				err := c.Login("", "")
				assert.Error(t, err)
			},
		},

		{
			"login",
			func(t *testing.T) {
				err := c.Login(config.Login, config.Password)
				assert.NoError(t, err)
			},
		},

		{
			"user list",
			func(t *testing.T) {
				us, err := c.UserList()
				assert.NoError(t, err)

				assert.Contains(t, us, config.Login)
			},
		},

		{
			"existing user",
			func(t *testing.T) {
				u, err := c.User(config.Login)
				assert.NoError(t, err)
				assert.NotNil(t, u)
			},
		},

		{
			"empty user",
			func(t *testing.T) {
				u, err := c.User("")
				assert.Error(t, err)
				assert.Empty(t, u)
			},
		},

		{
			"TestNonExistingUser",
			func(t *testing.T) {
				_, err := c.User(config.NotExistingUser)
				assert.Error(t, err)
			},
		},

		{
			"TestUserSearch",
			func(t *testing.T) {
				us, err := c.UserSearch(config.Login)
				assert.NoError(t, err)
				assert.Contains(t, us, config.Login)
			},
		},

		{
			"TestUserCreate",
			func(t *testing.T) {
				err := c.UserCreate(config.NotExistingUser, password, nil)
				assert.NoError(t, err)
			},
		},
		//{
		//	"TestUserCreateFull",
		//	func(t *testing.T) {
		//		if err := initClient(); err != nil {
		//			return
		//		}
		//		username := fmt.Sprintf("%s-2", config.NotExistingUser)
		//		user := &types.User{
		//			ID:          username,
		//			Displayname: strings.ToUpper(username),
		//			Email:       "some@address.com",
		//			Address:     "Main Street, City",
		//			Twitter:     "@me",
		//			Phone:       "42 42 242 424",
		//			Website:     "my.site.com",
		//		}
		//		err := c.UserCreate(username, password, user)
		//		assert.NoError(t, err)
		//		u, err := c.User(username)
		//		assert.NoError(t, err)
		//		o := structs.Map(user)
		//		r := structs.Map(u)
		//		for k := range o {
		//			if ignoredUserField(k) {
		//				continue
		//			}
		//			assert.Equal(t, o[k], r[k])
		//		}
		//		// Clean up
		//		err = c.UserDelete(u.ID)
		//		assert.NoError(t, err)
		//	},
		//},
		//
		//{
		//	"TestUserUpdate",
		//	func(t *testing.T) {
		//		if err := initClient(); err != nil {
		//			return
		//		}
		//		username := fmt.Sprintf("%s-2", config.NotExistingUser)
		//		err := c.UserCreate(username, password, nil)
		//		assert.NoError(t, err)
		//		user := &types.User{
		//			ID:          username,
		//			Displayname: strings.ToUpper(username),
		//			Email:       "some@address.com",
		//			Address:     "Main Street, City",
		//			Twitter:     "@me",
		//			Phone:       "42 42 242 424",
		//			Website:     "my.site.com",
		//		}
		//		err = c.UserUpdate(user)
		//		assert.NoError(t, err)
		//		u, err := c.User(username)
		//		assert.NoError(t, err)
		//		o := structs.Map(user)
		//		r := structs.Map(u)
		//		for k := range o {
		//			if ignoredUserField(k) {
		//				continue
		//			}
		//			assert.Equal(t, o[k], r[k])
		//		}
		//		// Clean up
		//		err = c.UserDelete(u.ID)
		//		assert.NoError(t, err)
		//	},
		//},
		{
			"TestUserCreateExisting",
			func(t *testing.T) {
				err := c.UserCreate(config.NotExistingUser, password, nil)
				assert.Error(t, err)
			},
		},

		{
			"TestGroupList",
			func(t *testing.T) {
				gs, err := c.GroupList()
				assert.NoError(t, err)
				assert.Contains(t, gs, "admin")
			},
		},

		{
			"TestGroupCreate",
			func(t *testing.T) {
				err := c.GroupCreate(config.NotExistingGroup)
				assert.NoError(t, err)
			},
		},

		{
			"TestUserUpdateEmail",
			func(t *testing.T) {
				email := "my@mail.com"
				err := c.UserUpdateEmail(config.NotExistingUser, email)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, email, u.Email)
			},
		},

		{
			"TestUserUpdateDisplayName",
			func(t *testing.T) {
				displayName := "Display Name"
				err := c.UserUpdateDisplayName(config.NotExistingUser, displayName)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, displayName, u.Displayname)
			},
		},

		{
			"TestUserUpdatePhone",
			func(t *testing.T) {
				phone := "+33 42 42 42 42"
				err := c.UserUpdatePhone(config.NotExistingUser, phone)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, phone, u.Phone)
			},
		},

		{
			"TestUserUpdateAddress",
			func(t *testing.T) {
				address := "Main Street, Galifrey"
				err := c.UserUpdateAddress(config.NotExistingUser, address)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, address, u.Address)
			},
		},

		{
			"TestUserUpdateWebSite",
			func(t *testing.T) {
				website := "www.doctor.who"
				err := c.UserUpdateWebSite(config.NotExistingUser, website)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, website, u.Website)
			},
		},
		{
			"TestUserUpdateTwitter",
			func(t *testing.T) {
				twitter := "@doctorwho"
				err := c.UserUpdateTwitter(config.NotExistingUser, twitter)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.Equal(t, twitter, u.Twitter)
			},
		},
		{
			"TestUserUpdateQuota",
			func(t *testing.T) {
				quota := 1024 * 1024 * 1024
				err := c.UserUpdateQuota(config.NotExistingUser, quota)
				assert.NoError(t, err)
				// TODO : Find better verification : A never connected User does not have quota available
				//u, err := c.User(config.NotExistingUser)
				//assert.NoError(t, err)
				//assert.Equal(t, quota, u.Quota.Quota)
			},
		},
		{
			"TestUserUpdatePassword",
			func(t *testing.T) {
				password := "newcomplexpassword"
				err := c.UserUpdatePassword(config.NotExistingUser, password)
				assert.NoError(t, err)
			}},
		{
			"TestUserGroupAdd",
			func(t *testing.T) {
				err := c.UserGroupAdd(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				gs, err := c.UserGroupList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Contains(t, gs, config.NotExistingGroup)
			},
		},
		{
			"TestUserGroupSubAdminList",
			func(t *testing.T) {
				gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Empty(t, gs)
			},
		},
		{
			"TestUserGroupPromote",
			func(t *testing.T) {
				err := c.UserGroupPromote(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Contains(t, gs, config.NotExistingGroup)
			},
		},
		{
			"TestUserGroupDemote",
			func(t *testing.T) {
				err := c.UserGroupDemote(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				//gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
				//assert.NoError(t, err)
				//assert.Empty(t, gs)
			},
		},
		{
			"TestUserDisable",
			func(t *testing.T) {
				err := c.UserDisable(config.NotExistingUser)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.False(t, u.Enabled)
			},
		},
		{
			"TestUserEnable",
			func(t *testing.T) {
				err := c.UserEnable(config.NotExistingUser)
				assert.NoError(t, err)
				u, err := c.User(config.NotExistingUser)
				assert.NoError(t, err)
				if err != nil {
					t.Fail()
					return
				}
				assert.True(t, u.Enabled)
			},
		},
		{
			"TestGroupDelete",
			func(t *testing.T) {
				err := c.GroupDelete(config.NotExistingGroup)
				assert.NoError(t, err)
			},
		},
		{
			"TestUserDelete",
			func(t *testing.T) {
				err := c.UserDelete(config.NotExistingUser)
				assert.NoError(t, err)
			},
		},
		{
			"TestInvalidBaseRequest",
			func(t *testing.T) {
				c.baseURL = &url.URL{}
				_, err := c.baseRequest(http.MethodGet, routes.capabilities, nil, "admin", "invalid")
				c = nil
				assert.Error(t, err)
			},
		},
		{
			"TestShareList",
			func(t *testing.T) {
				if err := initClient(); err != nil {
					return
				}
				s, err := c.SharesList()
				assert.NoError(t, err)
				assert.NotNil(t, s)
			},
		},
		{
			"TestLogout",
			func(t *testing.T) {
				err := c.Logout()
				assert.NoError(t, err)
				assert.Nil(t, c.session.HTTPClient.Jar)
			},
		},
		{
			"TestLoggedIn",
			func(t *testing.T) {
				c := &Client{}
				c.capabilities = &types.Capabilities{}
				assert.False(t, c.loggedIn())
			},
		},
		{
			"TestLoginInvalidURL",
			func(t *testing.T) {
				c, _ = NewClient("")
				err := c.Login("", "")
				assert.Error(t, err)
			},
		},
		{
			"TestBaseRequest",
			func(t *testing.T) {
				c, _ = NewClient("")
				_, err := c.baseRequest(http.MethodGet, routes.capabilities, nil, "admin", "invalid")
				assert.Error(t, err)
			},
		},
	}

	groupFoldersTests = []struct {
		string
		test
	}{
		{
			"TestGroupFoldersCreate",
			func(t *testing.T) {
				// Recreate client
				var err error
				groupID, err = c.GroupFoldersCreate("API")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersList",
			func(t *testing.T) {
				gfs, err := c.GroupFoldersList()
				assert.NoError(t, err)
				assert.NotNil(t, gfs[groupID])
			},
		},
		{
			"TestGroupFolders",
			func(t *testing.T) {
				gf, err := c.GroupFolders(groupID)
				assert.NoError(t, err)
				assert.NotNil(t, gf)
			},
		},
		{
			"TestGroupFolderRename",
			func(t *testing.T) {
				err := c.GroupFoldersRename(groupID, "API_Renamed")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersAddGroup",
			func(t *testing.T) {
				err := c.GroupFoldersAddGroup(groupID, "admin")
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersSetGroupPermissions",
			func(t *testing.T) {
				err := c.GroupFoldersSetGroupPermissions(groupID, "admin", types.ReadPermission)
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFoldersSetQuota",
			func(t *testing.T) {
				err := c.GroupFoldersSetQuota(groupID, 100)
				assert.NoError(t, err)
			},
		},
		{
			"TestGroupFolderRemoveGroup",
			func(t *testing.T) {
				err := c.GroupFoldersRemoveGroup(groupID, "admin")
				assert.NoError(t, err)
			},
		},
	}
)

func TestClient(t *testing.T) {
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	for _, tt := range provisionningTests {
		t.Run(tt.string, tt.test)
	}
}

func TestGroupFolders(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	for _, tt := range groupFoldersTests {
		t.Run(tt.string, tt.test)
	}
}

func TestUserCreateWithoutPassword(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	// Nextcloud does not seems to like recreating a deleted user
	rand.Seed(time.Now().Unix())
	n := fmt.Sprintf("%s-%s", config.NotExistingUser, strconv.Itoa(rand.Int()))
	err := c.UserCreateWithoutPassword(n, config.Email, strings.Title(config.NotExistingUser))
	assert.NoError(t, err)
	err = c.UserDelete(n)
	assert.NoError(t, err)
}

func TestUserListDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	us, err := c.UserListDetails()
	assert.NoError(t, err)
	assert.Contains(t, us, config.Login)
}

func TestGroupListDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	gs, err := c.GroupListDetails()
	assert.NoError(t, err)
	assert.NotEmpty(t, gs)
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
	if err := yaml.Unmarshal(b, &config); err != nil {
		return err
	}
	// Override with env variables
	u := os.Getenv("NEXTCLOUD_URL")
	if u != "" {
		config.URL = u
	}
	p := os.Getenv("NEXTCLOUD_PASSWORD")
	if p != "" {
		config.Password = p
	}
	e := os.Getenv("NEXTCLOUD_EMAIL")
	if e != "" {
		config.Email = e
	}
	return nil
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
