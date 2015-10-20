// internal/github.com/mattn/go-runewidth/runewidth_test.go: Part of the Antha language
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

package runewidth

import (
	"testing"
)

var runewidthtests = []struct {
	in  rune
	out int
}{
	{'世', 2},
	{'界', 2},
	{'ｾ', 1},
	{'ｶ', 1},
	{'ｲ', 1},
	{'☆', 2}, // double width in ambiguous
	{'\x00', 0},
	{'\x01', 1},
	{'\u0300', 0},
}

func TestRuneWidth(t *testing.T) {
	c := NewCondition()
	c.EastAsianWidth = true
	for _, tt := range runewidthtests {
		if out := c.RuneWidth(tt.in); out != tt.out {
			t.Errorf("Width(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

var isambiguouswidthtests = []struct {
	in  rune
	out bool
}{
	{'世', false},
	{'■', true},
	{'界', false},
	{'○', true},
	{'㈱', false},
	{'①', true},
	{'②', true},
	{'③', true},
	{'④', true},
	{'⑤', true},
	{'⑥', true},
	{'⑦', true},
	{'⑧', true},
	{'⑨', true},
	{'⑩', true},
	{'⑪', true},
	{'⑫', true},
	{'⑬', true},
	{'⑭', true},
	{'⑮', true},
	{'⑯', true},
	{'⑰', true},
	{'⑱', true},
	{'⑲', true},
	{'⑳', true},
	{'☆', true},
}

func TestIsAmbiguousWidth(t *testing.T) {
	for _, tt := range isambiguouswidthtests {
		if out := IsAmbiguousWidth(tt.in); out != tt.out {
			t.Errorf("IsAmbiguousWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

var stringwidthtests = []struct {
	in  string
	out int
}{
	{"■㈱の世界①", 12},
	{"スター☆", 8},
}

func TestStringWidth(t *testing.T) {
	c := NewCondition()
	c.EastAsianWidth = true
	for _, tt := range stringwidthtests {
		if out := c.StringWidth(tt.in); out != tt.out {
			t.Errorf("StringWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}

func TestStringWidthInvalid(t *testing.T) {
	s := "こんにちわ\x00世界"
	if out := StringWidth(s); out != 14 {
		t.Errorf("StringWidth(%q) = %v, want %v", s, out, 14)
	}
}

func TestTruncate(t *testing.T) {
	s := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお"
	expected := "あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおお..."

	if out := Truncate(s, 80, "..."); out != expected {
		t.Errorf("Truncate(%q) = %v, want %v", s, out, expected)
	}
}

func TestTruncateNoNeeded(t *testing.T) {
	s := "あいうえおあい"
	expected := "あいうえおあい"

	if out := Truncate(s, 80, "..."); out != expected {
		t.Errorf("Truncate(%q) = %v, want %v", s, out, expected)
	}
}

var isneutralwidthtests = []struct {
	in  rune
	out bool
}{
	{'→', false},
	{'┊', false},
	{'┈', false},
	{'～', false},
	{'└', false},
	{'⣀', true},
	{'⣀', true},
}

func TestIsNeutralWidth(t *testing.T) {
	for _, tt := range isneutralwidthtests {
		if out := IsNeutralWidth(tt.in); out != tt.out {
			t.Errorf("IsNeutralWidth(%q) = %v, want %v", tt.in, out, tt.out)
		}
	}
}