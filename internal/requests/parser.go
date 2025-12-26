package requests

import (
	"encoding/json"
)

/*

{
    "links": {
        "createjobstatuses": "http://server_address/api/createjobstatuses",
        "jobs": "http://server_address/api/jobs",
        "principalmaps": "http://server_address/api/principalmaps",
        "proxies": "http://server_address/api/proxies",
        "servers": "http://server_address/api/servers",
        "subservers": "http://server_address/api/subservers",
        "switchovergroups": "http://server_address/api/switchovergroups"
    }
}

// this returns something only for queued jobs (not actually queued for migration, but for creation of the job, e.g. prechecks)
createjobstatuses: Status of a job creation.

// TODO: add filter options, filter here or use the API functionality?
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

func parseGetResponse[T any](variable T, reqBody []byte) (T, error) {
	err := json.Unmarshal(reqBody, &variable)
	return variable, err
}
