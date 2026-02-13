package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestNewClient(t *testing.T) {
	godotenv.Load("../.test_env")
	address := os.Getenv("CORE_ADDRESS")
	token := os.Getenv("ADMIN_TOKEN")

	if address == "" || token == "" {
		t.Skip("Test ENV variables not set not set")
	}

	client := gosmapi.NewClient(address, token)
	if client.BaseURL == "" {
		t.Error("expected BaseURL to be set")
	}
	if client.Token != token {
		t.Error("expected Token to match")
	}
	if client.HTTPClient == nil {
		t.Error("expected HTTPClient to be initialized")
	}
}
