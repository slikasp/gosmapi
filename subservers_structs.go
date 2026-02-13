package gosmapi

// GENERAL
type ShareSettings struct {
	Mode            ShareMode `json:"mode"`
	DefaultMode     string    `json:"defaultMode"`
	DataAccessShare any       `json:"dataAccessShare"`
	SharesAvailable []string  `json:"sharesAvailableForManualMode"`
}

type ExportSettings struct {
	Mode             ShareMode `json:"mode"`
	DefaultMode      string    `json:"defaultMode"`
	DataAccessExport string    `json:"dataAccessExport"`
	ExportsAvailable []string  `json:"exportsAvailableForManualMode"`
}

// TODO: check PathOverides possible output
type Subserver struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name             string         `json:"name"`
		ConnectionConfig DataConnection `json:"dataConnectionConfig"`
		PathOverrides    struct {
			SMB struct {
			} `json:"smb"`
			NFS struct {
			} `json:"nfs"`
		} `json:"pathOverrides"`
		DataAccess struct {
			SMB map[string]ShareSettings  `json:"smb"`
			NFS map[string]ExportSettings `json:"nfs"`
		} `json:"dataAccess"`
	} `json:"attributes"`
}

// Used in parsing input/response containing one subserver, usually when creating it
type SingleFileSubserverOutput struct {
	Data Subserver `json:"data"`
}

// Used in parsing response containing multiple subservers
type MultiFileSubserverOutput struct {
	Data []Subserver `json:"data"`
}

// ShareMode for Integrated server: Automatic, Manual, None
//
// ShareMode for Other server: Manual, None
type DataAccessShareInput struct {
	Mode            ShareMode `json:"mode"`
	DataAccessShare string    `json:"dataAccessShare"`
}

// ShareMode for Integrated server: Automatic, Manual, None
//
// ShareMode for Other server: Manual, None
type DataAccessExportInput struct {
	Mode             ShareMode `json:"mode"`
	DataAccessExport string    `json:"dataAccessExport"`
}

// PATCH

type editSubserverConnectionInput struct {
	Data struct {
		Attributes struct {
			ConnectionConfig DataConnection `json:"dataConnectionConfig"`
		} `json:"attributes"`
	} `json:"data"`
}

// TODO: make one type for both shares and exports
type EditSubserverDataAccess struct {
	SMB *map[string]DataAccessShareInput  `json:"smb"`
	NFS *map[string]DataAccessExportInput `json:"nfs"`
}

type editSubserverDataAccessInput struct {
	Data struct {
		Attributes struct {
			DataAccess EditSubserverDataAccess `json:"dataAccess"`
		} `json:"attributes"`
	} `json:"data"`
}
