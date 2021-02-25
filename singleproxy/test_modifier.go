package singleproxy

import (
	"net/http"
)

type Modifier struct {}

func (t *Modifier) ModifyRequest(req *http.Request) error {
	//fmt.Println("in test modifier")
	return nil
}
