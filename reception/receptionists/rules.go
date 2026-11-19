package receptionists

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Request model for creating a new agent rule.
type ReceptionCreateAgentRuleRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId     string      `json:"-" url:"-"`
	Title       StrippedStr `json:"title" xml:"title" required:"true"`
	Description StrippedStr `json:"description" xml:"description" required:"true"`
	// How strictly the receptionist follows a rule.
	Priority *types.RulePriority `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (r *ReceptionCreateAgentRuleRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateAgentRuleRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionCreateAgentRuleRequest) GetTitle() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Title
}

func (r *ReceptionCreateAgentRuleRequest) GetDescription() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionCreateAgentRuleRequest) GetPriority() (value *types.RulePriority) {
	if r == nil {
		return
	}
	return r.Priority
}

func (r *ReceptionCreateAgentRuleRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateAgentRuleRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionCreateAgentRuleRequest) SetTitle(title StrippedStr) {
	r.Title = title
}

func (r *ReceptionCreateAgentRuleRequest) SetDescription(description StrippedStr) {
	r.Description = description
}

func (r *ReceptionCreateAgentRuleRequest) SetPriority(priority *types.RulePriority) {
	r.Priority = priority
}

func (r *ReceptionCreateAgentRuleRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateAgentRuleRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateAgentRuleRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateAgentRuleRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionCreateAgentRuleRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionCreateAgentRuleRequest(tmp)
	return nil
}

func (r *ReceptionCreateAgentRuleRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateAgentRuleRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

// Request model for updating an existing agent rule.
type ReceptionUpdateAgentRuleRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the rule. Returned when you add one, and listed in the receptionist's `config.agent_rules`.
	RuleId string `json:"-" url:"-"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId     string      `json:"-" url:"-"`
	Title       StrippedStr `json:"title" xml:"title" required:"true"`
	Description StrippedStr `json:"description" xml:"description" required:"true"`
	// How strictly the receptionist follows a rule.
	Priority *types.RulePriority `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (r *ReceptionUpdateAgentRuleRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateAgentRuleRequest) GetRuleId() (value string) {
	if r == nil {
		return
	}
	return r.RuleId
}

func (r *ReceptionUpdateAgentRuleRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionUpdateAgentRuleRequest) GetTitle() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Title
}

func (r *ReceptionUpdateAgentRuleRequest) GetDescription() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionUpdateAgentRuleRequest) GetPriority() (value *types.RulePriority) {
	if r == nil {
		return
	}
	return r.Priority
}

func (r *ReceptionUpdateAgentRuleRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateAgentRuleRequest) SetRuleId(ruleId string) {
	r.RuleId = ruleId
}

func (r *ReceptionUpdateAgentRuleRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionUpdateAgentRuleRequest) SetTitle(title StrippedStr) {
	r.Title = title
}

func (r *ReceptionUpdateAgentRuleRequest) SetDescription(description StrippedStr) {
	r.Description = description
}

func (r *ReceptionUpdateAgentRuleRequest) SetPriority(priority *types.RulePriority) {
	r.Priority = priority
}

func (r *ReceptionUpdateAgentRuleRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateAgentRuleRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateAgentRuleRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateAgentRuleRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionUpdateAgentRuleRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionUpdateAgentRuleRequest(tmp)
	return nil
}

func (r *ReceptionUpdateAgentRuleRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateAgentRuleRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	deleteRulesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteRulesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteRulesRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteRulesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteRulesRequestFieldXiWorkspaceId)
}
