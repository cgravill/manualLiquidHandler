// server/main.go: Part of the Antha language
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

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/antha-lang/manualLiquidHandler/ExtendedLiquidhandlingDriver"
	"github.com/antha-lang/manualLiquidHandler/cli"

	"os"

	"flag"

	"github.com/antha-lang/antha/bvendor/google.golang.org/grpc"
	"github.com/antha-lang/manualLiquidHandler"
	"io/ioutil"
)

var (
	port int
	view string
)

func main() {
	flag.IntVar(&port, "port", 50051, "Sepcify the port at which the server will be listening")
	flag.StringVar(&view, "view", "cli", "Specify the wished view to display the messages: cli | cui")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var x manualLiquidHandler.ManualExecuter
	if view == "cui" {
		//set log output to nil so that we don't bother cui interface
		log.SetOutput(ioutil.Discard)
		x = cli.NewCUI()
	} else if view == "cli" {
		x = cli.NewReadWriterExecutor(
			os.Stdin,
			os.Stdout,
		)
	} else {
		log.Printf("Unknow view given: %s\n", view)
		os.Exit(1)
	}

	x.Init()

	manual := manualLiquidHandler.NewManual(x)

	s := grpc.NewServer()
	ExtendedLiquidhandlingDriver.RegisterExtendedLiquidhandlingDriverServer(s, manual)
	go func() {
		log.Println("Listening at :", port)
		s.Serve(lis)
	}()
	x.Close()
}
