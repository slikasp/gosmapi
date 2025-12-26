package requests

type servers struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Name                      string `json:"name"`
			ServerType                string `json:"serverType"`
			ActualServerType          string `json:"actualServerType"`
			DiscoveryConnectionConfig struct {
				Smb struct {
					HostAddresses []string `json:"hostAddresses"`
					Username      string   `json:"username"`
				} `json:"smb"`
			} `json:"discoveryConnectionConfig"`
			DiscoveryScheduling struct {
				Strategy string `json:"strategy"`
			} `json:"discoveryScheduling"`
			ThrottlingSchedule struct {
				TimeZoneID    string `json:"timeZoneId"`
				BusinessHours struct {
					ThreadCountLimits struct {
						Scan     int `json:"scan"`
						Commands int `json:"commands"`
					} `json:"threadCountLimits"`
					TimeIntervals struct {
					} `json:"timeIntervals"`
				} `json:"businessHours"`
				DefaultThreadCountLimits struct {
					Scan     int `json:"scan"`
					Commands int `json:"commands"`
				} `json:"defaultThreadCountLimits"`
			} `json:"throttlingSchedule"`
			ExcludePatterns struct {
				Name            string   `json:"name"`
				ExcludePatterns []string `json:"excludePatterns"`
			} `json:"excludePatterns"`
		} `json:"attributes"`
		Relationships struct {
			SubServers struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"subServers"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
