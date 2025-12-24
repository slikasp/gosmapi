package requests

import (
	"fmt"
	"strings"
	"testing"
)

// add variables in order to test other values
type APIResponse[T any] struct {
	Links struct {
		Self string
	}
	Data []T
}

func getTester[T any](t *testing.T, endpoint string) {
	t.Helper()

	// TODO: pass server IP via *config
	link := fmt.Sprintf("http://%s/api/%s", "10.10.113.3", endpoint)
	body, _ := getRequestBody(link)
	var response APIResponse[T]
	response, _ = parse(response, body)
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
	getTester[jobs](t, "jobs")
}
func TestGetPrincipalmaps(t *testing.T) {
	getTester[principalmaps](t, "principalmaps")
}
func TestGetServers(t *testing.T) {
	getTester[servers](t, "servers")
}
func TestGetProxies(t *testing.T) {
	getTester[proxies](t, "proxies")
}
func TestGetSubservers(t *testing.T) {
	getTester[subservers](t, "subservers")
}
func TestGetSwitchovergroups(t *testing.T) {
	getTester[switchovergroups](t, "switchovergroups")
}
