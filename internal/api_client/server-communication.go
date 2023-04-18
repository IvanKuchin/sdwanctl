package apiclient

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

func sendRequestToServer(cfg *configreader.Config, req_type, urlPath, content_type, req_body string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(req_type, urlPath, strings.NewReader(req_body))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Type", content_type)
	req.Header.Add("X-XSRF-TOKEN", cfg.XSRFToken)

	cookie := http.Cookie{Name: "JSESSIONID", Value: cfg.BearerToken, MaxAge: 30}
	req.AddCookie(&cookie)

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

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	return content, nil
}
