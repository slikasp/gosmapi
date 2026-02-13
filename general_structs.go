package gosmapi

// TODO: figure out if I want to use the below to unify Proxy & Subserver assignment
// Should work for assigning subservers to proxies and proxies to subservers
// type Object struct {
// 	ID   string     `json:"id"`
// 	Type ObjectType `json:"type"`
// }
// type AssignObject struct {
// 	Data []Object `json:"data"`
// }

// Used for output only
type Connection struct {
	Management *FileConnection `json:"managementCredentials"`
	SMB        *FileConnection `json:"smb"`
	NFS        *FileConnection `json:"nfs"`
}

type ManagementConnection struct {
	Management FileConnection `json:"managementCredentials"`
}

// Either SMB or NFS or both are required
type DataConnection struct {
	SMB *FileConnection `json:"smb"`
	NFS *FileConnection `json:"nfs"`
}

// No Username or Password for NFS
type FileConnection struct {
	Addresses []string `json:"hostAddresses"`
	Username  *string  `json:"username,omitempty"`
	Password  *string  `json:"password,omitempty"`
}

// RELATIONSHIPS

type Assigned struct {
	ID   string     `json:"id"`
	Type ObjectType `json:"type"`
}

type AssignedInput struct {
	Data []Assigned `json:"data"`
}

type AssignedOutput struct {
	Data []Assigned `json:"data"`
}
