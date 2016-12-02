// util.go: Part of the Antha language
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
	"fmt"

	"github.com/antha-lang/antha/antha/anthalib/wtype"
)

// DecodeGenericPlate takes an interface{} that MUST be a json string representing a plate and will identify the
// proper type of plate and return an instance of it.
func DecodeGenericPlate(plate interface{}) (interface{}, error) {
	if p, ok := guessAddPlateToPlateType(plate); ok != nil {
		return nil, fmt.Errorf("AAH")
	} else {
		return p, nil
	}
}

func guessAddPlateToPlateType(plate interface{}) (interface{}, error) {
	if plate == nil {
		return nil, nil
	}
	switch p := plate.(type) {
	case string:
		var temp map[string]interface{}
		if err := json.Unmarshal([]byte(p), &temp); err != nil {
			return nil, err
		}
		//analyse what we got here
		if _, ok := temp["Welltype"]; ok { //wtype.LHPlate
			var ret wtype.LHPlate
			if err := json.Unmarshal([]byte(p), &ret); !ok {
				return nil, err
			}
			return ret, nil
		} else {
			if _, ok := temp["AsWell"]; ok {
				if _, ok := temp["TipXStart"]; ok { //wtype.LHTipbox
					var ret wtype.LHTipbox
					if err := json.Unmarshal([]byte(p), &ret); !ok {
						return nil, err
					}
					return ret, nil
				} else if _, ok := temp["WellXStart"]; ok { //wtype.LHTipwaste
					var ret wtype.LHTipwaste
					if err := json.Unmarshal([]byte(p), &ret); !ok {
						return nil, err
					}
					return ret, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Could not find suitable type for plate.")
}

//ToMultiLevelMessage returns a MultiLevelMessage from different kinds of types, string, []string and map[string]string
// This is a recursive method that will visit nested structures if present
func ToMultiLevelMessage(what interface{}) []MultiLevelMessage {
	ret := make([]MultiLevelMessage, 0)
	switch tw := what.(type) {
	case string:
		ret = append(ret, *NewMultiLevelMessage(tw, nil))
	case []string:
		for _, v := range tw {
			sthg := ToMultiLevelMessage(v)
			ret = append(ret, sthg...)
		}
	case map[string]string: //TODO jmanart implement a better way using reflect pkg
		for k, v := range tw {
			ret = append(ret, *NewMultiLevelMessage(k, ToMultiLevelMessage(v)))
		}
	}
	return ret
}
