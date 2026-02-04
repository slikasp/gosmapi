package gosmapi

// Defined in StorageMAP documentation

// ServerType in documentation as of 7.4, but Object servers are not added to API yet
type FileServerType string

const (
	AmazonEfs  FileServerType = "AMAZON_EFS"
	AmazonFsxn FileServerType = "AMAZON_FSXN"
	AzureFiles FileServerType = "AZURE_FILES"
	HpeAlletra FileServerType = "HPE_ALLETRA"
	HitachiHdi FileServerType = "HITACHI_HDI"
	Hnas       FileServerType = "HNAS"
	Netapp     FileServerType = "NETAPP" // Can be used during creation if exact NetApp type is not known
	Netapp7m   FileServerType = "NETAPP_7M"
	NetappCdot FileServerType = "NETAPP_CDOT"
	OtherNas   FileServerType = "OTHER_NAS"
	PowerScale FileServerType = "POWERSCALE"
	PowerStore FileServerType = "POWERSTORE"
	Unity      FileServerType = "UNITY"
	Vnx        FileServerType = "VNX"
)

type ShareMode string

const (
	ShareAutomatic ShareMode = "AUTOMATIC"
	ShareManual    ShareMode = "MANUAL"
	ShareNone      ShareMode = "NONE"
)

// My own for better readability

type RequestType string

const (
	GetRequest    RequestType = "GET"
	PostRequest   RequestType = "POST"
	PatchRequest  RequestType = "PATCH"
	DeleteRequest RequestType = "DELETE"
)

type ObjectType string

const (
	ServerObject    ObjectType = "servers"
	SubserverObject ObjectType = "subservers"
	ProxyObject     ObjectType = "proxies"
)

type DiscoveryStrategy string

const (
	StrategyDefault DiscoveryStrategy = "DEFAULT"
	StrategyManual  DiscoveryStrategy = "MANUAL"
	StrategyCron    DiscoveryStrategy = "CRON"
)
