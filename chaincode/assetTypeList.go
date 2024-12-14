package main

import (
	"github.com/hyperledger-labs/passport-chain/chaincode/assettypes"
	"github.com/hyperledger-labs/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Person,
	assettypes.Book,
	assettypes.Library,
	assettypes.Secret,
}
