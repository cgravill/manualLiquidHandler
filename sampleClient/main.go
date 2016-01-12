// sampleClient/main.go: Part of the Antha language
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
	"sync"

	"github.com/antha-lang/antha/bvendor/google.golang.org/grpc"
	"github.com/antha-lang/manualLiquidHandler/ExtendedLiquidhandlingDriver"
	"github.com/antha-lang/manualLiquidHandler/internal/golang.org/x/net/context"
)

const (
	port = ":50051"
	host = "localhost"
)

func main() {

	conn, err := grpc.Dial(fmt.Sprintf("%s%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ExtendedLiquidhandlingDriver.NewExtendedLiquidhandlingDriverClient(conn)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		// Contact the server and print out its response.
		r, err := c.Go(context.Background(), &ExtendedLiquidhandlingDriver.GoRequest{})
		if err != nil {
			log.Printf("could not execute Stop: %v", err)
		} else {
			log.Printf("Done 1: %v", r)
		}
		wg.Done()
	}()
	go func() {
		// Contact the server and print out its response.
		r, err := c.Stop(context.Background(), &ExtendedLiquidhandlingDriver.StopRequest{})
		if err != nil {
			log.Printf("could not execute Stop: %v", err)
		} else {
			log.Printf("Done 2: %v", r)
		}
		wg.Done()
	}()
	go func() {
		// Contact the server and print out its response.
		r, err := c.RemoveAllPlates(context.Background(), &ExtendedLiquidhandlingDriver.RemoveAllPlatesRequest{})
		if err != nil {
			log.Printf("could not execute RemoveAllPlates: %v", err)
		} else {
			log.Printf("Done 3: %v", r)
		}
		wg.Done()
	}()

	wg.Wait()
}
