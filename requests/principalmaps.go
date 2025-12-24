package requests

type principalmaps struct {
	Data []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"attributes"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
