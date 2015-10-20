// cli/readWriter.go: Part of the Antha language
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
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"

	"github.com/antha-lang/manualLiquidHandler"
)

// ReadWriterExecutor Executor implementation using a reader and writer to write read the messages and feedback
type ReadWriterExecutor struct {
	r      io.Reader
	w      io.Writer
	close  chan interface{}
	cmdIn  chan manualLiquidHandler.CLICommandRequest
	cmdOut map[string]chan manualLiquidHandler.CLICommandResult
}

//NewReadWriterExecutor instantiates a new ReadWriterExecutor using the given io.Reader and io.Writer
func NewReadWriterExecutor(r io.Reader, w io.Writer) *ReadWriterExecutor {
	close := make(chan interface{})
	cmdIn := make(chan manualLiquidHandler.CLICommandRequest)
	cmdOut := make(map[string]chan manualLiquidHandler.CLICommandResult)
	return &ReadWriterExecutor{
		r,
		w,
		close,
		cmdIn,
		cmdOut,
	}
}

//Execute executes a commandRequest via the CUI and waits for an answer to that specific command. Responses are handled
// via request id. It is the caller responsibility to avoid duplicity
func (rw *ReadWriterExecutor) Execute(r *manualLiquidHandler.CLICommandRequest) manualLiquidHandler.CLICommandResult {
	rw.cmdOut[r.Id] = make(chan manualLiquidHandler.CLICommandResult)
	rw.cmdIn <- *r
	return <-rw.cmdOut[r.Id]
}

//Close waits for the user to exit the interface, then returns
func (rw *ReadWriterExecutor) Close() {
	<-rw.close
	return
}

func (rw *ReadWriterExecutor) quit() {
	close(rw.cmdIn)
	close(rw.close)
}

//Init starts background routines to wait for ctrl-c capturing and exiting and
// requests handling
func (rw *ReadWriterExecutor) Init() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			rw.quit()
		}
	}()
	go func() {
		for r := range rw.cmdIn {
			bw := bufio.NewWriter(rw.w)
			fmt.Fprintln(bw, r.Id, " > ", r.Message)
			fmt.Fprint(bw, ">> 'y' or Write error: ")
			bw.Flush()
			rr := bufio.NewReader(rw.r)
			res, _ := rr.ReadString('\n')
			res = strings.TrimSpace(res)
			var err error
			if res != "y" {
				err = errors.New(res)
			} else {
				err = nil
			}
			rw.cmdOut[r.Id] <- *manualLiquidHandler.NewCLICommandResult(r.Id, err)
			close(rw.cmdOut[r.Id])
		}
	}()
	return nil
}

func (rw *ReadWriterExecutor) asExecutor() manualLiquidHandler.ManualExecuter {
	return manualLiquidHandler.ManualExecuter(rw)
}