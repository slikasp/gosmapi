package gosmapi

// TODO: finish job creation before making this

type Createjobstatus struct {
	ID         string     `json:"id"`
	Type       ObjectType `json:"type"`
	Attributes struct {
		Status JobStatus `json:"status"`
		Errors struct {
		} `json:"errors"`
	} `json:"attributes"`
}

type singleCreatejobstatusOutput struct {
	Data Createjobstatus `json:"data"`
}
