package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResellerWorkspaceOnboardingResponse struct {
	OnboardingCompletedAtUnix *int `json:"onboarding_completed_at_unix" xml:"onboarding_completed_at_unix" required:"true"`
	// The receptionist onboarding created or promoted.
	CustomerFacingAgentId string `json:"customer_facing_agent_id" xml:"customer_facing_agent_id" required:"true"`
	extraProperties       map[string]interface{}
}

func (r *ResellerWorkspaceOnboardingResponse) GetOnboardingCompletedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.OnboardingCompletedAtUnix
}

func (r *ResellerWorkspaceOnboardingResponse) GetCustomerFacingAgentId() (value string) {
	if r == nil {
		return
	}
	return r.CustomerFacingAgentId
}

func (r *ResellerWorkspaceOnboardingResponse) SetOnboardingCompletedAtUnix(onboardingCompletedAtUnix *int) {
	r.OnboardingCompletedAtUnix = onboardingCompletedAtUnix
}

func (r *ResellerWorkspaceOnboardingResponse) SetCustomerFacingAgentId(customerFacingAgentId string) {
	r.CustomerFacingAgentId = customerFacingAgentId
}

func (r *ResellerWorkspaceOnboardingResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerWorkspaceOnboardingResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerWorkspaceOnboardingResponse to string"
	}
	return string(jsonData)
}

func (r *ResellerWorkspaceOnboardingResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerWorkspaceOnboardingResponse) MarshalJSON() ([]byte, error) {
	type embed ResellerWorkspaceOnboardingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
