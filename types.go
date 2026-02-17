package gosmapi

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
	NFS            JobProtocol = "NFS"
	SMB            JobProtocol = "SMB"
	Multiprotocol  JobProtocol = "MULTIPROTOCOL"
	SymlinkOverNFS JobProtocol = "SMB_SYMLINKS_OVER_NFS"
	// object?
)

type NfsConstraints string

const (
	NfsV3     NfsConstraints = "V3"
	NfsV4     NfsConstraints = "V4"
	NfsV3ToV4 NfsConstraints = "V3_TO_V4"
)

type JobType string

const (
	NasMigrationJob JobType = "NAS_MIGRATION"
)

type JobStatus string

const (
	Queued JobStatus = "QUEUED"
	Failed JobStatus = "FAILED"
	// other?
)

type JobPhase string

const (
	FirstScan JobPhase = "FIRST_SCAN"
)

type DigestAlgorithm string

const (
	Default DigestAlgorithm = "DEFAULT"
	MD5     DigestAlgorithm = "MD5"
	SHA1    DigestAlgorithm = "SHA1"
	SHA256  DigestAlgorithm = "SHA256"
	SHA512  DigestAlgorithm = "SHA512"
	XXH128  DigestAlgorithm = "XXH128"
)

type MinimumAge string

const (
	DefaultAge    MinimumAge = "DEFAULT"
	NoneAge       MinimumAge = "ZERO_SECONDS"
	SecFifteenAge MinimumAge = "FIFTEEN_SECONDS"
	SecThirtyAge  MinimumAge = "THIRTY_SECONDS"
	MinOneAge     MinimumAge = "ONE_MINUTE"
	MinFiveAge    MinimumAge = "FIVE_MINUTES"
	MinTenAge     MinimumAge = "TEN_MINUTES"
	MinThirtyAge  MinimumAge = "THIRTY_MINUTES"
	HourOneAge    MinimumAge = "ONE_HOUR"
	HourTwoAge    MinimumAge = "TWO_HOURS"
	HourThreeAge  MinimumAge = "THREE_HOURS"
	HourSixAge    MinimumAge = "SIX_HOURS"
	HourNineAge   MinimumAge = "NINE_HOURS"
	HourTwelveAge MinimumAge = "TWELVE_HOURS"
	DayOneAge     MinimumAge = "ONE_DAY"
	DayTwoAge     MinimumAge = "TWO_DAYS"
	DayThreeAge   MinimumAge = "THREE_DAYS"
	WeekOneAge    MinimumAge = "ONE_WEEK"
)

type CocMode string

const (
	Full         CocMode = "FULL"
	None         CocMode = "NONE"
	MigratedOnly CocMode = "MIGRATED_ONLY"
)

type SmbSymlinkTargetMode string

const (
	Normal        SmbSymlinkTargetMode = "NORMAL"
	CreateOverNfs SmbSymlinkTargetMode = "CREATE_OVER_NFS"
	NfsOnly       SmbSymlinkTargetMode = "NFS_ONLY"
)

type CopyRootDirMode string

const (
	CopySecurity         CopyRootDirMode = "COPY_SECURITY"
	DontCopySecurity     CopyRootDirMode = "DONT_COPY_SECURITY"
	ConvertToExplicit    CopyRootDirMode = "COPY_SECURITY_INSTANTIATED"
	ConsolidateRoot      CopyRootDirMode = "CONSOLIDATE_AT_ROOT_LEVEL"
	ConsolidateSubfolder CopyRootDirMode = "CONSOLIDATE_AT_SUB_FOLDER_LEVEL"
)

type OperationRestrictions string

const (
	NoRestrictions OperationRestrictions = "NO_RESTRICTIONS"
	NoDeletes      OperationRestrictions = "NO_DELETES"
	NoUpdates      OperationRestrictions = "NO_DELETES_OR_UPDATES"
)

type MupScanErrorMode string

const (
	ReportAll           MupScanErrorMode = "REPORT_ERROR"
	CopyFilesReportDirs MupScanErrorMode = "CONTINUE_ON_NFS_ONLY_FOR_FILES"
	CopyAll             MupScanErrorMode = "CONTINUE_ON_NFS_ONLY_FOR_FILES_AND_DIRS"
)

type AipMode string

const (
	Enabled  AipMode = "ENABLED"
	Disabled AipMode = "DISABLED"
	Strict   AipMode = "STRICT"
)
