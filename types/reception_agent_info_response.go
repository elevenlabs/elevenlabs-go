package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The business profile, with the current time in the business's timezone.
type ReceptionAgentInfoResponse struct {
	// Name of the business.
	BusinessName *string `json:"business_name,omitempty" xml:"business_name,omitempty"`
	// ISO 3166-1 alpha-2 country code where the business is based (e.g. 'US', 'PL', 'DE').
	BusinessCountry *CountryCode `json:"business_country,omitempty" xml:"business_country,omitempty"`
	// What the business does, given to the receptionist.
	BusinessDescription *string `json:"business_description,omitempty" xml:"business_description,omitempty"`
	// IANA timezone of the business.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Currency code for all prices (ISO 4217)
	Currency       *ReceptionAgentInfoResponseCurrency `json:"currency,omitempty" xml:"currency,omitempty"`
	OperatingHours *WorkingHours                       `json:"operating_hours,omitempty" xml:"operating_hours,omitempty"`
	// Ambient sound played while any receptionist runs a tool (e.g. checking availability or looking up a client), filling the silence callers would otherwise hear. One sound for the whole workspace, because custom webhook and MCP tools are shared by every receptionist. Off when null.
	ToolCallSound *ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// ID of the business profile.
	Id string `json:"id" xml:"id" required:"true"`
	// When the business profile was created, as a Unix timestamp in seconds.
	CreatedAtUnix int `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	// When the business profile was last updated, as a Unix timestamp in seconds.
	UpdatedAtUnix   int           `json:"updated_at_unix" xml:"updated_at_unix" required:"true"`
	WorkspaceNow    *WorkspaceNow `json:"workspace_now,omitempty" xml:"workspace_now,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionAgentInfoResponse) GetBusinessName() (value *string) {
	if r == nil {
		return
	}
	return r.BusinessName
}

func (r *ReceptionAgentInfoResponse) GetBusinessCountry() (value *CountryCode) {
	if r == nil {
		return
	}
	return r.BusinessCountry
}

func (r *ReceptionAgentInfoResponse) GetBusinessDescription() (value *string) {
	if r == nil {
		return
	}
	return r.BusinessDescription
}

func (r *ReceptionAgentInfoResponse) GetTimezone() (value *string) {
	if r == nil {
		return
	}
	return r.Timezone
}

func (r *ReceptionAgentInfoResponse) GetCurrency() (value *ReceptionAgentInfoResponseCurrency) {
	if r == nil {
		return
	}
	return r.Currency
}

func (r *ReceptionAgentInfoResponse) GetOperatingHours() (value *WorkingHours) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *ReceptionAgentInfoResponse) GetToolCallSound() (value *ToolCallSoundType) {
	if r == nil {
		return
	}
	return r.ToolCallSound
}

func (r *ReceptionAgentInfoResponse) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionAgentInfoResponse) GetCreatedAtUnix() (value int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionAgentInfoResponse) GetUpdatedAtUnix() (value int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionAgentInfoResponse) GetWorkspaceNow() (value *WorkspaceNow) {
	if r == nil {
		return
	}
	return r.WorkspaceNow
}

func (r *ReceptionAgentInfoResponse) SetBusinessName(businessName *string) {
	r.BusinessName = businessName
}

func (r *ReceptionAgentInfoResponse) SetBusinessCountry(businessCountry *CountryCode) {
	r.BusinessCountry = businessCountry
}

func (r *ReceptionAgentInfoResponse) SetBusinessDescription(businessDescription *string) {
	r.BusinessDescription = businessDescription
}

func (r *ReceptionAgentInfoResponse) SetTimezone(timezone *string) {
	r.Timezone = timezone
}

func (r *ReceptionAgentInfoResponse) SetCurrency(currency *ReceptionAgentInfoResponseCurrency) {
	r.Currency = currency
}

func (r *ReceptionAgentInfoResponse) SetOperatingHours(operatingHours *WorkingHours) {
	r.OperatingHours = operatingHours
}

func (r *ReceptionAgentInfoResponse) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	r.ToolCallSound = toolCallSound
}

func (r *ReceptionAgentInfoResponse) SetId(id string) {
	r.Id = id
}

func (r *ReceptionAgentInfoResponse) SetCreatedAtUnix(createdAtUnix int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionAgentInfoResponse) SetUpdatedAtUnix(updatedAtUnix int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionAgentInfoResponse) SetWorkspaceNow(workspaceNow *WorkspaceNow) {
	r.WorkspaceNow = workspaceNow
}

func (r *ReceptionAgentInfoResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionAgentInfoResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionAgentInfoResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionAgentInfoResponse) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionAgentInfoResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionAgentInfoResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
