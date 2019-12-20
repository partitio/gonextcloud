package gonextcloud

import (
	"net/http"
	"strconv"
	"time"
)

//passwords contains some passwords app available actions
type passwords struct {
	c *client
}

func (p *passwords) List() ([]Password, error) {
	res, err := p.c.baseRequest(http.MethodGet, routes.passwords, nil)
	if err != nil {
		return nil, err
	}
	var r []Password
	err = res.JSON(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type passwordTime struct {
	time.Time
}

func (pt *passwordTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	t := time.Unix(i, 0)
	pt.Time = t
	return
}

type Password struct {
	Id string `json:"id"`
	Label string `json:"label"`
	Username string `json:"username"`
	Password string `json:"password"`
	Url string `json:"url"`
	Notes string `json:"notes"`
	Status int `json:"status"`
	StatusCode string `json:"statusCode"`
	Hash string `json:"hash"`
	Folder string `json:"foler"`
	Revision string `json:"revision"`
	Share string `json:"share"`
	CseType string `json:"cseType"`
	SseType string `json:"ssetype"`
	Favorite bool `json:"favorite"`
	Editable bool `json:"editable"`
	Edited passwordTime `json:"edited"`
	Created passwordTime `json:"created"`
	Updated passwordTime `json:"updated"`
}
