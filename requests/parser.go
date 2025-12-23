package requests

import (
	"encoding/json"
)

/*
createjobstatuses:Status of a job creation.

jobs: Job, for example, a migration.

principalmaps: An uploaded principal map.

proxies: A proxy accesses the data on servers.
Proxies are assigned to sub-servers, and StorageMAP sends the operation request to these proxies for processing.

servers: Server definition

subservers: The sub-server represents a single tenant of a multi-tenancy system.
For example, a NetApp cDOT (clustered Data ONTAP) server is subdivided in one or more Storage Virtual Machines or SVM. Each SVM will be represented by a single sub-server in StorageMAP.
The data access is configured on sub-server level in StorageMAP.

switchovergroups: Switchover group reference
*/

func parse[T any](variable T, reqBody []byte) (T, error) {
	err := json.Unmarshal(reqBody, &variable)
	return variable, err
}
