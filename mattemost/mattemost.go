/*
  Copyright 2021 Davide Madrisan <davide.madrisan@gmail.com>

  Licensed under the Mozilla Public License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      https://www.mozilla.org/en-US/MPL/2.0/

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

// Package mattermost implements the API v4 calls to Mattemost.
package mattermost

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// forgeAPIv4URL returns the Mattermost APIv4 URL for the given endpoint.
func forgeAPIv4URL(baseUrl, endpoint string) string {
	var url = fmt.Sprintf("%s/api/v4/%s",
		strings.TrimRight(baseUrl, "/"),
		strings.TrimLeft(endpoint, "/"))
	return url
}

// forgeBearerAuthentication returns the string to be sent to Mattermost for a Bearer Authentication.
func forgeBearerAuthentication(accessToken string) string {
	return "Bearer " + accessToken
}

// getAccessToken returns the Mattermost token set at command-line or via the environment variable MATTERMOST_ACCESS_TOKEN.
func getAccessToken() (string, error) {
	accessToken := viper.GetString("access-token")
	if accessToken == "" {
		return "", fmt.Errorf("the Mattermost Access Token has not been set")
	}
	return accessToken, nil
}

// getUrl returns the Mattermost URL set at command-line or via the environment variable MATTERMOST_URL.
func getUrl() (string, error) {
	baseUrl := viper.GetString("url")
	if baseUrl == "" {
		return "", fmt.Errorf("the Mattermost URL has not been set")
	}
	return baseUrl, nil
}

// queryAPIv4 makes a query to Mattermost using its REST API v4.
func queryAPIv4(method, endpoint string, payload io.Reader) (interface{}, error) {
	baseUrl, err := getUrl()
	if err != nil {
		return nil, err
	}

	accessToken, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	var bearer = forgeBearerAuthentication(accessToken)
	var url = forgeAPIv4URL(baseUrl, endpoint)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		statusCodeText := http.StatusText(response.StatusCode)
		return nil, fmt.Errorf("the HTTP query to %s has ended with a %d (\"%s\") code",
			url, response.StatusCode, statusCodeText)
	}

	// Read body
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Get makes a query of type GET to Mattermost.
func Get(endpoint string) (interface{}, error) {
	response, err := queryAPIv4(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Post makes a query of type POST to Mattermost.
func Post(endpoint string, payload io.Reader) (interface{}, error) {
	response, err := queryAPIv4(http.MethodPost, endpoint, payload)
	if err != nil {
		return nil, err
	}

	return response, nil
}
