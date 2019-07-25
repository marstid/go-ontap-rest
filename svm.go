package ontap

import (
	"encoding/json"
	"fmt"
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

	return "", fmt.Errorf("0 - Storage VM with name %s not found", name)
}

// Return a list of svm objects
func (c *Client) GetStorageVM(uuid string) (svm StorageVM, err error) {
	uri := "/api/svm/svms/" + uuid

	data, err := c.clientGet(uri)
	if err != nil {
		return svm, err
	}

	var resp StorageVM
	json.Unmarshal(data, &resp)
	if err != nil {
		return svm, err
	}

	return svm, nil
}

// Return a list of svm objects
func (c *Client) GetStorageVMsShort() (svms []Svm, err error) {

	uri := "/api/svm/svms?order_by=name"

	data, err := c.clientGet(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return svms, err
	}

	var resp SvmResponse
	json.Unmarshal(data, &resp)
	if err != nil {
		return svms, err
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


	for createJob.State == "running" {
		time.Sleep(time.Second)
		createJob, err = c.GetJob(url)
	}

	if createJob.State == "failure" {
		return fmt.Errorf("%d - %s", createJob.Code, createJob.Message)
	}

	return nil
}

func (c *Client) DeleteStorageVM(uuid string) (err error){
	uri := "/api/svm/svms/" + uuid
	//fmt.Println(uri)
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