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

package cmd

import (
	"io"
	"os"
	"strings"
	"testing"

	mattermost "github.com/madrisan/go-mattermost-notify/mattemost"
	"github.com/spf13/viper"
)

func TestGetAttachmentColor(t *testing.T) {
	t.Parallel()

	cases := []struct {
		level         string
		colorShouldBe string
	}{
		{
			"critical",
			COLOR_CRITICAL,
		},
		{
			"info",
			COLOR_INFO,
		},
		{
			"success",
			COLOR_SUCCESS,
		},
		{
			"warning",
			COLOR_WARNING,
		},
		{
			"unknown",
			COLOR_DEFAULT,
		},
	}

	t.Run("get_attachment_color", func(t *testing.T) {
		t.Parallel()

		for _, tc := range cases {
			t.Run(tc.level, func(t *testing.T) {
				v := getAttachmentColor(tc.level)
				if v != tc.colorShouldBe {
					t.Error("For", tc.level,
						"expected", tc.colorShouldBe, "got", v,
					)
				}
			})
		}
	})
}

func TestGetKV(t *testing.T) {
	t.Parallel()

	cases := []struct {
		data     interface{}
		key      string
		shouldBe string
	}{
		{
			map[string]interface{}{"white": "#FFFFFF", "black": "#000000"},
			"white",
			"#FFFFFF",
		},
		{
			map[string]interface{}{"white": "#FFFFFF", "black": "#000000"},
			"black",
			"#000000",
		},
	}

	t.Run("get_kv", func(t *testing.T) {
		t.Parallel()

		for _, tc := range cases {
			t.Run(tc.key, func(t *testing.T) {
				value, err := getKV(tc.data, tc.key)
				if err != nil {
					t.Error("For", tc.key, "getKV has failed:", err)
				} else if strings.Compare(value, tc.shouldBe) != 0 {
					t.Error("For", tc.key, "expected",
						tc.shouldBe, "got", value,
					)
				}
			})
		}
	})
}

func TestEnvVariables(t *testing.T) {
	oldMattermostGet := mattermost.Get
	oldMattermostPost := mattermost.Post
	defer func() {
		mattermostGet = oldMattermostGet
		mattermostPost = oldMattermostPost
	}()

	mattermostGet = func(endpoint string) (interface{}, error) {
		return map[string]interface{}{}, nil
	}
	mattermostPost = func(endpoint string, payload io.Reader) (interface{}, error) {
		return map[string]interface{}{}, nil
	}

	envVariables := []struct {
		name  string
		value string
		vname string
	}{
		{
			"MATTERMOST_URL",
			"http://example.com/mattermost",
			"url",
		},
		{
			"MATTERMOST_ACCESS_TOKEN",
			"2bff151e935e4017a5222076c6f77311",
			"access-token",
		},
	}

	for _, tc := range envVariables {
		if err := os.Setenv(tc.name, tc.value); err != nil {
			t.Skip("cannot set the environment variable", tc.name)
			return
		}
	}

	args := []string{
		"post",
		"--author", "that's me",
		"--channel", "7trmbhd8xg9tmiagqfx1fzhhjo",
		"--message", "this is a dumb text",
		"--title", "testing viber",
	}
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("The rootCmd.Execute has failed: %s", err)
	}

	for _, tc := range envVariables {
		v := viper.GetString(tc.vname)
		if v != tc.value {
			t.Fatalf("For \"%s\" expected \"%s\" got \"%s\"",
				tc.name, tc.value, v,
			)
		}
	}
}
