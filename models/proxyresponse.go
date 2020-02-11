package models

import "net/http"

type Proxyresponse struct {
	Header http.Header
	Body   string
}
