package requests

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/pauslik/gosmapi/internal/config"
)

// This struct has parts of the response json that fits to every endpoint
// add variables in order to test other values
type APIResponse[T any] struct {
	Links struct {
		Self string
	}
	Data []T
}

// TODO: remake the cofig so the variables are taken from config file
var testConfig Config

func init() {
	user := SMuser{
		Name:     "admin",
		Token:    "qIKgyRASxdrSPEhqW36VDGffINp5b4",
		UserRole: "admin",
	}
	core := SMserver{
		Address:    "10.10.113.3",
		ServerRole: "core",
	}
	testConfig = Config{
		User: user,
		Core: core,
	}
}

func getTester[T any](t *testing.T, coreAddress string, endpoint string) {
	t.Helper()

	link := fmt.Sprintf("http://%s/api/%s", coreAddress, endpoint)
	body, _ := getRequestBody(link)
	var response APIResponse[T]
	response, _ = parseGetResponse(response, body)
	if strings.Compare(response.Links.Self, link) != 0 {
		t.Errorf("Bad request response:\n  >expected %v\n  >got: %v", link, response.Links.Self)
		t.Fail()
	}
	if len(response.Data) == 0 {
		t.Errorf("No items returned/configured")
		t.Fail()
	}
}

// func TestGetJobsOld(t *testing.T) {
// 	t.Log()
// 	// TODO: pass server IP via *config
// 	link := fmt.Sprintf("http://%s/api/servers", "10.10.113.3")
// 	body, _ := getRequestBody(link)
// 	jobs := jobs{}
// 	jobs, _ = parse(jobs, body)
// 	if strings.Compare(jobs.Links.Self, link) != 0 {
// 		t.Errorf("Bad request response:\n  >expected %v\n  >got: %v", link, jobs.Links.Self)
// 		t.Fail()
// 	}
// 	if len(jobs.Data) == 0 {
// 		t.Errorf("No jobs returned/configured")
// 		t.Fail()
// 	}
// }

func TestGetJobs(t *testing.T) {
	getTester[jobs](t, testConfig.Core.Address, "jobs")
}
func TestGetPrincipalmaps(t *testing.T) {
	getTester[principalmaps](t, testConfig.Core.Address, "principalmaps")
}
func TestGetServers(t *testing.T) {
	getTester[servers](t, testConfig.Core.Address, "servers")
}
func TestGetProxies(t *testing.T) {
	getTester[proxies](t, testConfig.Core.Address, "proxies")
}
func TestGetSubservers(t *testing.T) {
	getTester[subservers](t, testConfig.Core.Address, "subservers")
}
func TestGetSwitchovergroups(t *testing.T) {
	getTester[switchovergroups](t, testConfig.Core.Address, "switchovergroups")
}
