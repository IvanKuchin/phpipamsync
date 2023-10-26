package api_client

import (
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type http_headers map[string]string

func sendRequestToServer(req_type, urlPath string, headers http_headers, req_body string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(req_type, urlPath, strings.NewReader(req_body))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Error sending request:", resp.Status)
		return nil, errors.New(resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	return content, nil
}
