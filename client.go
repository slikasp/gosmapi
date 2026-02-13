package gosmapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Holds connection details to the StorageMAP core server
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// Create new connection with below details
//
// address: server FQDN or IP address
//
// token: token string from 'Add Token' in Configuration -> Users -> admin (API Token)
func NewClient(address, token string) *Client {
	httpClient := http.DefaultClient
	apiURL := fmt.Sprintf("http://%s/api", address)
	return &Client{BaseURL: apiURL, Token: token, HTTPClient: httpClient}
}

func (c *Client) makeRequest(
	ctx context.Context,
	method RequestType,
	endpoint Endpoint,
	elements []string,
	parameters map[string]string,
	in interface{},
	out interface{},
) error {
	// build the path out of provided variables
	path := buildRequestPath(endpoint, elements, parameters)

	// make new reader
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	// create the request variable and set headers
	req, err := http.NewRequestWithContext(ctx, string(method), c.BaseURL+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	// debugging input
	// test, _ := httputil.DumpRequest(req, true)
	// fmt.Print(string(test))

	// make the request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	// debugging output
	// fmt.Print(string(data))

	// Handle error responses
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("API error: %d: %s", res.StatusCode, string(data))
	}

	// Handle 204 response with no body
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	// Try unmarshaling the body
	if out != nil {
		err = json.Unmarshal(data, out)
		if err != nil {
			return fmt.Errorf("JSON unmarshal failed: %v", err)
		}
	}

	return nil
}
