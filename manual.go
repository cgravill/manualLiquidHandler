// manual.go: Part of the Antha language
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
	"fmt"

	liquidhandlingDriver "github.com/antha-lang/manualLiquidHandler/ExtendedLiquidhandlingDriver"
	"github.com/antha-lang/manualLiquidHandler/internal/github.com/twinj/uuid"
	"github.com/antha-lang/manualLiquidHandler/internal/golang.org/x/net/context"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
)

//Manual struct that implements the liquidHandlerDriver and routes all functions via the underlying ManualExecuter
type Manual struct {
	x ManualExecuter
}

//NewManual creates a Manual struct wrapping the given Executer
func NewManual(x ManualExecuter) *Manual {
	return &Manual{x}
}

func (m *Manual) AddPlateTo(c context.Context, r *liquidhandlingDriver.AddPlateToRequest) (*liquidhandlingDriver.AddPlateToReply, error) {
	fullArg := r.GetArg_2()
	if fullArg == nil {
		return nil, fmt.Errorf("Plate was nil. Plate contents are mandatory.")
	}
	p, err := DecodeGenericPlate(fullArg.Arg_1)
	if err != nil {
		panic(err)
	}
	var description string
	var dest string
	//Extract Plate id, destination
	switch what := p.(type) {
	case wtype.LHPlate:
		description = fmt.Sprintf("Plate %s", what.ID)
	case wtype.LHTipbox:
		description = fmt.Sprintf("Tip Box %s", what.ID)
	case wtype.LHTipwaste:
		description = fmt.Sprintf("Tip Waste %s", what.ID)
	}

	dest = r.Arg_3

	mess := make([]string, 0)
	mess = append(mess, description)
	mess = append(mess, dest)
	res := m.translateCall("Add Plate", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.AddPlateToReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"ADDPLATETO ACK",
		},
	}, nil
}

func (m *Manual) Aspirate(c context.Context, r *liquidhandlingDriver.AspirateRequest) (*liquidhandlingDriver.AspirateReply, error) {
	vols := r.GetArg_1().Arg_1
	whats := r.GetArg_6().Arg_1
	if len(vols) != len(whats) {
		return nil, fmt.Errorf("Expecting the same size for volumes (%d) and reagents (%d)", len(vols), len(whats))
	}
	mess := make([]string, 0)
	for i := 0; i < len(vols); i++ {
		mess = append(mess, fmt.Sprintf("Aspirate %g ul of %s", vols[i], whats[i]))
	}
	res := m.translateCall("Aspirate", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.AspirateReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"ASPIRATE ACK",
		},
	}, nil
}

func (m *Manual) Close(c context.Context, r *liquidhandlingDriver.CloseRequest) (*liquidhandlingDriver.CloseReply, error) {
	//res := m.translateCall("CLOSE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.CloseReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"CLOSE ACK",
		},
	}, nil
}

func (m *Manual) Dispense(c context.Context, r *liquidhandlingDriver.DispenseRequest) (*liquidhandlingDriver.DispenseReply, error) {
	vols := r.GetArg_1().Arg_1
	whats := r.GetArg_6().Arg_1
	if len(vols) != len(whats) {
		return nil, fmt.Errorf("Expecting the same size for volumes (%d) and reagents (%d)", len(vols), len(whats))
	}
	mess := make([]string, 0)
	for i := 0; i < len(vols); i++ {
		mess = append(mess, fmt.Sprintf("Dispense %g ul of %s", vols[i], whats[i]))
	}

	res := m.translateCall("Dispense", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.DispenseReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"DISPENSE ACK",
		},
	}, nil
}
func (m *Manual) Finalize(c context.Context, r *liquidhandlingDriver.FinalizeRequest) (*liquidhandlingDriver.FinalizeReply, error) {
	res := m.translateCall("Finalize", fmt.Sprintf("The protocol has finished."))
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.FinalizeReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"FINALIZE ACK",
		},
	}, nil
}
func (m *Manual) GetCapabilities(c context.Context, r *liquidhandlingDriver.GetCapabilitiesRequest) (*liquidhandlingDriver.GetCapabilitiesReply, error) {
	//res := m.translateCall("GETCAPABILITIES", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GetCapabilitiesReply{
		nil,
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GETCAPABILITIES ACK",
		},
	}, nil
}
func (m *Manual) GetCurrentPosition(c context.Context, r *liquidhandlingDriver.GetCurrentPositionRequest) (*liquidhandlingDriver.GetCurrentPositionReply, error) {
	//res := m.translateCall("STOP", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GetCurrentPositionReply{
		"",
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GETCURRENTPOSITION ACK",
		},
	}, nil
}
func (m *Manual) GetHeadState(c context.Context, r *liquidhandlingDriver.GetHeadStateRequest) (*liquidhandlingDriver.GetHeadStateReply, error) {
	//res := m.translateCall("GETHEADSTATE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GetHeadStateReply{
		"",
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GETHEADSTATE ACK",
		},
	}, nil
}
func (m *Manual) GetPositionState(c context.Context, r *liquidhandlingDriver.GetPositionStateRequest) (*liquidhandlingDriver.GetPositionStateReply, error) {
	//res := m.translateCall("GETPOSITIONSTATE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GetPositionStateReply{
		"",
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GETPOSITIONSTATE ACK",
		},
	}, nil
}
func (m *Manual) GetStatus(c context.Context, r *liquidhandlingDriver.GetStatusRequest) (*liquidhandlingDriver.GetStatusReply, error) {
	//res := m.translateCall("GETSTATUS", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GetStatusReply{
		nil,
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GETSTATUS ACK",
		},
	}, nil
}
func (m *Manual) Go(c context.Context, r *liquidhandlingDriver.GoRequest) (*liquidhandlingDriver.GoReply, error) {
	//res := m.translateCall("GO", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.GoReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"GO ACK",
		},
	}, nil
}
func (m *Manual) Initialize(c context.Context, r *liquidhandlingDriver.InitializeRequest) (*liquidhandlingDriver.InitializeReply, error) {
	//res := m.translateCall("INITIALIZE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.InitializeReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"INITIALIZE ACK",
		},
	}, nil
}
func (m *Manual) LightsOff(c context.Context, r *liquidhandlingDriver.LightsOffRequest) (*liquidhandlingDriver.LightsOffReply, error) {
	//res := m.translateCall("LIGHTSOFF", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.LightsOffReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"LIGHTSOFF ACK",
		},
	}, nil
}
func (m *Manual) LightsOn(c context.Context, r *liquidhandlingDriver.LightsOnRequest) (*liquidhandlingDriver.LightsOnReply, error) {
	//res := m.translateCall("LIGHTSON", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.LightsOnReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"LIGHTSON ACK",
		},
	}, nil
}
func (m *Manual) LoadAdaptor(c context.Context, r *liquidhandlingDriver.LoadAdaptorRequest) (*liquidhandlingDriver.LoadAdaptorReply, error) {
	//res := m.translateCall("LOADADAPTOR", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.LoadAdaptorReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"LOADADAPTOR ACK",
		},
	}, nil
}
func (m *Manual) LoadHead(c context.Context, r *liquidhandlingDriver.LoadHeadRequest) (*liquidhandlingDriver.LoadHeadReply, error) {
	//res := m.translateCall("LOADHEAD", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.LoadHeadReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"LOADHEAD ACK",
		},
	}, nil
}

func (m *Manual) LoadTips(c context.Context, r *liquidhandlingDriver.LoadTipsRequest) (*liquidhandlingDriver.LoadTipsReply, error) {
	positions := r.GetArg_5().Arg_1
	mess := make([]string, 0)
	for _, v := range positions {
		mess = append(mess, fmt.Sprintf("Unload Tip at position %s", v))
	}
	res := m.translateCall("Load Tips", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.LoadTipsReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"LOADTIPS ACK",
		},
	}, nil
}

func (m *Manual) Message(c context.Context, r *liquidhandlingDriver.MessageRequest) (*liquidhandlingDriver.MessageReply, error) {
	//res := m.translateCall("MESSAGE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.MessageReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"MESSAGE ACK",
		},
	}, nil
}

func (m *Manual) Mix(c context.Context, r *liquidhandlingDriver.MixRequest) (*liquidhandlingDriver.MixReply, error) {
	vols := r.GetArg_2().Arg_1
	/*
		fvols := r.GetArg_3().Arg_1
		if len(vols) != len(fvols) {
			return nil, fmt.Errorf("Expecting the same length for vols (%d) and final vols (%d)", len(vols), len(fvols))
		}
	*/
	mess := make([]string, 0)
	for i := 0; i < len(vols); i++ {
		mess = append(mess, fmt.Sprintf("Mix the contents using a volume of %g ul", vols[i]))
	}
	res := m.translateCall("Mix", r)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.MixReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"MIX ACK",
		},
	}, nil
}
func (m *Manual) Move(c context.Context, r *liquidhandlingDriver.MoveRequest) (*liquidhandlingDriver.MoveReply, error) {
	//TODO save positions of movement so that they can be attached to aspirate / dispense .. actions
	decks := r.GetArg_1().Arg_1
	wells := r.GetArg_2().Arg_1
	if len(decks) != len(wells) {
		return nil, fmt.Errorf("Expecting the same size for decks (%d) and wells (%d)", len(decks), len(wells))
	}
	mess := make([]string, 0)
	for i := 0; i < len(decks); i++ {
		mess = append(mess, fmt.Sprintf("Move to deck %s well %s", decks[i], wells[i]))
	}
	res := m.translateCall("Move ", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.MoveReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"MOVE ACK",
		},
	}, nil
}
func (m *Manual) MoveRaw(c context.Context, r *liquidhandlingDriver.MoveRawRequest) (*liquidhandlingDriver.MoveRawReply, error) {
	//res := m.translateCall("MOVERAW", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.MoveRawReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"MOVERAW ACK",
		},
	}, nil
}
func (m *Manual) Open(c context.Context, r *liquidhandlingDriver.OpenRequest) (*liquidhandlingDriver.OpenReply, error) {
	//res := m.translateCall("OPEN", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.OpenReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"OPEN ACK",
		},
	}, nil
}
func (m *Manual) RemoveAllPlates(c context.Context, r *liquidhandlingDriver.RemoveAllPlatesRequest) (*liquidhandlingDriver.RemoveAllPlatesReply, error) {
	//This call has no further information.
	mess := make([]string, 0)
	mess = append(mess, fmt.Sprintf("Remove all plates from the deck."))
	res := m.translateCall("Remove All Plates", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.RemoveAllPlatesReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"REMOVEALLPLATES ACK",
		},
	}, nil
}
func (m *Manual) RemovePlateAt(c context.Context, r *liquidhandlingDriver.RemovePlateAtRequest) (*liquidhandlingDriver.RemovePlateAtReply, error) {
	//res := m.translateCall("REMOVEPLATEAT", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.RemovePlateAtReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"REMOVEPLATEAT ACK",
		},
	}, nil
}
func (m *Manual) ResetPistons(c context.Context, r *liquidhandlingDriver.ResetPistonsRequest) (*liquidhandlingDriver.ResetPistonsReply, error) {
	//res := m.translateCall("RESETPISTONS", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.ResetPistonsReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"RESETPISTONS ACK",
		},
	}, nil
}
func (m *Manual) SetDriveSpeed(c context.Context, r *liquidhandlingDriver.SetDriveSpeedRequest) (*liquidhandlingDriver.SetDriveSpeedReply, error) {
	//res := m.translateCall("SETDRIVESPEED", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.SetDriveSpeedReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"SETDRIVESPEED ACK",
		},
	}, nil
}
func (m *Manual) SetPipetteSpeed(c context.Context, r *liquidhandlingDriver.SetPipetteSpeedRequest) (*liquidhandlingDriver.SetPipetteSpeedReply, error) {
	//res := m.translateCall("SETPIPETTESPEED", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.SetPipetteSpeedReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"SETPIPETTESPEED ACK",
		},
	}, nil
}
func (m *Manual) SetPositionState(c context.Context, r *liquidhandlingDriver.SetPositionStateRequest) (*liquidhandlingDriver.SetPositionStateReply, error) {
	//res := m.translateCall("SETPOSITIONSTATE", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.SetPositionStateReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"SETPOSITIONSTATE ACK",
		},
	}, nil
}
func (m *Manual) Stop(c context.Context, r *liquidhandlingDriver.StopRequest) (*liquidhandlingDriver.StopReply, error) {
	//res := m.translateCall("STOP", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.StopReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"Stop ACK",
		},
	}, nil
}
func (m *Manual) UnloadAdaptor(c context.Context, r *liquidhandlingDriver.UnloadAdaptorRequest) (*liquidhandlingDriver.UnloadAdaptorReply, error) {
	//res := m.translateCall("UNLOADADAPTOR", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.UnloadAdaptorReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"UNLOADADAPTOR ACK",
		},
	}, nil
}
func (m *Manual) UnloadHead(c context.Context, r *liquidhandlingDriver.UnloadHeadRequest) (*liquidhandlingDriver.UnloadHeadReply, error) {
	//res := m.translateCall("UNLOADHEAD", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.UnloadHeadReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"UNLOADHEAD ACK",
		},
	}, nil
}
func (m *Manual) UnloadTips(c context.Context, r *liquidhandlingDriver.UnloadTipsRequest) (*liquidhandlingDriver.UnloadTipsReply, error) {
	positions := r.GetArg_5().Arg_1
	mess := make([]string, 0)
	for _, v := range positions {
		mess = append(mess, fmt.Sprintf("Unload Tip at position %s", v))
	}
	res := m.translateCall("Unload Tips", mess)
	if res.Error != nil {
		return nil, res.Error
	}
	return &liquidhandlingDriver.UnloadTipsReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"UNLOADTIPS ACK",
		},
	}, nil
}
func (m *Manual) UpdateMetaData(c context.Context, r *liquidhandlingDriver.UpdateMetaDataRequest) (*liquidhandlingDriver.UpdateMetaDataReply, error) {
	//	if p1 := r.GetArg_1(); p1 != nil {
	//		if p2 := p1.GetArg_1(); p2 != nil {
	//			var lhp liquidhandling.LHProperties
	//			lhp = client.DecodeLHProperties(p2)
	//		}
	//	}
	//	res := m.translateCall("UPDATEMETADATA", r)
	//	if res.Error != nil {
	//		return nil, res.Error
	//	}
	return &liquidhandlingDriver.UpdateMetaDataReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"UPDATEMETADATA ACK",
		},
	}, nil
}
func (m *Manual) Wait(c context.Context, r *liquidhandlingDriver.WaitRequest) (*liquidhandlingDriver.WaitReply, error) {
	//res := m.translateCall("WAIT", r)
	//if res.Error != nil {
	//	return nil, res.Error
	//}
	return &liquidhandlingDriver.WaitReply{
		&liquidhandlingDriver.CommandStatusMessage{
			true,
			0,
			"WAIT ACK",
		},
	}, nil
}

func (m *Manual) translateCall(message string, more interface{}) CLICommandResult {
	id := uuid.NewV4().String()
	req := NewCLICommandRequest(id, *NewMultiLevelMessage(message, ToMultiLevelMessage(more)))
	return m.x.Execute(req)
}