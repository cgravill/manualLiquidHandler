// internal/github.com/bradfitz/http2/gotrack_test.go: Part of the Antha language
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
	"fmt"
	"strings"
	"testing"
)

func TestGoroutineLock(t *testing.T) {
	DebugGoroutines = true
	g := newGoroutineLock()
	g.check()

	sawPanic := make(chan interface{})
	go func() {
		defer func() { sawPanic <- recover() }()
		g.check() // should panic
	}()
	e := <-sawPanic
	if e == nil {
		t.Fatal("did not see panic from check in other goroutine")
	}
	if !strings.Contains(fmt.Sprint(e), "wrong goroutine") {
		t.Errorf("expected on see panic about running on the wrong goroutine; got %v", e)
	}
}