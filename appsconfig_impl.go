package gonextcloud

import (
	"net/http"
	"sync"

	req "github.com/levigross/grequests"
)

//appsConfig contains all apps Configuration available actions
type appsConfig struct {
	c *client
}

//List lists all the available apps
func (a *appsConfig) List() (apps []string, err error) {
	res, err := a.c.baseOcsRequest(http.MethodGet, routes.appsConfig, nil)
	if err != nil {
		return nil, err
	}
	var r appConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//Keys returns the app's config keys
func (a *appsConfig) Keys(id string) (keys []string, err error) {
	res, err := a.c.baseOcsRequest(http.MethodGet, routes.appsConfig, nil, id)
	if err != nil {
		return nil, err
	}
	var r appConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//Value get the config value for the given app's key
func (a *appsConfig) Value(id, key string) (string, error) {
	res, err := a.c.baseOcsRequest(http.MethodGet, routes.appsConfig, nil, id, key)
	if err != nil {
		return "", err
	}
	var r appcConfigValueResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//SetValue set the config value for the given app's key
func (a *appsConfig) SetValue(id, key, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"value": value,
		},
	}
	_, err := a.c.baseOcsRequest(http.MethodPost, routes.appsConfig, ro, id, key)
	return err
}

//DeleteValue delete the config value and (!! be careful !!) the key
func (a *appsConfig) DeleteValue(id, key, value string) error {
	_, err := a.c.baseOcsRequest(http.MethodDelete, routes.appsConfig, nil, id, key)
	return err
}

//Get returns all apps AppConfigDetails
func (a *appsConfig) Get() (map[string]map[string]string, error) {
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
func (a *appsConfig) Details(appID string) (map[string]string, error) {
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
