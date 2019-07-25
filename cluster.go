package ontap

import (
	"encoding/json"
	"fmt"
)



func (c *Client) GetCluster() (cluster Cluster, err error) {
	// Todo check error handling faulty url
	uri := "/api/cluster"
	//uri := "/api/foobar"
	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return cluster, err
	}

	err = json.Unmarshal(data, &cluster)
	if err != nil {
		return cluster, err
	}

	return cluster, err

}
