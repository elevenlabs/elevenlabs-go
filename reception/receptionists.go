package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ReceptionUpdateCustomerFacingAgentRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId string `json:"-" url:"-"`
	// Receptionist settings to change; omitted settings keep their value.
	Config *types.ReceptionCustomerFacingConfigPatch `json:"config,omitempty" xml:"config,omitempty"`
	// New display name for the receptionist.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// New website widget appearance.
	Widget *types.ReceptionWidgetAppearance `json:"widget,omitempty" xml:"widget,omitempty"`
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetConfig() (value *types.ReceptionCustomerFacingConfigPatch) {
	if r == nil {
		return
	}
	return r.Config
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetName() (value *string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetWidget() (value *types.ReceptionWidgetAppearance) {
	if r == nil {
		return
	}
	return r.Widget
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) SetConfig(config *types.ReceptionCustomerFacingConfigPatch) {
	r.Config = config
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) SetName(name *string) {
	r.Name = name
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) SetWidget(widget *types.ReceptionWidgetAppearance) {
	r.Widget = widget
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateCustomerFacingAgentRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateCustomerFacingAgentRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *ReceptionUpdateCustomerFacingAgentRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateCustomerFacingAgentRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listReceptionistsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListReceptionistsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListReceptionistsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReceptionistsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listReceptionistsRequestFieldXiWorkspaceId)
}

var (
	activateReceptionistsRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ActivateReceptionistsRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (a *ActivateReceptionistsRequest) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *ActivateReceptionistsRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	a.XiWorkspaceId = xiWorkspaceId
	a.require(activateReceptionistsRequestFieldXiWorkspaceId)
}
