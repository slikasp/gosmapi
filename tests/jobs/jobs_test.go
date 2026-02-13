package tests

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

func TestJobs(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	jobs, err := client.Jobs(context.Background())
	if err != nil {
		t.Fatalf("Jobs failed: %v", err)
	}

	if len(jobs) == 0 {
		t.Errorf("No jobs returned.")
	}
}

func TestJob(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	jobID := os.Getenv("NFS_JOB_ID")

	job, err := client.Job(context.Background(), jobID)
	if err != nil {
		t.Fatalf("Job failed: %v", err)
	}

	if job.ID != jobID {
		t.Errorf("Wrong job returned %s", job.ID)
	}
}

func TestCreateJobFail(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	source := os.Getenv("SUBSERVER_POWERSCALE_ID")
	destination := os.Getenv("SUBSERVER_OTHER_ID")

	nfsVersion := gosmapi.NfsV3
	options := gosmapi.JobOptions{
		Configuration: &gosmapi.JobConfiguration{
			Protocol:       gosmapi.NfsProtocol,
			NfsConstraints: &nfsVersion,
		},
	}
	attributes := gosmapi.CreateJobAttributes{
		JobType:         string(gosmapi.NasMigrationJob),
		SourcePath:      "/ifs/home/api/postman/source",
		DestinationPath: "/ifs/home/api/postman/destination",
		Options:         options,
	}

	createjobstatus, err := client.CreateJob(context.Background(), source, destination, attributes)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}

	if createjobstatus.Attributes.Status != "FAILED" {
		t.Errorf("Unexpected job status %s", createjobstatus.Attributes.Status)
	}
}

func TestCreateJobSuccess(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	source := os.Getenv("SUBSERVER_POWERSCALE_ID")
	destination := os.Getenv("SUBSERVER_OTHER_ID")

	nfsVersion := gosmapi.NfsV3
	options := gosmapi.JobOptions{
		Configuration: &gosmapi.JobConfiguration{
			Protocol:       gosmapi.NfsProtocol,
			NfsConstraints: &nfsVersion,
		},
	}
	attributes := gosmapi.CreateJobAttributes{
		JobType:         string(gosmapi.NasMigrationJob),
		SourcePath:      "/ifs/home/api/postman/s2",
		DestinationPath: "/ifs/home/api/postman/d2",
		Options:         options,
	}

	createjobstatus, err := client.CreateJob(context.Background(), source, destination, attributes)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}

	if createjobstatus.Attributes.Status != "QUEUED" {
		t.Errorf("Unexpected job status %s. ID: %s", createjobstatus.Attributes.Status, createjobstatus.ID)
	}
}
