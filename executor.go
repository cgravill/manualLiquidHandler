// executor.go: Part of the Antha language
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
	"bytes"
	"fmt"
)

//Manual Executer interface to be implemented to execute actions as a manual driver
type ManualExecuter interface {
	//Init make whatever initializations in the background needed for the interaction
	Init() error
	//Execute a single request
	Execute(r *CLICommandRequest) CLICommandResult
	//Close actively exit the executor
	Close()
}

//MultiLevelMessage represents an aggregating struct to hold indentable strings
type MultiLevelMessage struct {
	Message  string
	Children []MultiLevelMessage
}

//LeveledString prints in the given buffer an indented version of the MultiLevelMessage using the level string as padding
func (m *MultiLevelMessage) LeveledString(level string, out *bytes.Buffer) {
	out.WriteString(fmt.Sprintf("%s%s\n", level, m.Message))
	for _, v := range m.Children {
		v.LeveledString(level+level, out)
	}
}

//String will return the whole MultiLevelMessage indented with spaces in a multiline string
func (m *MultiLevelMessage) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("%s\n", m.Message))
	for _, v := range m.Children {
		v.LeveledString("  ", &out)
	}
	return out.String()
}
func NewMultiLevelMessage(message string, children []MultiLevelMessage) *MultiLevelMessage {
	mlm := new(MultiLevelMessage)
	mlm.Message = message
	mlm.Children = children
	return mlm
}
func (m *MultiLevelMessage) ChildrenText() string {
	var out bytes.Buffer
	for _, v := range m.Children {
		v.LeveledString("  ", &out)
	}
	return out.String()
}

//CommandRequest represents a request to command line in which a question is made, a set of options are presented and
// one of those options must be selected. Expected is the option which should be selected for the command to execute
// successfully
type CLICommandRequest struct {
	Id      string
	Message MultiLevelMessage
}

//NewCLICommandRequest instantiates a new command request with the given id and the specified MultiLevelMessage
func NewCLICommandRequest(id string, message MultiLevelMessage) *CLICommandRequest {
	cr := new(CLICommandRequest)
	cr.Id = id
	cr.Message = message
	return cr
}

//CommandResult is the result of a CommandRequest. It has a bool representing whether the action was performed successfully
// and an Answer holding a string with a textual representation of the result should it be needed
type CLICommandResult struct {
	Id    string
	Error error
}

//NewCLICommandResult returns a new result for the specified request id with the given error
func NewCLICommandResult(id string, err error) *CLICommandResult {
	cr := new(CLICommandResult)
	cr.Id = id
	cr.Error = err
	return cr
}