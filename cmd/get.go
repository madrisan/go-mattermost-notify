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
	"github.com/spf13/cobra"
)

// getCmd represents the get CLI command.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send a Get query to Mattermost",
	Long: `Send a Get query to Mattermost using its REST APIv4 interface.

See the Mattermost API documentation:
  https://api.mattermost.com/

Example:
  get /bots
  get /channels
  get /users/me`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			handleError("%v", "An endpoint must be specified in the command-line arguments")
		}
		response, err := mattermostGet(args[0])
		if err != nil {
			handleError("%v", err)
		}
		prettyPrint(response)
	},
}

// init initializes the post command flags.
func init() {
	rootCmd.AddCommand(getCmd)
}
