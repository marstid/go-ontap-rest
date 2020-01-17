package ontap

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Return svm uuid from name
func (c *Client) GetStorageVmUUIDByName(name string) (uuid string, err error) {
	uri := "/api/svm/svms?name=" + name
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

	//return "", fmt.Errorf("0 - Storage VM with name %s not found", name)
	return "", &apiError{0,"0 - Storage VM with name " + name + "%s not found"}
}

// Return a list of svm objects
func (c *Client) GetStorageVM(uuid string) (svm StorageVM, err error) {
	uri := "/api/svm/svms/" + uuid

	data, err := c.clientGet(uri)
	if err != nil {
		return svm,  &apiError{1, err.Error()}
	}

	var resp StorageVM
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return svm, &apiError{2, err.Error()}
	}

	return svm, nil
}

// Return a list of svm objects
func (c *Client) GetStorageVMsShort() (svms []Svm, err error) {

	uri := "/api/svm/svms?order_by=name"

	data, err := c.clientGet(uri)
	if err != nil {
		//fmt.Println("Error: " + err.Error())
		return svms, &apiError{1, err.Error()}
	}

	var resp SvmResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return svms, &apiError{2, err.Error()}
	}

	for _, v := range resp.Records {
		nvm := Svm{Name: v.Name, UUID: v.UUID, SelfLink: v.Links.Self.Href}
		svms = append(svms, nvm)
	}

	return svms, err

}

func (c *Client) CreateStorageVM(name, comment, ipSpaceUUID string) (err error){
	uri := "/api/svm/svms"

	var payload map[string]interface{}
	payload = make( map[string]interface{})
	payload["name"] = name
	payload["comment"] = comment
	if ipSpaceUUID != "" {
		payload["ipspace.uuid"] = ipSpaceUUID
	}
	jsonPayload, _ := json.Marshal(payload)
	data, err := c.clientPost(uri, jsonPayload)
	if err != nil {
		//fmt.Println("Error: " + err.Error())
		return &apiError{1, err.Error()}
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
		return &apiError{int64(createJob.Code),  createJob.Message}
		//return fmt.Errorf("%d - %s", createJob.Code, createJob.Message)
	}

	return nil
}

func (c *Client) DeleteStorageVM(uuid string) (err error){
	uri := "/api/svm/svms/" + uuid

	data, err := c.clientDelete(uri)
	if err != nil {
		if strings.Contains(err.Error(),"Error-4"){
			return &apiError{4, fmt.Sprintf("SVM with UUID \"%s\" not found", uuid)}
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
		return &apiError{int64(deleteJob.Code),  deleteJob.Message}
	}

	return nil
}