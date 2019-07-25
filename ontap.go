package ontap

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	UserName string
	Password string
	Host     string
	Debug    bool
	TrustSSL bool
	TimeOut  time.Duration
}

func NewClient(user, password, host string, debug, ssl bool) (client *Client, error error) {

	return &Client{
		UserName: user,
		Password: password,
		Host:     host,
		Debug:    debug,
		TrustSSL: ssl,
		TimeOut:  3,
	}, error
}

func (c *Client) clientGet(uri string) (data []byte, err error) {

	url := "https://" + c.Host + uri

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) clientPost(uri string, json []byte) (data []byte, err error) {

	url := "https://" + c.Host + uri

	payload := bytes.NewReader(json)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) clientDelete(uri string) (data []byte, err error) {

	url := "https://" + c.Host + uri

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.UserName, c.Password)
	req.Header.Set("Content-Type", "application/hal+json")
	req.Header.Set("UserAgent", "go-ontap-rest")

	httpClient := &http.Client{
		Timeout: time.Second * c.TimeOut,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.TrustSSL,
			},
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {

		var jsonError string
		if strings.Contains(string(body), "message") {
			var jec HttpError
			json.Unmarshal(body, &jec)
			jsonError = jsonError + jec.Error.Code + " - API error - " + jec.Error.Message
			fmt.Println(string(body))
			return nil, fmt.Errorf("%s", jsonError)
		}
		return nil, fmt.Errorf("%s", body)

	}

	return body, nil

}
