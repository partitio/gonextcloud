package gonextcloud

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
)

func TestUserUpdate(t *testing.T) {
	if err := initClient(); err != nil {
		return
	}
	username := fmt.Sprintf("%s-2", config.NotExistingUser)
	err := c.Users().Create(username, password, nil)
	if err != nil {
		t.FailNow()
	}
	err = c.Groups().Create(config.NotExistingGroup)
	if err != nil {
		t.FailNow()
	}
	user := &UserDetails{
		ID:          username,
		Displayname: strings.ToUpper(username),
		Email:       "some@mail.com",
		Quota: Quota{
			// Unlimited
			Quota: -3,
		},
		Groups: []string{config.NotExistingGroup},
	}
	s := time.Now()
	err = c.Users().Update(user)
	e := time.Now().Sub(s)
	fmt.Println(e.String())
	assert.NoError(t, err)
	u, err := c.Users().Get(username)
	assert.NoError(t, err)
	o := structs.Map(user)
	r := structs.Map(u)
	for k := range o {
		if ignoredUserField(k) {
			continue
		}
		assert.Equal(t, o[k], r[k])
	}
	// Clean up
	err = c.Users().Delete(username)
	assert.NoError(t, err)
	err = c.Groups().Delete(config.NotExistingGroup)
	assert.NoError(t, err)

}
