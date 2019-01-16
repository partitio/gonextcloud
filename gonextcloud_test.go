package gonextcloud

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
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
				us, err := c.Users().List()
				assert.NoError(t, err)
				assert.Contains(t, us, config.Login)
			},
		},

		{
			"existing user",
			func(t *testing.T) {
				u, err := c.Users().Get(config.Login)
				assert.NoError(t, err)
				assert.NotNil(t, u)
			},
		},

		{
			"empty user",
			func(t *testing.T) {
				u, err := c.Users().Get("")
				assert.Error(t, err)
				assert.Empty(t, u)
			},
		},

		{
			"TestNonExistingUser",
			func(t *testing.T) {
				_, err := c.Users().Get(config.NotExistingUser)
				assert.Error(t, err)
			},
		},

		{
			"TestUserSearch",
			func(t *testing.T) {
				us, err := c.Users().Search(config.Login)
				assert.NoError(t, err)
				assert.Contains(t, us, config.Login)
			},
		},

		{
			"TestUserCreate",
			func(t *testing.T) {
				err := c.Users().Create(config.NotExistingUser, password, nil)
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
		//		user := &types.Users{
		//			ID:          username,
		//			Displayname: strings.ToUpper(username),
		//			Email:       "some@address.com",
		//			Address:     "Main Street, City",
		//			Twitter:     "@me",
		//			Phone:       "42 42 242 424",
		//			Website:     "my.site.com",
		//		}
		//		err := c.Users().Create(username, password, user)
		//		assert.NoError(t, err)
		//		u, err := c.Users().Get(username)
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
		//		err = c.Users().Delete(u.ID)
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
		//		err := c.Users().Create(username, password, nil)
		//		assert.NoError(t, err)
		//		user := &types.UserDetails{
		//			ID:          username,
		//			Displayname: strings.ToUpper(username),
		//			Email:       "some@address.com",
		//			Address:     "Main Street, City",
		//			Twitter:     "@me",
		//			Phone:       "42 42 242 424",
		//			Website:     "my.site.com",
		//			Quota: types.Quota{
		//				// Unlimited
		//				Quota: -3,
		//			},
		//		}
		//		err = c.Users().Update(user)
		//		assert.NoError(t, err)
		//		u, err := c.Users().Get(username)
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
		//		err = c.Users().Delete(u.ID)
		//		assert.NoError(t, err)
		//	},
		//},
		{
			"TestUserCreateExisting",
			func(t *testing.T) {
				err := c.Users().Create(config.Login, password, nil)
				assert.Error(t, err)
			},
		},

		{
			"TestGroupList",
			func(t *testing.T) {
				gs, err := c.Groups().List()
				assert.NoError(t, err)
				assert.Contains(t, gs, "admin")
			},
		},

		{
			"TestGroupCreate",
			func(t *testing.T) {
				err := c.Groups().Create(config.NotExistingGroup)
				assert.NoError(t, err)
			},
		},

		{
			"TestUserUpdateEmail",
			func(t *testing.T) {
				email := "my@mail.com"
				err := c.Users().UpdateEmail(config.NotExistingUser, email)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().UpdateDisplayName(config.NotExistingUser, displayName)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().UpdatePhone(config.NotExistingUser, phone)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().UpdateAddress(config.NotExistingUser, address)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().UpdateWebSite(config.NotExistingUser, website)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().UpdateTwitter(config.NotExistingUser, twitter)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				quota := int64(1024 * 1024 * 1024)
				err := c.Users().UpdateQuota(config.NotExistingUser, quota)
				assert.NoError(t, err)
				// TODO : Find better verification : A never connected Users does not have quota available
				//u, err := c.Users(config.NotExistingUser)
				//assert.NoError(t, err)
				//assert.Equal(t, quota, u.Quota.Quota)
			},
		},
		{
			"TestUserUpdatePassword",
			func(t *testing.T) {
				password := "newcomplexpassword"
				err := c.Users().UpdatePassword(config.NotExistingUser, password)
				assert.NoError(t, err)
			}},
		{
			"TestUserGroupAdd",
			func(t *testing.T) {
				err := c.Users().GroupAdd(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				gs, err := c.Users().GroupList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Contains(t, gs, config.NotExistingGroup)
			},
		},
		{
			"TestUserGroupSubAdminList",
			func(t *testing.T) {
				gs, err := c.Users().GroupSubAdminList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Empty(t, gs)
			},
		},
		{
			"TestUserGroupPromote",
			func(t *testing.T) {
				err := c.Users().GroupPromote(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				gs, err := c.Users().GroupSubAdminList(config.NotExistingUser)
				assert.NoError(t, err)
				assert.Contains(t, gs, config.NotExistingGroup)
			},
		},
		{
			"TestUserGroupDemote",
			func(t *testing.T) {
				err := c.Users().GroupDemote(config.NotExistingUser, config.NotExistingGroup)
				assert.NoError(t, err)
				//gs, err := c.UserGroupSubAdminList(config.NotExistingUser)
				//assert.NoError(t, err)
				//assert.Empty(t, gs)
			},
		},
		{
			"TestUserDisable",
			func(t *testing.T) {
				err := c.Users().Disable(config.NotExistingUser)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Users().Enable(config.NotExistingUser)
				assert.NoError(t, err)
				u, err := c.Users().Get(config.NotExistingUser)
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
				err := c.Groups().Delete(config.NotExistingGroup)
				assert.NoError(t, err)
			},
		},
		{
			"TestUserDelete",
			func(t *testing.T) {
				err := c.Users().Delete(config.NotExistingUser)
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
				s, err := c.Shares().List()
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
)

func TestClient(t *testing.T) {
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	for _, tt := range provisionningTests {
		t.Run(tt.string, tt.test)
	}
}

func TestUserCreateWithoutPassword(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	if c.version.Major < 14 {
		t.SkipNow()
	}
	// Nextcloud does not seems to like recreating a deleted user
	err := c.Users().CreateWithoutPassword(config.NotExistingUser, config.Email, strings.Title(config.NotExistingUser), "100024", "en", "admin")
	assert.NoError(t, err)
	err = c.Users().Delete(config.NotExistingUser)
	assert.NoError(t, err)
}

func TestUserCreateBatchWithoutPassword(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	if c.version.Major < 14 {
		t.SkipNow()
	}
	var us []types.User
	for i := 0; i < 5; i++ {
		u := fmt.Sprintf(config.NotExistingUser+"_%d", i)
		us = append(us, types.User{
			Username:    u,
			DisplayName: strings.Title(u),
			Groups:      []string{"admin"},
			Email:       config.Email,
			Language:    "fr",
			Quota:       "100024",
		})
	}
	err := c.Users().CreateBatchWithoutPassword(us)
	assert.NoError(t, err)

	// Cleaning
	var wg sync.WaitGroup
	for _, u := range us {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			c.Users().Delete(n)
		}(u.Username)
	}
	wg.Wait()
}

func TestUserListDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	us, err := c.Users().ListDetails()
	assert.NoError(t, err)
	assert.Contains(t, us, config.Login)
}

func TestGroupListDetails(t *testing.T) {
	c = nil
	if err := initClient(); err != nil {
		t.Fatal(err)
	}
	gs, err := c.Groups().ListDetails()
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
	config.NotExistingUser = fmt.Sprintf("%s-%s", config.NotExistingUser, strconv.Itoa(rand.Int()))
	return nil
}

func initClient() error {
	if c == nil {
		rand.Seed(time.Now().Unix())
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
