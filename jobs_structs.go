package gosmapi

type JobsNfsOptions struct {
	CopyNamedAttributesMode any `json:"copyNamedAttributesMode"`
	CopyACL                 any `json:"copyAcl"`
	CopyOwner               any `json:"copyOwner"`
	CopyGroup               any `json:"copyGroup"`
	CopyPermissions         any `json:"copyPermissions"`
	CopySymlinksPermissions any `json:"copySymlinksPermissions"`
	CopySocketMode          any `json:"copySocketFilesMode"`
	CopyPipeMode            any `json:"copyPipeFilesMode"`
	CopyCharDeviceMode      any `json:"copyCharacterDeviceFilesMode"`
	CopyBlockDeviceMode     any `json:"copyBlockDeviceFileMode"`
	CopyPermissionModifier  any `json:"copyPermissionModifier"`
}

type JobsSmbOptions struct {
	DesiredOwner                   any `json:"desiredOwner"`
	DesiredGroup                   any `json:"desiredGroup"`
	CopySacl                       any `json:"copySacl"`
	CleanInvalidSid                any `json:"cleanInvalidSecurityDescriptors"`
	CopyUserAttributesMode         any `json:"copyUserDefinedAttributesMode"`
	ReplaceCreatorOwnerGroup       any `json:"replaceCreatorOwnerGroup"`
	CopyOwner                      any `json:"copyOwner"`
	CopyGroup                      any `json:"copyGroup"`
	CopyDacl                       any `json:"copyDacl"`
	CopyHiddenAttribute            any `json:"copyHiddenAttribute"`
	CopyNotContentIndexedAttribute any `json:"copyNotContentIndexedAttribute"`
	CopyReadOnlyAttribute          any `json:"copyReadOnlyAttribute"`
	CopySystemAttribute            any `json:"copySystemAttribute"`
	CopyTemporaryAttribute         any `json:"copyTemporaryAttribute"`
	CopyOfflineAttribute           any `json:"copyOfflineAttribute"`
	CopyArchiveAttribute           any `json:"copyArchiveAttribute"`
}

// TODO: define actual type, similar to discovery
type IterationScheduling struct {
	Strategy string `json:"strategy"`
}

type JobConfiguration struct {
	Protocol       JobProtocol     `json:"protocol"`
	NfsConstraints *NfsConstraints `json:"nfsConstraints"`
}

// TODO: add specific types for all string types
type JobOptions struct {
	Configuration          *JobConfiguration      `json:"configuration"`
	IterationScheduling    *IterationScheduling   `json:"iterationScheduling"`
	StartFirstIterationNow *bool                  `json:"startFirstIterationNow"`
	DigestAlgorithm        *DigestAlgorithm       `json:"digestAlgorithm"`
	CocMode                *CocMode               `json:"chainOfCustodyMode"`
	SmbSymlinkTargetMode   *SmbSymlinkTargetMode  `json:"smbSymlinkTargetMode"`
	MinimumAge             *MinimumAge            `json:"minimumAge"`
	SkipFilesPatterns      []string               `json:"skipFilesPatterns"`
	ExcludePatterns        []string               `json:"excludeDirectoriesPatterns"`
	CopyRootDirMode        *CopyRootDirMode       `json:"copyRootDirectoryMode"`
	OperationRestrictions  *OperationRestrictions `json:"operationRestrictions"`
	VerifySourceAfterCopy  *bool                  `json:"verifySourceAfterCopy"`
	PreserveAccessTime     *bool                  `json:"preserveAccessTime"`
	FollowJunctionsSource  *bool                  `json:"followJunctionsSource"`
	FollowJunctionsTarget  *bool                  `json:"followJunctionsTarget"`
	MupScanErrorMode       *MupScanErrorMode      `json:"multiProtocolScanErrorMode"`
	NfsOptions             *JobsNfsOptions        `json:"nfsOptions"`
	SmbOptions             *JobsSmbOptions        `json:"smbOptions"`
	AipMode                *AipMode               `json:"advancedIntegrityProtectionMode"`
}

type Job struct {
	ID         string     `json:"id"`
	Type       ObjectType `json:"type"`
	Attributes struct {
		JobType         JobType    `json:"jobType"`
		SourcePath      string     `json:"sourcePath"`
		DestinationPath string     `json:"destinationPath"`
		Options         JobOptions `json:"options"`
		SpecialUseCase  string     `json:"specialUseCase"`
		Status          struct {
			Phase                           JobPhase `json:"phase"`
			LastCompleteIterationErrorCount int      `json:"lastCompleteIterationErrorCount"`
		} `json:"status"`
	} `json:"attributes"`
}

// JobType: "NAS_MIGRATION" for migration
//
// Options:Configuration:Protocol required
type CreateJobAttributes struct {
	JobType         string     `json:"jobType"`
	SourcePath      string     `json:"sourcePath"`
	DestinationPath string     `json:"destinationPath"`
	Options         JobOptions `json:"options"`
}

type CreateJobRelationships struct {
	SourceSubserver struct {
		Data struct {
			ID   string     `json:"id"`
			Type ObjectType `json:"type"`
		} `json:"data"`
	} `json:"sourceSubServer"`
	DestinationSubserver struct {
		Data struct {
			ID   string     `json:"id"`
			Type ObjectType `json:"type"`
		} `json:"data"`
	} `json:"destinationSubServer"`
}

type createJob struct {
	Data struct {
		Type          ObjectType             `json:"type"`
		Attributes    CreateJobAttributes    `json:"attributes"`
		Relationships CreateJobRelationships `json:"relationships"`
	} `json:"data"`
}

type editJob struct {
	Data struct {
		Attributes struct {
			Options JobOptions `json:"options"`
		} `json:"attributes"`
	} `json:"data"`
}

type singleJobOutput struct {
	Data Job `json:"data"`
}

type multipleJobOutput struct {
	Data []Job `json:"data"`
}
