package ontap

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Creates Aggregates

func (c *Client) CreateAggregate(node, name string, disks int) (err error) {
	uri := "/api/storage/aggregates"

	payload := AggrCreateData{}
	payload.Node.Name = node
	payload.Name = name
	payload.BlockStorage.Primary.DiskCount = strconv.Itoa(disks)

	jsonPayload, _ := json.Marshal(payload)
	data, err := c.clientPost(uri, jsonPayload)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	job := result["job"].(map[string]interface{})
	link := job["_links"].(map[string]interface{})
	href := link["self"].(map[string]interface{})
	url := href["href"].(string)

	createJob, err := c.GetJob(url)

	//count := 1
	for createJob.State == "running" {
		time.Sleep(time.Second)
		createJob, err = c.GetJob(url)
		//fmt.Println(count)
		//count++
	}

	if createJob.State == "failure" {
		return fmt.Errorf("%d - %s", createJob.Code, createJob.Message)
	}

	return nil

}

func (c *Client) DeleteAggregate(uuid string) (err error) {
	uri := "/api/storage/aggregates/" + uuid
	data, err := c.clientDelete(uri)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	job := result["job"].(map[string]interface{})
	link := job["_links"].(map[string]interface{})
	href := link["self"].(map[string]interface{})
	url := href["href"].(string)

	deleteJob, err := c.GetJob(url)

	for deleteJob.State == "running" {
		time.Sleep(time.Second)
		deleteJob, err = c.GetJob(url)
	}

	if deleteJob.State == "failure" {
		return fmt.Errorf("%d - %s", deleteJob.Code, deleteJob.Message)
		//return fmt.Errorf("Delete failed: %s", deleteJob.Message)
	}

	return nil
}

func (c *Client) GetAggregate(uuid string) (aggr AggrRecord, err error) {
	uri := "/api/storage/aggregates/" + uuid
	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return aggr, err
	}

	var resp AggrRecord
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return aggr, err
	}

	return aggr, err

}

func (c *Client) GetAggregates() (aggrs []AggrRecord, err error) {
	uri := "/api/storage/aggregates"
	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	var resp AggrResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	for _, v := range resp.Records {
		r, err := c.GetAggregate(v.UUID)
		if err != nil {
			return nil, err
		}
		aggrs = append(aggrs, r)

	}

	return

}

func (c *Client) GetAggregateUUID(name string) (uuid string, err error) {
	uri := "/api/storage/aggregates?name=" + name
	data, err := c.clientGet(uri)
	if err != nil {
		return "", err
	}

	//fmt.Println(string(data))
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	records := result["records"].([]interface{})
	for _, v := range records {
		rec := v.(map[string]interface{})
		if rec["name"] == name {
			return rec["uuid"].(string), nil
		}

	}

	return "", fmt.Errorf("0 - Aggregate with name %s not found", name)

}
