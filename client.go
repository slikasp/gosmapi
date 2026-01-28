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

func (c *Client) makeRequest(ctx context.Context, method, path string, in interface{}, out interface{}) error {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
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

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)

	// debugging output
	// fmt.Print(string(data))

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("api error: %d: %s", res.StatusCode, string(data))
	}

	if out != nil {
		return json.Unmarshal(data, out)
	}
	return nil
}
