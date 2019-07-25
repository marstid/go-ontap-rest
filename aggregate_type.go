package ontap

import "time"

type AggrResponse struct {
	Records []struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		Node struct {
			UUID  string `json:"uuid"`
			Name  string `json:"name"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"node"`
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

type AggrRecord struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Node struct {
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"node"`
	HomeNode struct {
		UUID  string `json:"uuid"`
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"home_node"`
	Space struct {
		BlockStorage struct {
			Size                 int64 `json:"size"`
			Available            int64 `json:"available"`
			Used                 int   `json:"used"`
			FullThresholdPercent int   `json:"full_threshold_percent"`
		} `json:"block_storage"`
		CloudStorage struct {
			Used int `json:"used"`
		} `json:"cloud_storage"`
		Efficiency struct {
			Savings     int     `json:"savings"`
			Ratio       float64 `json:"ratio"`
			LogicalUsed int     `json:"logical_used"`
		} `json:"efficiency"`
		EfficiencyWithoutSnapshots struct {
			Savings     int `json:"savings"`
			Ratio       int `json:"ratio"`
			LogicalUsed int `json:"logical_used"`
		} `json:"efficiency_without_snapshots"`
	} `json:"space"`
	State          string    `json:"state"`
	SnaplockType   string    `json:"snaplock_type"`
	CreateTime     time.Time `json:"create_time"`
	DataEncryption struct {
		SoftwareEncryptionEnabled bool `json:"software_encryption_enabled"`
		DriveProtectionEnabled    bool `json:"drive_protection_enabled"`
	} `json:"data_encryption"`
	BlockStorage struct {
		Primary struct {
			DiskCount     int    `json:"disk_count"`
			DiskClass     string `json:"disk_class"`
			RaidType      string `json:"raid_type"`
			RaidSize      int    `json:"raid_size"`
			ChecksumStyle string `json:"checksum_style"`
		} `json:"primary"`
		HybridCache struct {
			Enabled bool `json:"enabled"`
		} `json:"hybrid_cache"`
		Mirror struct {
			Enabled bool   `json:"enabled"`
			State   string `json:"state"`
		} `json:"mirror"`
	} `json:"block_storage"`
	Plexes []struct {
		Name  string `json:"name"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"plexes"`
	CloudStorage struct {
		AttachEligible bool `json:"attach_eligible"`
	} `json:"cloud_storage"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type AggrCreateData struct {
	Node struct {
		Name string `json:"name"`
	} `json:"node"`
	Name         string `json:"name"`
	BlockStorage struct {
		Primary struct {
			DiskCount string `json:"disk_count"`
		} `json:"primary"`
	} `json:"block_storage"`
}
