package http

import "net/http"

func GetHeader(url string, header http.Header) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = header
	return http.DefaultClient.Do(req)
}

func Header(data map[string]string) http.Header {
	header := make(http.Header)
	for k, v := range data {
		header.Add(k, v)
	}
	return header
}
