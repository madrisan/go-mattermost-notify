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
	"fmt"
	"os"

	mattermost "github.com/madrisan/go-mattermost-notify/mattermost"
	"github.com/spf13/cobra"

	"github.com/madrisan/go-mattermost-notify/config"
)

// getCmd represents the get CLI command.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send a Get query to Mattermost",
	Long: `Send a Get query to Mattermost using its REST APIv4 interface.

See the Mattermost API documentation:
  https://api.mattermost.com/`,
	Example: `  get /bots
  get /channels
  get /users/me`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("An endpoint must be specified in the command-line arguments")
		}
		var opts = config.Options{
			ConnectionTimeout: mattermostConnectionTimeout,
			SkipTLSVerify:     mattermostSkipTLSVerify,
		}
		response, err := mattermostGet(args[0], opts)
		if err != nil {
			return err
		}

		mattermost.PrettyPrint(os.Stdout, response)
		return nil
	},
}

// init initializes the post command flags.
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&mattermostSkipTLSVerify,
		"insecure", "i", false, "ignore SSL/TLS certificate check")
}
