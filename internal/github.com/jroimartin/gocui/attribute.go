// internal/github.com/jroimartin/gocui/attribute.go: Part of the Antha language
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

// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocui

import "github.com/antha-lang/manualLiquidHandler/internal/github.com/nsf/termbox-go"

// Attributes can be combined using bitwise OR (|). Note that it is not
// possible to combine multiple color attributes.
type Attribute termbox.Attribute

// Color attributes.
const (
	ColorDefault Attribute = Attribute(termbox.ColorDefault)
	ColorBlack             = Attribute(termbox.ColorBlack)
	ColorRed               = Attribute(termbox.ColorRed)
	ColorGreen             = Attribute(termbox.ColorGreen)
	ColorYellow            = Attribute(termbox.ColorYellow)
	ColorBlue              = Attribute(termbox.ColorBlue)
	ColorMagenta           = Attribute(termbox.ColorMagenta)
	ColorCyan              = Attribute(termbox.ColorCyan)
	ColorWhite             = Attribute(termbox.ColorWhite)
)

// Text style attributes.
const (
	AttrBold      Attribute = Attribute(termbox.AttrBold)
	AttrUnderline           = Attribute(termbox.AttrUnderline)
	AttrReverse             = Attribute(termbox.AttrReverse)
)