package receptionists

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

// Request model for creating a new procedure.
type ReceptionCreateProcedureRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId string      `json:"-" url:"-"`
	Name    StrippedStr `json:"name" xml:"name" required:"true"`
	Trigger StrippedStr `json:"trigger" xml:"trigger" required:"true"`
	// Step-by-step instructions, as markdown.
	Body *string `json:"body,omitempty" xml:"body,omitempty" maxLength:"40000"`
}

func (r *ReceptionCreateProcedureRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionCreateProcedureRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionCreateProcedureRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCreateProcedureRequest) GetTrigger() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Trigger
}

func (r *ReceptionCreateProcedureRequest) GetBody() (value *string) {
	if r == nil {
		return
	}
	return r.Body
}

func (r *ReceptionCreateProcedureRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionCreateProcedureRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionCreateProcedureRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionCreateProcedureRequest) SetTrigger(trigger StrippedStr) {
	r.Trigger = trigger
}

func (r *ReceptionCreateProcedureRequest) SetBody(body *string) {
	r.Body = body
}

func (r *ReceptionCreateProcedureRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionCreateProcedureRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCreateProcedureRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionCreateProcedureRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionCreateProcedureRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionCreateProcedureRequest(tmp)
	return nil
}

func (r *ReceptionCreateProcedureRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionCreateProcedureRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

// Request model for updating an existing procedure.
type ReceptionUpdateProcedureRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// ID of the procedure. Returned when you add one; list them with `GET /v1/reception/customer-facing-agents/{agent_id}/procedures`.
	ProcedureId string `json:"-" url:"-"`
	// ID of the receptionist. List them with `GET /v1/reception/customer-facing-agents`.
	AgentId string      `json:"-" url:"-"`
	Name    StrippedStr `json:"name" xml:"name" required:"true"`
	Trigger StrippedStr `json:"trigger" xml:"trigger" required:"true"`
	// Step-by-step instructions, as markdown.
	Body *string `json:"body,omitempty" xml:"body,omitempty" maxLength:"40000"`
}

func (r *ReceptionUpdateProcedureRequest) GetXiWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.XiWorkspaceId
}

func (r *ReceptionUpdateProcedureRequest) GetProcedureId() (value string) {
	if r == nil {
		return
	}
	return r.ProcedureId
}

func (r *ReceptionUpdateProcedureRequest) GetAgentId() (value string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionUpdateProcedureRequest) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionUpdateProcedureRequest) GetTrigger() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Trigger
}

func (r *ReceptionUpdateProcedureRequest) GetBody() (value *string) {
	if r == nil {
		return
	}
	return r.Body
}

func (r *ReceptionUpdateProcedureRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
}

func (r *ReceptionUpdateProcedureRequest) SetProcedureId(procedureId string) {
	r.ProcedureId = procedureId
}

func (r *ReceptionUpdateProcedureRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
}

func (r *ReceptionUpdateProcedureRequest) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionUpdateProcedureRequest) SetTrigger(trigger StrippedStr) {
	r.Trigger = trigger
}

func (r *ReceptionUpdateProcedureRequest) SetBody(body *string) {
	r.Body = body
}

func (r *ReceptionUpdateProcedureRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReceptionUpdateProcedureRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionUpdateProcedureRequest to string"
	}
	return string(jsonData)
}

func (r *ReceptionUpdateProcedureRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionUpdateProcedureRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionUpdateProcedureRequest(tmp)
	return nil
}

func (r *ReceptionUpdateProcedureRequest) MarshalJSON() ([]byte, error) {
	type embed ReceptionUpdateProcedureRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

var (
	listProceduresRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListProceduresRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListProceduresRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProceduresRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listProceduresRequestFieldXiWorkspaceId)
}

var (
	deleteProceduresRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteProceduresRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteProceduresRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteProceduresRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteProceduresRequestFieldXiWorkspaceId)
}
