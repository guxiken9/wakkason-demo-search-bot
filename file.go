package main

import "github.com/go-resty/resty/v2"

func downloadFile(url string) (*resty.Response, error) {

	c := resty.New()

	// ファイルをダウンロード
	resp, err := c.R().Get(url)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return resp, err
}
