package ontap

type SvmResponse struct {
	Records []struct {
		Name  string `json:"name"`
		UUID  string `json:"uuid"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"records"`
	NumRecords int `json:"num_records"`
	Links      struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type StorageVM struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Subtype    string `json:"subtype"`
	Language   string `json:"language"`
	Aggregates []struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	} `json:"aggregates"`
	State   string `json:"state"`
	Comment string `json:"comment"`
	Ipspace struct {
		Name  string `json:"name"`
		UUID  string `json:"uuid"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"ipspace"`
	SnapshotPolicy struct {
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"snapshot_policy"`
	Nsswitch struct {
		Hosts    []string `json:"hosts"`
		Group    []string `json:"group"`
		Passwd   []string `json:"passwd"`
		Netgroup []string `json:"netgroup"`
		Namemap  []string `json:"namemap"`
	} `json:"nsswitch"`
	Nis struct {
		Enabled bool `json:"enabled"`
	} `json:"nis"`
	Ldap struct {
		Enabled bool `json:"enabled"`
	} `json:"ldap"`
	Nfs struct {
		Enabled bool `json:"enabled"`
	} `json:"nfs"`
	Cifs struct {
		Enabled bool `json:"enabled"`
	} `json:"cifs"`
	Iscsi struct {
		Enabled bool `json:"enabled"`
	} `json:"iscsi"`
	Fcp struct {
		Enabled bool `json:"enabled"`
	} `json:"fcp"`
	Nvme struct {
		Enabled bool `json:"enabled"`
	} `json:"nvme"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type Svm struct {
	Name     string
	UUID     string
	SelfLink string
}
