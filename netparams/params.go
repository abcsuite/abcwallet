// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import "github.com/abcsuite/abcd/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	JSONRPCClientPort string
	JSONRPCServerPort string
	GRPCServerPort    string
}

// MainNetParams contains parameters specific running abcwallet and
// abcd on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:            &chaincfg.MainNetParams,
	JSONRPCClientPort: "9528",
	JSONRPCServerPort: "9520",
	GRPCServerPort:    "9111",
}

// TestNet2Params contains parameters specific running abcwallet and
// abcd on the test network (version 2) (wire.TestNet2).
var TestNet2Params = Params{
	Params:            &chaincfg.TestNet2Params,
	JSONRPCClientPort: "19529",
	JSONRPCServerPort: "19520",
	GRPCServerPort:    "19111",
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:            &chaincfg.SimNetParams,
	JSONRPCClientPort: "19556",
	JSONRPCServerPort: "19557",
	GRPCServerPort:    "19558",
}
