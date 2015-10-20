// internal/github.com/kylelemons/godebug/pretty/reflect_test.go: Part of the Antha language
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
	"reflect"
	"testing"
	"time"
)

func TestVal2nodeDefault(t *testing.T) {
	tests := []struct {
		desc string
		raw  interface{}
		want node
	}{
		{
			"nil",
			(*int)(nil),
			rawVal("nil"),
		},
		{
			"string",
			"zaphod",
			stringVal("zaphod"),
		},
		{
			"slice",
			[]string{"a", "b"},
			list{stringVal("a"), stringVal("b")},
		},
		{
			"map",
			map[string]string{
				"zaphod": "beeblebrox",
				"ford":   "prefect",
			},
			keyvals{
				{"ford", stringVal("prefect")},
				{"zaphod", stringVal("beeblebrox")},
			},
		},
		{
			"map of [2]int",
			map[[2]int]string{
				[2]int{-1, 2}: "school",
				[2]int{0, 0}:  "origin",
				[2]int{1, 3}:  "home",
			},
			keyvals{
				{"[-1,2]", stringVal("school")},
				{"[0,0]", stringVal("origin")},
				{"[1,3]", stringVal("home")},
			},
		},
		{
			"struct",
			struct{ Zaphod, Ford string }{"beeblebrox", "prefect"},
			keyvals{
				{"Zaphod", stringVal("beeblebrox")},
				{"Ford", stringVal("prefect")},
			},
		},
		{
			"int",
			3,
			rawVal("3"),
		},
	}

	for _, test := range tests {
		if got, want := DefaultConfig.val2node(reflect.ValueOf(test.raw)), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %#v, want %#v", test.desc, got, want)
		}
	}
}

func TestVal2node(t *testing.T) {
	tests := []struct {
		desc string
		raw  interface{}
		cfg  *Config
		want node
	}{
		{
			"struct default",
			struct{ Zaphod, Ford, foo string }{"beeblebrox", "prefect", "BAD"},
			DefaultConfig,
			keyvals{
				{"Zaphod", stringVal("beeblebrox")},
				{"Ford", stringVal("prefect")},
			},
		},
		{
			"struct w/ IncludeUnexported",
			struct{ Zaphod, Ford, foo string }{"beeblebrox", "prefect", "GOOD"},
			&Config{
				IncludeUnexported: true,
			},
			keyvals{
				{"Zaphod", stringVal("beeblebrox")},
				{"Ford", stringVal("prefect")},
				{"foo", stringVal("GOOD")},
			},
		},
		{
			"time default",
			struct{ Date time.Time }{time.Unix(1234567890, 0).UTC()},
			DefaultConfig,
			keyvals{
				{"Date", keyvals{}}, // empty struct, it has unexported fields
			},
		},
		{
			"time w/ PrintStringers",
			struct{ Date time.Time }{time.Unix(1234567890, 0).UTC()},
			&Config{
				PrintStringers: true,
			},
			keyvals{
				{"Date", stringVal("2009-02-13 23:31:30 +0000 UTC")},
			},
		},
	}

	for _, test := range tests {
		if got, want := test.cfg.val2node(reflect.ValueOf(test.raw)), test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %#v, want %#v", test.desc, got, want)
		}
	}
}