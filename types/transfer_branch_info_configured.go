package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferBranchInfoConfigured struct {
	branchReason    string
	BranchId        string `json:"branch_id" xml:"branch_id" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferBranchInfoConfigured) GetBranchId() (value string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TransferBranchInfoConfigured) SetBranchId(branchId string) {
	t.BranchId = branchId
}

func (t *TransferBranchInfoConfigured) BranchReason() string {
	return t.branchReason
}

func (t *TransferBranchInfoConfigured) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferBranchInfoConfigured) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferBranchInfoConfigured to string"
	}
	return string(jsonData)
}

func (t *TransferBranchInfoConfigured) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferBranchInfoConfigured
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferBranchInfoConfigured(tmp)
	var literals struct {
		BranchReason *string `json:"branch_reason"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.BranchReason != nil && *literals.BranchReason != "configured" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "configured", *literals.BranchReason)
	}
	t.branchReason = "configured"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "branch_reason")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferBranchInfoConfigured) MarshalJSON() ([]byte, error) {
	type embed TransferBranchInfoConfigured
	var marshaler = struct {
		embed
		BranchReason string `json:"branch_reason"`
	}{
		embed:        embed(*t),
		BranchReason: "configured",
	}
	return json.Marshal(marshaler)
}
