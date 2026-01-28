package gosmapi

import "testing"

func TestBuildRequestPathEscapesSegments(t *testing.T) {
	path := buildRequestPath([]string{"servers", "space here"}, nil)

	if path != "/servers/space%20here" {
		t.Fatalf("expected escaped path, got %q", path)
	}
}

func TestBuildRequestPathEscapesQueryParams(t *testing.T) {
	params := map[string]string{
		"include": "sub servers",
	}
	path := buildRequestPath([]string{"servers"}, params)

	if path != "/servers?include=sub+servers" {
		t.Fatalf("expected escaped params, got %q", path)
	}
}

func TestBuildRequestPathEncodesMultipleParams(t *testing.T) {
	params := map[string]string{
		"filter":  "name",
		"include": "sub servers",
	}
	path := buildRequestPath([]string{"servers"}, params)

	if path != "/servers?filter=name&include=sub+servers" && path != "/servers?include=sub+servers&filter=name" {
		t.Fatalf("expected encoded params, got %q", path)
	}
}
