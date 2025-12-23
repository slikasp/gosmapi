package requests

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetServers(t *testing.T) {
	t.Log()

	// TODO: pass server IP via *config
	link := fmt.Sprintf("http://%s/api/servers", "10.10.113.3")
	body, _ := getRequestBody(link)
	servers := servers{}
	servers, _ = parse(servers, body)
	if strings.Compare(servers.Links.Self, link) != 0 {
		t.Errorf("Bad request response:\n  >expected %v\n  >got: %v", link, servers.Links.Self)
		t.Fail()
	}
	if len(servers.Data) == 0 {
		t.Errorf("No servers returned/configured")
		t.Fail()
	}
}
