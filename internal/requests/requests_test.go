package requests

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/pauslik/gosmapi/internal/config"
)

// This struct has parts of the response json that fits to every endpoint
// add variables in correct structure to test be able to test other values
type APIResponse[T any] struct {
	Links struct {
		Self string
	}
	Data []T
}

// TODO: remake the cofig so the variables are taken from config file
var testConfig Config

func init() {
	testConfig.SetAdmin("admin", "qIKgyRASxdrSPEhqW36VDGffINp5b4")
	testConfig.SetCore("10.10.113.3")
	testConfig.NewClient(10)
}

// TODO: handle errors coming from requests and parsers
func getTester[T any](t *testing.T, cfg *Config, path string) {
	t.Helper()

	link := fmt.Sprintf("http://%s/api/%s", cfg.Core.Address, path)
	body, _ := getRequest(cfg, link)

	var response APIResponse[T]
	response, _ = parseResponse(response, body)

	if strings.Compare(response.Links.Self, link) != 0 {
		t.Errorf("Bad request response:\n  >expected %v\n  >got: %v", link, response.Links.Self)
		t.Fail()
	}
	if len(response.Data) == 0 {
		t.Errorf("No items returned/configured")
		t.Fail()
	}
}

func TestGetJobs(t *testing.T) {
	getTester[jobs](t, &testConfig, "jobs")
}
func TestGetPrincipalmaps(t *testing.T) {
	getTester[principalmaps](t, &testConfig, "principalmaps")
}

func TestGetServers(t *testing.T) {
	t.Log()

	response, _ := GetServers(&testConfig)

	if len(response.Data) == 0 {
		t.Errorf("No items returned/configured")
		t.Fail()
	}
}

func TestGetProxies(t *testing.T) {
	getTester[proxies](t, &testConfig, "proxies")
}
func TestGetSubservers(t *testing.T) {
	getTester[subservers](t, &testConfig, "subservers")
}
func TestGetSwitchovergroups(t *testing.T) {
	getTester[switchovergroups](t, &testConfig, "switchovergroups")
}
