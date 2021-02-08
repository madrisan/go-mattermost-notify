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

package mattermost

import (
	"bytes"
	"strings"
	"testing"
)

func TestForgeAPIv4URL(t *testing.T) {
	t.Parallel()

	cases := []struct {
		baseUrl     string
		endpoint    string
		urlShouldBe string
	}{
		{
			"http://example.com/mattermost",
			"/users/me",
			"http://example.com/mattermost/api/v4/users/me",
		},
		{
			"http://example.com/mattermost/",
			"/users/me",
			"http://example.com/mattermost/api/v4/users/me",
		},
	}

	t.Run("apiv4_url", func(t *testing.T) {
		t.Parallel()

		for _, tc := range cases {
			t.Run(tc.endpoint, func(t *testing.T) {
				v := forgeAPIv4URL(tc.baseUrl, tc.endpoint)
				if v != tc.urlShouldBe {
					t.Error("For", tc.baseUrl, "and", tc.endpoint,
						"expected", tc.urlShouldBe, "got", v,
					)
				}
			})
		}
	})
}

func TestForgeBearerAuthentication(t *testing.T) {
	accessToken := "2bff151e935e4017a5222076c6f77311"
	shouldBe := "Bearer " + accessToken

	t.Run("bearer_auth", func(t *testing.T) {
		v := forgeBearerAuthentication(accessToken)
		if v != shouldBe {
			t.Error("For", accessToken,
				"expected", shouldBe, "got", v,
			)
		}
	})
}

func TestPrettyPrint(t *testing.T) {
	t.Parallel()

	cases := []struct {
		dataType string
		data     interface{}
		shouldBe string
	}{
		{
			"map",
			map[string]string{"white": "#FFFFFF", "black": "#000000"},
			"{\n  \"black\": \"#000000\",\n  \"white\": \"#FFFFFF\"\n}\n",
		},
		{
			"array",
			[]string{"Italy", "France", "Poland"},
			"[\n  \"Italy\",\n  \"France\",\n  \"Poland\"\n]\n",
		},
	}

	t.Run("pretty_print", func(t *testing.T) {
		t.Parallel()

		for _, tc := range cases {
			t.Run(tc.dataType, func(t *testing.T) {
				var outbuf bytes.Buffer
				err := PrettyPrint(&outbuf, tc.data)
				if err != nil {
					t.Error("For", tc.dataType, "PrettyPrint has failed")
				} else if strings.Compare(outbuf.String(), tc.shouldBe) != 0 {
					t.Error("For", tc.dataType, "expected",
						tc.shouldBe, "got", outbuf.String(),
					)
				}
			})
		}
	})
}
