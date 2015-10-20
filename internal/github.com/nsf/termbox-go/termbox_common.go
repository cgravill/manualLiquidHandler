// internal/github.com/nsf/termbox-go/termbox_common.go: Part of the Antha language
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

package termbox

// private API, common OS agnostic part

type cellbuf struct {
	width  int
	height int
	cells  []Cell
}

func (this *cellbuf) init(width, height int) {
	this.width = width
	this.height = height
	this.cells = make([]Cell, width*height)
}

func (this *cellbuf) resize(width, height int) {
	if this.width == width && this.height == height {
		return
	}

	oldw := this.width
	oldh := this.height
	oldcells := this.cells

	this.init(width, height)
	this.clear()

	minw, minh := oldw, oldh

	if width < minw {
		minw = width
	}
	if height < minh {
		minh = height
	}

	for i := 0; i < minh; i++ {
		srco, dsto := i*oldw, i*width
		src := oldcells[srco : srco+minw]
		dst := this.cells[dsto : dsto+minw]
		copy(dst, src)
	}
}

func (this *cellbuf) clear() {
	for i := range this.cells {
		c := &this.cells[i]
		c.Ch = ' '
		c.Fg = foreground
		c.Bg = background
	}
}

const cursor_hidden = -1

func is_cursor_hidden(x, y int) bool {
	return x == cursor_hidden || y == cursor_hidden
}