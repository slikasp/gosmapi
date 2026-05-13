# GoSMapi

Current version: 0.3.0

A Datadobi StorageMAP API wrapper written in Go. This project is mainly for myself to use in other projects/apps that interact with StorageMAP and a way to learn Go while writing it.

Wrapper Functionality is dependent on API capabilities, currently there is no support for Object and Analytics parts of the software.


## Installation

```bash
go get github.com/slikasp/gosmapi
```

Then import in your Go code:

```go
import "github.com/slikasp/gosmapi"
```

## Using GoSMapi

Current version of GoSMapi supports the latest version of StorageMAP core server, so I would highly recommend to have them upgraded before use. If, for some reason, you cannot upgrade, see below for supported older versions.

API for job management was introduced in version 5.13.0 of StorageMAP (DobiMigrate back then), but server management only added in 6.7.0, some minor bugs resolved in 6.7.8 (and few previous releases).

### Create a new client

`smapi := Client.New("<core_address>", "<admin_token>")`

This creates a client with all relevant connection details, you can create multiple clients if you have multiple core servers or multiple API users (authorisation is not yet available, only use admin user)

### servers

#### GET

- `smapi.FileServers()`

Get all File Servers

- `smapi.FileServer(serverID)`

Get details of one File Server, including subservers

#### POST

- `smapi.AddIntegratedFileServer(server)`

Configure integrated File Server using the management address

- `smapi.AddOtherFileServer(server)`

Create generic File Server. Instead of Management, you have to specify at least one of Nfs or Smb to connect directly to data interface.
 
 #### PATCH

- `smapi.EditIntegratedFileServer(serverID, changes)`

Modify integrated File Server.

- `smapi.EditOtherFileServer(serverID, changes)`

Modify generic File Server.

#### DELETE

(not available yet)



### subservers

#### GET

- `smapi.Subservers()`

Get all Subservers

- `smapi.Subserver(subserverID)`

Get details of one Subserver

- `smapi.SubserverParent(subserverID)`

Get parent Server of the Subserver

- `smapi.SubserverProxies(subServerID)`

Get Proxies assigned to Subserver


#### POST

- `smapi.SubserverAddProxies(subserverID, proxies)`

Assign Proxies to a Subserver

#### PATCH

- `smapi.EditSubserverConnection(subserverID, connection)`

Modify Subserver SMB/NFS connection

- `smapi.EditDataAccess(subServerID, access)`

Modify Subserver data access via Shares/Exports

- `smapi.SubserverSetProxies(subserverID, proxies)`

Modify Proxies assigned to a Subserver


// TODO: add struct and function for path overrides

#### DELETE

- `smapi.SubserverRemoveProxies(subserverID, proxies)`

Unassign Proxies from a Subserver


### proxies

#### GET

- `smapi.Proxies()`

Get all Proxies

- `smapi.Proxy(proxyID)`

Get details of one Proxy

- `smapi.ProxySubservers(proxyID)`

Get all Subservers assigned to a Proxy

#### POST

- `smapi.ProxyAddSubservers(proxyID, subservers)`

Assign Proxy to Subservers

#### PATCH

- `smapi.ProxySetSubservers(proxyID, subservers)`

Modify Subservers where Proxy is assigned

#### DELETE

- `smapi.ProxyRemoveSubservers(proxyID, subservers)`

Unassign Proxy from Subservers


### jobs

#### GET

- `smapi.Jobs()`

Get all jobs

- `smapi.Job(jobID)`

Get details of one job

#### POST

- `smapi.CreateJob(sourceSubserverID, destinationSubserverID, attributes)`

Create a new job

#### PATCH

(editing of jobs is not added yet so you need to create it with options provided in attributes)

#### DELETE

(not available yet)

### createjobstatuses

#### GET

- `smapi.Createjobstatuses()`

Get all createjobstatuses

- `smapi.Createjobstatus(createjobstatusID)`

Get details of one createjobstatus (returns job struct if successful)

#### POST
#### PATCH
#### DELETE

### switchovergroups.go (not yet available in current goSMapi version)

#### GET
#### POST
#### PATCH
#### DELETE

### principalmaps.go (not yet available in current goSMapi version)

#### GET
#### POST
#### PATCH
#### DELETE


## Examples

- From 0 to a running migration

## Built-in types

Is this needed in README? Can be found in types.go file

## External packages used

Testing:
github.com/joho/godotenv

Shoutout to Matt Holt for his wornderful tool (https://mholt.github.io/json-to-go/) that helped dealing with the JSON soup.
