// util_test.go: Part of the Antha language
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

package manualLiquidHandler

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/antha-lang/antha/antha/anthalib/wtype"
)

func TestDecodeGenericPlateLHTipBox(t *testing.T) {
	shp := wtype.NewShape("cylinder", "mm", 7.3, 7.3, 51.2)
	w := wtype.NewLHWell("Gilson20Tipbox", "", "A1", "ul", 20.0, 1.0, shp, 0, 7.3, 7.3, 46.0, 0.0, "mm")
	w.Extra["InnerL"] = 5.5
	w.Extra["InnerW"] = 5.5
	w.Extra["Tipeffectiveheight"] = 34.6
	tip := wtype.NewLHTip("gilson", "Gilson20", 0.5, 20.0, "ul")
	lhTipBox := *wtype.NewLHTipbox(8, 12, 60.13, "Gilson", "DL10 Tip Rack (PIPETMAX 8x20)", tip, w, 9.0, 9.0, 0.0, 0.0, 28.93)

	enc, err := json.Marshal(lhTipBox)
	if err != nil {
		t.Fatal(err)
	}

	out, err := DecodeGenericPlate(string(enc))
	if err != nil {
		t.Fatal(err)
	}

	if reflect.TypeOf(out) != reflect.TypeOf(lhTipBox) {
		t.Fatal("expecting output type ", reflect.TypeOf(lhTipBox), " got ", reflect.TypeOf(out))
	} else if !reflect.DeepEqual(out, lhTipBox) {
		t.Fatal("The input and output contents are not the same")
	}
}

func TestDecodeGenericPlateLHTipWaste(t *testing.T) {
	shp := wtype.NewShape("box", "mm", 123.0, 80.0, 92.0)

	w := wtype.NewLHWell("Gilsontipwaste", "", "A1", "ul", 800000.0, 800000.0, shp, 0, 123.0, 80.0, 92.0, 0.0, "mm")
	lht := *wtype.NewLHTipwaste(200, "gilsontipwaste", "gilson", 92.0, w, 49.5, 31.5, 0.0)

	enc, err := json.Marshal(lht)
	if err != nil {
		t.Fatal(err)
	}

	out, err := DecodeGenericPlate(string(enc))
	if err != nil {
		t.Fatal(err)
	}

	if reflect.TypeOf(out) != reflect.TypeOf(lht) {
		t.Fatal("expecting output type ", reflect.TypeOf(lht), " got ", reflect.TypeOf(out))
	} else if !reflect.DeepEqual(out, lht) {
		t.Fatal("The input and output contents are not the same")
	}
}

func TestDecodeGenericPlateLHPlate(t *testing.T) { //TODO jmanart, this really should have better set of tests
	swshp := wtype.NewShape("box", "mm", 8.2, 8.2, 41.3)
	welltype := wtype.NewLHWell("DSW96", "", "", "ul", 2000, 25, swshp, 3, 8.2, 8.2, 41.3, 4.7, "mm")
	plate := *wtype.NewLHPlate("DSW96", "Unknown", 8, 12, 44.1, "mm", welltype, 9, 9, 0.0, 0.0, 0.0)

	enc, err := json.Marshal(plate)
	if err != nil {
		t.Fatal(err)
	}
	out, err := DecodeGenericPlate(string(enc))
	if err != nil {
		t.Fatal(err)
	}

	if reflect.TypeOf(out) != reflect.TypeOf(plate) {
		t.Fatal("expecting output type ", reflect.TypeOf(plate), " got ", reflect.TypeOf(out))
	} else if !reflect.DeepEqual(out, plate) {
		//let's compare the json output
		wanted, _ := json.Marshal(plate)
		got, _ := json.Marshal(out)
		if strings.TrimSpace(string(wanted)) != strings.TrimSpace(string(got)) {
			//		wanted, _ := json.MarshalIndent(plate, "", "\t")
			//		got, _ := json.MarshalIndent(out, "", "\t")
			//		fmt.Println(pretty.Compare(out, plate))
			//cannot use this because it generates a closed loop inside the library
			//		fmt.Println(diff.Diff(string(got), string(wanted)))
			t.Fatalf(
				//			"The input and output contents are not the same.",
				"The input and output contents are not the same; Wanted \n%s\n got \n%s\n",
				string(got),
				string(wanted),
			)
		}
	}
}

func TestToMultiLevelMessage(t *testing.T) {
	t1 := make([]string, 0)
	t1 = append(t1, "Hi")
	t1 = append(t1, "Bye")
	r := ToMultiLevelMessage(t1)
	//check contents
	if len(r) != 2 {
		t.Fatal("Wrong lenght. Got ", len(r), " Expected ", len(t1))
	}
	if r[0].Message != t1[0] {
		t.Fatal("Wrong contents. Got ", r[0].Message, " Expected ", t1[0])
	}
	if r[1].Message != t1[1] {
		t.Fatal("Wrong contents. Got ", r[1].Message, " Expected ", t1[1])
	}

	t2 := make(map[string]string)
	t2["Hi"] = "Hello My Dear Friend"
	t2["GoodBye"] = "I'll miss you but not much"
	r2 := ToMultiLevelMessage(t2)
	//check contents
	if len(r2) != 2 {
		t.Fatal("Wrong length. Got ", len(r2), " Expected ", len(t2))
	}
	//TODO Contents are random because of maps, find easy way to check contents
}
