// moreutils.go: Part of the Antha language
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

	material "github.com/antha-lang/antha/antha/anthalib/material"
	wtype "github.com/antha-lang/antha/antha/anthalib/wtype"
	wunit "github.com/antha-lang/antha/antha/anthalib/wunit"
	driver "github.com/antha-lang/antha/microArch/driver"
	liquidhandling "github.com/antha-lang/antha/microArch/driver/liquidhandling"
	"github.com/antha-lang/antha/microArch/factory"
	pb "github.com/antha-lang/manualLiquidHandler/ExtendedLiquidhandlingDriver"

	"log"
)

type dummyserver struct {
}

func (s *dummyserver) GetCapabilities() (*pb.GetCapabilitiesReply, error) {
	r_1 := factory.GetLiquidhandlerByType("Manual")
	r_2 := driver.CommandStatus{
		true,
		0,
		"GETCAPABILITIES ACK",
	}

	ret := pb.GetCapabilitiesReply{
		EncodeLHProperties((liquidhandling.LHProperties)(*r_1)),
		EncodeCommandStatus((driver.CommandStatus)(r_2)),
	}
	return &ret, nil
}

func Encodeinterface(arg interface{}) *pb.AnyMessage {
	s, _ := json.Marshal(arg)
	ret := pb.AnyMessage{string(s)}
	return &ret
}
func Decodeinterface(msg *pb.AnyMessage) interface{} {
	var v interface{}
	json.Unmarshal([]byte(msg.Arg_1), &v)
	return v
}

func EncodeCommandStatus(arg driver.CommandStatus) *pb.CommandStatusMessage {
	ret := pb.CommandStatusMessage{(bool)(arg.OK), int64(arg.Errorcode), (string)(arg.Msg)}
	return &ret
}
func DecodeCommandStatus(arg *pb.CommandStatusMessage) driver.CommandStatus {
	ret := driver.CommandStatus{(bool)(arg.Arg_1), (int)(arg.Arg_2), (string)(arg.Arg_3)}
	return ret
}
func EncodePtrToLHProperties(arg *liquidhandling.LHProperties) *pb.PtrToLHPropertiesMessage {
	var ret pb.PtrToLHPropertiesMessage
	if arg == nil {
		ret = pb.PtrToLHPropertiesMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHPropertiesMessage{
			EncodeLHProperties(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHProperties(arg *pb.PtrToLHPropertiesMessage) *liquidhandling.LHProperties {
	if arg == nil {
		log.Println("Arg for PtrToLHProperties was nil")
		return nil
	}

	ret := DecodeLHProperties(arg.Arg_1)
	return &ret
}
func EncodeMapstringinterfaceMessage(arg map[string]interface{}) *pb.MapstringAnyMessageMessage {
	a := make([]*pb.MapstringAnyMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringinterfaceMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringAnyMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringinterfaceMessageFieldEntry(k string, v interface{}) pb.MapstringAnyMessageMessageFieldEntry {
	ret := pb.MapstringAnyMessageMessageFieldEntry{
		(string)(k),
		Encodeinterface(v),
	}
	return ret
}
func DecodeMapstringinterfaceMessage(arg *pb.MapstringAnyMessageMessage) map[string]interface{} {
	a := make(map[(string)](interface{}), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringinterfaceMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringinterfaceMessageFieldEntry(arg *pb.MapstringAnyMessageMessageFieldEntry) (string, interface{}) {
	k := (string)(arg.Key)
	v := Decodeinterface(arg.Value)
	return k, v
}
func EncodeArrayOfstring(arg []string) *pb.ArrayOfstring {
	a := make([]string, len(arg))
	for i, v := range arg {
		a[i] = (string)(v)
	}
	ret := pb.ArrayOfstring{
		a,
	}
	return &ret
}
func DecodeArrayOfstring(arg *pb.ArrayOfstring) []string {
	ret := make(([]string), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = (string)(v)
	}
	return ret
}
func EncodeArrayOfbool(arg []bool) *pb.ArrayOfbool {
	a := make([]bool, len(arg))
	for i, v := range arg {
		a[i] = (bool)(v)
	}
	ret := pb.ArrayOfbool{
		a,
	}
	return &ret
}
func DecodeArrayOfbool(arg *pb.ArrayOfbool) []bool {
	ret := make(([]bool), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = (bool)(v)
	}
	return ret
}
func EncodeArrayOfint(arg []int) *pb.ArrayOfint64 {
	a := make([]int64, len(arg))
	for i, v := range arg {
		a[i] = int64(v)
	}
	ret := pb.ArrayOfint64{
		a,
	}
	return &ret
}
func DecodeArrayOfint(arg *pb.ArrayOfint64) []int {
	ret := make(([]int), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = (int)(v)
	}
	return ret
}
func EncodeArrayOffloat64(arg []float64) *pb.ArrayOfdouble {
	a := make([]float64, len(arg))
	for i, v := range arg {
		a[i] = (float64)(v)
	}
	ret := pb.ArrayOfdouble{
		a,
	}
	return &ret
}
func DecodeArrayOffloat64(arg *pb.ArrayOfdouble) []float64 {
	ret := make(([]float64), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = (float64)(v)
	}
	return ret
}
func EncodeLHProperties(arg liquidhandling.LHProperties) *pb.LHPropertiesMessage {
	ret := pb.LHPropertiesMessage{(string)(arg.ID), int64(arg.Nposns), EncodeMapstringPtrToLHPositionMessage(arg.Positions), EncodeMapstringinterfaceMessage(arg.PlateLookup), EncodeMapstringstringMessage(arg.PosLookup), EncodeMapstringstringMessage(arg.PlateIDLookup), EncodeMapstringPtrToLHPlateMessage(arg.Plates), EncodeMapstringPtrToLHTipboxMessage(arg.Tipboxes), EncodeMapstringPtrToLHTipwasteMessage(arg.Tipwastes), EncodeMapstringPtrToLHPlateMessage(arg.Wastes), EncodeMapstringPtrToLHPlateMessage(arg.Washes), EncodeMapstringstringMessage(arg.Devices), (string)(arg.Model), (string)(arg.Mnfr), (string)(arg.LHType), (string)(arg.TipType), EncodeArrayOfPtrToLHHead(arg.Heads), EncodeArrayOfPtrToLHHead(arg.HeadsLoaded), EncodeArrayOfPtrToLHAdaptor(arg.Adaptors), EncodeArrayOfPtrToLHTip(arg.Tips), EncodeArrayOfstring(arg.Tip_preferences), EncodeArrayOfstring(arg.Input_preferences), EncodeArrayOfstring(arg.Output_preferences), EncodeArrayOfstring(arg.Tipwaste_preferences), EncodeArrayOfstring(arg.Waste_preferences), EncodeArrayOfstring(arg.Wash_preferences), EncodePtrToLHChannelParameter(arg.CurrConf), EncodeArrayOfPtrToLHChannelParameter(arg.Cnfvol), EncodeMapstringCoordinatesMessage(arg.Layout), int64(arg.MaterialType)}
	return &ret
}
func DecodeLHProperties(arg *pb.LHPropertiesMessage) liquidhandling.LHProperties {
	ret := liquidhandling.LHProperties{(string)(arg.Arg_1), (int)(arg.Arg_2), (map[string]*wtype.LHPosition)(DecodeMapstringPtrToLHPositionMessage(arg.Arg_3)), (map[string]interface{})(DecodeMapstringinterfaceMessage(arg.Arg_4)), (map[string]string)(DecodeMapstringstringMessage(arg.Arg_5)), (map[string]string)(DecodeMapstringstringMessage(arg.Arg_6)), (map[string]*wtype.LHPlate)(DecodeMapstringPtrToLHPlateMessage(arg.Arg_7)), (map[string]*wtype.LHTipbox)(DecodeMapstringPtrToLHTipboxMessage(arg.Arg_8)), (map[string]*wtype.LHTipwaste)(DecodeMapstringPtrToLHTipwasteMessage(arg.Arg_9)), (map[string]*wtype.LHPlate)(DecodeMapstringPtrToLHPlateMessage(arg.Arg_10)), (map[string]*wtype.LHPlate)(DecodeMapstringPtrToLHPlateMessage(arg.Arg_11)), (map[string]string)(DecodeMapstringstringMessage(arg.Arg_12)), (string)(arg.Arg_13), (string)(arg.Arg_14), (string)(arg.Arg_15), (string)(arg.Arg_16), ([]*wtype.LHHead)(DecodeArrayOfPtrToLHHead(arg.Arg_17)), ([]*wtype.LHHead)(DecodeArrayOfPtrToLHHead(arg.Arg_18)), ([]*wtype.LHAdaptor)(DecodeArrayOfPtrToLHAdaptor(arg.Arg_19)), ([]*wtype.LHTip)(DecodeArrayOfPtrToLHTip(arg.Arg_20)), ([]string)(DecodeArrayOfstring(arg.Arg_21)), ([]string)(DecodeArrayOfstring(arg.Arg_22)), ([]string)(DecodeArrayOfstring(arg.Arg_23)), ([]string)(DecodeArrayOfstring(arg.Arg_24)), ([]string)(DecodeArrayOfstring(arg.Arg_25)), ([]string)(DecodeArrayOfstring(arg.Arg_26)), nil, (*wtype.LHChannelParameter)(DecodePtrToLHChannelParameter(arg.Arg_27)), ([]*wtype.LHChannelParameter)(DecodeArrayOfPtrToLHChannelParameter(arg.Arg_28)), (map[string]wtype.Coordinates)(DecodeMapstringCoordinatesMessage(arg.Arg_29)), (material.MaterialType)(arg.Arg_30)}
	return ret
}
func EncodePtrToLHPosition(arg *wtype.LHPosition) *pb.PtrToLHPositionMessage {
	var ret pb.PtrToLHPositionMessage
	if arg == nil {
		ret = pb.PtrToLHPositionMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHPositionMessage{
			EncodeLHPosition(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHPosition(arg *pb.PtrToLHPositionMessage) *wtype.LHPosition {
	if arg == nil {
		log.Println("Arg for PtrToLHPosition was nil")
		return nil
	}

	ret := DecodeLHPosition(arg.Arg_1)
	return &ret
}
func EncodePtrToLHPlate(arg *wtype.LHPlate) *pb.PtrToLHPlateMessage {
	var ret pb.PtrToLHPlateMessage
	if arg == nil {
		ret = pb.PtrToLHPlateMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHPlateMessage{
			EncodeLHPlate(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHPlate(arg *pb.PtrToLHPlateMessage) *wtype.LHPlate {
	if arg == nil {
		log.Println("Arg for PtrToLHPlate was nil")
		return nil
	}

	ret := DecodeLHPlate(arg.Arg_1)
	return &ret
}
func EncodeArrayOfPtrToLHChannelParameter(arg []*wtype.LHChannelParameter) *pb.ArrayOfPtrToLHChannelParameterMessage {
	a := make([]*pb.PtrToLHChannelParameterMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHChannelParameter(v)
	}
	ret := pb.ArrayOfPtrToLHChannelParameterMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHChannelParameter(arg *pb.ArrayOfPtrToLHChannelParameterMessage) []*wtype.LHChannelParameter {
	ret := make(([]*wtype.LHChannelParameter), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHChannelParameter(v)
	}
	return ret
}
func EncodeLHTipbox(arg wtype.LHTipbox) *pb.LHTipboxMessage {
	ret := pb.LHTipboxMessage{EncodePtrToGenericSolid(arg.GenericSolid), (string)(arg.ID), (string)(arg.Boxname), (string)(arg.Type), (string)(arg.Mnfr), int64(arg.Nrows), int64(arg.Ncols), (float64)(arg.Height), EncodePtrToLHTip(arg.Tiptype), EncodePtrToLHWell(arg.AsWell), int64(arg.NTips), EncodeArrayOfArrayOfPtrToLHTip(arg.Tips), (float64)(arg.TipXOffset), (float64)(arg.TipYOffset), (float64)(arg.TipXStart), (float64)(arg.TipYStart), (float64)(arg.TipZStart)}
	return &ret
}
func DecodeLHTipbox(arg *pb.LHTipboxMessage) wtype.LHTipbox {
	ret := wtype.LHTipbox{(*wtype.GenericSolid)(DecodePtrToGenericSolid(arg.Arg_1)), (string)(arg.Arg_2), (string)(arg.Arg_3), (string)(arg.Arg_4), (string)(arg.Arg_5), (int)(arg.Arg_6), (int)(arg.Arg_7), (float64)(arg.Arg_8), (*wtype.LHTip)(DecodePtrToLHTip(arg.Arg_9)), (*wtype.LHWell)(DecodePtrToLHWell(arg.Arg_10)), (int)(arg.Arg_11), ([][]*wtype.LHTip)(DecodeArrayOfArrayOfPtrToLHTip(arg.Arg_12)), (float64)(arg.Arg_13), (float64)(arg.Arg_14), (float64)(arg.Arg_15), (float64)(arg.Arg_16), (float64)(arg.Arg_17)}
	return ret
}
func EncodeMapstringPtrToLHTipboxMessage(arg map[string]*wtype.LHTipbox) *pb.MapstringPtrToLHTipboxMessageMessage {
	a := make([]*pb.MapstringPtrToLHTipboxMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringPtrToLHTipboxMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringPtrToLHTipboxMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringPtrToLHTipboxMessageFieldEntry(k string, v *wtype.LHTipbox) pb.MapstringPtrToLHTipboxMessageMessageFieldEntry {
	ret := pb.MapstringPtrToLHTipboxMessageMessageFieldEntry{
		(string)(k),
		EncodePtrToLHTipbox(v),
	}
	return ret
}
func DecodeMapstringPtrToLHTipboxMessage(arg *pb.MapstringPtrToLHTipboxMessageMessage) map[string]*wtype.LHTipbox {
	a := make(map[(string)](*wtype.LHTipbox), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringPtrToLHTipboxMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringPtrToLHTipboxMessageFieldEntry(arg *pb.MapstringPtrToLHTipboxMessageMessageFieldEntry) (string, *wtype.LHTipbox) {
	k := (string)(arg.Key)
	v := DecodePtrToLHTipbox(arg.Value)
	return k, v
}
func EncodeLHHead(arg wtype.LHHead) *pb.LHHeadMessage {
	ret := pb.LHHeadMessage{(string)(arg.Name), (string)(arg.Manufacturer), (string)(arg.ID), EncodePtrToLHAdaptor(arg.Adaptor), EncodePtrToLHChannelParameter(arg.Params)}
	return &ret
}
func DecodeLHHead(arg *pb.LHHeadMessage) wtype.LHHead {
	ret := wtype.LHHead{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), (*wtype.LHAdaptor)(DecodePtrToLHAdaptor(arg.Arg_4)), (*wtype.LHChannelParameter)(DecodePtrToLHChannelParameter(arg.Arg_5))}
	return ret
}
func EncodePtrToLHHead(arg *wtype.LHHead) *pb.PtrToLHHeadMessage {
	var ret pb.PtrToLHHeadMessage
	if arg == nil {
		ret = pb.PtrToLHHeadMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHHeadMessage{
			EncodeLHHead(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHHead(arg *pb.PtrToLHHeadMessage) *wtype.LHHead {
	if arg == nil {
		log.Println("Arg for PtrToLHHead was nil")
		return nil
	}

	ret := DecodeLHHead(arg.Arg_1)
	return &ret
}
func EncodeArrayOfPtrToLHTip(arg []*wtype.LHTip) *pb.ArrayOfPtrToLHTipMessage {
	a := make([]*pb.PtrToLHTipMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHTip(v)
	}
	ret := pb.ArrayOfPtrToLHTipMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHTip(arg *pb.ArrayOfPtrToLHTipMessage) []*wtype.LHTip {
	ret := make(([]*wtype.LHTip), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHTip(v)
	}
	return ret
}
func EncodeLHPlate(arg wtype.LHPlate) *pb.LHPlateMessage {
	ret := pb.LHPlateMessage{EncodePtrToGenericEntity(arg.GenericEntity), (string)(arg.ID), (string)(arg.Inst), (string)(arg.Loc), (string)(arg.PlateName), (string)(arg.Type), (string)(arg.Mnfr), int64(arg.WlsX), int64(arg.WlsY), int64(arg.Nwells), EncodeMapstringPtrToLHWellMessage(arg.HWells), (float64)(arg.Height), (string)(arg.Hunit), EncodeArrayOfArrayOfPtrToLHWell(arg.Rows), EncodeArrayOfArrayOfPtrToLHWell(arg.Cols), EncodePtrToLHWell(arg.Welltype), EncodeMapstringPtrToLHWellMessage(arg.Wellcoords), (float64)(arg.WellXOffset), (float64)(arg.WellYOffset), (float64)(arg.WellXStart), (float64)(arg.WellYStart), (float64)(arg.WellZStart)}
	return &ret
}
func DecodeLHPlate(arg *pb.LHPlateMessage) wtype.LHPlate {
	ret := wtype.LHPlate{(*wtype.GenericEntity)(DecodePtrToGenericEntity(arg.Arg_1)), (string)(arg.Arg_2), (string)(arg.Arg_3), (string)(arg.Arg_4), (string)(arg.Arg_5), (string)(arg.Arg_6), (string)(arg.Arg_7), (int)(arg.Arg_8), (int)(arg.Arg_9), (int)(arg.Arg_10), (map[string]*wtype.LHWell)(DecodeMapstringPtrToLHWellMessage(arg.Arg_11)), (float64)(arg.Arg_12), (string)(arg.Arg_13), ([][]*wtype.LHWell)(DecodeArrayOfArrayOfPtrToLHWell(arg.Arg_14)), ([][]*wtype.LHWell)(DecodeArrayOfArrayOfPtrToLHWell(arg.Arg_15)), (*wtype.LHWell)(DecodePtrToLHWell(arg.Arg_16)), (map[string]*wtype.LHWell)(DecodeMapstringPtrToLHWellMessage(arg.Arg_17)), (float64)(arg.Arg_18), (float64)(arg.Arg_19), (float64)(arg.Arg_20), (float64)(arg.Arg_21), (float64)(arg.Arg_22)}
	return ret
}
func EncodeMapstringPtrToLHTipwasteMessage(arg map[string]*wtype.LHTipwaste) *pb.MapstringPtrToLHTipwasteMessageMessage {
	a := make([]*pb.MapstringPtrToLHTipwasteMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringPtrToLHTipwasteMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringPtrToLHTipwasteMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringPtrToLHTipwasteMessageFieldEntry(k string, v *wtype.LHTipwaste) pb.MapstringPtrToLHTipwasteMessageMessageFieldEntry {
	ret := pb.MapstringPtrToLHTipwasteMessageMessageFieldEntry{
		(string)(k),
		EncodePtrToLHTipwaste(v),
	}
	return ret
}
func DecodeMapstringPtrToLHTipwasteMessage(arg *pb.MapstringPtrToLHTipwasteMessageMessage) map[string]*wtype.LHTipwaste {
	a := make(map[(string)](*wtype.LHTipwaste), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringPtrToLHTipwasteMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringPtrToLHTipwasteMessageFieldEntry(arg *pb.MapstringPtrToLHTipwasteMessageMessageFieldEntry) (string, *wtype.LHTipwaste) {
	k := (string)(arg.Key)
	v := DecodePtrToLHTipwaste(arg.Value)
	return k, v
}
func EncodeArrayOfPtrToLHAdaptor(arg []*wtype.LHAdaptor) *pb.ArrayOfPtrToLHAdaptorMessage {
	a := make([]*pb.PtrToLHAdaptorMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHAdaptor(v)
	}
	ret := pb.ArrayOfPtrToLHAdaptorMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHAdaptor(arg *pb.ArrayOfPtrToLHAdaptorMessage) []*wtype.LHAdaptor {
	ret := make(([]*wtype.LHAdaptor), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHAdaptor(v)
	}
	return ret
}
func EncodeCoordinates(arg wtype.Coordinates) *pb.CoordinatesMessage {
	ret := pb.CoordinatesMessage{(float64)(arg.X), (float64)(arg.Y), (float64)(arg.Z)}
	return &ret
}
func DecodeCoordinates(arg *pb.CoordinatesMessage) wtype.Coordinates {
	ret := wtype.Coordinates{(float64)(arg.Arg_1), (float64)(arg.Arg_2), (float64)(arg.Arg_3)}
	return ret
}
func EncodeLHPosition(arg wtype.LHPosition) *pb.LHPositionMessage {
	ret := pb.LHPositionMessage{(string)(arg.ID), (string)(arg.Name), int64(arg.Num), EncodeArrayOfLHDevice(arg.Extra), (float64)(arg.Maxh)}
	return &ret
}
func DecodeLHPosition(arg *pb.LHPositionMessage) wtype.LHPosition {
	ret := wtype.LHPosition{(string)(arg.Arg_1), (string)(arg.Arg_2), (int)(arg.Arg_3), ([]wtype.LHDevice)(DecodeArrayOfLHDevice(arg.Arg_4)), (float64)(arg.Arg_5)}
	return ret
}
func EncodeMapstringPtrToLHPositionMessage(arg map[string]*wtype.LHPosition) *pb.MapstringPtrToLHPositionMessageMessage {
	a := make([]*pb.MapstringPtrToLHPositionMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringPtrToLHPositionMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringPtrToLHPositionMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringPtrToLHPositionMessageFieldEntry(k string, v *wtype.LHPosition) pb.MapstringPtrToLHPositionMessageMessageFieldEntry {
	ret := pb.MapstringPtrToLHPositionMessageMessageFieldEntry{
		(string)(k),
		EncodePtrToLHPosition(v),
	}
	return ret
}
func DecodeMapstringPtrToLHPositionMessage(arg *pb.MapstringPtrToLHPositionMessageMessage) map[string]*wtype.LHPosition {
	a := make(map[(string)](*wtype.LHPosition), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringPtrToLHPositionMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringPtrToLHPositionMessageFieldEntry(arg *pb.MapstringPtrToLHPositionMessageMessageFieldEntry) (string, *wtype.LHPosition) {
	k := (string)(arg.Key)
	v := DecodePtrToLHPosition(arg.Value)
	return k, v
}
func EncodeMapstringPtrToLHPlateMessage(arg map[string]*wtype.LHPlate) *pb.MapstringPtrToLHPlateMessageMessage {
	a := make([]*pb.MapstringPtrToLHPlateMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringPtrToLHPlateMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringPtrToLHPlateMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringPtrToLHPlateMessageFieldEntry(k string, v *wtype.LHPlate) pb.MapstringPtrToLHPlateMessageMessageFieldEntry {
	ret := pb.MapstringPtrToLHPlateMessageMessageFieldEntry{
		(string)(k),
		EncodePtrToLHPlate(v),
	}
	return ret
}
func DecodeMapstringPtrToLHPlateMessage(arg *pb.MapstringPtrToLHPlateMessageMessage) map[string]*wtype.LHPlate {
	a := make(map[(string)](*wtype.LHPlate), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringPtrToLHPlateMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringPtrToLHPlateMessageFieldEntry(arg *pb.MapstringPtrToLHPlateMessageMessageFieldEntry) (string, *wtype.LHPlate) {
	k := (string)(arg.Key)
	v := DecodePtrToLHPlate(arg.Value)
	return k, v
}
func EncodeArrayOfPtrToLHHead(arg []*wtype.LHHead) *pb.ArrayOfPtrToLHHeadMessage {
	a := make([]*pb.PtrToLHHeadMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHHead(v)
	}
	ret := pb.ArrayOfPtrToLHHeadMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHHead(arg *pb.ArrayOfPtrToLHHeadMessage) []*wtype.LHHead {
	ret := make(([]*wtype.LHHead), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHHead(v)
	}
	return ret
}
func EncodePtrToLHTip(arg *wtype.LHTip) *pb.PtrToLHTipMessage {
	var ret pb.PtrToLHTipMessage
	if arg == nil {
		ret = pb.PtrToLHTipMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHTipMessage{
			EncodeLHTip(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHTip(arg *pb.PtrToLHTipMessage) *wtype.LHTip {
	if arg == nil {
		log.Println("Arg for PtrToLHTip was nil")
		return nil
	}

	ret := DecodeLHTip(arg.Arg_1)
	return &ret
}
func EncodePtrToLHTipwaste(arg *wtype.LHTipwaste) *pb.PtrToLHTipwasteMessage {
	var ret pb.PtrToLHTipwasteMessage
	if arg == nil {
		ret = pb.PtrToLHTipwasteMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHTipwasteMessage{
			EncodeLHTipwaste(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHTipwaste(arg *pb.PtrToLHTipwasteMessage) *wtype.LHTipwaste {
	if arg == nil {
		log.Println("Arg for PtrToLHTipwaste was nil")
		return nil
	}

	ret := DecodeLHTipwaste(arg.Arg_1)
	return &ret
}
func EncodeLHTip(arg wtype.LHTip) *pb.LHTipMessage {
	ret := pb.LHTipMessage{(string)(arg.ID), (string)(arg.Type), (string)(arg.Mnfr), (bool)(arg.Dirty), EncodePtrToVolume(arg.MaxVol), EncodePtrToVolume(arg.MinVol)}
	return &ret
}
func DecodeLHTip(arg *pb.LHTipMessage) wtype.LHTip {
	ret := wtype.LHTip{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), (bool)(arg.Arg_4), (*wunit.Volume)(DecodePtrToVolume(arg.Arg_5)), (*wunit.Volume)(DecodePtrToVolume(arg.Arg_6))}
	return ret
}
func EncodeLHChannelParameter(arg wtype.LHChannelParameter) *pb.LHChannelParameterMessage {
	ret := pb.LHChannelParameterMessage{(string)(arg.ID), (string)(arg.Name), EncodePtrToVolume(arg.Minvol), EncodePtrToVolume(arg.Maxvol), EncodePtrToFlowRate(arg.Minspd), EncodePtrToFlowRate(arg.Maxspd), int64(arg.Multi), (bool)(arg.Independent), int64(arg.Orientation), int64(arg.Head)}
	return &ret
}
func DecodeLHChannelParameter(arg *pb.LHChannelParameterMessage) wtype.LHChannelParameter {
	ret := wtype.LHChannelParameter{(string)(arg.Arg_1), (string)(arg.Arg_2), (*wunit.Volume)(DecodePtrToVolume(arg.Arg_3)), (*wunit.Volume)(DecodePtrToVolume(arg.Arg_4)), (*wunit.FlowRate)(DecodePtrToFlowRate(arg.Arg_5)), (*wunit.FlowRate)(DecodePtrToFlowRate(arg.Arg_6)), (int)(arg.Arg_7), (bool)(arg.Arg_8), (int)(arg.Arg_9), (int)(arg.Arg_10)}
	return ret
}
func EncodePtrToLHChannelParameter(arg *wtype.LHChannelParameter) *pb.PtrToLHChannelParameterMessage {
	var ret pb.PtrToLHChannelParameterMessage
	if arg == nil {
		ret = pb.PtrToLHChannelParameterMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHChannelParameterMessage{
			EncodeLHChannelParameter(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHChannelParameter(arg *pb.PtrToLHChannelParameterMessage) *wtype.LHChannelParameter {
	if arg == nil {
		log.Println("Arg for PtrToLHChannelParameter was nil")
		return nil
	}

	ret := DecodeLHChannelParameter(arg.Arg_1)
	return &ret
}
func EncodeMapstringCoordinatesMessage(arg map[string]wtype.Coordinates) *pb.MapstringCoordinatesMessageMessage {
	a := make([]*pb.MapstringCoordinatesMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringCoordinatesMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringCoordinatesMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringCoordinatesMessageFieldEntry(k string, v wtype.Coordinates) pb.MapstringCoordinatesMessageMessageFieldEntry {
	ret := pb.MapstringCoordinatesMessageMessageFieldEntry{
		(string)(k),
		EncodeCoordinates(v),
	}
	return ret
}
func DecodeMapstringCoordinatesMessage(arg *pb.MapstringCoordinatesMessageMessage) map[string]wtype.Coordinates {
	a := make(map[(string)](wtype.Coordinates), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringCoordinatesMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringCoordinatesMessageFieldEntry(arg *pb.MapstringCoordinatesMessageMessageFieldEntry) (string, wtype.Coordinates) {
	k := (string)(arg.Key)
	v := DecodeCoordinates(arg.Value)
	return k, v
}
func EncodeLHAdaptor(arg wtype.LHAdaptor) *pb.LHAdaptorMessage {
	ret := pb.LHAdaptorMessage{(string)(arg.Name), (string)(arg.ID), (string)(arg.Manufacturer), EncodePtrToLHChannelParameter(arg.Params), int64(arg.Ntipsloaded), EncodePtrToLHTip(arg.Tiptypeloaded)}
	return &ret
}
func DecodeLHAdaptor(arg *pb.LHAdaptorMessage) wtype.LHAdaptor {
	ret := wtype.LHAdaptor{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), (*wtype.LHChannelParameter)(DecodePtrToLHChannelParameter(arg.Arg_4)), (int)(arg.Arg_5), (*wtype.LHTip)(DecodePtrToLHTip(arg.Arg_6))}
	return ret
}
func EncodePtrToLHAdaptor(arg *wtype.LHAdaptor) *pb.PtrToLHAdaptorMessage {
	var ret pb.PtrToLHAdaptorMessage
	if arg == nil {
		ret = pb.PtrToLHAdaptorMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHAdaptorMessage{
			EncodeLHAdaptor(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHAdaptor(arg *pb.PtrToLHAdaptorMessage) *wtype.LHAdaptor {
	if arg == nil {
		log.Println("Arg for PtrToLHAdaptor was nil")
		return nil
	}

	ret := DecodeLHAdaptor(arg.Arg_1)
	return &ret
}
func EncodePtrToLHTipbox(arg *wtype.LHTipbox) *pb.PtrToLHTipboxMessage {
	var ret pb.PtrToLHTipboxMessage
	if arg == nil {
		ret = pb.PtrToLHTipboxMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHTipboxMessage{
			EncodeLHTipbox(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHTipbox(arg *pb.PtrToLHTipboxMessage) *wtype.LHTipbox {
	if arg == nil {
		log.Println("Arg for PtrToLHTipbox was nil")
		return nil
	}

	ret := DecodeLHTipbox(arg.Arg_1)
	return &ret
}
func EncodeLHTipwaste(arg wtype.LHTipwaste) *pb.LHTipwasteMessage {
	ret := pb.LHTipwasteMessage{(string)(arg.ID), (string)(arg.Type), (string)(arg.Mnfr), int64(arg.Capacity), int64(arg.Contents), (float64)(arg.Height), (float64)(arg.WellXStart), (float64)(arg.WellYStart), (float64)(arg.WellZStart), EncodePtrToLHWell(arg.AsWell)}
	return &ret
}
func DecodeLHTipwaste(arg *pb.LHTipwasteMessage) wtype.LHTipwaste {
	ret := wtype.LHTipwaste{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), (int)(arg.Arg_4), (int)(arg.Arg_5), (float64)(arg.Arg_6), (float64)(arg.Arg_7), (float64)(arg.Arg_8), (float64)(arg.Arg_9), (*wtype.LHWell)(DecodePtrToLHWell(arg.Arg_10))}
	return ret
}
func EncodeMapstringstringMessage(arg map[string]string) *pb.MapstringstringMessage {
	a := make([]*pb.MapstringstringMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringstringMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringstringMessage{
		a,
	}
	return &ret
}
func EncodeMapstringstringMessageFieldEntry(k string, v string) pb.MapstringstringMessageFieldEntry {
	ret := pb.MapstringstringMessageFieldEntry{
		(string)(k),
		(string)(v),
	}
	return ret
}
func DecodeMapstringstringMessage(arg *pb.MapstringstringMessage) map[string]string {
	a := make(map[(string)](string), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringstringMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringstringMessageFieldEntry(arg *pb.MapstringstringMessageFieldEntry) (string, string) {
	k := (string)(arg.Key)
	v := (string)(arg.Value)
	return k, v
}
func EncodeFlowRate(arg wunit.FlowRate) *pb.FlowRateMessage {
	ret := pb.FlowRateMessage{EncodeConcreteMeasurement(arg.ConcreteMeasurement)}
	return &ret
}
func DecodeFlowRate(arg *pb.FlowRateMessage) wunit.FlowRate {
	ret := wunit.FlowRate{(wunit.ConcreteMeasurement)(DecodeConcreteMeasurement(arg.Arg_1))}
	return ret
}
func EncodeArrayOfPtrToLHWell(arg []*wtype.LHWell) *pb.ArrayOfPtrToLHWellMessage {
	a := make([]*pb.PtrToLHWellMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHWell(v)
	}
	ret := pb.ArrayOfPtrToLHWellMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHWell(arg *pb.ArrayOfPtrToLHWellMessage) []*wtype.LHWell {
	ret := make(([]*wtype.LHWell), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHWell(v)
	}
	return ret
}
func EncodeLHDevice(arg wtype.LHDevice) *pb.LHDeviceMessage {
	ret := pb.LHDeviceMessage{(string)(arg.ID), (string)(arg.Name), (string)(arg.Mnfr)}
	return &ret
}
func DecodeLHDevice(arg *pb.LHDeviceMessage) wtype.LHDevice {
	ret := wtype.LHDevice{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3)}
	return ret
}
func EncodePtrToFlowRate(arg *wunit.FlowRate) *pb.PtrToFlowRateMessage {
	var ret pb.PtrToFlowRateMessage
	if arg == nil {
		ret = pb.PtrToFlowRateMessage{
			nil,
		}
	} else {
		ret = pb.PtrToFlowRateMessage{
			EncodeFlowRate(*arg),
		}
	}
	return &ret
}
func DecodePtrToFlowRate(arg *pb.PtrToFlowRateMessage) *wunit.FlowRate {
	if arg == nil {
		log.Println("Arg for PtrToFlowRate was nil")
		return nil
	}

	ret := DecodeFlowRate(arg.Arg_1)
	return &ret
}
func EncodeMapstringPtrToLHWellMessage(arg map[string]*wtype.LHWell) *pb.MapstringPtrToLHWellMessageMessage {
	a := make([]*pb.MapstringPtrToLHWellMessageMessageFieldEntry, 0, len(arg))
	for k, v := range arg {
		fe := EncodeMapstringPtrToLHWellMessageFieldEntry(k, v)
		a = append(a, &fe)
	}
	ret := pb.MapstringPtrToLHWellMessageMessage{
		a,
	}
	return &ret
}
func EncodeMapstringPtrToLHWellMessageFieldEntry(k string, v *wtype.LHWell) pb.MapstringPtrToLHWellMessageMessageFieldEntry {
	ret := pb.MapstringPtrToLHWellMessageMessageFieldEntry{
		(string)(k),
		EncodePtrToLHWell(v),
	}
	return ret
}
func DecodeMapstringPtrToLHWellMessage(arg *pb.MapstringPtrToLHWellMessageMessage) map[string]*wtype.LHWell {
	a := make(map[(string)](*wtype.LHWell), len(arg.MapField))
	for _, fe := range arg.MapField {
		k, v := DecodeMapstringPtrToLHWellMessageFieldEntry(fe)
		a[k] = v
	}
	return a
}
func DecodeMapstringPtrToLHWellMessageFieldEntry(arg *pb.MapstringPtrToLHWellMessageMessageFieldEntry) (string, *wtype.LHWell) {
	k := (string)(arg.Key)
	v := DecodePtrToLHWell(arg.Value)
	return k, v
}
func EncodeArrayOfLHDevice(arg []wtype.LHDevice) *pb.ArrayOfLHDeviceMessage {
	a := make([]*pb.LHDeviceMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodeLHDevice(v)
	}
	ret := pb.ArrayOfLHDeviceMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfLHDevice(arg *pb.ArrayOfLHDeviceMessage) []wtype.LHDevice {
	ret := make(([]wtype.LHDevice), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodeLHDevice(v)
	}
	return ret
}
func EncodePtrToVolume(arg *wunit.Volume) *pb.PtrToVolumeMessage {
	var ret pb.PtrToVolumeMessage
	if arg == nil {
		ret = pb.PtrToVolumeMessage{
			nil,
		}
	} else {
		ret = pb.PtrToVolumeMessage{
			EncodeVolume(*arg),
		}
	}
	return &ret
}
func DecodePtrToVolume(arg *pb.PtrToVolumeMessage) *wunit.Volume {
	if arg == nil {
		log.Println("Arg for PtrToVolume was nil")
		return nil
	}

	ret := DecodeVolume(arg.Arg_1)
	return &ret
}
func EncodeArrayOfArrayOfPtrToLHTip(arg [][]*wtype.LHTip) *pb.ArrayOfArrayOfPtrToLHTipMessage {
	a := make([]*pb.ArrayOfPtrToLHTipMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodeArrayOfPtrToLHTip(v)
	}
	ret := pb.ArrayOfArrayOfPtrToLHTipMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfArrayOfPtrToLHTip(arg *pb.ArrayOfArrayOfPtrToLHTipMessage) [][]*wtype.LHTip {
	ret := make(([][]*wtype.LHTip), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodeArrayOfPtrToLHTip(v)
	}
	return ret
}
func EncodeArrayOfArrayOfPtrToLHWell(arg [][]*wtype.LHWell) *pb.ArrayOfArrayOfPtrToLHWellMessage {
	a := make([]*pb.ArrayOfPtrToLHWellMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodeArrayOfPtrToLHWell(v)
	}
	ret := pb.ArrayOfArrayOfPtrToLHWellMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfArrayOfPtrToLHWell(arg *pb.ArrayOfArrayOfPtrToLHWellMessage) [][]*wtype.LHWell {
	ret := make(([][]*wtype.LHWell), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodeArrayOfPtrToLHWell(v)
	}
	return ret
}
func EncodePtrToGenericSolid(arg *wtype.GenericSolid) *pb.PtrToGenericSolidMessage {
	var ret pb.PtrToGenericSolidMessage
	if arg == nil {
		ret = pb.PtrToGenericSolidMessage{
			nil,
		}
	} else {
		ret = pb.PtrToGenericSolidMessage{
			EncodeGenericSolid(*arg),
		}
	}
	return &ret
}
func DecodePtrToGenericSolid(arg *pb.PtrToGenericSolidMessage) *wtype.GenericSolid {
	if arg == nil {
		log.Println("Arg for PtrToGenericSolid was nil")
		return nil
	}

	ret := DecodeGenericSolid(arg.Arg_1)
	return &ret
}
func EncodeGenericSolid(arg wtype.GenericSolid) *pb.GenericSolidMessage {
	ret := pb.GenericSolidMessage{EncodeGenericPhysical(arg.GenericPhysical), EncodePtrToShape(arg.Myshape)}
	return &ret
}
func DecodeGenericSolid(arg *pb.GenericSolidMessage) wtype.GenericSolid {
	ret := wtype.GenericSolid{(wtype.GenericPhysical)(DecodeGenericPhysical(arg.Arg_1)), (*wtype.Shape)(DecodePtrToShape(arg.Arg_2))}
	return ret
}
func EncodePtrToLHWell(arg *wtype.LHWell) *pb.PtrToLHWellMessage {
	var ret pb.PtrToLHWellMessage
	if arg == nil {
		ret = pb.PtrToLHWellMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHWellMessage{
			EncodeLHWell(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHWell(arg *pb.PtrToLHWellMessage) *wtype.LHWell {
	if arg == nil {
		log.Println("Arg for PtrToLHWell was nil")
		return nil
	}

	ret := DecodeLHWell(arg.Arg_1)
	return &ret
}
func EncodePtrToGenericEntity(arg *wtype.GenericEntity) *pb.PtrToGenericEntityMessage {
	var ret pb.PtrToGenericEntityMessage
	if arg == nil {
		ret = pb.PtrToGenericEntityMessage{
			nil,
		}
	} else {
		ret = pb.PtrToGenericEntityMessage{
			EncodeGenericEntity(*arg),
		}
	}
	return &ret
}
func DecodePtrToGenericEntity(arg *pb.PtrToGenericEntityMessage) *wtype.GenericEntity {
	if arg == nil {
		log.Println("Arg for PtrToGenericEntity was nil")
		return nil
	}

	ret := DecodeGenericEntity(arg.Arg_1)
	return &ret
}
func EncodeVolume(arg wunit.Volume) *pb.VolumeMessage {
	ret := pb.VolumeMessage{EncodeConcreteMeasurement(arg.ConcreteMeasurement)}
	return &ret
}
func DecodeVolume(arg *pb.VolumeMessage) wunit.Volume {
	ret := wunit.Volume{(wunit.ConcreteMeasurement)(DecodeConcreteMeasurement(arg.Arg_1))}
	return ret
}
func EncodeLHWell(arg wtype.LHWell) *pb.LHWellMessage {
	ret := pb.LHWellMessage{(string)(arg.ID), (string)(arg.Inst), (string)(arg.Plateinst), (string)(arg.Plateid), (string)(arg.Platetype), (string)(arg.Crds), (float64)(arg.Vol), (string)(arg.Vunit), EncodeArrayOfPtrToLHComponent(arg.WContents), (float64)(arg.Rvol), (float64)(arg.Currvol), EncodePtrToShape(arg.WShape), int64(arg.Bottom), (float64)(arg.Xdim), (float64)(arg.Ydim), (float64)(arg.Zdim), (float64)(arg.Bottomh), (string)(arg.Dunit), EncodeMapstringinterfaceMessage(arg.Extra)}
	return &ret
}
func DecodeLHWell(arg *pb.LHWellMessage) wtype.LHWell {
	ret := wtype.LHWell{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), (string)(arg.Arg_4), (string)(arg.Arg_5), (string)(arg.Arg_6), (float64)(arg.Arg_7), (string)(arg.Arg_8), ([]*wtype.LHComponent)(DecodeArrayOfPtrToLHComponent(arg.Arg_9)), (float64)(arg.Arg_10), (float64)(arg.Arg_11), (*wtype.Shape)(DecodePtrToShape(arg.Arg_12)), (int)(arg.Arg_13), (float64)(arg.Arg_14), (float64)(arg.Arg_15), (float64)(arg.Arg_16), (float64)(arg.Arg_17), (string)(arg.Arg_18), (map[string]interface{})(DecodeMapstringinterfaceMessage(arg.Arg_19)), nil}
	return ret
}
func EncodeGenericEntity(arg wtype.GenericEntity) *pb.GenericEntityMessage {
	ret := pb.GenericEntityMessage{EncodePtrToGenericSolid(arg.GenericSolid), EncodePtrToConcreteLocation(arg.Loc)}
	return &ret
}
func DecodeGenericEntity(arg *pb.GenericEntityMessage) wtype.GenericEntity {
	ret := wtype.GenericEntity{(*wtype.GenericSolid)(DecodePtrToGenericSolid(arg.Arg_1)), (*wtype.ConcreteLocation)(DecodePtrToConcreteLocation(arg.Arg_2))}
	return ret
}
func EncodeGenericPhysical(arg wtype.GenericPhysical) *pb.GenericPhysicalMessage {
	ret := pb.GenericPhysicalMessage{EncodeGenericMatter(arg.GenericMatter), (string)(arg.Myname), EncodeMass(arg.Mymass), EncodeVolume(arg.Myvol), EncodeTemperature(arg.Mytemp)}
	return &ret
}
func DecodeGenericPhysical(arg *pb.GenericPhysicalMessage) wtype.GenericPhysical {
	ret := wtype.GenericPhysical{(wtype.GenericMatter)(DecodeGenericMatter(arg.Arg_1)), (string)(arg.Arg_2), (wunit.Mass)(DecodeMass(arg.Arg_3)), (wunit.Volume)(DecodeVolume(arg.Arg_4)), (wunit.Temperature)(DecodeTemperature(arg.Arg_5))}
	return ret
}
func EncodeLHComponent(arg wtype.LHComponent) *pb.LHComponentMessage {
	ret := pb.LHComponentMessage{EncodePtrToGenericPhysical(arg.GenericPhysical), (string)(arg.ID), (string)(arg.Inst), int64(arg.Order), (string)(arg.CName), (string)(arg.Type), (float64)(arg.Vol), (float64)(arg.Conc), (string)(arg.Vunit), (string)(arg.Cunit), (float64)(arg.Tvol), (string)(arg.Loc), (float64)(arg.Smax), (float64)(arg.Visc), (float64)(arg.StockConcentration), (string)(arg.Destination), EncodeMapstringinterfaceMessage(arg.Extra)}
	return &ret
}
func DecodeLHComponent(arg *pb.LHComponentMessage) wtype.LHComponent {
	ret := wtype.LHComponent{(*wtype.GenericPhysical)(DecodePtrToGenericPhysical(arg.Arg_1)), (string)(arg.Arg_2), (string)(arg.Arg_3), (int)(arg.Arg_4), (string)(arg.Arg_5), (string)(arg.Arg_6), (float64)(arg.Arg_7), (float64)(arg.Arg_8), (string)(arg.Arg_9), (string)(arg.Arg_10), (float64)(arg.Arg_11), (string)(arg.Arg_12), (float64)(arg.Arg_13), (float64)(arg.Arg_14), (float64)(arg.Arg_15), nil, (string)(arg.Arg_16), (map[string]interface{})(DecodeMapstringinterfaceMessage(arg.Arg_17))}
	return ret
}
func EncodePtrToLHComponent(arg *wtype.LHComponent) *pb.PtrToLHComponentMessage {
	var ret pb.PtrToLHComponentMessage
	if arg == nil {
		ret = pb.PtrToLHComponentMessage{
			nil,
		}
	} else {
		ret = pb.PtrToLHComponentMessage{
			EncodeLHComponent(*arg),
		}
	}
	return &ret
}
func DecodePtrToLHComponent(arg *pb.PtrToLHComponentMessage) *wtype.LHComponent {
	if arg == nil {
		log.Println("Arg for PtrToLHComponent was nil")
		return nil
	}

	ret := DecodeLHComponent(arg.Arg_1)
	return &ret
}
func EncodeConcreteLocation(arg wtype.ConcreteLocation) *pb.ConcreteLocationMessage {
	ret := pb.ConcreteLocationMessage{(string)(arg.ID), (string)(arg.Inst), (string)(arg.Name), EncodeArrayOfPtrToConcreteLocation(arg.Psns), EncodePtrToConcreteLocation(arg.Cntr), EncodePtrToShape(arg.Shap)}
	return &ret
}
func DecodeConcreteLocation(arg *pb.ConcreteLocationMessage) wtype.ConcreteLocation {
	ret := wtype.ConcreteLocation{(string)(arg.Arg_1), (string)(arg.Arg_2), (string)(arg.Arg_3), ([]*wtype.ConcreteLocation)(DecodeArrayOfPtrToConcreteLocation(arg.Arg_4)), (*wtype.ConcreteLocation)(DecodePtrToConcreteLocation(arg.Arg_5)), (*wtype.Shape)(DecodePtrToShape(arg.Arg_6))}
	return ret
}
func EncodePtrToConcreteLocation(arg *wtype.ConcreteLocation) *pb.PtrToConcreteLocationMessage {
	var ret pb.PtrToConcreteLocationMessage
	if arg == nil {
		ret = pb.PtrToConcreteLocationMessage{
			nil,
		}
	} else {
		ret = pb.PtrToConcreteLocationMessage{
			EncodeConcreteLocation(*arg),
		}
	}
	return &ret
}
func DecodePtrToConcreteLocation(arg *pb.PtrToConcreteLocationMessage) *wtype.ConcreteLocation {
	if arg == nil {
		log.Println("Arg for PtrToConcreteLocation was nil")
		return nil
	}

	ret := DecodeConcreteLocation(arg.Arg_1)
	return &ret
}
func EncodePtrToShape(arg *wtype.Shape) *pb.PtrToShapeMessage {
	var ret pb.PtrToShapeMessage
	if arg == nil {
		ret = pb.PtrToShapeMessage{
			nil,
		}
	} else {
		ret = pb.PtrToShapeMessage{
			EncodeShape(*arg),
		}
	}
	return &ret
}
func DecodePtrToShape(arg *pb.PtrToShapeMessage) *wtype.Shape {
	if arg == nil {
		log.Println("Arg for PtrToShape was nil")
		return nil
	}

	ret := DecodeShape(arg.Arg_1)
	return &ret
}
func EncodeConcreteMeasurement(arg wunit.ConcreteMeasurement) *pb.ConcreteMeasurementMessage {
	ret := pb.ConcreteMeasurementMessage{(float64)(arg.Mvalue), EncodePtrToGenericPrefixedUnit(arg.Munit)}
	return &ret
}
func DecodeConcreteMeasurement(arg *pb.ConcreteMeasurementMessage) wunit.ConcreteMeasurement {
	ret := wunit.ConcreteMeasurement{(float64)(arg.Arg_1), (*wunit.GenericPrefixedUnit)(DecodePtrToGenericPrefixedUnit(arg.Arg_2))}
	return ret
}
func EncodeArrayOfPtrToLHComponent(arg []*wtype.LHComponent) *pb.ArrayOfPtrToLHComponentMessage {
	a := make([]*pb.PtrToLHComponentMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToLHComponent(v)
	}
	ret := pb.ArrayOfPtrToLHComponentMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToLHComponent(arg *pb.ArrayOfPtrToLHComponentMessage) []*wtype.LHComponent {
	ret := make(([]*wtype.LHComponent), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToLHComponent(v)
	}
	return ret
}
func EncodeShape(arg wtype.Shape) *pb.ShapeMessage {
	ret := pb.ShapeMessage{(string)(arg.ShapeName), (string)(arg.LengthUnit), (float64)(arg.H), (float64)(arg.W), (float64)(arg.D)}
	return &ret
}
func DecodeShape(arg *pb.ShapeMessage) wtype.Shape {
	ret := wtype.Shape{(string)(arg.Arg_1), (string)(arg.Arg_2), (float64)(arg.Arg_3), (float64)(arg.Arg_4), (float64)(arg.Arg_5)}
	return ret
}
func EncodeArrayOfPtrToConcreteLocation(arg []*wtype.ConcreteLocation) *pb.ArrayOfPtrToConcreteLocationMessage {
	a := make([]*pb.PtrToConcreteLocationMessage, len(arg))
	for i, v := range arg {
		a[i] = EncodePtrToConcreteLocation(v)
	}
	ret := pb.ArrayOfPtrToConcreteLocationMessage{
		a,
	}
	return &ret
}
func DecodeArrayOfPtrToConcreteLocation(arg *pb.ArrayOfPtrToConcreteLocationMessage) []*wtype.ConcreteLocation {
	ret := make(([]*wtype.ConcreteLocation), len(arg.Arg_1))
	for i, v := range arg.Arg_1 {
		ret[i] = DecodePtrToConcreteLocation(v)
	}
	return ret
}
func EncodePtrToGenericPrefixedUnit(arg *wunit.GenericPrefixedUnit) *pb.PtrToGenericPrefixedUnitMessage {
	var ret pb.PtrToGenericPrefixedUnitMessage
	if arg == nil {
		ret = pb.PtrToGenericPrefixedUnitMessage{
			nil,
		}
	} else {
		ret = pb.PtrToGenericPrefixedUnitMessage{
			EncodeGenericPrefixedUnit(*arg),
		}
	}
	return &ret
}
func DecodePtrToGenericPrefixedUnit(arg *pb.PtrToGenericPrefixedUnitMessage) *wunit.GenericPrefixedUnit {
	if arg == nil {
		log.Println("Arg for PtrToGenericPrefixedUnit was nil")
		return nil
	}

	ret := DecodeGenericPrefixedUnit(arg.Arg_1)
	return &ret
}
func EncodeTemperature(arg wunit.Temperature) *pb.TemperatureMessage {
	ret := pb.TemperatureMessage{EncodeConcreteMeasurement(arg.ConcreteMeasurement)}
	return &ret
}
func DecodeTemperature(arg *pb.TemperatureMessage) wunit.Temperature {
	ret := wunit.Temperature{(wunit.ConcreteMeasurement)(DecodeConcreteMeasurement(arg.Arg_1))}
	return ret
}
func EncodePtrToGenericPhysical(arg *wtype.GenericPhysical) *pb.PtrToGenericPhysicalMessage {
	var ret pb.PtrToGenericPhysicalMessage
	if arg == nil {
		ret = pb.PtrToGenericPhysicalMessage{
			nil,
		}
	} else {
		ret = pb.PtrToGenericPhysicalMessage{
			EncodeGenericPhysical(*arg),
		}
	}
	return &ret
}
func DecodePtrToGenericPhysical(arg *pb.PtrToGenericPhysicalMessage) *wtype.GenericPhysical {
	if arg == nil {
		log.Println("Arg for PtrToGenericPhysical was nil")
		return nil
	}

	ret := DecodeGenericPhysical(arg.Arg_1)
	return &ret
}
func EncodeGenericPrefixedUnit(arg wunit.GenericPrefixedUnit) *pb.GenericPrefixedUnitMessage {
	ret := pb.GenericPrefixedUnitMessage{EncodeGenericUnit(arg.GenericUnit), EncodeSIPrefix(arg.SPrefix)}
	return &ret
}
func DecodeGenericPrefixedUnit(arg *pb.GenericPrefixedUnitMessage) wunit.GenericPrefixedUnit {
	ret := wunit.GenericPrefixedUnit{(wunit.GenericUnit)(DecodeGenericUnit(arg.Arg_1)), (wunit.SIPrefix)(DecodeSIPrefix(arg.Arg_2))}
	return ret
}
func EncodeGenericMatter(arg wtype.GenericMatter) *pb.GenericMatterMessage {
	ret := pb.GenericMatterMessage{(string)(arg.Iname), EncodeTemperature(arg.Imp), EncodeTemperature(arg.Ibp), EncodeSpecificHeatCapacity(arg.Ishc)}
	return &ret
}
func DecodeGenericMatter(arg *pb.GenericMatterMessage) wtype.GenericMatter {
	ret := wtype.GenericMatter{(string)(arg.Arg_1), (wunit.Temperature)(DecodeTemperature(arg.Arg_2)), (wunit.Temperature)(DecodeTemperature(arg.Arg_3)), (wunit.SpecificHeatCapacity)(DecodeSpecificHeatCapacity(arg.Arg_4))}
	return ret
}
func EncodeMass(arg wunit.Mass) *pb.MassMessage {
	ret := pb.MassMessage{EncodeConcreteMeasurement(arg.ConcreteMeasurement)}
	return &ret
}
func DecodeMass(arg *pb.MassMessage) wunit.Mass {
	ret := wunit.Mass{(wunit.ConcreteMeasurement)(DecodeConcreteMeasurement(arg.Arg_1))}
	return ret
}
func EncodeGenericUnit(arg wunit.GenericUnit) *pb.GenericUnitMessage {
	ret := pb.GenericUnitMessage{(string)(arg.StrName), (string)(arg.StrSymbol), (float64)(arg.FltConversionfactor), (string)(arg.StrBaseUnit)}
	return &ret
}
func DecodeGenericUnit(arg *pb.GenericUnitMessage) wunit.GenericUnit {
	ret := wunit.GenericUnit{(string)(arg.Arg_1), (string)(arg.Arg_2), (float64)(arg.Arg_3), (string)(arg.Arg_4)}
	return ret
}
func EncodeSIPrefix(arg wunit.SIPrefix) *pb.SIPrefixMessage {
	ret := pb.SIPrefixMessage{(string)(arg.Name), (float64)(arg.Value)}
	return &ret
}
func DecodeSIPrefix(arg *pb.SIPrefixMessage) wunit.SIPrefix {
	ret := wunit.SIPrefix{(string)(arg.Arg_1), (float64)(arg.Arg_2)}
	return ret
}
func EncodeSpecificHeatCapacity(arg wunit.SpecificHeatCapacity) *pb.SpecificHeatCapacityMessage {
	ret := pb.SpecificHeatCapacityMessage{EncodeConcreteMeasurement(arg.ConcreteMeasurement)}
	return &ret
}
func DecodeSpecificHeatCapacity(arg *pb.SpecificHeatCapacityMessage) wunit.SpecificHeatCapacity {
	ret := wunit.SpecificHeatCapacity{(wunit.ConcreteMeasurement)(DecodeConcreteMeasurement(arg.Arg_1))}
	return ret
}
