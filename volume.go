package ontap

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"time"
)


func init(){
	ConfigLog()
}

func (c *Client) GetVolume(uuid string) (err error){

	return
}

func (c *Client) CreateVolume(name, comment, svm string, aggr []string, size int ) (err error){
	uri := "/api/storage/volumes"

	var payload map[string]interface{}
	payload = make( map[string]interface{})

	payload["name"] = name
	payload["svm.name"] = svm
	payload["size"] = size
	payload["comment"] = comment

	array := make([]map[string]interface{},len(aggr))
	for k, v := range aggr {
		array[k] = map[string]interface{}{
			`name`: v,
		}
	}

	//ag, _ := json.Marshal(payload)
	//fmt.Println(string(ag))
	payload["aggregates"] = array


	jsonPayload, _ := json.Marshal(payload)
	data, err := c.clientPost(uri, jsonPayload)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	job := result["job"].(map[string]interface{})
	link := job["_links"].(map[string]interface{})
	href := link["self"].(map[string]interface{})
	url := href["href"].(string)

	createJob, err := c.GetJob(url)


	for createJob.State == "running" {
		time.Sleep(time.Second)
		createJob, err = c.GetJob(url)
	}

	if createJob.State == "failure" {
		return fmt.Errorf("%d - %s", createJob.Code, createJob.Message)
	}

	return nil
}

func (c *Client) DeleteVolume(uuid string) (err error){

	uri := "/api/storage/volumes/" + uuid
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
	}

	return nil
}


func (c *Client) GetVolumeUUID(name string) (uuid string, err error) {
	uri := "/api/storage/volumes?name=" + name
	data, err := c.clientGet(uri)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal(data, &result)

	records := result["records"].([]interface{})
	for _, v := range records {
		rec := v.(map[string]interface{})
		if rec["name"] == name {
			return rec["uuid"].(string), nil
		}

	}

	return "", fmt.Errorf("0 - Volume with name %s not found", name)

}
