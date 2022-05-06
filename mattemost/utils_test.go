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
	"fmt"
	"strings"
	"testing"
)

func TestCreateMsgPayload(t *testing.T) {
	cases := []struct {
		attachmentColor     string
		mattermostChannelID string
		messageAuthor       string
		messageContent      string
		messageTitle        string
		shouldBe            []byte
	}{
		{
			"#FF0000",
			"azerty123",
			"Author",
			"This is a message",
			"Pretty simple Message",
			[]byte{
				123, 34, 99, 104, 97, 110, 110, 101, 108, 95, 105, 100, 34, 58, 34, 97,
				122, 101, 114, 116, 121, 49, 50, 51, 34, 44, 34, 112, 114, 111, 112, 115,
				34, 58, 123, 34, 97, 116, 116, 97, 99, 104, 109, 101, 110, 116, 115, 34,
				58, 91, 123, 34, 97, 117, 116, 104, 111, 114, 95, 110, 97, 109, 101, 34,
				58, 34, 65, 117, 116, 104, 111, 114, 34, 44, 34, 99, 111, 108, 111, 114,
				34, 58, 34, 35, 70, 70, 48, 48, 48, 48, 34, 44, 34, 116, 105, 116,
				108, 101, 34, 58, 34, 80, 114, 101, 116, 116, 121, 32, 115, 105, 109, 112,
				108, 101, 32, 77, 101, 115, 115, 97, 103, 101, 34, 44, 34, 116, 101, 120,
				116, 34, 58, 34, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 109, 101,
				115, 115, 97, 103, 101, 34, 125, 93, 125, 125,
			},
		}, {
			"#FFFFFF",
			"querty456",
			"Me",
			"Another message",
			"Message endeinf with an emotion :tada:",
			[]byte{
				123, 34, 99, 104, 97, 110, 110, 101, 108, 95, 105, 100, 34, 58, 34, 113,
				117, 101, 114, 116, 121, 52, 53, 54, 34, 44, 34, 112, 114, 111, 112, 115,
				34, 58, 123, 34, 97, 116, 116, 97, 99, 104, 109, 101, 110, 116, 115, 34,
				58, 91, 123, 34, 97, 117, 116, 104, 111, 114, 95, 110, 97, 109, 101, 34,
				58, 34, 77, 101, 34, 44, 34, 99, 111, 108, 111, 114, 34, 58, 34, 35,
				70, 70, 70, 70, 70, 70, 34, 44, 34, 116, 105, 116, 108, 101, 34, 58,
				34, 77, 101, 115, 115, 97, 103, 101, 32, 101, 110, 100, 101, 105, 110, 102,
				32, 119, 105, 116, 104, 32, 97, 110, 32, 101, 109, 111, 116, 105, 111, 110,
				32, 58, 116, 97, 100, 97, 58, 34, 44, 34, 116, 101, 120, 116, 34, 58,
				34, 65, 110, 111, 116, 104, 101, 114, 32, 109, 101, 115, 115, 97, 103,
				101, 34, 125, 93, 125, 125,
			},
		},
	}

	t.Run("payload", func(t *testing.T) {
		for i, tc := range cases {
			t.Run(fmt.Sprint(i), func(t *testing.T) {
				v, err := CreateMsgPayload(
					tc.attachmentColor,
					tc.mattermostChannelID,
					tc.messageAuthor,
					tc.messageContent,
					tc.messageTitle,
				)
				if err != nil {
					t.Error("For case", i, "CreateMsgPayload has failed", err)
				}
				if !bytes.Equal(v, tc.shouldBe) {
					t.Error("For case", i,
						"expected", tc.shouldBe, "got", v,
					)
				}
			})
		}
	})

}

func TestForgeAPIv4URL(t *testing.T) {
	t.Parallel()

	cases := []struct {
		baseURL     string
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
				v := forgeAPIv4URL(tc.baseURL, tc.endpoint)
				if v != tc.urlShouldBe {
					t.Error("For", tc.baseURL, "and", tc.endpoint,
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
