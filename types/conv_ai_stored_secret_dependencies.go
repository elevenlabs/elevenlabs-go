package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConvAiStoredSecretDependencies struct {
	Tools []*ConvAiStoredSecretDependenciesToolsItem `json:"tools" xml:"tools" required:"true"`
	// Whether there are more tool dependents beyond the returned preview
	ToolsHasMore *bool                                       `json:"tools_has_more,omitempty" xml:"tools_has_more,omitempty"`
	Agents       []*ConvAiStoredSecretDependenciesAgentsItem `json:"agents" xml:"agents" required:"true"`
	// Whether there are more agent dependents beyond the returned preview
	AgentsHasMore *bool                             `json:"agents_has_more,omitempty" xml:"agents_has_more,omitempty"`
	PhoneNumbers  []*DependentPhoneNumberIdentifier `json:"phone_numbers,omitempty" xml:"phone_numbers,omitempty"`
	// Whether there are more phone number dependents beyond the returned preview
	PhoneNumbersHasMore *bool                                           `json:"phone_numbers_has_more,omitempty" xml:"phone_numbers_has_more,omitempty"`
	McpServers          []*ConvAiStoredSecretDependenciesMcpServersItem `json:"mcp_servers,omitempty" xml:"mcp_servers,omitempty"`
	Others              []SecretDependencyType                          `json:"others" xml:"others" required:"true"`
	extraProperties     map[string]interface{}
}

func (c *ConvAiStoredSecretDependencies) GetTools() (value []*ConvAiStoredSecretDependenciesToolsItem) {
	if c == nil {
		return
	}
	return c.Tools
}

func (c *ConvAiStoredSecretDependencies) GetToolsHasMore() (value *bool) {
	if c == nil {
		return
	}
	return c.ToolsHasMore
}

func (c *ConvAiStoredSecretDependencies) GetAgents() (value []*ConvAiStoredSecretDependenciesAgentsItem) {
	if c == nil {
		return
	}
	return c.Agents
}

func (c *ConvAiStoredSecretDependencies) GetAgentsHasMore() (value *bool) {
	if c == nil {
		return
	}
	return c.AgentsHasMore
}

func (c *ConvAiStoredSecretDependencies) GetPhoneNumbers() (value []*DependentPhoneNumberIdentifier) {
	if c == nil {
		return
	}
	return c.PhoneNumbers
}

func (c *ConvAiStoredSecretDependencies) GetPhoneNumbersHasMore() (value *bool) {
	if c == nil {
		return
	}
	return c.PhoneNumbersHasMore
}

func (c *ConvAiStoredSecretDependencies) GetMcpServers() (value []*ConvAiStoredSecretDependenciesMcpServersItem) {
	if c == nil {
		return
	}
	return c.McpServers
}

func (c *ConvAiStoredSecretDependencies) GetOthers() (value []SecretDependencyType) {
	if c == nil {
		return
	}
	return c.Others
}

func (c *ConvAiStoredSecretDependencies) SetTools(tools []*ConvAiStoredSecretDependenciesToolsItem) {
	c.Tools = tools
}

func (c *ConvAiStoredSecretDependencies) SetToolsHasMore(toolsHasMore *bool) {
	c.ToolsHasMore = toolsHasMore
}

func (c *ConvAiStoredSecretDependencies) SetAgents(agents []*ConvAiStoredSecretDependenciesAgentsItem) {
	c.Agents = agents
}

func (c *ConvAiStoredSecretDependencies) SetAgentsHasMore(agentsHasMore *bool) {
	c.AgentsHasMore = agentsHasMore
}

func (c *ConvAiStoredSecretDependencies) SetPhoneNumbers(phoneNumbers []*DependentPhoneNumberIdentifier) {
	c.PhoneNumbers = phoneNumbers
}

func (c *ConvAiStoredSecretDependencies) SetPhoneNumbersHasMore(phoneNumbersHasMore *bool) {
	c.PhoneNumbersHasMore = phoneNumbersHasMore
}

func (c *ConvAiStoredSecretDependencies) SetMcpServers(mcpServers []*ConvAiStoredSecretDependenciesMcpServersItem) {
	c.McpServers = mcpServers
}

func (c *ConvAiStoredSecretDependencies) SetOthers(others []SecretDependencyType) {
	c.Others = others
}

func (c *ConvAiStoredSecretDependencies) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiStoredSecretDependencies) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiStoredSecretDependencies to string"
	}
	return string(jsonData)
}

func (c *ConvAiStoredSecretDependencies) UnmarshalJSON(data []byte) error {
	type unmarshaler ConvAiStoredSecretDependencies
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConvAiStoredSecretDependencies(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiStoredSecretDependencies) MarshalJSON() ([]byte, error) {
	type embed ConvAiStoredSecretDependencies
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
