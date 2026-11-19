package procedures

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type UpdateProcedureDraftRequestModel struct {
	// The procedure ID
	ProcedureId string `json:"-" url:"-"`
	// Branch ID to get the procedure draft from
	BranchId string `json:"-" url:"-"`
	// Agent ID to get the procedure draft from
	AgentId string `json:"-" url:"-"`
	// Procedure name
	Name string `json:"name" xml:"name" required:"true" maxLength:"200"`
	// Procedure content
	Content string              `json:"content" xml:"content" required:"true" maxLength:"50000"`
	Type    types.ProcedureType `json:"type" xml:"type" required:"true"`
	// When the agent should use this procedure. Empty string means this is a sub-procedure that should only start when another procedure references it. Omit this field (or send null) to keep the stored trigger, which is empty on create. Send `""` to clear it. `description` is accepted as an alias.
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
}

func (u *UpdateProcedureDraftRequestModel) GetProcedureId() (value string) {
	if u == nil {
		return
	}
	return u.ProcedureId
}

func (u *UpdateProcedureDraftRequestModel) GetBranchId() (value string) {
	if u == nil {
		return
	}
	return u.BranchId
}

func (u *UpdateProcedureDraftRequestModel) GetAgentId() (value string) {
	if u == nil {
		return
	}
	return u.AgentId
}

func (u *UpdateProcedureDraftRequestModel) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateProcedureDraftRequestModel) GetContent() (value string) {
	if u == nil {
		return
	}
	return u.Content
}

func (u *UpdateProcedureDraftRequestModel) GetType() (value types.ProcedureType) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UpdateProcedureDraftRequestModel) GetTrigger() (value *string) {
	if u == nil {
		return
	}
	return u.Trigger
}

func (u *UpdateProcedureDraftRequestModel) SetProcedureId(procedureId string) {
	u.ProcedureId = procedureId
}

func (u *UpdateProcedureDraftRequestModel) SetBranchId(branchId string) {
	u.BranchId = branchId
}

func (u *UpdateProcedureDraftRequestModel) SetAgentId(agentId string) {
	u.AgentId = agentId
}

func (u *UpdateProcedureDraftRequestModel) SetName(name string) {
	u.Name = name
}

func (u *UpdateProcedureDraftRequestModel) SetContent(content string) {
	u.Content = content
}

func (u *UpdateProcedureDraftRequestModel) SetType(type_ types.ProcedureType) {
	u.Type = type_
}

func (u *UpdateProcedureDraftRequestModel) SetTrigger(trigger *string) {
	u.Trigger = trigger
}

func (u *UpdateProcedureDraftRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateProcedureDraftRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateProcedureDraftRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateProcedureDraftRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateProcedureDraftRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateProcedureDraftRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
