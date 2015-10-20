// internal/golang.org/x/net/http2/buffer.go: Part of the Antha language
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
	"errors"
)

// buffer is an io.ReadWriteCloser backed by a fixed size buffer.
// It never allocates, but moves old data as new data is written.
type buffer struct {
	buf    []byte
	r, w   int
	closed bool
	err    error // err to return to reader
}

var (
	errReadEmpty   = errors.New("read from empty buffer")
	errWriteClosed = errors.New("write on closed buffer")
	errWriteFull   = errors.New("write on full buffer")
)

// Read copies bytes from the buffer into p.
// It is an error to read when no data is available.
func (b *buffer) Read(p []byte) (n int, err error) {
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	if b.closed && b.r == b.w {
		err = b.err
	} else if b.r == b.w && n == 0 {
		err = errReadEmpty
	}
	return n, err
}

// Len returns the number of bytes of the unread portion of the buffer.
func (b *buffer) Len() int {
	return b.w - b.r
}

// Write copies bytes from p into the buffer.
// It is an error to write more data than the buffer can hold.
func (b *buffer) Write(p []byte) (n int, err error) {
	if b.closed {
		return 0, errWriteClosed
	}

	// Slide existing data to beginning.
	if b.r > 0 && len(p) > len(b.buf)-b.w {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}

	// Write new data.
	n = copy(b.buf[b.w:], p)
	b.w += n
	if n < len(p) {
		err = errWriteFull
	}
	return n, err
}

// Close marks the buffer as closed. Future calls to Write will
// return an error. Future calls to Read, once the buffer is
// empty, will return err.
func (b *buffer) Close(err error) {
	if !b.closed {
		b.closed = true
		b.err = err
	}
}