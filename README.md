# GoSMapi

A Datadobi StorageMAP API wrapper written in Go. This project is mainly for myself to use in other projects/apps that interact with StorageMAP and a way to learn Go while writing it.

Wrapper Functionality is dependent on API capabilities, currently there is no support for Object and Analytics parts of the software.

Current version: 0.1.0

## Installation

```bash
go get github.com/pauslik/gosmapi
```

Then import in your Go code:

```go
import "github.com/pauslik/gosmapi"
```

## Using GoSMapi

### Create a new client

`smapi := Client.New("<core_address>", "<admin_token>")`

This creates a client with all relevant connection details, you can create multiple clients if you have multiple core servers or multiple API users (authorisation is not yet available, only use admin user)

### Servers actions (v0.1)

- Get details of one File Server, including subservers: `smapi.FileServersGet()`

- Get all File Servers: `smapi.FileServersGetAll()`

- Create integrated File Server: `smapi.FileServerAddIntegrated()`
TODO: create method
TODO: add tests

- Create generic File Server: `smapi.FileServerAddOther()`
Instead of Management, you have to specify at least one of Nfs or Smb to connect directly to data interface.
Define 

- Modify integrated File Server: `smapi.FileServerEditIntegrated()`
TODO: create method
TODO: add tests

- Modify generic File Server: `smapi.FileServerEditOther()`

- Remove File Server: NOT AVAILABLE IN CURRENT STORAGEMAP VERSION

### subservers.go (not yet available in current version) (v0.2)
### proxies.go (not yet available in current version) (v0.3)
### jobs.go (not yet available in current version) (v0.4)
### createjobstatuses.go (not yet available in current version) TBC
### switchovergroups.go (not yet available in current version) TBC
### principalmaps.go (not yet available in current version) TBC

## Examples

- From 0 to a running migration

## Built-in types

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

