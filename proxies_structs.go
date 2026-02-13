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

type singleProxyOutput struct {
	Data Proxy `json:"data"`
}

type multipleProxyOutput struct {
	Data []Proxy `json:"data"`
}
