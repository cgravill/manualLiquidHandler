// internal/golang.org/x/net/http2/buffer_test.go: Part of the Antha language
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
	"io"
	"reflect"
	"testing"
)

var bufferReadTests = []struct {
	buf      buffer
	read, wn int
	werr     error
	wp       []byte
	wbuf     buffer
}{
	{
		buffer{[]byte{'a', 0}, 0, 1, false, nil},
		5, 1, nil, []byte{'a'},
		buffer{[]byte{'a', 0}, 1, 1, false, nil},
	},
	{
		buffer{[]byte{'a', 0}, 0, 1, true, io.EOF},
		5, 1, io.EOF, []byte{'a'},
		buffer{[]byte{'a', 0}, 1, 1, true, io.EOF},
	},
	{
		buffer{[]byte{0, 'a'}, 1, 2, false, nil},
		5, 1, nil, []byte{'a'},
		buffer{[]byte{0, 'a'}, 2, 2, false, nil},
	},
	{
		buffer{[]byte{0, 'a'}, 1, 2, true, io.EOF},
		5, 1, io.EOF, []byte{'a'},
		buffer{[]byte{0, 'a'}, 2, 2, true, io.EOF},
	},
	{
		buffer{[]byte{}, 0, 0, false, nil},
		5, 0, errReadEmpty, []byte{},
		buffer{[]byte{}, 0, 0, false, nil},
	},
	{
		buffer{[]byte{}, 0, 0, true, io.EOF},
		5, 0, io.EOF, []byte{},
		buffer{[]byte{}, 0, 0, true, io.EOF},
	},
}

func TestBufferRead(t *testing.T) {
	for i, tt := range bufferReadTests {
		read := make([]byte, tt.read)
		n, err := tt.buf.Read(read)
		if n != tt.wn {
			t.Errorf("#%d: wn = %d want %d", i, n, tt.wn)
			continue
		}
		if err != tt.werr {
			t.Errorf("#%d: werr = %v want %v", i, err, tt.werr)
			continue
		}
		read = read[:n]
		if !reflect.DeepEqual(read, tt.wp) {
			t.Errorf("#%d: read = %+v want %+v", i, read, tt.wp)
		}
		if !reflect.DeepEqual(tt.buf, tt.wbuf) {
			t.Errorf("#%d: buf = %+v want %+v", i, tt.buf, tt.wbuf)
		}
	}
}

var bufferWriteTests = []struct {
	buf       buffer
	write, wn int
	werr      error
	wbuf      buffer
}{
	{
		buf: buffer{
			buf: []byte{},
		},
		wbuf: buffer{
			buf: []byte{},
		},
	},
	{
		buf: buffer{
			buf: []byte{1, 'a'},
		},
		write: 1,
		wn:    1,
		wbuf: buffer{
			buf: []byte{0, 'a'},
			w:   1,
		},
	},
	{
		buf: buffer{
			buf: []byte{'a', 1},
			r:   1,
			w:   1,
		},
		write: 2,
		wn:    2,
		wbuf: buffer{
			buf: []byte{0, 0},
			w:   2,
		},
	},
	{
		buf: buffer{
			buf:    []byte{},
			r:      1,
			closed: true,
		},
		write: 5,
		werr:  errWriteClosed,
		wbuf: buffer{
			buf:    []byte{},
			r:      1,
			closed: true,
		},
	},
	{
		buf: buffer{
			buf: []byte{},
		},
		write: 5,
		werr:  errWriteFull,
		wbuf: buffer{
			buf: []byte{},
		},
	},
}

func TestBufferWrite(t *testing.T) {
	for i, tt := range bufferWriteTests {
		n, err := tt.buf.Write(make([]byte, tt.write))
		if n != tt.wn {
			t.Errorf("#%d: wrote %d bytes; want %d", i, n, tt.wn)
			continue
		}
		if err != tt.werr {
			t.Errorf("#%d: error = %v; want %v", i, err, tt.werr)
			continue
		}
		if !reflect.DeepEqual(tt.buf, tt.wbuf) {
			t.Errorf("#%d: buf = %+v; want %+v", i, tt.buf, tt.wbuf)
		}
	}
}