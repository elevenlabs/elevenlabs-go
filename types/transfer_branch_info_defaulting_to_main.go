package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferBranchInfoDefaultingToMain struct {
	branchReason    string
	BranchId        string `json:"branch_id" xml:"branch_id" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferBranchInfoDefaultingToMain) GetBranchId() (value string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TransferBranchInfoDefaultingToMain) SetBranchId(branchId string) {
	t.BranchId = branchId
}

func (t *TransferBranchInfoDefaultingToMain) BranchReason() string {
	return t.branchReason
}

func (t *TransferBranchInfoDefaultingToMain) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferBranchInfoDefaultingToMain) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferBranchInfoDefaultingToMain to string"
	}
	return string(jsonData)
}

func (t *TransferBranchInfoDefaultingToMain) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferBranchInfoDefaultingToMain
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferBranchInfoDefaultingToMain(tmp)
	var literals struct {
		BranchReason *string `json:"branch_reason"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.BranchReason != nil && *literals.BranchReason != "defaulting_to_main" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "defaulting_to_main", *literals.BranchReason)
	}
	t.branchReason = "defaulting_to_main"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "branch_reason")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferBranchInfoDefaultingToMain) MarshalJSON() ([]byte, error) {
	type embed TransferBranchInfoDefaultingToMain
	var marshaler = struct {
		embed
		BranchReason string `json:"branch_reason"`
	}{
		embed:        embed(*t),
		BranchReason: "defaulting_to_main",
	}
	return json.Marshal(marshaler)
}
