package requests

type ShareSettings struct {
	Mode                         string   `json:"mode"`
	DefaultMode                  string   `json:"defaultMode"`
	DataAccessShare              any      `json:"dataAccessShare"`
	SharesAvailableForManualMode []string `json:"sharesAvailableForManualMode"`
}

type ExportSettings struct {
	Mode                          string   `json:"mode"`
	DefaultMode                   string   `json:"defaultMode"`
	DataAccessExport              string   `json:"dataAccessExport"`
	ExportsAvailableForManualMode []string `json:"exportsAvailableForManualMode"`
}

type subservers struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Name                 string `json:"name"`
			DataConnectionConfig struct {
				Smb struct {
					HostAddresses []string `json:"hostAddresses"`
					Username      string   `json:"username"`
				} `json:"smb"`
				Nfs any `json:"nfs"`
			} `json:"dataConnectionConfig"`
			PathOverrides struct {
				Smb struct {
				} `json:"smb"`
				Nfs struct {
				} `json:"nfs"`
			} `json:"pathOverrides"`
			DataAccess struct {
				Smb map[string]ShareSettings `json:"smb"`
				Nfs map[string]ShareSettings `json:"nfs"`
			} `json:"dataAccess"`
		} `json:"attributes"`
		Relationships struct {
			Server struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"server"`
			AssignedProxies struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"assignedProxies"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
