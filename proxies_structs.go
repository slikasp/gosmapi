package gosmapi

type Proxy struct {
	ID         string     `json:"id"`
	Type       ObjectType `json:"type"`
	Attributes struct {
		Status          string `json:"status"`
		Type            string `json:"type"`
		Address         string `json:"address"`
		SoftwareVersion string `json:"softwareVersion"`
	} `json:"attributes"`
}

type ProxyOutput struct {
	Data Proxy `json:"data"`
}

type ProxiesOutput struct {
	Data []Proxy `json:"data"`
}
