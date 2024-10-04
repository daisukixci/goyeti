package goyeti

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/yeti-platform/goyeti/models"
)

var _ models.ClientI = (*YetiClient)(nil)

type YetiClient struct {
	models.ClientConfig
	url     string
	client  *http.Client
	apikey  string
	token   string
	headers http.Header
}

func NewYetiClient(host string, port int, apikey string) *YetiClient {
	return &YetiClient{
		ClientConfig: models.ClientConfig{
			Host:    host,
			Port:    port,
			Timeout: time.Second * 5,
			Proto:   "https",
		},
		apikey: apikey,
	}
}

func parseResponse(res *http.Response, receiver *map[string]interface{}) error {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not parse response body: %w", err)
	}
	err = json.Unmarshal(resBody, receiver)
	if err != nil {
		return fmt.Errorf("could not unmarshal response body: %w", err)
	}
	return nil
}

func (c *YetiClient) doRequest(method string, endpoint string, header http.Header, data string) (*http.Response, error) {
	var body io.Reader

	if data != "" {
		body = strings.NewReader(data)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.url, endpoint), body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	for k, v := range header {
		req.Header.Add(k, v[0])
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return res, fmt.Errorf("received bad status: %s", res.Status)
	}
	return res, nil
}

func (c *YetiClient) getToken() (string, error) {
	var authResponse models.ServerResponseAuth

	res, err := c.Query("api/v2/auth/api-token", http.MethodPost, "")
	if err != nil {
		return "", err
	}

	err = mapstructure.Decode(res, &authResponse)
	if err != nil {
		return "", err
	}

	return authResponse.AccessToken, nil
}

func (c *YetiClient) checkAuth(token string) (bool, error) {
	res, err := c.Query("api/v2/auth/me", http.MethodGet, "")
	if err != nil {
		return false, err
	}
	return len(res) != 0, nil
}

func (c *YetiClient) Init() error {
	c.url = fmt.Sprintf("%s://%s:%d", c.Proto, c.Host, c.Port)
	c.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
			Proxy:           http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		},
		Timeout: c.Timeout,
	}
	c.headers = http.Header{
		"Content-Type":  []string{"application/json"},
		"x-yeti-apikey": []string{c.apikey},
	}
	token, err := c.getToken()
	if err != nil {
		return fmt.Errorf("could not authenticate: %s", err)
	}
	c.token = token
	c.headers.Add("authorization", fmt.Sprintf("Bearer %s", c.token))

	authed, err := c.checkAuth(c.token)
	if err != nil || !authed {
		return fmt.Errorf("could not authenticate: %s", err)
	}

	return nil
}

func (c *YetiClient) ObservablesSearch(search string, searchType string) (models.ServerResponseObservables, error) {
	var result models.ServerResponseObservables

	body := fmt.Sprintf(`{"query": {"value": "%s"},  "type": "%s",  "sorting": [],   "count": 50,  "page": 0}`, search, searchType)
	res, err := c.Query("api/v2/observables/search", http.MethodPost, body)
	if err != nil {
		return result, err
	}
	err = mapstructure.Decode(res, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *YetiClient) Query(endpoint string, method string, data string) (map[string]interface{}, error) {
	var result map[string]interface{}

	res, err := c.doRequest(method, endpoint, c.headers, data)
	if err != nil || res.StatusCode != http.StatusOK {
		return result, err
	}
	err = parseResponse(res, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *YetiClient) Close() error {
	defer c.client.CloseIdleConnections()
	res, err := c.doRequest(http.MethodPost, "api/v2/auth/logout", c.headers, "")
	if err != nil || res.StatusCode != http.StatusOK {
		return fmt.Errorf("could not logout: %s %s", res.Status, err)
	}
	return nil
}