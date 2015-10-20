// internal/github.com/nsf/termbox-go/syscalls_darwin_amd64.go: Part of the Antha language
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
// cgo -godefs syscalls.go

package termbox

type syscall_Termios struct {
	Iflag     uint64
	Oflag     uint64
	Cflag     uint64
	Lflag     uint64
	Cc        [20]uint8
	Pad_cgo_0 [4]byte
	Ispeed    uint64
	Ospeed    uint64
}

const (
	syscall_IGNBRK = 0x1
	syscall_BRKINT = 0x2
	syscall_PARMRK = 0x8
	syscall_ISTRIP = 0x20
	syscall_INLCR  = 0x40
	syscall_IGNCR  = 0x80
	syscall_ICRNL  = 0x100
	syscall_IXON   = 0x200
	syscall_OPOST  = 0x1
	syscall_ECHO   = 0x8
	syscall_ECHONL = 0x10
	syscall_ICANON = 0x100
	syscall_ISIG   = 0x80
	syscall_IEXTEN = 0x400
	syscall_CSIZE  = 0x300
	syscall_PARENB = 0x1000
	syscall_CS8    = 0x300
	syscall_VMIN   = 0x10
	syscall_VTIME  = 0x11

	syscall_TCGETS = 0x40487413
	syscall_TCSETS = 0x80487414
)