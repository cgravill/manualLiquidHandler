// internal/github.com/nsf/termbox-go/syscalls.go: Part of the Antha language
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

// +build ignore

package termbox

/*
#include <termios.h>
#include <sys/ioctl.h>
*/
import "C"

type syscall_Termios C.struct_termios

const (
	syscall_IGNBRK = C.IGNBRK
	syscall_BRKINT = C.BRKINT
	syscall_PARMRK = C.PARMRK
	syscall_ISTRIP = C.ISTRIP
	syscall_INLCR  = C.INLCR
	syscall_IGNCR  = C.IGNCR
	syscall_ICRNL  = C.ICRNL
	syscall_IXON   = C.IXON
	syscall_OPOST  = C.OPOST
	syscall_ECHO   = C.ECHO
	syscall_ECHONL = C.ECHONL
	syscall_ICANON = C.ICANON
	syscall_ISIG   = C.ISIG
	syscall_IEXTEN = C.IEXTEN
	syscall_CSIZE  = C.CSIZE
	syscall_PARENB = C.PARENB
	syscall_CS8    = C.CS8
	syscall_VMIN   = C.VMIN
	syscall_VTIME  = C.VTIME

	// on darwin change these to (on *bsd too?):
	// C.TIOCGETA
	// C.TIOCSETA
	syscall_TCGETS = C.TCGETS
	syscall_TCSETS = C.TCSETS
)