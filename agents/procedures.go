package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type CreateProcedureRequestModel struct {
	// Branch ID to get the procedure draft from
	BranchId string `json:"-" url:"-"`
	// Agent ID to get the procedure draft from
	AgentId string `json:"-" url:"-"`
	// Procedure name
	Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"200"`
	// Initial procedure content
	Content *string              `json:"content,omitempty" xml:"content,omitempty"`
	Type    *types.ProcedureType `json:"type,omitempty" xml:"type,omitempty"`
	// When the agent should use this procedure. Empty string means this is a sub-procedure that should only start when another procedure references it. Omit this field (or send null) to keep the stored trigger, which is empty on create. Send `""` to clear it. `description` is accepted as an alias.
	Trigger *string `json:"trigger,omitempty" xml:"trigger,omitempty"`
	// Procedure ID of the folder to create this procedure in, or null for root.
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
}

func (c *CreateProcedureRequestModel) GetBranchId() (value string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *CreateProcedureRequestModel) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateProcedureRequestModel) GetName() (value *string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateProcedureRequestModel) GetContent() (value *string) {
	if c == nil {
		return
	}
	return c.Content
}

func (c *CreateProcedureRequestModel) GetType() (value *types.ProcedureType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CreateProcedureRequestModel) GetTrigger() (value *string) {
	if c == nil {
		return
	}
	return c.Trigger
}

func (c *CreateProcedureRequestModel) GetFolderParentId() (value *string) {
	if c == nil {
		return
	}
	return c.FolderParentId
}

func (c *CreateProcedureRequestModel) SetBranchId(branchId string) {
	c.BranchId = branchId
}

func (c *CreateProcedureRequestModel) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *CreateProcedureRequestModel) SetName(name *string) {
	c.Name = name
}

func (c *CreateProcedureRequestModel) SetContent(content *string) {
	c.Content = content
}

func (c *CreateProcedureRequestModel) SetType(type_ *types.ProcedureType) {
	c.Type = type_
}

func (c *CreateProcedureRequestModel) SetTrigger(trigger *string) {
	c.Trigger = trigger
}

func (c *CreateProcedureRequestModel) SetFolderParentId(folderParentId *string) {
	c.FolderParentId = folderParentId
}

func (c *CreateProcedureRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateProcedureRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateProcedureRequestModel to string"
	}
	return string(jsonData)
}

func (c *CreateProcedureRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateProcedureRequestModel) MarshalJSON() ([]byte, error) {
	type embed CreateProcedureRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

var (
	listProceduresRequestFieldAgentVersionId = big.NewInt(1 << 0)
)

type ListProceduresRequest struct {
	AgentVersionId *string `json:"-" url:"agent_version_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListProceduresRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAgentVersionId sets the AgentVersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListProceduresRequest) SetAgentVersionId(agentVersionId *string) {
	l.AgentVersionId = agentVersionId
	l.require(listProceduresRequestFieldAgentVersionId)
}

var (
	getProceduresRequestFieldVersionId      = big.NewInt(1 << 0)
	getProceduresRequestFieldAgentVersionId = big.NewInt(1 << 1)
)

type GetProceduresRequest struct {
	VersionId      *string `json:"-" url:"version_id,omitempty"`
	AgentVersionId *string `json:"-" url:"agent_version_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetProceduresRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetProceduresRequest) SetVersionId(versionId *string) {
	g.VersionId = versionId
	g.require(getProceduresRequestFieldVersionId)
}

// SetAgentVersionId sets the AgentVersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetProceduresRequest) SetAgentVersionId(agentVersionId *string) {
	g.AgentVersionId = agentVersionId
	g.require(getProceduresRequestFieldAgentVersionId)
}
