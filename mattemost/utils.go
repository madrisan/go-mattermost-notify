/*
  Copyright 2021-2022 Davide Madrisan <davide.madrisan@gmail.com>

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
	"strings"

	"github.com/spf13/viper"
)

// forgeAPIv4URL returns the Mattermost APIv4 URL for the given endpoint.
func forgeAPIv4URL(baseURL, endpoint string) string {
	var url = fmt.Sprintf("%s/api/v4/%s",
		strings.TrimRight(baseURL, "/"),
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
func getURL() (string, error) {
	baseURL := viper.GetString("url")
	if baseURL == "" {
		return "", fmt.Errorf("the Mattermost URL has not been set")
	}
	return baseURL, nil
}

// CreateMsgPayload forges the payload containing the message to be posted to Mattermost
func CreateMsgPayload(
	attachmentColor,
	mattermostChannelID,
	messageAuthor, messageContent, messageTitle string) ([]byte, error) {

	type MsgAttachment struct {
		Author string `json:"author_name"`
		Color  string `json:"color"`
		Title  string `json:"title"`
		Text   string `json:"text"`
	}

	type MsgProperties struct {
		Attachments []MsgAttachment `json:"attachments"`
	}

	// MsgPayload is used to create the JSON payload used when posting a message to Mattermost.
	type MsgPayload struct {
		ID         string        `json:"channel_id"`
		Properties MsgProperties `json:"props"`
	}

	data := MsgPayload{
		ID: mattermostChannelID,
		Properties: MsgProperties{
			[]MsgAttachment{
				{
					Author: messageAuthor,
					Color:  attachmentColor,
					Title:  messageTitle,
					Text:   messageContent,
				},
			},
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// PrettyPrint prints the result of a Mattermost query in a pretty JSON format.
func PrettyPrint(w io.Writer, v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Fprintln(w, string(b))
	}
	return
}
