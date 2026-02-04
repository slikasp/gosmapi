package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestServer(t *testing.T) {
	godotenv.Load(".env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))
	serverID := os.Getenv("SERVER_POWERSCALE_ID")

	server, err := client.FileServer(context.Background(), serverID)
	if err != nil {
		t.Fatalf("GetServer failed: %v", err)
	}

	if server.Attributes.Name != "_postman_powerscale" {
		t.Logf("Wrong server returned %s", server.Attributes.Name)
	}
}

func TestServers(t *testing.T) {
	godotenv.Load(".env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	servers, err := client.FileServers(context.Background())
	if err != nil {
		t.Fatalf("GetServers failed: %v", err)
	}

	if len(servers) == 0 {
		t.Logf("No servers returned")
	}
}

func TestFileServerGeneric(t *testing.T) {
	godotenv.Load(".env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	// Creation

	genericNFS := gosmapi.AddOtherFileServerAttributes{}
	genericSMB := gosmapi.AddOtherFileServerAttributes{}
	genericBoth := gosmapi.AddOtherFileServerAttributes{}

	// NFS
	genericNFS.ServerType = gosmapi.OtherNas
	genericNFS.ActualServerType = gosmapi.OtherNas
	genericNFS.Name = "gosmapi_generic_nfs"
	genericNFS.ConnectionConfig.NFS = &gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
	}
	// SMB
	smbUser := os.Getenv("SMB_USER_NAME")
	smbPass := os.Getenv("SMB_USER_PASS")
	genericSMB.ServerType = gosmapi.OtherNas
	genericSMB.ActualServerType = gosmapi.OtherNas
	genericSMB.Name = "gosmapi_generic_smb"
	genericSMB.ConnectionConfig.SMB = &gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
		Username:  &smbUser,
		Password:  &smbPass,
	}
	// SMB and NFS
	genericBoth.ServerType = gosmapi.OtherNas
	genericBoth.ActualServerType = gosmapi.PowerScale
	genericBoth.Name = "gosmapi_generic_both"
	genericBoth.ConnectionConfig.NFS = &gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
	}
	genericBoth.ConnectionConfig.SMB = &gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
		Username:  &smbUser,
		Password:  &smbPass,
	}

	serverNFS, err := client.AddOtherFileServer(context.Background(), genericNFS)
	if err != nil {
		t.Fatalf("AddOtherFileServer failed: %v", err)
	}
	if serverNFS.Attributes.Name != genericNFS.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverNFS.Attributes.Name, genericNFS.Name)
	}

	serverSMB, err := client.AddOtherFileServer(context.Background(), genericSMB)
	if err != nil {
		t.Fatalf("AddOtherFileServer failed: %v", err)
	}
	if serverSMB.Attributes.Name != genericSMB.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverSMB.Attributes.Name, genericSMB.Name)
	}

	serverBoth, err := client.AddOtherFileServer(context.Background(), genericBoth)
	if err != nil {
		t.Fatalf("AddOtherFileServer failed: %v", err)
	}
	if serverBoth.Attributes.Name != genericBoth.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverBoth.Attributes.Name, genericBoth.Name)
	}

	// Modification

	modifyConnectionSMB := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
		Username:  &smbUser,
		Password:  &smbPass,
	}
	modifyConnectionNFS := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_OTHER_ADDRESS")},
	}
	modifySchedulingManual := gosmapi.Scheduling{
		Strategy: gosmapi.StrategyManual,
	}
	modifySchedulingCron := gosmapi.Scheduling{
		Strategy: gosmapi.StrategyCron,
		CronSchedule: gosmapi.CronSchedule{
			Name:         nil,
			CronPatterns: []string{"0 * */4 * *"},
			TimeZoneID:   "UTC",
		},
	}
	modifyTz1 := "UTC+1"
	modifyTz2 := "UTC+2"
	modifyBusinessTime := gosmapi.TimeIntervals{
		Monday:    []string{"00:00-02:00"},
		Tuesday:   []string{"02:00-04:00"},
		Wednesday: []string{"04:00-06:00"},
		Thursday:  []string{"06:00-08:00"},
		Friday:    []string{"12:00-14:00", "08:00-10:00", "04:00-06:00"},
		Saturday:  []string{"00:00-02:00", "22:00-00:00"},
		Sunday:    []string{"00:00-00:00"},
	}
	modifyThreads := gosmapi.ThreadCount{
		Scan:     1,
		Commands: 2,
	}
	modifyExclude := gosmapi.Exclude{
		Name:     "all files",
		Patterns: []string{"**.**"},
	}

	mod1 := gosmapi.EditOtherFileServerAttributes{}
	mod1.Name = "gosmapi_generic_nfs"
	mod1.ConnectionConfig.SMB = &modifyConnectionSMB
	mod1.Scheduling = &modifySchedulingCron
	mod1.Throttling.TimeZoneID = &modifyTz1
	mod1.Throttling.ThreadCount = &modifyThreads
	mod1.Throttling.BusinessHours.TimeIntervals = &modifyBusinessTime

	serverNFSmod1, err := client.EditOtherFileServer(context.Background(), mod1, serverNFS.ID)
	if err != nil {
		t.Fatalf("EditOtherFileServer for %s failed: %v", serverNFS.Attributes.Name, err)
	}
	if serverNFSmod1.Attributes.Name != mod1.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverNFSmod1.Attributes.Name, mod1.Name)
	}

	mod2 := gosmapi.EditOtherFileServerAttributes{}
	mod2.Name = "gosmapi_modified_SMB"
	mod2.ConnectionConfig.NFS = &modifyConnectionNFS
	mod2.Scheduling = &modifySchedulingManual
	mod2.Throttling.TimeZoneID = &modifyTz2
	mod2.Throttling.ThreadCount = &modifyThreads
	mod2.Exclude = &modifyExclude

	serverSMBmod2, err := client.EditOtherFileServer(context.Background(), mod2, serverSMB.ID)
	if err != nil {
		t.Fatalf("EditOtherFileServer for %s failed: %v", serverSMB.Attributes.Name, err)
	}
	if serverSMBmod2.Attributes.Throttling.TimeZoneID != modifyTz2 {
		t.Errorf("Expected the same timezone in response: %s vs %s", serverSMBmod2.Attributes.Throttling.TimeZoneID, modifyTz2)
	}

	// Deletion

}

func TestFileServerIntegrated(t *testing.T) {
	godotenv.Load(".env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	powerscale := gosmapi.AddIntegratedFileServerAttributes{}

	// PowerScale / Isilon
	powerscaleUser := os.Getenv("SERVER_POWERSCALE_USER")
	powerscalePass := os.Getenv("SERVER_POWERSCALE_PASS")
	powerscale.ServerType = gosmapi.PowerScale
	powerscale.Name = "gosmapi_integrated_powerscale"
	powerscale.ConnectionConfig.Management = &gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_POWERSCALE_ADDRESS")},
		Username:  &powerscaleUser,
		Password:  &powerscalePass,
	}

	// Creation

	// PowerScale
	// t.Log(powerscale)
	serverPowerscale, err := client.AddIntegratedFileServer(context.Background(), powerscale)
	if err != nil {
		t.Fatalf("AddIntegratedFileServer failed: %v", err)
	}
	if serverPowerscale.Attributes.Name != powerscale.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverPowerscale.Attributes.Name, powerscale.Name)
	}

	// Modification
	modifiedUser := "admin"
	modifyConnectionManagement := gosmapi.FileConnection{
		Addresses: []string{os.Getenv("SERVER_POWERSCALE_ADDRESS")},
		Username:  &modifiedUser,
		Password:  &powerscalePass,
	}
	// modifySchedulingManual := gosmapi.Discovery{
	// 	Strategy: gosmapi.StrategyManual,
	// }
	modifySchedulingCron := gosmapi.Scheduling{
		Strategy: gosmapi.StrategyCron,
		CronSchedule: gosmapi.CronSchedule{
			Name:         nil,
			CronPatterns: []string{"0 * */4 * *"},
			TimeZoneID:   "UTC",
		},
	}
	modifyTz := "UTC+1"
	modifyBusinessTime := gosmapi.TimeIntervals{
		Monday:    []string{"00:00-02:00"},
		Tuesday:   []string{"02:00-04:00"},
		Wednesday: []string{"04:00-06:00"},
		Thursday:  []string{"06:00-08:00"},
		Friday:    []string{"12:00-14:00", "08:00-10:00", "04:00-06:00"},
		Saturday:  []string{"00:00-02:00", "22:00-00:00"},
		Sunday:    []string{"00:00-00:00"},
	}
	modifyThreads := gosmapi.ThreadCount{
		Scan:     1,
		Commands: 2,
	}

	mod1 := gosmapi.EditIntegratedFileServerAttributes{}
	mod1.Name = "gosmapi_integrated_powerscale_modified"
	mod1.ConnectionConfig.Management = &modifyConnectionManagement
	mod1.Scheduling = &modifySchedulingCron
	mod1.Throttling.TimeZoneID = &modifyTz
	mod1.Throttling.BusinessHours.TimeIntervals = &modifyBusinessTime
	mod1.Throttling.ThreadCount = &modifyThreads

	serverPowerscaleMod1, err := client.EditIntegratedFileServer(context.Background(), mod1, serverPowerscale.ID)
	if err != nil {
		t.Fatalf("EditIntegratedFileServer for %s failed: %v", serverPowerscaleMod1.Attributes.Name, err)
	}
	if serverPowerscaleMod1.Attributes.Name != mod1.Name {
		t.Errorf("Expected the same server name in response: %s vs %s", serverPowerscaleMod1.Attributes.Name, mod1.Name)
	}

	// TODO add tests for every modified value

	// Deletion

}
