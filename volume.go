package ontap

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"strings"
	"time"
)


func init(){
	ConfigLog()
}

func (c *Client) GetVolume(uuid string) (vol Volume, err error){
	uri := "/api/storage/volumes/" + uuid
	data, err := c.clientGet(uri)
	if err != nil {
		if strings.Contains(err.Error(),"Error-4"){
			return vol, &apiError{4, fmt.Sprintf("Volume with UUID: %s not found", uuid)}
		}
		return vol, &apiError{1, err.Error()}
	}

	var result Volume
	err = json.Unmarshal(data, &result)
	if err != nil {
		return vol, &apiError{2, err.Error()}
	}

	return result, nil
}
// Size in GB

func (c *Client) CreateVolume(name, comment, svm string, aggr []string, size int ) ( err error){
	uri := "/api/storage/volumes"

	var payload map[string]interface{}
	payload = make( map[string]interface{})

	payload["name"] = name
	payload["svm.name"] = svm
	payload["size"] = size * 1024 * 1024 * 1024
	payload["comment"] = comment

	array := make([]map[string]interface{},len(aggr))
	for k, v := range aggr {
		array[k] = map[string]interface{}{
			`name`: v,
		}
	}

	payload["aggregates"] = array

	jsonPayload, _ := json.Marshal(payload)
	data, err := c.clientPost(uri, jsonPayload)
	if err != nil {
		log.Error(err.Error())
		return  err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return &apiError{2, err.Error()}
	}

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
		return  fmt.Errorf("%d - %s", createJob.Code, createJob.Message)
	}

	return nil
}

func (c *Client) DeleteVolume(uuid string) (err error){

	uri := "/api/storage/volumes/" + uuid
	data, err := c.clientDelete(uri)
	if err != nil {
		if strings.Contains(err.Error(),"Error-4"){
			return &apiError{4, fmt.Sprintf("Volume with UUID: %s not found", uuid)}
		}
		return &apiError{1, err.Error()}
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return &apiError{2, err.Error()}
	}

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
		return &apiError{int64(deleteJob.Code), deleteJob.Message}
		//return fmt.Errorf("%d - %s", deleteJob.Code, deleteJob.Message)
	}

	return nil
}

// Size in GB

func (c *Client) UpdateVolume(uuid, name, comment string, size int) (err error){

	uri := "/api/storage/volumes/" + uuid
	var payload map[string]interface{}
	payload = make( map[string]interface{})

	payload["name"] = name
	payload["size"] = size * 1024 * 1024 * 1024
	payload["comment"] = comment

	jsonPayload, _ := json.Marshal(payload)
	data, err := c.clientPatch(uri, jsonPayload)
	if err != nil {
		if strings.Contains(err.Error(),"Error-4"){
			return &apiError{4, fmt.Sprintf("Volume with UUID: %s not found", uuid)}
		}
		return &apiError{1, err.Error()}
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return &apiError{2, err.Error()}
	}

	job := result["job"].(map[string]interface{})
	link := job["_links"].(map[string]interface{})
	href := link["self"].(map[string]interface{})
	url := href["href"].(string)

	updateJob, err := c.GetJob(url)


	for updateJob.State == "running" {
		time.Sleep(time.Second)
		updateJob, err = c.GetJob(url)
	}

	if updateJob.State == "failure" {
		return  fmt.Errorf("%d - %s", updateJob.Code, updateJob.Message)
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
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", &apiError{2, err.Error()}
	}

	records := result["records"].([]interface{})
	for _, v := range records {
		rec := v.(map[string]interface{})
		if rec["name"] == name {
			return rec["uuid"].(string), nil
		}
	}

	return "", &apiError{4, fmt.Sprintf("Volume with name \"%s\" not found", name)}
}
