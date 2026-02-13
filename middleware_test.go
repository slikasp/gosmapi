package gosmapi

import "testing"

func TestBuildRequestPathEscapesSegments(t *testing.T) {
	path := buildRequestPath(ServersEndpoint, []string{"space here"}, nil)

	if path != "/servers/space%20here" {
		t.Fatalf("expected escaped path, got %q", path)
	}
}

func TestBuildRequestPathEscapesQueryParams(t *testing.T) {
	params := map[string]string{
		"include": "sub servers",
	}
	path := buildRequestPath(ServersEndpoint, nil, params)

	if path != "/servers?include=sub+servers" {
		t.Fatalf("expected escaped params, got %q", path)
	}
}

func TestBuildRequestPathEncodesMultipleParams(t *testing.T) {
	params := map[string]string{
		"filter":  "name",
		"include": "sub servers",
	}
	path := buildRequestPath(ServersEndpoint, nil, params)

	if path != "/servers?filter=name&include=sub+servers" && path != "/servers?include=sub+servers&filter=name" {
		t.Fatalf("expected encoded params, got %q", path)
	}
}

func TestBuildRequestPathNoParams(t *testing.T) {
	path := buildRequestPath(ServersEndpoint, nil, nil)

	if path != "/servers" {
		t.Fatalf("expected encoded params, got %q", path)
	}
}
