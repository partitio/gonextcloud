package gonextcloud

import (
	req "github.com/levigross/grequests"
	"gitlab.adphi.fr/partitio/Nextcloud-Partitio/gonextcloud/types"
	"net/http"
	"sync"
)

//AppsConfigList lists all the available apps
func (c *Client) AppsConfigList() (apps []string, err error) {
	res, err := c.baseRequest(http.MethodGet, routes.appsConfig, nil)
	if err != nil {
		return nil, err
	}
	var r types.AppConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//AppsConfigKeys returns the app's config keys
func (c *Client) AppsConfigKeys(id string) (keys []string, err error) {
	res, err := c.baseRequest(http.MethodGet, routes.appsConfig, nil, id)
	if err != nil {
		return nil, err
	}
	var r types.AppConfigResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//AppsConfigValue get the config value for the given app's key
func (c *Client) AppsConfigValue(id, key string) (string, error) {
	res, err := c.baseRequest(http.MethodGet, routes.appsConfig, nil, id, key)
	if err != nil {
		return "", err
	}
	var r types.AppcConfigValueResponse
	res.JSON(&r)
	return r.Ocs.Data.Data, nil
}

//AppsConfigSetValue set the config value for the given app's key
func (c *Client) AppsConfigSetValue(id, key, value string) error {
	ro := &req.RequestOptions{
		Data: map[string]string{
			"value": value,
		},
	}
	_, err := c.baseRequest(http.MethodPost, routes.appsConfig, ro, id, key)
	return err
}

//AppsConfigDeleteValue delete the config value and (!! be careful !!) the key
func (c *Client) AppsConfigDeleteValue(id, key, value string) error {
	_, err := c.baseRequest(http.MethodDelete, routes.appsConfig, nil, id, key)
	return err
}

//AppsConfig returns all apps AppConfigDetails
func (c *Client) AppsConfig() (map[string]map[string]string, error) {
	config := map[string]map[string]string{}
	m := sync.Mutex{}
	appsIDs, err := c.AppsConfigList()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(appsIDs))
	for i := range appsIDs {
		go func(id string) {
			defer wg.Done()
			d, err := c.AppsConfigDetails(id)
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

//AppsConfigDetails returns all the config's key, values pair of the app
func (c *Client) AppsConfigDetails(appID string) (map[string]string, error) {
	config := map[string]string{}
	m := sync.Mutex{}
	var err error
	var ks []string
	ks, err = c.AppsConfigKeys(appID)
	if err != nil {
		return config, err
	}
	var wg sync.WaitGroup
	wg.Add(len(ks))
	for i := range ks {
		go func(key string) {
			defer wg.Done()
			v, err := c.AppsConfigValue(appID, key)
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
