package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationInitiationClientDataConfigInput struct {
	ConversationConfigOverride *ConversationConfigClientOverrideConfigInput `json:"conversation_config_override,omitempty" xml:"conversation_config_override,omitempty"`
	// Whether to include custom LLM extra body
	CustomLlmExtraBody *bool `json:"custom_llm_extra_body,omitempty" xml:"custom_llm_extra_body,omitempty"`
	// Whether to enable conversation initiation client data from webhooks
	EnableConversationInitiationClientDataFromWebhook *bool `json:"enable_conversation_initiation_client_data_from_webhook,omitempty" xml:"enable_conversation_initiation_client_data_from_webhook,omitempty"`
	// Whether clients may pass starting_workflow_node_id in initiation client data; if false, sending it fails conversation start.
	EnableStartingWorkflowNodeIdFromClient *bool `json:"enable_starting_workflow_node_id_from_client,omitempty" xml:"enable_starting_workflow_node_id_from_client,omitempty"`
	// Whether clients may pass procedure_ids and start_procedure_ids in initiation client data. procedure_ids selects which of the agent's procedures are available. start_procedure_ids starts free-form procedures when the conversation starts. If false, sending either field fails conversation start.
	EnableProcedureIdsFromClient *bool `json:"enable_procedure_ids_from_client,omitempty" xml:"enable_procedure_ids_from_client,omitempty"`
	extraProperties              map[string]interface{}
}

func (c *ConversationInitiationClientDataConfigInput) GetConversationConfigOverride() (value *ConversationConfigClientOverrideConfigInput) {
	if c == nil {
		return
	}
	return c.ConversationConfigOverride
}

func (c *ConversationInitiationClientDataConfigInput) GetCustomLlmExtraBody() (value *bool) {
	if c == nil {
		return
	}
	return c.CustomLlmExtraBody
}

func (c *ConversationInitiationClientDataConfigInput) GetEnableConversationInitiationClientDataFromWebhook() (value *bool) {
	if c == nil {
		return
	}
	return c.EnableConversationInitiationClientDataFromWebhook
}

func (c *ConversationInitiationClientDataConfigInput) GetEnableStartingWorkflowNodeIdFromClient() (value *bool) {
	if c == nil {
		return
	}
	return c.EnableStartingWorkflowNodeIdFromClient
}

func (c *ConversationInitiationClientDataConfigInput) GetEnableProcedureIdsFromClient() (value *bool) {
	if c == nil {
		return
	}
	return c.EnableProcedureIdsFromClient
}

func (c *ConversationInitiationClientDataConfigInput) SetConversationConfigOverride(conversationConfigOverride *ConversationConfigClientOverrideConfigInput) {
	c.ConversationConfigOverride = conversationConfigOverride
}

func (c *ConversationInitiationClientDataConfigInput) SetCustomLlmExtraBody(customLlmExtraBody *bool) {
	c.CustomLlmExtraBody = customLlmExtraBody
}

func (c *ConversationInitiationClientDataConfigInput) SetEnableConversationInitiationClientDataFromWebhook(enableConversationInitiationClientDataFromWebhook *bool) {
	c.EnableConversationInitiationClientDataFromWebhook = enableConversationInitiationClientDataFromWebhook
}

func (c *ConversationInitiationClientDataConfigInput) SetEnableStartingWorkflowNodeIdFromClient(enableStartingWorkflowNodeIdFromClient *bool) {
	c.EnableStartingWorkflowNodeIdFromClient = enableStartingWorkflowNodeIdFromClient
}

func (c *ConversationInitiationClientDataConfigInput) SetEnableProcedureIdsFromClient(enableProcedureIdsFromClient *bool) {
	c.EnableProcedureIdsFromClient = enableProcedureIdsFromClient
}

func (c *ConversationInitiationClientDataConfigInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConfigInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConfigInput to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConfigInput) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
