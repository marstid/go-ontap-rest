package ontap

import "time"

// Generated - https://mholt.github.io/json-to-go/
type Cluster struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Version struct {
		Full       string `json:"full"`
		Generation int    `json:"generation"`
		Major      int    `json:"major"`
		Minor      int    `json:"minor"`
	} `json:"version"`
	ManagementInterfaces []struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		IP   struct {
			Address string `json:"address"`
		} `json:"ip"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"management_interfaces"`
	Metric struct {
		Timestamp time.Time `json:"timestamp"`
		Duration  string    `json:"duration"`
		Status    string    `json:"status"`
		Latency   struct {
			Other int `json:"other"`
			Total int `json:"total"`
			Read  int `json:"read"`
			Write int `json:"write"`
		} `json:"latency"`
		Iops struct {
			Read  int `json:"read"`
			Write int `json:"write"`
			Other int `json:"other"`
			Total int `json:"total"`
		} `json:"iops"`
		Throughput struct {
			Read  int `json:"read"`
			Write int `json:"write"`
			Other int `json:"other"`
			Total int `json:"total"`
		} `json:"throughput"`
	} `json:"metric"`
	Statistics struct {
		Timestamp  time.Time `json:"timestamp"`
		Status     string    `json:"status"`
		LatencyRaw struct {
			Other int `json:"other"`
			Total int `json:"total"`
			Read  int `json:"read"`
			Write int `json:"write"`
		} `json:"latency_raw"`
		IopsRaw struct {
			Read  int `json:"read"`
			Write int `json:"write"`
			Other int `json:"other"`
			Total int `json:"total"`
		} `json:"iops_raw"`
		ThroughputRaw struct {
			Read  int `json:"read"`
			Write int `json:"write"`
			Other int `json:"other"`
			Total int `json:"total"`
		} `json:"throughput_raw"`
	} `json:"statistics"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
