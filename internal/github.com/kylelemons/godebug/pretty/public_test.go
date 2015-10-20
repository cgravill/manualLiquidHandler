// internal/github.com/kylelemons/godebug/pretty/public_test.go: Part of the Antha language
// Copyright (C) 2015 The Antha authors. All rights reserved.
// 
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
// 
// For more information relating to the software or licensing issues please
// contact license@antha-lang.org or write to the Antha team c/o 
// Synthace Ltd. The London Bioscience Innovation Centre
// 2 Royal College St, London NW1 0NH UK

// Copyright 2013 Google Inc.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pretty

import (
	"testing"
)

func TestDiff(t *testing.T) {
	type example struct {
		Name    string
		Age     int
		Friends []string
	}

	tests := []struct {
		desc      string
		got, want interface{}
		diff      string
	}{
		{
			desc: "basic struct",
			got: example{
				Name: "Zaphd",
				Age:  42,
				Friends: []string{
					"Ford Prefect",
					"Trillian",
					"Marvin",
				},
			},
			want: example{
				Name: "Zaphod",
				Age:  42,
				Friends: []string{
					"Ford Prefect",
					"Trillian",
				},
			},
			diff: ` {
- Name:    "Zaphd",
+ Name:    "Zaphod",
  Age:     42,
  Friends: [
            "Ford Prefect",
            "Trillian",
-           "Marvin",
           ],
 }`,
		},
	}

	for _, test := range tests {
		if got, want := Compare(test.got, test.want), test.diff; got != want {
			t.Errorf("%s:", test.desc)
			t.Errorf("  got:  %q", got)
			t.Errorf("  want: %q", want)
		}
	}
}

func TestSkipZeroFields(t *testing.T) {
	type example struct {
		Name    string
		Species string
		Age     int
		Friends []string
	}

	tests := []struct {
		desc      string
		got, want interface{}
		diff      string
	}{
		{
			desc: "basic struct",
			got: example{
				Name:    "Zaphd",
				Species: "Betelgeusian",
				Age:     42,
			},
			want: example{
				Name:    "Zaphod",
				Species: "Betelgeusian",
				Age:     42,
				Friends: []string{
					"Ford Prefect",
					"Trillian",
					"",
				},
			},
			diff: ` {
- Name:    "Zaphd",
+ Name:    "Zaphod",
  Species: "Betelgeusian",
  Age:     42,
+ Friends: [
+           "Ford Prefect",
+           "Trillian",
+           "",
+          ],
 }`,
		},
	}

	cfg := *CompareConfig
	cfg.SkipZeroFields = true

	for _, test := range tests {
		if got, want := cfg.Compare(test.got, test.want), test.diff; got != want {
			t.Errorf("%s:", test.desc)
			t.Errorf("  got:  %q", got)
			t.Errorf("  want: %q", want)
		}
	}
}