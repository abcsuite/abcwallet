// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2017 The Aero Blockchain developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/abcsuite/abcd/abcjson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]abcjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*abcjson.CreateMultiSigResult)(nil)}},
	{"dumpprivkey", returnsString},
	{"getaccount", returnsString},
	{"getaccountaddress", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", append(returnsNumber, returnsNumber[0])},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*abcjson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*abcjson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"gettickets", []interface{}{(*abcjson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*abcjson.GetTransactionResult)(nil)}},
	{"getvotechoices", []interface{}{(*abcjson.GetVoteChoicesResult)(nil)}},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listlockunspent", []interface{}{(*[]abcjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]abcjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]abcjson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*abcjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*abcjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"redeemmultisigout", []interface{}{(*abcjson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*abcjson.RedeemMultiSigOutResult)(nil)}},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*abcjson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*abcjson.SignRawTransactionsResult)(nil)}},
	{"validateaddress", []interface{}{(*abcjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]abcjson.VersionResult)(nil)}},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"createnewaccount", nil},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*abcjson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"renameaccount", nil},
	{"walletislocked", returnsBool},
	{"walletinfo", []interface{}{(*abcjson.WalletInfoResult)(nil)}},

	// TODO Alphabetize
	{"purchaseticket", returnsString},
	{"sendtossrtx", returnsString},
	{"sendtosstx", returnsString},
	{"sendtossgen", returnsString},
	{"generatevote", []interface{}{(*abcjson.GenerateVoteResult)(nil)}},
	{"getstakeinfo", []interface{}{(*abcjson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"setticketfee", returnsBool},
	{"getwalletfee", returnsNumber},
	{"addticket", nil},
	{"listscripts", []interface{}{(*abcjson.ListScriptsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*abcjson.StakePoolUserInfoResult)(nil)}},
	{"ticketsforaddress", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
