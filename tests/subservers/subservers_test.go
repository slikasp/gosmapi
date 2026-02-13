package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestSubserver(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	subServerID := os.Getenv("SUBSERVER_POWERSCALE_ID")

	server, err := client.Subserver(context.Background(), subServerID)
	if err != nil {
		t.Fatalf("SubServer failed: %v", err)
	}

	if server.Attributes.Name != "System" {
		t.Logf("Wrong server returned %s", server.Attributes.Name)
	}
}

func TestSubserverAccess(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	// Create server
	powerscaleUser := os.Getenv("SERVER_POWERSCALE_USER")
	powerscalePass := os.Getenv("SERVER_POWERSCALE_PASS")
	powerscale := gosmapi.AddIntegratedFileServerAttributes{}
	powerscale.ServerType = gosmapi.PowerScale
	powerscale.Name = "gosmapi_integrated_access"
	powerscale.ConnectionConfig.Management = gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_POWERSCALE_ADDRESS")},
		Username:  &powerscaleUser,
		Password:  &powerscalePass,
	}
	serverPowerscale, err := client.AddIntegratedFileServer(context.Background(), powerscale)
	if err != nil {
		t.Fatalf("AddIntegratedFileServer failed: %v", err)
	}
	systemSubserver := serverPowerscale.Relationships.SubServers.Data[0].ID

	// Configure subserver access

	smbUser := os.Getenv("SMB_USER_NAME")
	smbPass := os.Getenv("SMB_USER_PASS")
	modConnection := gosmapi.DataConnection{}
	modSMB := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SUBSERVER_POWERSCALE_ADDRESS")},
		Username:  &smbUser,
		Password:  &smbPass,
	}
	modNFS := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SUBSERVER_POWERSCALE_ADDRESS")},
	}
	modConnection.SMB = &modSMB
	modConnection.NFS = &modNFS

	modServerPowerscale, err := client.EditSubserverConnection(context.Background(), systemSubserver, modConnection)
	if err != nil {
		t.Fatalf("EditSubServerAccess for %s failed: %v", serverPowerscale.Attributes.Name, err)
	}

	smbAddress := modServerPowerscale.Attributes.ConnectionConfig.SMB.Addresses[0]
	nfsAddress := modServerPowerscale.Attributes.ConnectionConfig.NFS.Addresses[0]
	if smbAddress != nfsAddress {
		t.Fatalf("SMB and NFS Subserver addresses must match: %s != %s", smbAddress, nfsAddress)
	}

	// Configure share

	modData1 := gosmapi.EditSubserverDataAccess{}
	share := gosmapi.DataAccessShareInput{
		Mode:            gosmapi.ManualShare,
		DataAccessShare: "ifs$",
	}
	modData1.SMB = &map[string]gosmapi.DataAccessShareInput{
		"/ifs": share,
	}

	mod1ServerPowerscale, err := client.EditDataAccess(context.Background(), systemSubserver, modData1)
	if err != nil {
		t.Fatalf("EditDataAccess for %s failed: %v", serverPowerscale.Attributes.Name, err)
	}

	smbShareMode := mod1ServerPowerscale.Attributes.DataAccess.SMB["/ifs"].Mode
	if smbShareMode != share.Mode {
		t.Fatalf("SMB share mode not changed: %s != %s", smbShareMode, share.Mode)
	}

	// Configure export

	modData2 := gosmapi.EditSubserverDataAccess{}
	export := gosmapi.DataAccessExportInput{
		Mode:             gosmapi.AutomaticShare,
		DataAccessExport: "/ifs/home/Testing/mil_dir/Subfolder1",
	}
	modData2.NFS = &map[string]gosmapi.DataAccessExportInput{
		"/ifs/home/Testing/mil_dir/Subfolder1": export,
	}

	mod2ServerPowerscale, err := client.EditDataAccess(context.Background(), systemSubserver, modData2)
	if err != nil {
		t.Fatalf("EditDataAccess for %s failed: %v", serverPowerscale.Attributes.Name, err)
	}

	nfsExportMode := mod2ServerPowerscale.Attributes.DataAccess.NFS["/ifs/home/Testing/mil_dir/Subfolder1"].Mode
	if nfsExportMode != export.Mode {
		t.Fatalf("NFS export mode not changed: %s != %s", nfsExportMode, export.Mode)
	}

	// Unconfigure access

	// modSMB.Addresses = []string{}
	modNFS.Addresses = []string{}
	// send SMB as to null to remove IP address and User configuration
	// for NFS you can either send empty Addresses or null
	modConnection.SMB = nil
	modConnection.NFS = &modNFS

	unmodServerPowerscale, err := client.EditSubserverConnection(context.Background(), systemSubserver, modConnection)
	if err != nil {
		t.Fatalf("EditSubServerAccess for %s failed: %v", unmodServerPowerscale.Attributes.Name, err)
	}

	unmodSMB := unmodServerPowerscale.Attributes.ConnectionConfig.SMB
	if unmodSMB != nil {
		t.Fatalf("Addresses still configured: %v", unmodSMB)
	}
}

func TestSubServerProxies(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	// Create server
	powerscaleUser := os.Getenv("SERVER_POWERSCALE_USER")
	powerscalePass := os.Getenv("SERVER_POWERSCALE_PASS")
	powerscale := gosmapi.AddIntegratedFileServerAttributes{}
	powerscale.ServerType = gosmapi.PowerScale
	powerscale.Name = "gosmapi_integrated_proxies"
	powerscale.ConnectionConfig.Management = gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_POWERSCALE_ADDRESS")},
		Username:  &powerscaleUser,
		Password:  &powerscalePass,
	}
	serverPowerscale, err := client.AddIntegratedFileServer(context.Background(), powerscale)
	if err != nil {
		t.Fatalf("AddIntegratedFileServer failed: %v", err)
	}
	systemSubserver := serverPowerscale.Relationships.SubServers.Data[0].ID

	// Configure subserver access

	smbUser := os.Getenv("SMB_USER_NAME")
	smbPass := os.Getenv("SMB_USER_PASS")
	modConnection := gosmapi.DataConnection{}
	modSMB := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SUBSERVER_POWERSCALE_ADDRESS")},
		Username:  &smbUser,
		Password:  &smbPass,
	}
	modNFS := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SUBSERVER_POWERSCALE_ADDRESS")},
	}
	modConnection.SMB = &modSMB
	modConnection.NFS = &modNFS

	subserverPowerscale, err := client.EditSubserverConnection(context.Background(), systemSubserver, modConnection)
	if err != nil {
		t.Fatalf("EditSubServerAccess for %s failed: %v", subserverPowerscale.Attributes.Name, err)
	}

	smbAddress := subserverPowerscale.Attributes.ConnectionConfig.SMB.Addresses[0]
	nfsAddress := subserverPowerscale.Attributes.ConnectionConfig.NFS.Addresses[0]
	if smbAddress != nfsAddress {
		t.Fatalf("SMB and NFS Subserver addresses must match: %s != %s", smbAddress, nfsAddress)
	}

	// Assign proxies

	nfsProxyID := "NFS:NFS_Proxy"
	universalProxyID := "UNIVERSAL:builtin"
	modProxies := []string{nfsProxyID, universalProxyID}

	err = client.SubserverSetProxies(context.Background(), systemSubserver, modProxies)
	if err != nil {
		t.Fatalf("SubserverSetProxies for %s failed: %v", subserverPowerscale.Attributes.Name, err)
	}

	// Check proxies

	proxies, err := client.SubserverProxies(context.Background(), subserverPowerscale.ID)
	if err != nil {
		t.Fatalf("SubserverProxies for %s failed: %v", subserverPowerscale.Attributes.Name, err)
	}

	if len(proxies) == 0 {
		t.Fatalf("No proxies assigned for %s", subserverPowerscale.Attributes.Name)
	}

	// Unassign proxies
	err = client.SubserverRemoveProxies(context.Background(), subserverPowerscale.ID, modProxies)
	if err != nil {
		t.Fatalf("SubserverRemoveProxies for %s failed: %v", subserverPowerscale.Attributes.Name, err)
	}

	// Check proxies

	proxies, err = client.SubserverProxies(context.Background(), subserverPowerscale.ID)
	if err != nil {
		t.Fatalf("SubserverProxies for %s failed: %v", subserverPowerscale.Attributes.Name, err)
	}

	if len(proxies) > 0 {
		t.Fatalf("Proxies still assigned on %s", subserverPowerscale.Attributes.Name)
	}

}
