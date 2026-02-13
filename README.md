# GoSMapi

Current version: 0.3.0

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

### api/servers functions (v0.1)

#### GET

- `smapi.FileServers()`
Get details of one File Server, including subservers
- `smapi.FileServer(serverID)`
Get all File Servers

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

- No API support



### subservers.go (v0.2)

#### GET

- `smapi.SubServers()`
- `smapi.SubServer(subserverID)`
- `smapi.SubserverParent(subserverID)`
- `smapi.SubserverProxies(subServerID)`

#### POST

- `smapi.SubserverAddProxies(subserverID, proxies)`

#### PATCH

- `smapi.EditSubserverConnection(subserverID, connection)`
- `smapi.EditDataAccess(subServerID, access)`
- `smapi.SubserverSetProxies(subserverID, proxies)`

// TODO: add struct and function for path overrides

#### DELETE

- `smapi.SubserverRemoveProxies(subserverID, proxies)`



### proxies.go (v0.3)

#### GET

- `smapi.Proxies()`
- `smapi.Proxy(proxyID)`
- `smapi.ProxySubservers(proxyID)`

#### POST

- `smapi.ProxyAddSubservers(proxyID, subservers)`

#### PATCH

- `smapi.ProxySetSubservers(proxyID, subservers)`

#### DELETE

- `smapi.ProxyRemoveSubservers(proxyID, subservers)`


### jobs.go (not yet available in current version) (v0.4)

#### GET
#### POST
#### PATCH
#### DELETE

### createjobstatuses.go (not yet available in current version) (v0.4)

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

Testing:
github.com/joho/godotenv

Shoutout to Matt Holt for his wornderful tool (https://mholt.github.io/json-to-go/) that helped dealing with the JSON soup.