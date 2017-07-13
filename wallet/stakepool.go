// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"errors"

	"github.com/abcsuite/abcutil"
	"github.com/abcsuite/abcwallet/wallet/udb"
	"github.com/abcsuite/abcwallet/walletdb"
)

// StakePoolUserInfo returns the stake pool user information for a user
// identified by their P2SH voting address.
func (w *Wallet) StakePoolUserInfo(userAddress abcutil.Address) (*udb.StakePoolUser, error) {
	switch userAddress.(type) {
	case *abcutil.AddressPubKeyHash: // ok
	case *abcutil.AddressScriptHash: // ok
	default:
		return nil, errors.New("stake pool user address must be P2PKH or P2SH")
	}

	var user *udb.StakePoolUser
	err := walletdb.View(w.db, func(tx walletdb.ReadTx) error {
		stakemgrNs := tx.ReadBucket(wstakemgrNamespaceKey)
		var err error
		user, err = w.StakeMgr.StakePoolUserInfo(stakemgrNs, userAddress)
		return err
	})
	return user, err
}
