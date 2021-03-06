/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ordererledger_test

import (
	"io/ioutil"
	"os"

	"github.com/hyperledger/fabric/common/configtx/tool/provisional"
	. "github.com/hyperledger/fabric/orderer/ledger"
	jsonledger "github.com/hyperledger/fabric/orderer/ledger/json"
	cb "github.com/hyperledger/fabric/protos/common"
)

var genesisBlock = cb.NewBlock(0, nil)

func init() {
	testables = append(testables, &jsonLedgerTestEnv{})
}

type jsonLedgerTestFactory struct {
	location string
}

type jsonLedgerTestEnv struct {
}

func (env *jsonLedgerTestEnv) Initialize() (ledgerTestFactory, error) {
	var err error
	location, err := ioutil.TempDir("", "hyperledger")
	if err != nil {
		return nil, err
	}
	return &jsonLedgerTestFactory{location: location}, nil
}

func (env *jsonLedgerTestEnv) Name() string {
	return "jsonledger"
}

func (env *jsonLedgerTestFactory) Destroy() error {
	err := os.RemoveAll(env.location)
	return err
}

func (env *jsonLedgerTestFactory) Persistent() bool {
	return true
}

func (env *jsonLedgerTestFactory) New() (Factory, ReadWriter) {
	flf := jsonledger.New(env.location)
	fl, err := flf.GetOrCreate(provisional.TestChainID)
	if err != nil {
		panic(err)
	}
	fl.Append(genesisBlock)
	return flf, fl
}
