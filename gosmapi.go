package gosmapi

import (
	"fmt"
)

// How this should work:
// con := Connection.New() function that would create a new connection with config from file and http client
// then you can use functions like this:
// con.AddServer(...)
// con.AddSubserver(...)
// con.CreateJob(...)

func Hello() {
	fmt.Println("StorageMAP API Wrapper")
}
