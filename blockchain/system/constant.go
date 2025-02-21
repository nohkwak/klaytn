// Copyright 2023 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package system

import (
	"errors"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	contracts "github.com/klaytn/klaytn/contracts/system_contracts"
	"github.com/klaytn/klaytn/log"
)

var (
	logger = log.NewModuleLogger(log.Blockchain)

	// Canonical system contract names registered in Registry.
	AddressBookName = "AddressBook"
	GovParamName    = "GovParam"
	Kip103Name      = "TreasuryRebalance"
	Kip113Name      = "KIP113"

	AllContractNames = []string{
		AddressBookName,
		GovParamName,
		Kip103Name,
		Kip113Name,
	}

	// This is the keccak-256 hash of "eip1967.proxy.implementation" subtracted by 1 used in the
	// EIP-1967 proxy contract. See https://eips.ethereum.org/EIPS/eip-1967#implementation-slot
	ImplementationSlot = common.Hex2Bytes("360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")

	// Some system contracts are allocated at special addresses.
	AddressBookAddr = common.HexToAddress("0x0000000000000000000000000000000000000400") // TODO: replace contracts/reward/contract/utils.go
	RegistryAddr    = common.HexToAddress("0x0000000000000000000000000000000000000401")

	// System contract binaries to be injected at hardfork or used in testing.
	RegistryCode     = hexutil.MustDecode("0x" + contracts.RegistryBinRuntime)
	RegistryMockCode = hexutil.MustDecode("0x" + contracts.RegistryMockBinRuntime)
	Kip113Code       = hexutil.MustDecode("0x" + contracts.KIP113BinRuntime)
	Kip113MockCode   = hexutil.MustDecode("0x" + contracts.KIP113MockBinRuntime)

	ERC1967ProxyCode = hexutil.MustDecode("0x" + contracts.ERC1967ProxyBinRuntime)

	// Errors
	ErrRegistryNotInstalled = errors.New("Registry contract not installed")
	ErrKip113BadResult      = errors.New("KIP113 call returned bad data")
	ErrKip113BadPop         = errors.New("KIP113 PoP verification failed")
)
