package types

import (
	"fmt"
)

// Exit branch taken when the Genesys Bot Connector session ends.
//
// These values are the intent names registered in the Genesys bot list, so they are
// what the Call Bot Connector node branches on. Adding a value here without also
// registering the intent in Genesys would produce an exit the flow cannot handle.
type GenesysBotOutcome string

const (
	GenesysBotOutcomeSuccess  GenesysBotOutcome = "success"
	GenesysBotOutcomeEscalate GenesysBotOutcome = "escalate"
)

func NewGenesysBotOutcomeFromString(s string) (GenesysBotOutcome, error) {
	switch s {
	case "success":
		return GenesysBotOutcomeSuccess, nil
	case "escalate":
		return GenesysBotOutcomeEscalate, nil
	}
	var t GenesysBotOutcome
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GenesysBotOutcome) Ptr() *GenesysBotOutcome {
	return &e
}
