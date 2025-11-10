package client

import (
	"io"
	"log"
	"net/http"
)

type GetResult struct {
	IsError  bool
	Response []byte
}

func Get(skipCertVerify bool, url string) (*GetResult, error) {
	c := getNewClient(skipCertVerify)

	r, err := c.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	return &GetResult{IsError: r.StatusCode != http.StatusOK, Response: body}, nil
}
