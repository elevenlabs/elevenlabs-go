package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Fields to change on the business profile.
type ReceptionUpdateAgentInfoRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Name of the business.
	BusinessName *string `json:"business_name,omitempty" xml:"business_name,omitempty"`
	// ISO 3166-1 alpha-2 country code where the business is based (e.g. 'US', 'PL', 'DE').
	BusinessCountry *CountryCode `json:"business_country,omitempty" xml:"business_country,omitempty"`
	// What the business does, given to the receptionist.
	BusinessDescription *string `json:"business_description,omitempty" xml:"business_description,omitempty"`
	// IANA timezone of the business.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Currency code for all prices (ISO 4217)
	Currency       *types.ReceptionAgentInfoResponseCurrency `json:"currency,omitempty" xml:"currency,omitempty"`
	OperatingHours *types.WorkingHours                       `json:"operating_hours,omitempty" xml:"operating_hours,omitempty"`
	// Ambient sound played while any receptionist runs a tool (e.g. checking availability or looking up a client), filling the silence callers would otherwise hear. One sound for the whole workspace, because custom webhook and MCP tools are shared by every receptionist. Off when null.
	ToolCallSound *types.ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
}

func (r *ReceptionUpdateAgentInfoRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateAgentInfoRequest) GetBusinessName() (value *string) {
	if r == nil {
		return
	}
	return r.BusinessName
}

func (r *ReceptionUpdateAgentInfoRequest) GetBusinessCountry() (value *CountryCode) {
	if r == nil {
		return
	}
	return r.BusinessCountry
}

func (r *ReceptionUpdateAgentInfoRequest) GetBusinessDescription() (value *string) {
	if r == nil {
		return
	}
	return r.BusinessDescription
}

func (r *ReceptionUpdateAgentInfoRequest) GetTimezone() (value *string) {
	if r == nil {
		return
	}
	return r.Timezone
}

func (r *ReceptionUpdateAgentInfoRequest) GetCurrency() (value *types.ReceptionAgentInfoResponseCurrency) {
	if r == nil {
		return
	}
	return r.Currency
}

func (r *ReceptionUpdateAgentInfoRequest) GetOperatingHours() (value *types.WorkingHours) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *ReceptionUpdateAgentInfoRequest) GetToolCallSound() (value *types.ToolCallSoundType) {
	if r == nil {
		return
	}
	return r.ToolCallSound
}

func (r *ReceptionUpdateAgentInfoRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateAgentInfoRequest) SetBusinessName(businessName *string) {
	r.BusinessName = businessName
}

func (r *ReceptionUpdateAgentInfoRequest) SetBusinessCountry(businessCountry *CountryCode) {
	r.BusinessCountry = businessCountry
}

func (r *ReceptionUpdateAgentInfoRequest) SetBusinessDescription(businessDescription *string) {
	r.BusinessDescription = businessDescription
}

func (r *ReceptionUpdateAgentInfoRequest) SetTimezone(timezone *string) {
	r.Timezone = timezone
}

func (r *ReceptionUpdateAgentInfoRequest) SetCurrency(currency *types.ReceptionAgentInfoResponseCurrency) {
	r.Currency = currency
}

func (r *ReceptionUpdateAgentInfoRequest) SetOperatingHours(operatingHours *types.WorkingHours) {
	r.OperatingHours = operatingHours
}

func (r *ReceptionUpdateAgentInfoRequest) SetToolCallSound(toolCallSound *types.ToolCallSoundType) {
	r.ToolCallSound = toolCallSound
}

func (r *ReceptionUpdateAgentInfoRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateAgentInfoRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateAgentInfoRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateAgentInfoRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateAgentInfoRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateAgentInfoRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	getBusinessProfileRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetBusinessProfileRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetBusinessProfileRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetBusinessProfileRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getBusinessProfileRequestFieldXiWorkspaceId)
}
