// internal/github.com/bradfitz/http2/errors.go: Part of the Antha language
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

import "fmt"

// An ErrCode is an unsigned 32-bit error code as defined in the HTTP/2 spec.
type ErrCode uint32

const (
	ErrCodeNo                 ErrCode = 0x0
	ErrCodeProtocol           ErrCode = 0x1
	ErrCodeInternal           ErrCode = 0x2
	ErrCodeFlowControl        ErrCode = 0x3
	ErrCodeSettingsTimeout    ErrCode = 0x4
	ErrCodeStreamClosed       ErrCode = 0x5
	ErrCodeFrameSize          ErrCode = 0x6
	ErrCodeRefusedStream      ErrCode = 0x7
	ErrCodeCancel             ErrCode = 0x8
	ErrCodeCompression        ErrCode = 0x9
	ErrCodeConnect            ErrCode = 0xa
	ErrCodeEnhanceYourCalm    ErrCode = 0xb
	ErrCodeInadequateSecurity ErrCode = 0xc
	ErrCodeHTTP11Required     ErrCode = 0xd
)

var errCodeName = map[ErrCode]string{
	ErrCodeNo:                 "NO_ERROR",
	ErrCodeProtocol:           "PROTOCOL_ERROR",
	ErrCodeInternal:           "INTERNAL_ERROR",
	ErrCodeFlowControl:        "FLOW_CONTROL_ERROR",
	ErrCodeSettingsTimeout:    "SETTINGS_TIMEOUT",
	ErrCodeStreamClosed:       "STREAM_CLOSED",
	ErrCodeFrameSize:          "FRAME_SIZE_ERROR",
	ErrCodeRefusedStream:      "REFUSED_STREAM",
	ErrCodeCancel:             "CANCEL",
	ErrCodeCompression:        "COMPRESSION_ERROR",
	ErrCodeConnect:            "CONNECT_ERROR",
	ErrCodeEnhanceYourCalm:    "ENHANCE_YOUR_CALM",
	ErrCodeInadequateSecurity: "INADEQUATE_SECURITY",
	ErrCodeHTTP11Required:     "HTTP_1_1_REQUIRED",
}

func (e ErrCode) String() string {
	if s, ok := errCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error code 0x%x", uint32(e))
}

// ConnectionError is an error that results in the termination of the
// entire connection.
type ConnectionError ErrCode

func (e ConnectionError) Error() string { return fmt.Sprintf("connection error: %s", ErrCode(e)) }

// StreamError is an error that only affects one stream within an
// HTTP/2 connection.
type StreamError struct {
	StreamID uint32
	Code     ErrCode
}

func (e StreamError) Error() string {
	return fmt.Sprintf("stream error: stream ID %d; %v", e.StreamID, e.Code)
}

// 6.9.1 The Flow Control Window
// "If a sender receives a WINDOW_UPDATE that causes a flow control
// window to exceed this maximum it MUST terminate either the stream
// or the connection, as appropriate. For streams, [...]; for the
// connection, a GOAWAY frame with a FLOW_CONTROL_ERROR code."
type goAwayFlowError struct{}

func (goAwayFlowError) Error() string { return "connection exceeded flow control window size" }