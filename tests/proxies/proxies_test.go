package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestProxies(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	proxies, err := client.Proxies(context.Background())
	if err != nil {
		t.Fatalf("Proxies failed: %v", err)
	}

	if len(proxies) == 0 {
		t.Errorf("No proxies returned.")
	}
}

func TestProxy(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	proxyID := os.Getenv("UNIVERSAL_PROXY")

	proxy, err := client.Proxy(context.Background(), proxyID)
	if err != nil {
		t.Fatalf("Proxy failed: %v", err)
	}

	if proxy.ID != proxyID {
		t.Errorf("Wrong Proxy returned %s", proxy.ID)
	}
}

func TestProxiesAssignSubservers(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	proxyID := os.Getenv("NFS_PROXY")
	otherNasSubserver := os.Getenv("SUBSERVER_OTHER_ID")
	integratedSubserver := os.Getenv("SUBSERVER_POWERSCALE_ID")

	// set to empty before testing
	err := client.ProxySetSubservers(context.Background(), proxyID, nil)
	if err != nil {
		t.Fatalf("ProxySetSubservers failed: %v", err)
	}
	subservers, err := client.ProxySubservers(context.Background(), proxyID)
	if err != nil {
		t.Fatalf("ProxySubservers failed: %v", err)
	}
	if len(subservers) > 0 {
		t.Errorf("Proxies not empty before testing.")
	}

	// test setting subservers
	err = client.ProxySetSubservers(context.Background(), proxyID, []string{otherNasSubserver})
	if err != nil {
		t.Fatalf("ProxySetSubservers failed: %v", err)
	}
	subservers, err = client.ProxySubservers(context.Background(), proxyID)
	if err != nil {
		t.Fatalf("ProxySubservers failed: %v", err)
	}
	if subservers[0].ID != otherNasSubserver {
		t.Errorf("Incorrect proxy assigned. %v vs %v", subservers[0].ID, otherNasSubserver)
	}

	// test adding subservers
	err = client.ProxyAddSubservers(context.Background(), proxyID, []string{integratedSubserver})
	if err != nil {
		t.Fatalf("ProxySetSubservers failed: %v", err)
	}
	subservers, err = client.ProxySubservers(context.Background(), proxyID)
	if err != nil {
		t.Fatalf("ProxySubservers failed: %v", err)
	}
	if len(subservers) != 2 {
		t.Errorf("Bad number of proxies returned.")
	}

	// test removing subservers
	err = client.ProxyRemoveSubservers(context.Background(), proxyID, []string{otherNasSubserver})
	if err != nil {
		t.Fatalf("ProxySetSubservers failed: %v", err)
	}
	subservers, err = client.ProxySubservers(context.Background(), proxyID)
	if err != nil {
		t.Fatalf("ProxySubservers failed: %v", err)
	}
	if subservers[0].ID != integratedSubserver {
		t.Errorf("Incorrect proxy assigned. %v vs %v", subservers[0].ID, integratedSubserver)
	}
}
