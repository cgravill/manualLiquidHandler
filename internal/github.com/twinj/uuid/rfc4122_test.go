// internal/github.com/twinj/uuid/rfc4122_test.go: Part of the Antha language
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
 * Date: 16/02/14
 * Time: 11:29 AM
 ***************/

import (
	"fmt"
	"net/url"
	"testing"
)

var (
	goLang Name = "https://google.com/golang.org?q=golang"
)

const (
	generate = 1000000
)

func TestUUID_NewV1(t *testing.T) {
	u := NewV1()
	if u.Version() != 1 {
		t.Errorf("Expected correct version %d, but got %d", 1, u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected RFC4122 variant %x, but got %x", ReservedRFC4122, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
}

func TestUUID_NewV1Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV1()
	}
}

// Tests NewV3
func TestUUID_NewV3(t *testing.T) {
	u := NewV3(NamespaceURL, goLang)
	if u.Version() != 3 {
		t.Errorf("Expected correct version %d, but got %d", 3, u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected RFC4122 variant %x, but got %x", ReservedRFC4122, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV3(NamespaceURL, ur)
	if !Equal(u2, u) {
		t.Errorf("Expected UUIDs generated with same namespace and name to equal but got: %s and %s", u2, u)
	}

	// Different NS same name MUST NOT be equal
	u3 := NewV3(NamespaceDNS, ur)
	if Equal(u3, u) {
		t.Errorf("Expected UUIDs generated with different namespace and same name to be different but got: %s and %s", u3, u)
	}

	// Same NS different name MUST NOT be equal
	u4 := NewV3(NamespaceURL, u)
	if Equal(u4, u) {
		t.Errorf("Expected UUIDs generated with the same namespace and different names to be different but got: %s and %s", u4, u)
	}

	ids := []UUID{
		u, u2, u3, u4,
	}
	for j, id := range ids {
		i := NewV3(NamespaceURL, NewName(string(j), id))
		if Equal(id, i) {
			t.Errorf("Expected UUIDs generated with the same namespace and different names to be different but got: %s and %s", id, i)
		}
	}
}

func TestUUID_NewV3Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV3(NamespaceDNS, goLang)
	}
}

func TestUUID_NewV4(t *testing.T) {
	u := NewV4()
	if u.Version() != 4 {
		t.Errorf("Expected correct version %d, but got %d", 4, u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected RFC4122 variant %x, but got %x", ReservedRFC4122, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
}

func TestUUID_NewV4Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV4()
	}
}

// Tests NewV5
func TestUUID_NewV5(t *testing.T) {
	u := NewV5(NamespaceURL, goLang)
	if u.Version() != 5 {
		t.Errorf("Expected correct version %d, but got %d", 5, u.Version())
	}
	if u.Variant() != ReservedRFC4122 {
		t.Errorf("Expected RFC4122 variant %x, but got %x", ReservedRFC4122, u.Variant())
	}
	if !parseUUIDRegex.MatchString(u.String()) {
		t.Errorf("Expected string representation to be valid, given: %s", u.String())
	}
	ur, _ := url.Parse(string(goLang))

	// Same NS same name MUST be equal
	u2 := NewV5(NamespaceURL, ur)
	if !Equal(u2, u) {
		t.Errorf("Expected UUIDs generated with same namespace and name to equal but got: %s and %s", u2, u)
	}

	// Different NS same name MUST NOT be equal
	u3 := NewV5(NamespaceDNS, ur)
	if Equal(u3, u) {
		t.Errorf("Expected UUIDs generated with different namespace and same name to be different but got: %s and %s", u3, u)
	}

	// Same NS different name MUST NOT be equal
	u4 := NewV5(NamespaceURL, u)
	if Equal(u4, u) {
		t.Errorf("Expected UUIDs generated with the same namespace and different names to be different but got: %s and %s", u4, u)
	}

	ids := []UUID{
		u, u2, u3, u4,
	}
	for j, id := range ids {
		i := NewV5(NamespaceURL, NewName(string(j), id))
		if Equal(id, i) {
			t.Errorf("Expected UUIDs generated with the same namespace and different names to be different but got: %s and %s", id, i)
		}
	}
}

func TestUUID_NewV5Bulk(t *testing.T) {
	for i := 0; i < generate; i++ {
		NewV5(NamespaceDNS, goLang)
	}
}

// A small test to test uniqueness across all UUIDs created
func TestUUID_EachIsUnique(t *testing.T) {
	s := 1000
	ids := make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV1()
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V1 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV3(NamespaceDNS, NewName(string(i), Name(goLang)))
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V3 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV4()
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V4 UUID", u, ids[j])
			}
		}
	}
	ids = make([]UUID, s)
	for i := 0; i < s; i++ {
		u := NewV5(NamespaceDNS, NewName(string(i), Name(goLang)))
		ids[i] = u
		for j := 0; j < i; j++ {
			if Equal(ids[j], u) {
				t.Error("Should not create the same V5 UUID", u, ids[j])
			}
		}
	}
}

// Not really a test but used for visual verification of the defaults
func UUID_NamespaceDefaults() {
	fmt.Println(NamespaceDNS)
	fmt.Println(NamespaceURL)
	fmt.Println(NamespaceOID)
	fmt.Println(NamespaceX500)
}