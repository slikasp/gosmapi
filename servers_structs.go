package gosmapi

// GENERAL

// Attributes / Connection

// see ConnectionConfig in general_structs

// Attributes / Scheduling
type Scheduling struct {
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

// OUTPUT

type FileServer struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name             string     `json:"name"`
		ServerType       string     `json:"serverType"`
		ActualServerType string     `json:"actualServerType"`
		ConnectionConfig Connection `json:"discoveryConnectionConfig"`
		Scheduling       Scheduling `json:"discoveryScheduling"`
		Throttling       struct {
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
type SingleFileServerOutput struct {
	Data FileServer `json:"data"`
}

// Used in parsing response containing multiple servers
type MultiFileServerOutput struct {
	Data []FileServer `json:"data"`
}

// POST INPUT

// Minimal input to create an OtherNas file server.
//
// - ServerType has to be gosmapi.OtherNas.
//
// - ActualServerType can be the actual type of the server.
//
// - One or both of SMB/NFS ConnectionConfig is required.
type AddOtherFileServerAttributes struct {
	ServerType       FileServerType `json:"serverType"`
	ActualServerType FileServerType `json:"actualServerType"`
	Name             string         `json:"name"`
	ConnectionConfig DataConnection `json:"discoveryConnectionConfig"`
}

// Minimal inputs to create an Integrated file server.
//
// - Management ConnectionConfig is required.
type AddIntegratedFileServerAttributes struct {
	ServerType       FileServerType       `json:"serverType"`
	Name             string               `json:"name"`
	ConnectionConfig ManagementConnection `json:"discoveryConnectionConfig"`
}

type AddFileServerAttributes interface {
	AddOtherFileServerAttributes | AddIntegratedFileServerAttributes
}

type addFileServerInput[T AddFileServerAttributes] struct {
	Data struct {
		Attributes T `json:"attributes"`
	} `json:"data"`
}

// PATCH INPUT

// Input struct for OtherNas file server modification
//
// Data.Attributes.Name is mandatory (for some reason)
//
// All other fields are optional
type EditOtherFileServerAttributes struct {
	Name             string         `json:"name"`
	ConnectionConfig DataConnection `json:"discoveryConnectionConfig"`
	Scheduling       *Scheduling    `json:"discoveryScheduling"`
	Throttling       struct {
		TimeZoneID    *string `json:"timeZoneId"`
		BusinessHours struct {
			ThreadCount   *ThreadCount   `json:"threadCountLimits"`
			TimeIntervals *TimeIntervals `json:"timeIntervals"`
		} `json:"businessHours"`
		ThreadCount *ThreadCount `json:"defaultThreadCountLimits"`
	} `json:"throttlingSchedule"`
	Exclude *Exclude `json:"excludePatterns"`
}

// Input struct for Integrated file server modification
//
// All fields are optional
type EditIntegratedFileServerAttributes struct {
	Name             string               `json:"name"`
	ConnectionConfig ManagementConnection `json:"discoveryConnectionConfig"`
	Scheduling       *Scheduling          `json:"discoveryScheduling"`
	Throttling       struct {
		TimeZoneID    *string `json:"timeZoneId"`
		BusinessHours struct {
			ThreadCount   *ThreadCount   `json:"threadCountLimits"`
			TimeIntervals *TimeIntervals `json:"timeIntervals"`
		} `json:"businessHours"`
		ThreadCount *ThreadCount `json:"defaultThreadCountLimits"`
	} `json:"throttlingSchedule"`
}

type EditFileServerAttributes interface {
	EditOtherFileServerAttributes | EditIntegratedFileServerAttributes
}

type editFileServerInput[T EditFileServerAttributes] struct {
	Data struct {
		Attributes T `json:"attributes"`
	} `json:"data"`
}
