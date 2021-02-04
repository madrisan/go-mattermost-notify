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
	"testing"
)

//func getAttachmentColor(level string) string {
//                 "critical": "#FF0000",
//                 "info":     "#E0E0D1",
//                 "success":  "#00FF00",
//                 "warning":  "#FF8000",

func TestGetAttachmentColor(t *testing.T) {
	t.Parallel()

	cases := []struct {
		level         string
		colorShouldBe string
	}{
		{
			"critical",
			"#FF0000",
		},
		{
			"info",
			"#E0E0D1",
		},
		{
			"success",
			"#00FF00",
		},
		{
			"warning",
			"#FF8000",
		},
		{
			"unknown",
			"#E0E0D1",
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
