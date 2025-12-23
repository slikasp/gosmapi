package config

import (
	"strings"
	"testing"
)

func TestStructs(t *testing.T) {
	adm := user{
		name:     "admin",
		token:    "admin_token",
		userRole: ADMIN,
	}

	testCore := server{
		address:    "0.0.0.0",
		serverRole: CORE,
	}

	cfg := config{
		user: adm,
		core: testCore,
	}

	if strings.Compare("0.0.0.0", cfg.core.address) != 0 {
		t.Error("cfg.core.address mismatch\n")
		t.Fail()
	}
	if strings.Compare("admin_token", cfg.user.token) != 0 {
		t.Error("cfg.user.token mismatch\n")
		t.Fail()
	}
}

func TestFunctions(t *testing.T) {
	cfg := config{}

	cfg.SetAdmin("admin", "admin_token")
	cfg.SetCore("0.0.0.0")

	if strings.Compare("0.0.0.0", cfg.core.address) != 0 {
		t.Error("cfg.core.address mismatch\n")
		t.Fail()
	}
	if strings.Compare("admin_token", cfg.user.token) != 0 {
		t.Error("cfg.user.token mismatch\n")
		t.Fail()
	}
}
