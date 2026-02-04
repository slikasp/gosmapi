# GoSMapi

Current version: 0.2.0

A Datadobi StorageMAP API wrapper written in Go. This project is mainly for myself to use in other projects/apps that interact with StorageMAP and a way to learn Go while writing it.

Wrapper Functionality is dependent on API capabilities, currently there is no support for Object and Analytics parts of the software.


## Installation

```bash
go get github.com/pauslik/gosmapi
```

Then import in your Go code:

```go
import "github.com/pauslik/gosmapi"
```

## Using GoSMapi

Current version of GoSMapi supports the latest version of StorageMAP core server, so I would highly recommend to have them upgraded before use. If, for some reason, you cannot upgrade, see below for supported older versions.

API for job management was introduced in version 5.13.0 of StorageMAP (DobiMigrate back then), but server management only added in 6.7.0, some minor bugs resolved in 6.7.8 (and few previous releases).

### Create a new client

`smapi := Client.New("<core_address>", "<admin_token>")`

This creates a client with all relevant connection details, you can create multiple clients if you have multiple core servers or multiple API users (authorisation is not yet available, only use admin user)

### api/servers functions (v0.1.0)

#### GET

- `Client.FileServers()`
Get details of one File Server, including subservers
- `Client.FileServer(serverID)`
Get all File Servers

#### POST

- `Client.AddIntegratedFileServer(server)`
Configure integrated File Server using the management address
- `Client.AddOtherFileServer(server)`
Create generic File Server. Instead of Management, you have to specify at least one of Nfs or Smb to connect directly to data interface.
 
 #### PATCH

- `Client.EditIntegratedFileServer(changes)`
Modify integrated File Server.

- `Client.EditOtherFileServer(changes)`
Modify generic File Server.

#### DELETE

- No API support



### subservers.go (v0.2)

#### GET

- `Client.SubServers()`
- `Client.SubServer(subServerID)`
- `Client.SubServerParent(subServerID)`
- `Client.SubServerAssignedProxies(subServerID)`

#### POST

- No API support (subservers are auto-discovered)

#### PATCH

- `Client.EditSubServerAccess(access, subServerID)`
- `Client.EditDataAccess(access, subServerID)`
- `Client.AssignProxies(proxies, subServerID)`

// TODO: add struct and function for path overrides

#### DELETE

- No API support



### proxies.go (not yet available in current version) (v0.3)

#### GET
#### POST
#### PATCH
#### DELETE

### jobs.go (not yet available in current version) (v0.4)

#### GET
#### POST
#### PATCH
#### DELETE

### createjobstatuses.go (not yet available in current version) (v0.5)

#### GET
#### POST
#### PATCH
#### DELETE

### switchovergroups.go (not yet available in current version) TBC

#### GET
#### POST
#### PATCH
#### DELETE

### principalmaps.go (not yet available in current version) TBC

#### GET
#### POST
#### PATCH
#### DELETE


## Examples

- From 0 to a running migration

## Built-in types

Is this at all needed?

### FileServerType

- AmazonEfs
- AmazonFsxn
- AzureFiles
- HpeAlletra
- HitachiHdi
- Hnas
- Netapp
- Netapp7m
- NetappCdot
- OtherNas
- PowerScale
- PowerStore
- Unity
- Vnx

### ShareMode

- ShareAutomatic
- ShareManual
- ShareNone

### DiscoveryStrategy

- StrategyDefault
- StrategyManual
- StrategyCron

## External packages used

None? Feels like I did something wrong

Shoutout to Matt Holt for his wornderful tool (https://mholt.github.io/json-to-go/) that helped dealing with the JSON soup.