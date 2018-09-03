package gonextcloud

import (
	"github.com/partitio/gonextcloud/types"
	"net/http"
)

//Monitoring return nextcloud monitoring statistics
func (c *Client) Monitoring() (*types.Monitoring, error) {
	res, err := c.baseRequest(http.MethodGet, routes.monitor, nil)
	if err != nil {
		return nil, err
	}
	var m types.MonitoringResponse
	res.JSON(&m)
	return &m.Ocs.Data, nil
}
