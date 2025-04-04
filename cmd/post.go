/*
  Copyright 2021-2024 Davide Madrisan <d.madrisan@proton.me>

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
	"os"
	"reflect"
	"strings"
	"time"

	mattermost "github.com/madrisan/go-mattermost-notify/mattermost"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/madrisan/go-mattermost-notify/config"
)

var (
	// mattermostChannel contains the Mattermost Channel ID.
	mattermostChannel string
	// mattermostConnectionTimeout defines the maximum time in seconds allowed for Mattermost connections.
	mattermostConnectionTimeout time.Duration
	// mattermostSkipTLSVerify tells if the SSL/TLS certificate check must be ignored or not.
	mattermostSkipTLSVerify bool
	// mattermostTeam contains the Mattermost Team.
	mattermostTeam string
	// messageAuthor contains the author of the Mattermost post to be sent.
	messageAuthor string
	// messageContent contains the text message of the Mattermost post.
	messageContent string
	// messageLevel defines the criticity of the post message.
	// Can be "info" (the default), "success", "warning", or "critical".
	messageLevel string
	// messageTitle contains the title of the post message to be sent.
	messageTitle string
	// mattermostGet contains the pointer to the Get function in the mattermost package.
	// It's used to easily mockup the Mattermost server in the unit tests.
	mattermostGet = mattermost.Get
	// mattermostPost contains the pointer to the Post function in the mattermost package.
	// It's used to easily mockup the Mattermost server in the unit tests.
	mattermostPost = mattermost.Post
)

// The HTML colors used in the post message attachment.
const (
	colorCritical = "#FF0000" // The color code for critical messages.
	colorInfo     = "#E0E0D1" // The color code for informational messages.
	colorSuccess  = "#00FF00" // The color code for successful messages.
	colorWarning  = "#FF8000" // The color code for warning messages.
	colorDefault  = "#E0E0D1" // The default color.
)

// getAttachmentColor returns the HTLM color code to be used for the message assignment
// or COLOR_DEFAULT if the given level is invalid.
func getAttachmentColor(level string) string {
	var color = map[string]string{
		"critical": colorCritical,
		"info":     colorInfo,
		"success":  colorSuccess,
		"warning":  colorWarning,
	}

	if c, found := color[level]; found {
		return c
	}

	return colorDefault
}

// getKV returns the value of key in the JSON response data.
func getKV(response interface{}, key string) (string, error) {
	switch response.(type) {
	case map[string]interface{}:
		data := response.(map[string]interface{})
		value, found := data[key].(string)
		if !found {
			return "", fmt.Errorf("no such key: \"%s\"", key)
		}
		return value, nil
	}

	return "", fmt.Errorf("unexpected response format from Mattermost (%s)",
		reflect.TypeOf(response))
}

// getLoggedUsername returns the username of the logged Mattermost user.
func getLoggedUsername() (string, error) {
	var opts = config.Options{
		ConnectionTimeout: mattermostConnectionTimeout,
		SkipTLSVerify:     mattermostSkipTLSVerify,
	}

	response, err := mattermostGet("/users/me", opts)
	if err != nil {
		return "", err
	}

	username, err := getKV(response, "username")
	if err != nil {
		return "", err
	}

	return username, nil
}

// getLoggedUserID returns the Mattermost ID of the logged user.
func getLoggedUserID() (string, error) {
	username, err := getLoggedUsername()
	if err != nil {
		return "", err
	}

	id, err := getUserID(username)
	if err != nil {
		return "", err
	}

	return id, nil
}

// getUserID returns the Mattemost ID associated to the given user.
func getUserID(username string) (string, error) {
	var opts = config.Options{
		ConnectionTimeout: mattermostConnectionTimeout,
		SkipTLSVerify:     mattermostSkipTLSVerify,
	}

	endpoint := fmt.Sprintf("/users/username/%s", username)
	response, err := mattermostGet(endpoint, opts)
	if err != nil {
		return "", err
	}
	id, err := getKV(response, "id")
	if err != nil {
		return "", fmt.Errorf("cannot get the Mattermost ID of the current user %s: %v", username, err)
	}
	return id, nil
}

// postCmd represents the post CLI command.
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a message to a Mattermost channel or user",
	Long:  `Post a message to a Mattermost channel or user using its REST APIv4 interface.`,
	Example: `  post -c rybfbdi9ojy8xxxjjxc88kh3me -A CI -t "Job Status" -m "The job \#BEEF has failed :bug:" -l critical
  post -c @alice -A CI -t "Job Status" -m "The job \#BEEF ended successfully :tada:" -l success -s 3s`,
	RunE: func(cmd *cobra.Command, args []string) error {
		attachmentColor := getAttachmentColor(messageLevel)

		var mattermostChannelID string
		var opts = config.Options{
			ConnectionTimeout: mattermostConnectionTimeout,
			SkipTLSVerify:     mattermostSkipTLSVerify,
		}

		if opts.SkipTLSVerify {
			fmt.Fprintln(os.Stderr, os.Args[0], "Warning: SSL/TLS certificate check is disabled!")
		}

		if strings.HasPrefix(mattermostChannel, "@") {
			userIDFrom, err := getLoggedUserID()
			if err != nil {
				return err
			}

			userIDTo, err := getUserID(strings.TrimLeft(mattermostChannel, "@"))
			if err != nil {
				return err
			}

			payload, err := json.Marshal([]string{userIDFrom, userIDTo})
			if err != nil {
				return err
			}

			response, err := mattermostPost("/channels/direct", bytes.NewReader(payload), opts)
			if err != nil {
				return err
			}

			mattermostChannelID, err = getKV(response, "id")
			if err != nil {
				return fmt.Errorf("cannot get the Mattermost direct channel ID %v", err)
			}
		} else {
			mattermostChannelID = mattermostChannel
		}

		payload, err := mattermost.CreateMsgPayload(
			attachmentColor,
			mattermostChannelID,
			messageAuthor, messageContent, messageTitle)
		if err != nil {
			return err
		}

		response, err := mattermostPost("/posts", bytes.NewReader(payload), opts)
		if err != nil {
			return err
		}
		if !viper.GetBool("quiet") {
			mattermost.PrettyPrint(os.Stdout, response)
		}

		return nil
	},
}

// init initializes the post command flags.
func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&messageAuthor,
		"author", "A", "", "author of the message")
	postCmd.Flags().StringVarP(&mattermostChannel,
		"channel", "c", "", "Mattermost channel ID or username. Example: rybfbdi9ojy8xxxjjxc88kh3me or @alice")
	postCmd.Flags().BoolVarP(&mattermostSkipTLSVerify,
		"insecure", "i", false, "ignore SSL/TLS certificate check")
	postCmd.Flags().StringVarP(&messageLevel,
		"level", "l", "info", "criticity level. Can be info, success, warning, or critical")
	postCmd.Flags().StringVarP(&messageContent,
		"message", "m", "", "the (markdown-formatted) message to send to the Mattermost channel")
	postCmd.Flags().StringVarP(&mattermostTeam, "team", "T", "", "the Mattermost team")
	postCmd.Flags().DurationVarP(&mattermostConnectionTimeout,
		"timeout", "s", 10*time.Second, "the maximum time in seconds allowed for a Mattermost connection")
	postCmd.Flags().StringVarP(&messageTitle,
		"title", "t", "", "the title that will precede the text message")

	var requiredFlags = [...]string{
		"author",
		"channel",
		"message",
		"title",
	}

	for _, requiredFlag := range requiredFlags {
		err := postCmd.MarkFlagRequired(requiredFlag)
		checkErr(err)
	}
}
