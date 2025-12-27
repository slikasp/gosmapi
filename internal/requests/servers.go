package requests

import (
	"fmt"

	. "github.com/pauslik/gosmapi/internal/config"
)

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
				Nfs struct {
					HostAddresses []string `json:"hostAddresses"`
				} `json:"nfs"`
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

type serverAdd struct {
	Data struct {
		Attributes struct {
			ServerType                FileServerType `json:"serverType"`
			Name                      string         `json:"name"`
			DiscoveryConnectionConfig struct {
				Smb struct {
					HostAddresses []string `json:"hostAddresses"`
					Username      string   `json:"username"`
					Password      string   `json:"password"`
				} `json:"smb"`
				Nfs struct {
					HostAddresses []string `json:"hostAddresses"`
				} `json:"nfs"`
			} `json:"discoveryConnectionConfig"`
		} `json:"attributes"`
	} `json:"data"`
}

type FileServerType string

const (
	AmazonEfs  FileServerType = "AMAZON_EFS"
	AmazonFsxn FileServerType = "AMAZON_FSXN"
	AzureFiles FileServerType = "AZURE_FILES"
	HpeAlletra FileServerType = "HPE_ALLETRA"
	HitachiHdi FileServerType = "HITACHI_HDI"
	Hnas       FileServerType = "HNAS"
	Netapp     FileServerType = "NETAPP"
	Netapp7m   FileServerType = "NETAPP_7M"
	NetappCdot FileServerType = "NETAPP_CDOT"
	OtherNas   FileServerType = "OTHER_NAS"
	PowerScale FileServerType = "POWERSCALE"
	PowerStore FileServerType = "POWERSTORE"
	Unity      FileServerType = "UNITY"
	Vnx        FileServerType = "VNX"
)

func GetServers(cfg *Config) (servers, error) {
	var response servers

	link := fmt.Sprintf("http://%s/api/%s", cfg.Core.Address, "servers")

	body, err := getRequest(cfg, link)
	if err != nil {
		return response, err
	}

	response, err = parseResponse(response, body)
	if err != nil {
		return response, err
	}

	return response, nil
}

func AddFileServerGeneric(cfg *Config,
	srvType FileServerType,
	name, smbUser, smbPass string,
	smbAddr, nfsAddr []string) (servers, error) {
	var response servers

	link := fmt.Sprintf("http://%s/api/%s", cfg.Core.Address, "servers?include=subServers")

	var server serverAdd
	server.Data.Attributes.ServerType = srvType
	server.Data.Attributes.DiscoveryConnectionConfig.Smb.Username = smbUser
	server.Data.Attributes.DiscoveryConnectionConfig.Smb.Password = smbPass
	server.Data.Attributes.DiscoveryConnectionConfig.Smb.HostAddresses = smbAddr
	server.Data.Attributes.DiscoveryConnectionConfig.Nfs.HostAddresses = nfsAddr

	body, err := postRequest(cfg, link, server)
	if err != nil {
		return response, err
	}

	response, err = parseResponse(response, body)
	if err != nil {
		return response, err
	}

	return response, nil
}
