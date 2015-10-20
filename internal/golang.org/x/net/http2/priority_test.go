// internal/golang.org/x/net/http2/priority_test.go: Part of the Antha language
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

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// See https://code.google.com/p/go/source/browse/CONTRIBUTORS
// Licensed under the same terms as Go itself:
// https://code.google.com/p/go/source/browse/LICENSE

package http2

import (
	"testing"
)

func TestPriority(t *testing.T) {
	// A -> B
	// move A's parent to B
	streams := make(map[uint32]*stream)
	a := &stream{
		parent: nil,
		weight: 16,
	}
	streams[1] = a
	b := &stream{
		parent: a,
		weight: 16,
	}
	streams[2] = b
	adjustStreamPriority(streams, 1, PriorityParam{
		Weight:    20,
		StreamDep: 2,
	})
	if a.parent != b {
		t.Errorf("Expected A's parent to be B")
	}
	if a.weight != 20 {
		t.Errorf("Expected A's weight to be 20; got %d", a.weight)
	}
	if b.parent != nil {
		t.Errorf("Expected B to have no parent")
	}
	if b.weight != 16 {
		t.Errorf("Expected B's weight to be 16; got %d", b.weight)
	}
}

func TestPriorityExclusiveZero(t *testing.T) {
	// A B and C are all children of the 0 stream.
	// Exclusive reprioritization to any of the streams
	// should bring the rest of the streams under the
	// reprioritized stream
	streams := make(map[uint32]*stream)
	a := &stream{
		parent: nil,
		weight: 16,
	}
	streams[1] = a
	b := &stream{
		parent: nil,
		weight: 16,
	}
	streams[2] = b
	c := &stream{
		parent: nil,
		weight: 16,
	}
	streams[3] = c
	adjustStreamPriority(streams, 3, PriorityParam{
		Weight:    20,
		StreamDep: 0,
		Exclusive: true,
	})
	if a.parent != c {
		t.Errorf("Expected A's parent to be C")
	}
	if a.weight != 16 {
		t.Errorf("Expected A's weight to be 16; got %d", a.weight)
	}
	if b.parent != c {
		t.Errorf("Expected B's parent to be C")
	}
	if b.weight != 16 {
		t.Errorf("Expected B's weight to be 16; got %d", b.weight)
	}
	if c.parent != nil {
		t.Errorf("Expected C to have no parent")
	}
	if c.weight != 20 {
		t.Errorf("Expected C's weight to be 20; got %d", b.weight)
	}
}

func TestPriorityOwnParent(t *testing.T) {
	streams := make(map[uint32]*stream)
	a := &stream{
		parent: nil,
		weight: 16,
	}
	streams[1] = a
	b := &stream{
		parent: a,
		weight: 16,
	}
	streams[2] = b
	adjustStreamPriority(streams, 1, PriorityParam{
		Weight:    20,
		StreamDep: 1,
	})
	if a.parent != nil {
		t.Errorf("Expected A's parent to be nil")
	}
	if a.weight != 20 {
		t.Errorf("Expected A's weight to be 20; got %d", a.weight)
	}
	if b.parent != a {
		t.Errorf("Expected B's parent to be A")
	}
	if b.weight != 16 {
		t.Errorf("Expected B's weight to be 16; got %d", b.weight)
	}

}