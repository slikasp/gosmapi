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

// Management: Integrated server API connection
//
// SMB / NFS: Other NAS or Subserver data connection
//
// No Username or Password for NFS
type ConnectionConfig struct {
	Management *FileConnection `json:"managementCredentials"`
	SMB        *FileConnection `json:"smb"`
	NFS        *FileConnection `json:"nfs"`
}

type FileConnection struct {
	Addresses []string `json:"hostAddresses"`
	Username  *string  `json:"username"`
	Password  *string  `json:"password"`
}
