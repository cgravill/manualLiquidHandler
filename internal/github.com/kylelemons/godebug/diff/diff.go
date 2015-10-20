// internal/github.com/kylelemons/godebug/diff/diff.go: Part of the Antha language
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

// Package diff implements a linewise diff algorithm.
package diff

import (
	"bytes"
	"fmt"
	"strings"
)

// Chunk represents a piece of the diff.  A chunk will not have both added and
// deleted lines.  Equal lines are always after any added or deleted lines.
// A Chunk may or may not have any lines in it, especially for the first or last
// chunk in a computation.
type Chunk struct {
	Added   []string
	Deleted []string
	Equal   []string
}

// Diff returns a string containing a line-by-line unified diff of the linewise
// changes required to make A into B.  Each line is prefixed with '+', '-', or
// ' ' to indicate if it should be added, removed, or is correct respectively.
func Diff(A, B string) string {
	aLines := strings.Split(A, "\n")
	bLines := strings.Split(B, "\n")

	chunks := DiffChunks(aLines, bLines)

	buf := new(bytes.Buffer)
	for _, c := range chunks {
		for _, line := range c.Added {
			fmt.Fprintf(buf, "+%s\n", line)
		}
		for _, line := range c.Deleted {
			fmt.Fprintf(buf, "-%s\n", line)
		}
		for _, line := range c.Equal {
			fmt.Fprintf(buf, " %s\n", line)
		}
	}
	return strings.TrimRight(buf.String(), "\n")
}

// DiffChunks uses an O(D(N+M)) shortest-edit-script algorithm
// to compute the edits required from A to B and returns the
// edit chunks.
func DiffChunks(A, B []string) []Chunk {
	// algorithm: http://www.xmailserver.org/diff2.pdf

	N, M := len(A), len(B)
	MAX := N + M
	V := make([]int, 2*MAX+1)
	Vs := make([][]int, 0, 8)

	var D int
dLoop:
	for D = 0; D <= MAX; D++ {
		for k := -D; k <= D; k += 2 {
			var x int
			if k == -D || (k != D && V[MAX+k-1] < V[MAX+k+1]) {
				x = V[MAX+k+1]
			} else {
				x = V[MAX+k-1] + 1
			}
			y := x - k
			for x < N && y < M && A[x] == B[y] {
				x++
				y++
			}
			V[MAX+k] = x
			if x >= N && y >= M {
				Vs = append(Vs, append(make([]int, 0, len(V)), V...))
				break dLoop
			}
		}
		Vs = append(Vs, append(make([]int, 0, len(V)), V...))
	}
	if D == 0 {
		return nil
	}
	chunks := make([]Chunk, D+1)

	x, y := N, M
	for d := D; d > 0; d-- {
		V := Vs[d]
		k := x - y
		insert := k == -d || (k != d && V[MAX+k-1] < V[MAX+k+1])

		x1 := V[MAX+k]
		var x0, xM, kk int
		if insert {
			kk = k + 1
			x0 = V[MAX+kk]
			xM = x0
		} else {
			kk = k - 1
			x0 = V[MAX+kk]
			xM = x0 + 1
		}
		y0 := x0 - kk

		var c Chunk
		if insert {
			c.Added = B[y0:][:1]
		} else {
			c.Deleted = A[x0:][:1]
		}
		if xM < x1 {
			c.Equal = A[xM:][:x1-xM]
		}

		x, y = x0, y0
		chunks[d] = c
	}
	if x > 0 {
		chunks[0].Equal = A[:x]
	}
	return chunks
}