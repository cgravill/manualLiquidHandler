// internal/github.com/kylelemons/godebug/pretty/public.go: Part of the Antha language
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
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/antha-lang/manualLiquidHandler/internal/github.com/kylelemons/godebug/diff"
)

// A Config represents optional configuration parameters for formatting.
//
// Some options, notably ShortList, dramatically increase the overhead
// of pretty-printing a value.
type Config struct {
	// Verbosity options
	Compact  bool // One-line output. Overrides Diffable.
	Diffable bool // Adds extra newlines for more easily diffable output.

	// Field and value options
	IncludeUnexported bool // Include unexported fields in output
	PrintStringers    bool // Call String on a fmt.Stringer
	SkipZeroFields    bool // Skip struct fields that have a zero value.

	// Output transforms
	ShortList int // Maximum character length for short lists if nonzero.
}

// Default Config objects
var (
	// CompareConfig is the default configuration used for Compare.
	CompareConfig = &Config{
		Diffable:          true,
		IncludeUnexported: true,
	}

	// DefaultConfig is the default configuration used for all other top-level functions.
	DefaultConfig = &Config{}
)

func (cfg *Config) fprint(buf *bytes.Buffer, vals ...interface{}) {
	for i, val := range vals {
		if i > 0 {
			buf.WriteByte('\n')
		}
		cfg.val2node(reflect.ValueOf(val)).WriteTo(buf, "", cfg)
	}
}

// Print writes the DefaultConfig representation of the given values to standard output.
func Print(vals ...interface{}) {
	DefaultConfig.Print(vals...)
}

// Print writes the configured presentation of the given values to standard output.
func (cfg *Config) Print(vals ...interface{}) {
	fmt.Println(cfg.Sprint(vals...))
}

// Sprint returns a string representation of the given value according to the DefaultConfig.
func Sprint(vals ...interface{}) string {
	return DefaultConfig.Sprint(vals...)
}

// Sprint returns a string representation of the given value according to cfg.
func (cfg *Config) Sprint(vals ...interface{}) string {
	buf := new(bytes.Buffer)
	cfg.fprint(buf, vals...)
	return buf.String()
}

// Fprint writes the representation of the given value to the writer according to the DefaultConfig.
func Fprint(w io.Writer, vals ...interface{}) (n int64, err error) {
	return DefaultConfig.Fprint(w, vals...)
}

// Fprint writes the representation of the given value to the writer according to the cfg.
func (cfg *Config) Fprint(w io.Writer, vals ...interface{}) (n int64, err error) {
	buf := new(bytes.Buffer)
	cfg.fprint(buf, vals...)
	return buf.WriteTo(w)
}

// Compare returns a string containing a line-by-line unified diff of the
// values in got and want, using the CompareConfig.
//
// Each line in the output is prefixed with '+', '-', or ' ' to indicate if it
// should be added to, removed from, or is correct for the "got" value with
// respect to the "want" value.
func Compare(got, want interface{}) string {
	return CompareConfig.Compare(got, want)
}

// Compare returns a string containing a line-by-line unified diff of the
// values in got and want according to the cfg.
func (cfg *Config) Compare(got, want interface{}) string {
	diffCfg := *cfg
	diffCfg.Diffable = true
	return diff.Diff(cfg.Sprint(got), cfg.Sprint(want))
}