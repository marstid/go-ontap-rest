package ontap

import "time"

type Volume struct {
	UUID       string    `json:"uuid"`
	Comment    string    `json:"comment"`
	CreateTime time.Time `json:"create_time"`
	Language   string    `json:"language"`
	Name       string    `json:"name"`
	Size       int       `json:"size"`
	State      string    `json:"state"`
	Style      string    `json:"style"`
	Tiering    struct {
		Policy string `json:"policy"`
	} `json:"tiering"`
	Type       string `json:"type"`
	Aggregates []struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"aggregates"`
	Clone struct {
		IsFlexclone bool `json:"is_flexclone"`
	} `json:"clone"`
	Nas struct {
		ExportPolicy struct {
			Name string `json:"name"`
		} `json:"export_policy"`
	} `json:"nas"`
	SnapshotPolicy struct {
		Name string `json:"name"`
	} `json:"snapshot_policy"`
	Svm struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"svm"`
	Space struct {
		Size      int `json:"size"`
		Available int `json:"available"`
		Used      int `json:"used"`
	} `json:"space"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
