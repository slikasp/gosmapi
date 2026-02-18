package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/pauslik/gosmapi"
)

// move it somewhere
func ptrTo[T any](v T) *T {
	p := v
	return &p
}

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

	jobID := os.Getenv("MULTI_JOB_ID")

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
			Protocol:       gosmapi.NFS,
			NfsConstraints: &nfsVersion,
		},
	}
	attributes := gosmapi.CreateJobAttributes{
		JobType:         string(gosmapi.NasMigration),
		SourcePath:      "/ifs/home/api/postman/source/s1",
		DestinationPath: "/ifs/home/api/postman/destination/d1",
		Options:         options,
	}

	createjobstatus, err := client.CreateJob(context.Background(), source, destination, attributes)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}

	if createjobstatus.Attributes.Status != gosmapi.Failed {
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
			Protocol:       gosmapi.NFS,
			NfsConstraints: &nfsVersion,
		},
	}
	attributes := gosmapi.CreateJobAttributes{
		JobType:         string(gosmapi.NasMigration),
		SourcePath:      "/ifs/home/api/postman/source/s2",
		DestinationPath: "/ifs/home/api/postman/destination/d2",
		Options:         options,
	}

	createjobstatus, err := client.CreateJob(context.Background(), source, destination, attributes)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}

	if createjobstatus.Attributes.Status != gosmapi.Queued {
		t.Errorf("Unexpected job status %s. ID: %s", createjobstatus.Attributes.Status, createjobstatus.ID)
	}
}

func TestCreateJobWithOptions(t *testing.T) {
	godotenv.Load("../.test_env")
	client := gosmapi.NewClient(os.Getenv("CORE_ADDRESS"), os.Getenv("ADMIN_TOKEN"))

	source := os.Getenv("SUBSERVER_POWERSCALE_ID")
	destination := os.Getenv("SUBSERVER_OTHER_ID")

	nfsVersion := gosmapi.NfsV3
	config := gosmapi.JobConfiguration{
		Protocol:       gosmapi.Multiprotocol,
		NfsConstraints: &nfsVersion,
	}

	da := gosmapi.MD5Algo
	coc := gosmapi.MigratedOnlyCoc
	sym := gosmapi.CreateOverNfsSymlink
	age := gosmapi.SixHourAge
	root := gosmapi.ConvertToExplicit
	ops := gosmapi.NoDeletes

	options := gosmapi.JobOptions{
		Configuration:         &config,
		DigestAlgorithm:       &da,
		CocMode:               &coc,
		SmbSymlinkTargetMode:  &sym,
		MinimumAge:            &age,
		CopyRootDirMode:       &root,
		OperationRestrictions: &ops,
		VerifySourceAfterCopy: ptrTo(false),
		PreserveAccessTime:    ptrTo(true),
	}

	attributes := gosmapi.CreateJobAttributes{
		JobType:         string(gosmapi.NasMigration),
		SourcePath:      "/ifs/home/api/postman/source/s3",
		DestinationPath: "/ifs/home/api/postman/destination/d3",
		Options:         options,
	}

	createjobstatus, err := client.CreateJob(context.Background(), source, destination, attributes)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}
	if createjobstatus.Attributes.Status != "QUEUED" {
		t.Errorf("Unexpected job status %s. ID: %s", createjobstatus.Attributes.Status, createjobstatus.ID)
	}

	time.Sleep(2 * time.Second)
	// Use createjobstatus with createjobstatusID if SUCCESS to get the jobID
	tempjob, err := client.Createjobstatus(context.Background(), createjobstatus.ID)
	if err != nil {
		t.Fatalf("Createjobstatus failed: %v", err)
	}

	job, err := client.Job(context.Background(), tempjob.ID)
	if err != nil {
		t.Fatalf("Job failed: %v", err)
	}
	if *job.Attributes.Options.PreserveAccessTime != true {
		t.Fatalf("Options not set for job %v", job.ID)
	}
}
