// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build none

/*

   The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

       go run mkalloc.go genesis.json

*/
package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/hypnosisfoundation/go-hypnosis/consensus/dpos"
	"github.com/hypnosisfoundation/go-hypnosis/core"
	"github.com/hypnosisfoundation/go-hypnosis/rlp"
)

type allocItem struct {
	Addr    *big.Int
	Balance *big.Int
	Code    []byte
}

type allocList []allocItem

func (a allocList) Len() int           { return len(a) }
func (a allocList) Less(i, j int) bool { return a[i].Addr.Cmp(a[j].Addr) < 0 }
func (a allocList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func makelist(g *core.Genesis) allocList {
	a := make(allocList, 0, len(g.Alloc))
	for addr, account := range g.Alloc {
		if len(account.Storage) > 0 || account.Nonce != 0 {
			panic(fmt.Sprintf("can't encode account %x", addr))
		}
		bigAddr := new(big.Int).SetBytes(addr.Bytes())
		a = append(a, allocItem{bigAddr, account.Balance, account.Code})
	}
	sort.Sort(a)
	return a
}

func makealloc(g *core.Genesis) string {
	a := makelist(g)
	data, err := rlp.EncodeToBytes(a)
	if err != nil {
		panic(err)
	}
	return strconv.QuoteToASCII(string(data))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g2 := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(g2); err != nil {
		panic(err)
	}

	ga := new(core.GenesisAlloc)
	err = json.NewDecoder(strings.NewReader(dpos.GenesisAlloc)).Decode(ga)
	if err != nil {
		panic(err)
	}

	for i, v := range g2.Alloc {
		(*ga)[i] = v
	}
	g2.Alloc = *ga

	allocData := makealloc(g2)
	outputFile, err := os.OpenFile("./build/bin/genesisAllocRlpData", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	outputFile.Write([]byte(allocData))
}
