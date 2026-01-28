package gosmapi

// Attributes
// Attributes / Connection

// Attributes / Connection / Management
type Management struct {
	Addresses []string `json:"hostAddresses"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}

// TODO: same as Management, merge them?
// Attributes / Connection / SMB
type SMB struct {
	Addresses []string `json:"hostAddresses"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}

// TODO: same as Management but less fields, merge them?
// Attributes / Connection / NFS
type NFS struct {
	Addresses []string `json:"hostAddresses"`
}

// Attributes / Discovery
type Discovery struct {
	Strategy     DiscoveryStrategy `json:"strategy"`
	CronSchedule CronSchedule      `json:"cronSchedule"`
}
type CronSchedule struct {
	Name         any      `json:"name"`
	CronPatterns []string `json:"cronPatterns"`
	TimeZoneID   string   `json:"timeZoneId"`
}

// Attributes / Throttling / BusinessHours / ThreadCount
// Attributes / Throttling / ThreadCount
type ThreadCount struct {
	Scan     int `json:"scan"`
	Commands int `json:"commands"`
}

// Attributes / Throttling / BusinessHours / TimeIntervals
type TimeIntervals struct {
	Monday    []string `json:"monday"`
	Tuesday   []string `json:"tuesday"`
	Wednesday []string `json:"wednesday"`
	Thursday  []string `json:"thursday"`
	Friday    []string `json:"friday"`
	Saturday  []string `json:"saturday"`
	Sunday    []string `json:"sunday"`
}

// Attributes / Exclude
type Exclude struct {
	Name     string   `json:"name"`
	Patterns []string `json:"excludePatterns"`
}

type Srv struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name             string `json:"name"`
		ServerType       string `json:"serverType"`
		ActualServerType string `json:"actualServerType"`
		Connection       struct {
			Management Management `json:"managementCredentials"`
			SMB        SMB        `json:"smb"`
			NFS        NFS        `json:"nfs"`
		} `json:"discoveryConnectionConfig"`
		Discovery  Discovery `json:"discoveryScheduling"`
		Throttling struct {
			TimeZoneID    string `json:"timeZoneId"`
			BusinessHours struct {
				ThreadCount   ThreadCount   `json:"threadCountLimits"`
				TimeIntervals TimeIntervals `json:"timeIntervals"`
			} `json:"businessHours"`
			ThreadCount ThreadCount `json:"defaultThreadCountLimits"`
		} `json:"throttlingSchedule"`
		Exclude Exclude `json:"excludePatterns"`
	} `json:"attributes"`
	Relationships struct {
		SubServers struct {
			Data []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"subServers"`
	} `json:"relationships"`
}

// Used in parsing input/response containing one server, usually when creating it
type Server struct {
	Data struct{ Srv } `json:"data"`
}

// Used in parsing response containing multiple servers
type Servers struct {
	Data []struct{ Srv } `json:"data"`
}

// Minimal list of variables to create an OtherNas file server.
//
// - ServerType has to be gosmapi.OtherNas.
//
// - ActualServerType can be the actual type of the server.
//
// - One or both of SMB/NFS config is required.
type FileServerAddOtherRequest struct {
	Data struct {
		Attributes struct {
			ServerType       FileServerType `json:"serverType"`
			ActualServerType FileServerType `json:"actualServerType"`
			Name             string         `json:"name"`
			Connection       struct {
				SMB *SMB `json:"smb"`
				NFS *NFS `json:"nfs"`
			} `json:"discoveryConnectionConfig"`
		} `json:"attributes"`
	} `json:"data"`
}

// Input struct for generic file server modification
//
// Data.Attributes.Name is mandatory (for some reason)
//
// All other fields are optional
type FileServerEditOtherRequest struct {
	Data struct {
		Attributes struct {
			Name       string `json:"name"`
			Connection struct {
				SMB *SMB `json:"smb"`
				NFS *NFS `json:"nfs"`
			} `json:"discoveryConnectionConfig"`
			Discovery  *Discovery `json:"discoveryScheduling"`
			Throttling struct {
				TimeZoneID    *string `json:"timeZoneId"`
				BusinessHours struct {
					ThreadCount   *ThreadCount   `json:"threadCountLimits"`
					TimeIntervals *TimeIntervals `json:"timeIntervals"`
				} `json:"businessHours"`
				ThreadCount *ThreadCount `json:"defaultThreadCountLimits"`
			} `json:"throttlingSchedule"`
			Exclude *Exclude `json:"excludePatterns"`
		} `json:"attributes"`
	} `json:"data"`
}

// Minimal list of variables to create an Integrated file server.
type FileServerAddIntegratedRequest struct {
	Data struct {
		Attributes struct {
			ServerType FileServerType `json:"serverType"`
			Name       string         `json:"name"`
			Connection struct {
				Management Management `json:"managementCredentials"`
			} `json:"discoveryConnectionConfig"`
		} `json:"attributes"`
	} `json:"data"`
}

// Input struct for integrated file server modification
//
// All fields are optional
type FileServerEditIntegratedRequest struct {
	Data struct {
		Attributes struct {
			Name       string `json:"name"`
			Connection struct {
				Management *Management `json:"managementCredentials"`
			} `json:"discoveryConnectionConfig"`
			Discovery  *Discovery `json:"discoveryScheduling"`
			Throttling struct {
				TimeZoneID    *string `json:"timeZoneId"`
				BusinessHours struct {
					ThreadCount   *ThreadCount   `json:"threadCountLimits"`
					TimeIntervals *TimeIntervals `json:"timeIntervals"`
				} `json:"businessHours"`
				ThreadCount *ThreadCount `json:"defaultThreadCountLimits"`
			} `json:"throttlingSchedule"`
		} `json:"attributes"`
	} `json:"data"`
}
