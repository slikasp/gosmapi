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
	NasMigration JobType = "NAS_MIGRATION"
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

// Job General Options

// Default | MD5 | SHA1 | SHA256 | SHA512 | XXH128
type DigestAlgorithm string

const (
	DefaultAlgo DigestAlgorithm = "DEFAULT"
	MD5Algo     DigestAlgorithm = "MD5"
	SHA1Algo    DigestAlgorithm = "SHA1"
	SHA256Algo  DigestAlgorithm = "SHA256"
	SHA512Algo  DigestAlgorithm = "SHA512"
	XXH128Algo  DigestAlgorithm = "XXH128"
)

func (s DigestAlgorithm) String() string { return string(s) }

type MinimumAge string

const (
	DefaultAge    MinimumAge = "DEFAULT"
	NoneAge       MinimumAge = "ZERO_SECONDS"
	FifteenSecAge MinimumAge = "FIFTEEN_SECONDS"
	ThirtySecAge  MinimumAge = "THIRTY_SECONDS"
	OneMinAge     MinimumAge = "ONE_MINUTE"
	FiveMinAge    MinimumAge = "FIVE_MINUTES"
	TenMinAge     MinimumAge = "TEN_MINUTES"
	ThirtyMinAge  MinimumAge = "THIRTY_MINUTES"
	OneHourAge    MinimumAge = "ONE_HOUR"
	TwoHourAge    MinimumAge = "TWO_HOURS"
	ThreeHourAge  MinimumAge = "THREE_HOURS"
	SixHourAge    MinimumAge = "SIX_HOURS"
	NineHourAge   MinimumAge = "NINE_HOURS"
	TwelveHourAge MinimumAge = "TWELVE_HOURS"
	OneDayAge     MinimumAge = "ONE_DAY"
	TwoDayAge     MinimumAge = "TWO_DAYS"
	ThreeDayAge   MinimumAge = "THREE_DAYS"
	OneWeekAge    MinimumAge = "ONE_WEEK"
)

// Full | None | MigratedOnly
type CocMode string

const (
	FullCoc         CocMode = "FULL"
	NoneCoc         CocMode = "NONE"
	MigratedOnlyCoc CocMode = "MIGRATED_ONLY"
)

// Normal | CreateOverNfs | NfsOnly
type SmbSymlinkTargetMode string

const (
	NormalSymlink        SmbSymlinkTargetMode = "NORMAL"
	CreateOverNfsSymlink SmbSymlinkTargetMode = "CREATE_OVER_NFS"
	NfsOnlySymlink       SmbSymlinkTargetMode = "NFS_ONLY"
)

// CopySecurity | DontCopySecurity | ConvertToExplicit | ConsolidateRoot | ConsolidateSubfolder
type CopyRootDirMode string

const (
	CopySecurity         CopyRootDirMode = "COPY_SECURITY"
	DontCopySecurity     CopyRootDirMode = "DONT_COPY_SECURITY"
	ConvertToExplicit    CopyRootDirMode = "COPY_SECURITY_INSTANTIATED"
	ConsolidateRoot      CopyRootDirMode = "CONSOLIDATE_AT_ROOT_LEVEL"
	ConsolidateSubfolder CopyRootDirMode = "CONSOLIDATE_AT_SUB_FOLDER_LEVEL"
)

// NoRestrictions | NoDeletes | NoUpdates
type OperationRestrictions string

const (
	NoRestrictions OperationRestrictions = "NO_RESTRICTIONS"
	NoDeletes      OperationRestrictions = "NO_DELETES"
	NoUpdates      OperationRestrictions = "NO_DELETES_OR_UPDATES"
)

// ReportAll | CopyFilesReportDirs | CopyAll
type MupScanErrorMode string

const (
	ReportAll           MupScanErrorMode = "REPORT_ERROR"
	CopyFilesReportDirs MupScanErrorMode = "CONTINUE_ON_NFS_ONLY_FOR_FILES"
	CopyAll             MupScanErrorMode = "CONTINUE_ON_NFS_ONLY_FOR_FILES_AND_DIRS"
)

// Enabled | Disabled | Strict
type AipMode string

const (
	EnabledAip  AipMode = "ENABLED"
	DisabledAip AipMode = "DISABLED"
	StrictAip   AipMode = "STRICT"
)

// Job NFS Options

// does not work at the moment

// Job SMB Options

// does not work at the moment
