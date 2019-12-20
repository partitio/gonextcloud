package gonextcloud

import (
	"net/http"
)

//Monitoring return nextcloud monitoring statistics
func (c *client) Monitoring() (*Monitoring, error) {
	res, err := c.baseOcsRequest(http.MethodGet, routes.monitor, nil)
	if err != nil {
		return nil, err
	}
	var m monitoringResponse
	res.JSON(&m)
	return &m.Ocs.Data, nil
}
