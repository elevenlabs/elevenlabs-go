package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferBranchInfoTrafficSplit struct {
	branchReason      string
	BranchId          string  `json:"branch_id" xml:"branch_id" required:"true"`
	TrafficPercentage float64 `json:"traffic_percentage" xml:"traffic_percentage" required:"true"`
	extraProperties   map[string]interface{}
}

func (t *TransferBranchInfoTrafficSplit) GetBranchId() (value string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TransferBranchInfoTrafficSplit) GetTrafficPercentage() (value float64) {
	if t == nil {
		return
	}
	return t.TrafficPercentage
}

func (t *TransferBranchInfoTrafficSplit) SetBranchId(branchId string) {
	t.BranchId = branchId
}

func (t *TransferBranchInfoTrafficSplit) SetTrafficPercentage(trafficPercentage float64) {
	t.TrafficPercentage = trafficPercentage
}

func (t *TransferBranchInfoTrafficSplit) BranchReason() string {
	return t.branchReason
}

func (t *TransferBranchInfoTrafficSplit) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferBranchInfoTrafficSplit) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferBranchInfoTrafficSplit to string"
	}
	return string(jsonData)
}

func (t *TransferBranchInfoTrafficSplit) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferBranchInfoTrafficSplit
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferBranchInfoTrafficSplit(tmp)
	var literals struct {
		BranchReason *string `json:"branch_reason"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.BranchReason != nil && *literals.BranchReason != "traffic_split" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "traffic_split", *literals.BranchReason)
	}
	t.branchReason = "traffic_split"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "branch_reason")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferBranchInfoTrafficSplit) MarshalJSON() ([]byte, error) {
	type embed TransferBranchInfoTrafficSplit
	var marshaler = struct {
		embed
		BranchReason string `json:"branch_reason"`
	}{
		embed:        embed(*t),
		BranchReason: "traffic_split",
	}
	return json.Marshal(marshaler)
}
