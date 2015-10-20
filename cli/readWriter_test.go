// cli/readWriter_test.go: Part of the Antha language
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

package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/antha-lang/manualLiquidHandler"
)

func TestExecuteNoError(t *testing.T) {
	in := strings.NewReader("y\ny\n") //this is an enter, should be trimmed to an empty input, which is no error
	out := bytes.Buffer{}
	mlh := NewReadWriterExecutor(in, &out)
	mlh.Init()
	defer mlh.Close()

	ret := mlh.Execute(manualLiquidHandler.NewCLICommandRequest("test", *manualLiquidHandler.NewMultiLevelMessage("test Message", nil)))
	if ret.Id != "test" {
		t.Fatalf("Request and Result Ids do not match. Expected %s, got %s.", "test", ret.Id)
	}
	if ret.Error != nil {
		t.Fatalf("Unexpected Error in result: %v", ret.Error)
	}
	mlh.quit() //manually do this, would be like pressing ctrl-c
}

func TestExecuteError(t *testing.T) {
	errMsg := "TestError 666"
	in := strings.NewReader(fmt.Sprintf("%s\n", errMsg)) //this is an enter, should be trimmed to an empty input, which is no error
	out := bytes.Buffer{}
	mlh := NewReadWriterExecutor(in, &out)
	mlh.Init()
	defer mlh.Close()

	ret := mlh.Execute(manualLiquidHandler.NewCLICommandRequest("test", *manualLiquidHandler.NewMultiLevelMessage("test Message", nil)))
	if ret.Id != "test" {
		t.Fatalf("Request and Result Ids do not match. Expected %s, got %s.", "test", ret.Id)
	}
	if ret.Error == nil {
		t.Fatalf("Expecting Error on result, none found")
	}
	if ret.Error.Error() != errMsg {
		t.Fatalf("On Test Error message. Expecting %s, got %s", errMsg, ret.Error)
	}
	mlh.quit() //manually do this, would be like pressing ctrl-c
}