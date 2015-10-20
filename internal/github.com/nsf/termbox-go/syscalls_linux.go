// internal/github.com/nsf/termbox-go/syscalls_linux.go: Part of the Antha language
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

import "syscall"

type syscall_Termios syscall.Termios

const (
	syscall_IGNBRK = syscall.IGNBRK
	syscall_BRKINT = syscall.BRKINT
	syscall_PARMRK = syscall.PARMRK
	syscall_ISTRIP = syscall.ISTRIP
	syscall_INLCR  = syscall.INLCR
	syscall_IGNCR  = syscall.IGNCR
	syscall_ICRNL  = syscall.ICRNL
	syscall_IXON   = syscall.IXON
	syscall_OPOST  = syscall.OPOST
	syscall_ECHO   = syscall.ECHO
	syscall_ECHONL = syscall.ECHONL
	syscall_ICANON = syscall.ICANON
	syscall_ISIG   = syscall.ISIG
	syscall_IEXTEN = syscall.IEXTEN
	syscall_CSIZE  = syscall.CSIZE
	syscall_PARENB = syscall.PARENB
	syscall_CS8    = syscall.CS8
	syscall_VMIN   = syscall.VMIN
	syscall_VTIME  = syscall.VTIME

	syscall_TCGETS = syscall.TCGETS
	syscall_TCSETS = syscall.TCSETS
)