package config

import (
	"strings"
	"testing"
)

func TestStructs(t *testing.T) {
	adm := SMuser{
		Name:     "admin",
		Token:    "admin_token",
		UserRole: ADMIN,
	}

	testCore := SMserver{
		Address:    "0.0.0.0",
		ServerRole: CORE,
	}

	cfg := Config{
		User: adm,
		Core: testCore,
	}

	if strings.Compare("0.0.0.0", cfg.Core.Address) != 0 {
		t.Error("cfg.core.address mismatch\n")
		t.Fail()
	}
	if strings.Compare("admin_token", cfg.User.Token) != 0 {
		t.Error("cfg.user.token mismatch\n")
		t.Fail()
	}
}

func TestFunctions(t *testing.T) {
	cfg := Config{}

	cfg.SetAdmin("admin", "admin_token")
	cfg.SetCore("0.0.0.0")

	if strings.Compare("0.0.0.0", cfg.Core.Address) != 0 {
		t.Error("cfg.core.address mismatch\n")
		t.Fail()
	}
	if strings.Compare("admin_token", cfg.User.Token) != 0 {
		t.Error("cfg.user.token mismatch\n")
		t.Fail()
	}
}
