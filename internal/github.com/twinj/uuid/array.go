// internal/github.com/twinj/uuid/array.go: Part of the Antha language
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
 * Date: 1/02/14
 * Time: 10:08 AM
 ***************/

const (
	variantIndex = 8
	versionIndex = 6
)

// A clean UUID type for simpler UUID versions
type Array [length]byte

func (Array) Size() int {
	return length
}

func (o Array) Version() int {
	return int(o[versionIndex]) >> 4
}

func (o *Array) setVersion(pVersion int) {
	o[versionIndex] &= 0x0F
	o[versionIndex] |= byte(pVersion) << 4
}

func (o *Array) Variant() byte {
	return variant(o[variantIndex])
}

func (o *Array) setVariant(pVariant byte) {
	setVariant(&o[variantIndex], pVariant)
}

func (o *Array) Unmarshal(pData []byte) {
	copy(o[:], pData)
}

func (o *Array) Bytes() []byte {
	return o[:]
}

func (o Array) String() string {
	return formatter(&o, format)
}

func (o Array) Format(pFormat string) string {
	return formatter(&o, pFormat)
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant equivalent in the array to ReservedRFC4122.
func (o *Array) setRFC4122Variant() {
	o[variantIndex] &= 0x3F
	o[variantIndex] |= ReservedRFC4122
}

// Marshals the UUID bytes into a slice
func (o *Array) MarshalBinary() ([]byte, error) {
	return o.Bytes(), nil
}

// Un-marshals the data bytes into the UUID.
func (o *Array) UnmarshalBinary(pData []byte) error {
	return UnmarshalBinary(o, pData)
}