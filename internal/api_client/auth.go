package apiclient

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

// authenticates user and returns bearer token
func getBearerToken(cfg *configreader.Config) (string, error) {
	if cfg.BearerToken != "" {
		return cfg.BearerToken, nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Set up the HTTP client
	client := &http.Client{Transport: tr}

	data := url.Values{}
	data.Set("j_username", cfg.Login)
	data.Set("j_password", cfg.Password)
	req_body := data.Encode()

	req, err := http.NewRequest("POST", cfg.Server+"/j_security_check", strings.NewReader(req_body))
	if err != nil {
		log.Println("Error creating request:", err)
		return "", err
	}

	// Set any headers you need to include
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Collect cookies from the response
	cookies := resp.Cookies()

	// Check cookies
	bearer_token := ""
	for _, cookie := range cookies {
		if cookie.Name == "JSESSIONID" {
			bearer_token = cookie.Value
		}
	}

	if bearer_token == "" {
		error_message := "Error: authentication failed"
		log.Print("ERROR:", error_message)
		return "", errors.New(error_message)
	}

	// Print response headers
	// for k, v := range resp.Header {
	// 	fmt.Printf("Header field %q, Value %q\n", k, v)
	// }

	// Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("Error reading response body:", err)
	// 	return "", err
	// }

	return bearer_token, err
}

func getXSRFToken(cfg *configreader.Config) (string, error) {
	body, err := sendRequestToServer(cfg, "GET", cfg.Server+"/dataservice/client/token", "application/json", "")
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Login(cfg *configreader.Config) error {
	bearer_token, err := getBearerToken(cfg)
	if err != nil {
		return err
	}

	cfg.BearerToken = bearer_token

	xsrf_token, err := getXSRFToken(cfg)
	if err != nil {
		return err
	}

	cfg.XSRFToken = xsrf_token

	return nil
}

func Logout(cfg *configreader.Config) error {
	_, err := sendRequestToServer(cfg, "GET", cfg.Server+"/logout", "application/json", "")
	if err != nil {
		return err
	}

	return nil
}
