package requests

type proxies struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Status          string `json:"status"`
			Type            string `json:"type"`
			Address         string `json:"address"`
			SoftwareVersion string `json:"softwareVersion"`
		} `json:"attributes"`
		Relationships struct {
			AssignedSubServers struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"assignedSubServers"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
