/*
  Copyright 2021 Davide Madrisan <davide.madrisan@gmail.com>

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	mattermost "github.com/madrisan/go-mattermost-notify/mattemost"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mattermostChannel, mattermostTeam string
var messageAuthor, messageContent, messageLevel, messageTitle string

// getAttachmentColor returns the HTLM color code to be used for the message assignment
// or "#E0E0D1" if the given level is invalid.
func getAttachmentColor(level string) string {
	var color = map[string]string{
		"critical": "#FF0000",
		"info":     "#E0E0D1",
		"success":  "#00FF00",
		"warning":  "#FF8000",
	}

	if c, found := color[level]; found {
		return c
	}

	return color["info"]
}

// getUserID returns the Mattemost ID associated to the given user
// of to the current user if no user is specified (empty string).
func getUserID(username string) (string, error) {
	if username == "" {
		var found bool
		response, err := mattermost.Get("/users/me")
		if err != nil {
			return "", err
		}
		username, found = response["username"].(string)
		if !found {
			return "", fmt.Errorf("cannot get the Mattermost username of the current user")
		}
	}

	endpoint := fmt.Sprintf("/users/username/%s", username)
	response, err := mattermost.Get(endpoint)
	if err != nil {
		return "", err
	}

	id, found := response["id"]
	if !found {
		return "", fmt.Errorf("cannot get the Mattermost ID of the current user %s", username)
	}

	return id.(string), nil
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a message to a Mattermost channel or user",
	Long: `Post a message to a Mattermost channel or user using its REST APIv4 interface.

Example:
  post rybfbdi9ojy8xxxjjxc88kh3me --author CI --title "Job Status" --message "The job \#BEEF has failed :bug:" --level=critical`,
	Run: func(cmd *cobra.Command, args []string) {
		attachmentColor := getAttachmentColor(messageLevel)

		var mattermostChannelID string

		if strings.HasPrefix(mattermostChannel, "@") {
			userIDFrom, err := getUserID("")
			if err != nil {
				HandleError("%v", err)
			}

			userIDTo, err := getUserID(strings.TrimLeft(mattermostChannel, "@"))
			if err != nil {
				HandleError("%v", err)
			}

			var found bool
			payload, err := json.Marshal([]string{userIDFrom, userIDTo})
			if err != nil {
				HandleError("%v", err)
			}

			response, err := mattermost.Post("/channels/direct", bytes.NewReader(payload))
			if err != nil {
				HandleError("%v", err)
			}
			mattermostChannelID, found = response["id"].(string)
			if !found {
				HandleError("cannot get the Mattermost Channel ID")
			}
		} else {
			mattermostChannelID = mattermostChannel
		}

		type Attachment struct {
			Author string `json:"author_name"`
			Color  string `json:"color"`
			Title  string `json:"title"`
			Text   string `json:"text"`
		}

		type Properties struct {
			Attachments []Attachment `json:"attachments"`
		}

		type Payload struct {
			ID         string     `json:"channel_id"`
			Properties Properties `json:"props"`
		}

		data := Payload{
			ID: mattermostChannelID,
			Properties: Properties{
				[]Attachment{
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
			HandleError("%v", err)
		}
		response, err := mattermost.Post("/posts", bytes.NewReader(payload))
		if err != nil {
			HandleError("%v", err)
		}
		if !viper.GetBool("quiet") {
			prettyPrint(response)
		}
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&messageAuthor,
		"author", "A", "", "author of the message")
	postCmd.Flags().StringVarP(&mattermostChannel,
		"channel", "c", "", "Mattermost channel ID or username. Example: rybfbdi9ojy8xxxjjxc88kh3me or @alice")
	postCmd.Flags().StringVarP(&mattermostTeam, "team", "T", "", "the Mattermost team")
	postCmd.Flags().StringVarP(&messageLevel,
		"level", "l", "info", "criticity level. Can be info (default), success, warning, or critical")
	postCmd.Flags().StringVarP(&messageContent,
		"message", "m", "", "the (markdown-formatted) message to send to the Mattermost channel")
	postCmd.Flags().StringVarP(&messageTitle,
		"title", "t", "", "the title that will precede the text message")

	var requiredFlags = [...]string{
		"author",
		"channel",
		"message",
		"title",
	}

	for _, requiredFlag := range requiredFlags {
		if err := postCmd.MarkFlagRequired(requiredFlag); err != nil {
			HandleError("%v", err)
		}
	}
}
