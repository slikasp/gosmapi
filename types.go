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
	AutomaticShare ShareMode = "AUTOMATIC"
	ManualShare    ShareMode = "MANUAL"
	NoneShare      ShareMode = "NONE"
)

// My own for better readability

type RequestType string

const (
	GetRequest    RequestType = "GET"
	PostRequest   RequestType = "POST"
	PatchRequest  RequestType = "PATCH"
	DeleteRequest RequestType = "DELETE"
)

type Endpoint string

const (
	ServersEndpoint           Endpoint = "servers"
	SubserversEndpoint        Endpoint = "subservers"
	ProxiesEndpoint           Endpoint = "proxies"
	JobsEndpoint              Endpoint = "jobs"
	CreatejobstatusesEndpoint Endpoint = "createjobstatuses"
)

type ObjectType string

const (
	ServersObject           ObjectType = "servers"
	SubserversObject        ObjectType = "subservers"
	ProxiesObject           ObjectType = "proxies"
	JobsObject              ObjectType = "jobs"
	CreatejobstatusesObject ObjectType = "createjobstatuses"
)

type AssignElement string

const (
	AssignedProxiesElement    AssignElement = "assignedProxies"
	AssignedSubserversElement AssignElement = "assignedSubServers"
)

type DiscoveryStrategy string

const (
	DefaultStrategy DiscoveryStrategy = "DEFAULT"
	ManualStrategy  DiscoveryStrategy = "MANUAL"
	CronStrategy    DiscoveryStrategy = "CRON"
)

type JobProtocol string

const (
	NfsProtocol JobProtocol = "NFS"
	SmbProtocol JobProtocol = "SMB"
	// all other options?
	// object?
)

type NfsConstraints string

const (
	NfsV3 NfsConstraints = "V3"
	NfsV4 NfsConstraints = "V4"
)

type JobType string

const (
	NasMigrationJob JobType = "NAS_MIGRATION"
)

type JobStatus string

const (
	QueuedStatus JobStatus = "QUEUED"
	FailedStatus JobStatus = "FAILED"
	// other?
)

type JobPhase string

const (
	FirstScanPhase JobPhase = "FIRST_SCAN"
)
