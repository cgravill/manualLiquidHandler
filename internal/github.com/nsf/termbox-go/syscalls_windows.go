// internal/github.com/nsf/termbox-go/syscalls_windows.go: Part of the Antha language
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

// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -DUNICODE syscalls.go

package termbox

const (
	foreground_blue          = 0x1
	foreground_green         = 0x2
	foreground_red           = 0x4
	foreground_intensity     = 0x8
	background_blue          = 0x10
	background_green         = 0x20
	background_red           = 0x40
	background_intensity     = 0x80
	std_input_handle         = -0xa
	std_output_handle        = -0xb
	key_event                = 0x1
	mouse_event              = 0x2
	window_buffer_size_event = 0x4
	enable_window_input      = 0x8
	enable_mouse_input       = 0x10
	enable_extended_flags    = 0x80

	vk_f1          = 0x70
	vk_f2          = 0x71
	vk_f3          = 0x72
	vk_f4          = 0x73
	vk_f5          = 0x74
	vk_f6          = 0x75
	vk_f7          = 0x76
	vk_f8          = 0x77
	vk_f9          = 0x78
	vk_f10         = 0x79
	vk_f11         = 0x7a
	vk_f12         = 0x7b
	vk_insert      = 0x2d
	vk_delete      = 0x2e
	vk_home        = 0x24
	vk_end         = 0x23
	vk_pgup        = 0x21
	vk_pgdn        = 0x22
	vk_arrow_up    = 0x26
	vk_arrow_down  = 0x28
	vk_arrow_left  = 0x25
	vk_arrow_right = 0x27
	vk_backspace   = 0x8
	vk_tab         = 0x9
	vk_enter       = 0xd
	vk_esc         = 0x1b
	vk_space       = 0x20

	left_alt_pressed   = 0x2
	left_ctrl_pressed  = 0x8
	right_alt_pressed  = 0x1
	right_ctrl_pressed = 0x4
	shift_pressed      = 0x10

	generic_read            = 0x80000000
	generic_write           = 0x40000000
	console_textmode_buffer = 0x1
)