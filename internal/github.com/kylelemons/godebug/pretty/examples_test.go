// internal/github.com/kylelemons/godebug/pretty/examples_test.go: Part of the Antha language
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

// Copyright 2013 Google Inc.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pretty_test

import (
	"fmt"
	"github.com/antha-lang/manualLiquidHandler/internal/github.com/kylelemons/godebug/pretty"
)

func ExampleConfig_Sprint() {
	type Pair [2]int
	type Map struct {
		Name      string
		Players   map[string]Pair
		Obstacles map[Pair]string
	}

	m := Map{
		Name: "Rock Creek",
		Players: map[string]Pair{
			"player1": {1, 3},
			"player2": {0, -1},
		},
		Obstacles: map[Pair]string{
			Pair{0, 0}: "rock",
			Pair{2, 1}: "pond",
			Pair{1, 1}: "stream",
			Pair{0, 1}: "stream",
		},
	}

	// Specific output formats
	compact := &pretty.Config{
		Compact: true,
	}
	diffable := &pretty.Config{
		Diffable: true,
	}

	// Print out a summary
	fmt.Printf("Players: %s\n", compact.Sprint(m.Players))

	// Print diffable output
	fmt.Printf("Map State:\n%s", diffable.Sprint(m))

	// Output:
	// Players: {player1:[1,3],player2:[0,-1]}
	// Map State:
	// {
	//  Name:      "Rock Creek",
	//  Players:   {
	//              player1: [
	//                        1,
	//                        3,
	//                       ],
	//              player2: [
	//                        0,
	//                        -1,
	//                       ],
	//             },
	//  Obstacles: {
	//              [0,0]: "rock",
	//              [0,1]: "stream",
	//              [1,1]: "stream",
	//              [2,1]: "pond",
	//             },
	// }
}

func ExamplePrint() {
	type ShipManifest struct {
		Name     string
		Crew     map[string]string
		Androids int
		Stolen   bool
	}

	manifest := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Zaphod Beeblebrox": "Galactic President",
			"Trillian":          "Human",
			"Ford Prefect":      "A Hoopy Frood",
			"Arthur Dent":       "Along for the Ride",
		},
		Androids: 1,
		Stolen:   true,
	}

	pretty.Print(manifest)

	// Output:
	// {Name:     "Spaceship Heart of Gold",
	//  Crew:     {Arthur Dent:       "Along for the Ride",
	//             Ford Prefect:      "A Hoopy Frood",
	//             Trillian:          "Human",
	//             Zaphod Beeblebrox: "Galactic President"},
	//  Androids: 1,
	//  Stolen:   true}
}

func ExampleCompare() {
	type ShipManifest struct {
		Name     string
		Crew     map[string]string
		Androids int
		Stolen   bool
	}

	reported := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Zaphod Beeblebrox": "Galactic President",
			"Trillian":          "Human",
			"Ford Prefect":      "A Hoopy Frood",
			"Arthur Dent":       "Along for the Ride",
		},
		Androids: 1,
		Stolen:   true,
	}

	expected := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Rowan Artosok": "Captain",
		},
		Androids: 1,
		Stolen:   false,
	}

	fmt.Println(pretty.Compare(reported, expected))
	// Output:
	//  {
	//   Name:     "Spaceship Heart of Gold",
	//   Crew:     {
	// -            Arthur Dent:       "Along for the Ride",
	// -            Ford Prefect:      "A Hoopy Frood",
	// -            Trillian:          "Human",
	// -            Zaphod Beeblebrox: "Galactic President",
	// +            Rowan Artosok: "Captain",
	//             },
	//   Androids: 1,
	// - Stolen:   true,
	// + Stolen:   false,
	//  }
}