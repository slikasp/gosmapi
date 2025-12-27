package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "github.com/pauslik/gosmapi/internal/config"
)

// TODO: make request() function and make getRequest() postRequest() call that function instead of repeating code

func getRequest(cfg *Config, link string) ([]byte, error) {
	body := make([]byte, 0)

	// make the request
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return body, err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/vnd.api+json")
	bearer := fmt.Sprintf("Bearer %s", cfg.User.Token)
	req.Header.Set("Authorization", bearer)

	// send the request
	res, err := cfg.HttpClient.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	// read the request body
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func postRequest[T any](cfg *Config, link string, data T) ([]byte, error) {
	body := make([]byte, 0)

	// prepare input data struct for request
	jsonData, err := json.Marshal(data)
	if err != nil {
		return body, err
	}

	// make the request
	req, err := http.NewRequest("POST", link, bytes.NewReader(jsonData))
	if err != nil {
		return body, err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/vnd.api+json")
	bearer := fmt.Sprintf("Bearer %s", cfg.User.Token)
	req.Header.Set("Authorization", bearer)

	// send the request
	res, err := cfg.HttpClient.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	// read the request body
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}
