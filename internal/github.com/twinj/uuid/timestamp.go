// internal/github.com/twinj/uuid/timestamp.go: Part of the Antha language
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

package uuid

/****************
 * Date: 14/02/14
 * Time: 7:46 PM
 ***************/

import (
	"time"
)

const (
	// A tick is 100 ns
	ticksPerSecond = 10000000

	// Difference between
	gregorianToUNIXOffset uint64 = 0x01B21DD213814000

	// set the following to the number of 100ns ticks of the actual
	// resolution of your system's clock
	idsPerTimestamp = 1024
)

var (
	lastTimestamp    Timestamp
	idsThisTimestamp = idsPerTimestamp
)

// **********************************************  Timestamp

type Timestamp uint64

// TODO Create c version same as package runtime and time
func Now() (sec int64, nsec int32) {
	t := time.Now()
	sec = t.Unix()
	nsec = int32(t.Nanosecond())
	return
}

// Converts Unix formatted time to RFC4122 UUID formatted times
// UUID UTC base time is October 15, 1582.
// Unix base time is January 1, 1970.
// Converts time to 100 nanosecond ticks since epoch
// There are 1000000000 nanoseconds in a second,
// 1000000000 / 100 = 10000000 tiks per second
func timestamp() Timestamp {
	sec, nsec := Now()
	return Timestamp(uint64(sec)*ticksPerSecond +
		uint64(nsec)/100 + gregorianToUNIXOffset)
}

func (o Timestamp) Unix() time.Time {
	t := uint64(o) - gregorianToUNIXOffset
	return time.Unix(0, int64(t*100))
}

// Get time as 60-bit 100ns ticks since UUID epoch.
// Compensate for the fact that real clock resolution is
// less than 100ns.
func currentUUIDTimestamp() Timestamp {
	var timeNow Timestamp
	for {
		timeNow = timestamp()

		// if clock reading changed since last UUID generated
		if lastTimestamp != timeNow {
			// reset count of UUIDs with this timestamp
			idsThisTimestamp = 0
			lastTimestamp = timeNow
			break
		}
		if idsThisTimestamp < idsPerTimestamp {
			idsThisTimestamp++
			break
		}
		// going too fast for the clock; spin
	}
	// add the count of UUIDs to low order bits of the clock reading
	return timeNow + Timestamp(idsThisTimestamp)
}