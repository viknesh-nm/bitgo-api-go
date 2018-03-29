package bitgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseAPIURL = "https://www.bitgo.com/api/"
	v1         = "v1/"
	v2         = "v2/"
)

// Config -
type Config struct {
	AccessToken string
	BaseURL     string
}

// New -
func New(token string) *Config {
	return &Config{
		AccessToken: token,
		BaseURL:     baseAPIURL,
	}
}

// SendHTTPRequest -
func (c *Config) SendHTTPRequest(method, path string, data interface{}) error {
	method = strings.ToUpper(method)

	req, err := http.NewRequest(method, path, strings.NewReader(""))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		err := json.Unmarshal(contents, &data)
		if err != nil {
			return err
		}
		return nil
	}

	errg := Error{}

	err = json.Unmarshal(contents, &errg)
	if err != nil {
		return err
	}

	return errg
}
