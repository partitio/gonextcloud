package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"sync"
)

//AppsConfigI available methods
type AppsConfigI interface {
	List() (apps []string, err error)
	Keys(id string) (keys []string, err error)
	Value(id, key string) (string, error)
	SetValue(id, key, value string) error
	DeleteValue(id, key, value string) error
	Get() (map[string]map[string]string, error)
	Details(appID string) (map[string]string, error)
}

//AppsConfig contains all Apps Configuration available actions
type AppsConfig struct {
	c *Client
}

//List lists all the available apps
func (a *AppsConfig) List() (apps []string, err error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.appsConfig, nil)
	if err != nil {
		return nil, err
	}
	var r types.AppConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//Keys returns the app's config keys
func (a *AppsConfig) Keys(id string) (keys []string, err error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.appsConfig, nil, id)
	if err != nil {
		return nil, err
	}
	var r types.AppConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//Value get the config value for the given app's key
func (a *AppsConfig) Value(id, key string) (string, error) {
	res, err := a.c.baseRequest(http.MethodGet, routes.appsConfig, nil, id, key)
	if err != nil {
		return "", err
	}
	var r types.AppcConfigValueResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//SetValue set the config value for the given app's key
func (a *AppsConfig) SetValue(id, key, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"value": value,
		},
	}
	_, err := a.c.baseRequest(http.MethodPost, routes.appsConfig, ro, id, key)
	return err
}

//DeleteValue delete the config value and (!! be careful !!) the key
func (a *AppsConfig) DeleteValue(id, key, value string) error {
	_, err := a.c.baseRequest(http.MethodDelete, routes.appsConfig, nil, id, key)
	return err
}

//Get returns all apps AppConfigDetails
func (a *AppsConfig) Get() (map[string]map[string]string, error) {
	config := map[string]map[string]string{}
	m := sync.Mutex{}
	appsIDs, err := a.List()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(appsIDs))
	for i := range appsIDs {
		go func(id string) {
			defer wg.Done()
			d, err := a.Details(id)
			if err == nil {
				m.Lock()
				config[id] = d
				m.Unlock()
			}
		}(appsIDs[i])
	}
	wg.Wait()
	return config, err
}

//Details returns all the config's key, values pair of the app
func (a *AppsConfig) Details(appID string) (map[string]string, error) {
	config := map[string]string{}
	m := sync.Mutex{}
	var err error
	var ks []string
	ks, err = a.Keys(appID)
	if err != nil {
		return config, err
	}
	var wg sync.WaitGroup
	wg.Add(len(ks))
	for i := range ks {
		go func(key string) {
			defer wg.Done()
			v, err := a.Value(appID, key)
			if err == nil {
				m.Lock()
				config[key] = v
				m.Unlock()
			}
		}(ks[i])
	}
	wg.Wait()
	return config, err
}
