// internal/golang.org/x/net/http2/pipe.go: Part of the Antha language
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

// Copyright 2014 The Go Authors.
// See https://code.google.com/p/go/source/browse/CONTRIBUTORS
// Licensed under the same terms as Go itself:
// https://code.google.com/p/go/source/browse/LICENSE

package http2

import (
	"sync"
)

type pipe struct {
	b buffer
	c sync.Cond
	m sync.Mutex
}

// Read waits until data is available and copies bytes
// from the buffer into p.
func (r *pipe) Read(p []byte) (n int, err error) {
	r.c.L.Lock()
	defer r.c.L.Unlock()
	for r.b.Len() == 0 && !r.b.closed {
		r.c.Wait()
	}
	return r.b.Read(p)
}

// Write copies bytes from p into the buffer and wakes a reader.
// It is an error to write more data than the buffer can hold.
func (w *pipe) Write(p []byte) (n int, err error) {
	w.c.L.Lock()
	defer w.c.L.Unlock()
	defer w.c.Signal()
	return w.b.Write(p)
}

func (c *pipe) Close(err error) {
	c.c.L.Lock()
	defer c.c.L.Unlock()
	defer c.c.Signal()
	c.b.Close(err)
}