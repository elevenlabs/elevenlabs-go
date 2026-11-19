package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CompileProceduresResponseModel struct {
	Workflow        *AgentWorkflowResponseModel `json:"workflow" xml:"workflow" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CompileProceduresResponseModel) GetWorkflow() (value *AgentWorkflowResponseModel) {
	if c == nil {
		return
	}
	return c.Workflow
}

func (c *CompileProceduresResponseModel) SetWorkflow(workflow *AgentWorkflowResponseModel) {
	c.Workflow = workflow
}

func (c *CompileProceduresResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CompileProceduresResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CompileProceduresResponseModel to string"
	}
	return string(jsonData)
}

func (c *CompileProceduresResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CompileProceduresResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CompileProceduresResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CompileProceduresResponseModel) MarshalJSON() ([]byte, error) {
	type embed CompileProceduresResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
