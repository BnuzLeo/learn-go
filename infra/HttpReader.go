package infra

import (
	"net/http"
	"net/http/httputil"
)

type HttpReader struct {
}

func (h HttpReader) Download(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Can't get this url : " + url
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return "dump response failed."
	}
	return string(result)
}
