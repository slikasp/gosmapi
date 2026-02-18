package test

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestCreatejobstatuses(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	createjobstatuses, err := client.Createjobstatuses(context.Background())
	if err != nil {
		t.Fatalf("Createjobstatuses failed: %v", err)
	}

	if len(createjobstatuses) == 0 {
		t.Errorf("No Createjobstatuses returned.")
	}
}

func TestCreatejobstatus(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	createjobstatusError := "3780cbf3-4fad-4249-b747-5424f2a90071"
	createjobstatusSuccess := "39890380-5f28-4452-97f2-cef09371e567"

	cjse, err := client.Createjobstatus(context.Background(), createjobstatusError)
	if err != nil {
		t.Fatalf("Createjobstatus failed: %v", err)
	}
	if cjse.Type != "createjobstatuses" {
		t.Errorf("Wrong type returned %s", cjse.Type)
	}

	cjss, err := client.Createjobstatus(context.Background(), createjobstatusSuccess)
	if err != nil {
		t.Fatalf("Createjobstatus failed: %v", err)
	}
	if cjss.Type != "jobs" {
		t.Errorf("Wrong type returned %s", cjss.ID)
	}
}
