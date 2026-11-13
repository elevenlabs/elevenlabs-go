package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TransferToAgentToolResultSuccessModelOutputBranchInfo struct {
	BranchReason     string
	Configured       *TransferBranchInfoConfigured
	DefaultingToMain *TransferBranchInfoDefaultingToMain
	TrafficSplit     *TransferBranchInfoTrafficSplit
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) GetConfigured() *TransferBranchInfoConfigured {
	if t == nil {
		return nil
	}
	if t.Configured == nil {
		return nil
	}
	return t.Configured
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) GetDefaultingToMain() *TransferBranchInfoDefaultingToMain {
	if t == nil {
		return nil
	}
	if t.DefaultingToMain == nil {
		return nil
	}
	return t.DefaultingToMain
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) GetTrafficSplit() *TransferBranchInfoTrafficSplit {
	if t == nil {
		return nil
	}
	if t.TrafficSplit == nil {
		return nil
	}
	return t.TrafficSplit
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) GetBranchReason() string {
	if t == nil {
		return ""
	}
	return t.BranchReason
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		BranchReason string `json:"branch_reason"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.BranchReason = unmarshaler.BranchReason
	if unmarshaler.BranchReason == "" {
		return fmt.Errorf("%T did not include discriminant branch_reason", t)
	}
	switch unmarshaler.BranchReason {
	case "configured":
		value := new(TransferBranchInfoConfigured)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Configured = value
	case "defaulting_to_main":
		value := new(TransferBranchInfoDefaultingToMain)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.DefaultingToMain = value
	case "traffic_split":
		value := new(TransferBranchInfoTrafficSplit)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.TrafficSplit = value
	}
	return nil
}

func (t TransferToAgentToolResultSuccessModelOutputBranchInfo) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Configured != nil {
		return internal.MarshalJSONWithDiscriminant(t.Configured, "branch_reason", "configured")
	}
	if t.DefaultingToMain != nil {
		return internal.MarshalJSONWithDiscriminant(t.DefaultingToMain, "branch_reason", "defaulting_to_main")
	}
	if t.TrafficSplit != nil {
		return internal.MarshalJSONWithDiscriminant(t.TrafficSplit, "branch_reason", "traffic_split")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TransferToAgentToolResultSuccessModelOutputBranchInfoVisitor interface {
	VisitConfigured(*TransferBranchInfoConfigured) error
	VisitDefaultingToMain(*TransferBranchInfoDefaultingToMain) error
	VisitTrafficSplit(*TransferBranchInfoTrafficSplit) error
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) Accept(visitor TransferToAgentToolResultSuccessModelOutputBranchInfoVisitor) error {
	if t.Configured != nil {
		return visitor.VisitConfigured(t.Configured)
	}
	if t.DefaultingToMain != nil {
		return visitor.VisitDefaultingToMain(t.DefaultingToMain)
	}
	if t.TrafficSplit != nil {
		return visitor.VisitTrafficSplit(t.TrafficSplit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TransferToAgentToolResultSuccessModelOutputBranchInfo) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Configured != nil {
		fields = append(fields, "configured")
	}
	if t.DefaultingToMain != nil {
		fields = append(fields, "defaulting_to_main")
	}
	if t.TrafficSplit != nil {
		fields = append(fields, "traffic_split")
	}
	if len(fields) == 0 {
		if t.BranchReason != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.BranchReason)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.BranchReason != "" {
		field := fields[0]
		if t.BranchReason != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.BranchReason,
				t,
			)
		}
	}
	return nil
}
