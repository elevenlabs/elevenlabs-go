package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Read the workspace's current Reception plan and its hard limits.
//
// Call this before answering any plan, pricing, trial, or limit question so
// the assistant stops guessing (e.g. telling a paying owner they're on the
// trial, or hallucinating concurrent-call / receptionist / voice-cloning
// caps). The tool exposes only the shape of the plan — a display name;
// whether it's trial / paid / untiered; the entitlement caps (max
// receptionists, phone numbers, active locations, team members, concurrent
// calls, voice-cloning quota) alongside the current counts; and whether
// there's a payment issue. It deliberately does NOT return list prices or
// dollar amounts, nor the trial / renewal end date: pricing and renewal
// dates are a product decision and live in the Settings > Billing plan
// picker.
type GetWorkspacePlanParams struct {
	extraProperties map[string]interface{}
}

func (g *GetWorkspacePlanParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetWorkspacePlanParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWorkspacePlanParams to string"
	}
	return string(jsonData)
}

func (g *GetWorkspacePlanParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetWorkspacePlanParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetWorkspacePlanParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetWorkspacePlanParams) MarshalJSON() ([]byte, error) {
	type embed GetWorkspacePlanParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
