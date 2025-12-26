package requests

import (
	"fmt"
	"io"
	"net/http"
)

// TODO: add *config input
func getRequestBody(link string) ([]byte, error) {
	client := http.DefaultClient
	body := make([]byte, 0)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return body, err
	}

	// TODO: make this to be read from a config file and passed in a function as *config
	req.Header.Set("Content-Type", "application/vnd.api+json")
	bearer := fmt.Sprintf("Bearer %s", "qIKgyRASxdrSPEhqW36VDGffINp5b4")
	req.Header.Set("Authorization", bearer)

	res, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}
