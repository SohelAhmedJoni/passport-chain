package main

import (
	"github.com/hyperledger-labs/passport-chain/chaincode/eventtypes"
	"github.com/hyperledger-labs/cc-tools/events"
)

var eventTypeList = []events.Event{
	eventtypes.CreateLibraryLog,
}
