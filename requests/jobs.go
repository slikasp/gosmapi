package requests

type jobs struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			JobType         string `json:"jobType"`
			SourcePath      string `json:"sourcePath"`
			DestinationPath string `json:"destinationPath"`
			Options         struct {
				Configuration struct {
					Protocol string `json:"protocol"`
				} `json:"configuration"`
				IterationScheduling struct {
					Strategy string `json:"strategy"`
				} `json:"iterationScheduling"`
				StartFirstIterationNow     bool     `json:"startFirstIterationNow"`
				DigestAlgorithm            string   `json:"digestAlgorithm"`
				ChainOfCustodyMode         string   `json:"chainOfCustodyMode"`
				SmbSymlinkTargetMode       string   `json:"smbSymlinkTargetMode"`
				MinimumAge                 string   `json:"minimumAge"`
				SkipFilesPatterns          []any    `json:"skipFilesPatterns"`
				ExcludeDirectoriesPatterns []string `json:"excludeDirectoriesPatterns"`
				CopyRootDirectoryMode      string   `json:"copyRootDirectoryMode"`
				OperationRestrictions      string   `json:"operationRestrictions"`
				VerifySourceAfterCopy      bool     `json:"verifySourceAfterCopy"`
				PreserveAccessTime         bool     `json:"preserveAccessTime"`
				FollowJunctionsSource      bool     `json:"followJunctionsSource"`
				FollowJunctionsTarget      bool     `json:"followJunctionsTarget"`
				MultiProtocolScanErrorMode string   `json:"multiProtocolScanErrorMode"`
				NfsOptions                 struct {
					CopyNamedAttributesMode      any `json:"copyNamedAttributesMode"`
					CopyACL                      any `json:"copyAcl"`
					CopyOwner                    any `json:"copyOwner"`
					CopyGroup                    any `json:"copyGroup"`
					CopyPermissions              any `json:"copyPermissions"`
					CopySymlinksPermissions      any `json:"copySymlinksPermissions"`
					CopySocketFilesMode          any `json:"copySocketFilesMode"`
					CopyPipeFilesMode            any `json:"copyPipeFilesMode"`
					CopyCharacterDeviceFilesMode any `json:"copyCharacterDeviceFilesMode"`
					CopyBlockDeviceFileMode      any `json:"copyBlockDeviceFileMode"`
					CopyPermissionModifier       any `json:"copyPermissionModifier"`
				} `json:"nfsOptions"`
				SmbOptions struct {
					DesiredOwner                    any `json:"desiredOwner"`
					DesiredGroup                    any `json:"desiredGroup"`
					CopySacl                        any `json:"copySacl"`
					CleanInvalidSecurityDescriptors any `json:"cleanInvalidSecurityDescriptors"`
					CopyUserDefinedAttributesMode   any `json:"copyUserDefinedAttributesMode"`
					ReplaceCreatorOwnerGroup        any `json:"replaceCreatorOwnerGroup"`
					CopyOwner                       any `json:"copyOwner"`
					CopyGroup                       any `json:"copyGroup"`
					CopyDacl                        any `json:"copyDacl"`
					CopyHiddenAttribute             any `json:"copyHiddenAttribute"`
					CopyNotContentIndexedAttribute  any `json:"copyNotContentIndexedAttribute"`
					CopyReadOnlyAttribute           any `json:"copyReadOnlyAttribute"`
					CopySystemAttribute             any `json:"copySystemAttribute"`
					CopyTemporaryAttribute          any `json:"copyTemporaryAttribute"`
					CopyOfflineAttribute            any `json:"copyOfflineAttribute"`
					CopyArchiveAttribute            any `json:"copyArchiveAttribute"`
				} `json:"smbOptions"`
				AdvancedIntegrityProtectionMode string `json:"advancedIntegrityProtectionMode"`
			} `json:"options"`
			SpecialUseCase string `json:"specialUseCase"`
			Status         struct {
				Phase                           string `json:"phase"`
				LastCompleteIterationErrorCount int    `json:"lastCompleteIterationErrorCount"`
			} `json:"status"`
		} `json:"attributes"`
		Relationships struct {
			UIDMapping struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"uidMapping"`
			GidMapping struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"gidMapping"`
			DestinationSubServer struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"destinationSubServer"`
			SwitchoverGroup struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"switchoverGroup"`
			SidMapping struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"sidMapping"`
			SourceSubServer struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"sourceSubServer"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
