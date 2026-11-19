package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowFeaturesUsageCommonModel struct {
	Enabled             *bool                     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	ToolNode            *FeatureStatusCommonModel `json:"tool_node,omitempty" xml:"tool_node,omitempty"`
	StandaloneAgentNode *FeatureStatusCommonModel `json:"standalone_agent_node,omitempty" xml:"standalone_agent_node,omitempty"`
	PhoneNumberNode     *FeatureStatusCommonModel `json:"phone_number_node,omitempty" xml:"phone_number_node,omitempty"`
	EndNode             *FeatureStatusCommonModel `json:"end_node,omitempty" xml:"end_node,omitempty"`
	OverrideAgentNode   *FeatureStatusCommonModel `json:"override_agent_node,omitempty" xml:"override_agent_node,omitempty"`
	extraProperties     map[string]interface{}
}

func (w *WorkflowFeaturesUsageCommonModel) GetEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.Enabled
}

func (w *WorkflowFeaturesUsageCommonModel) GetToolNode() (value *FeatureStatusCommonModel) {
	if w == nil {
		return
	}
	return w.ToolNode
}

func (w *WorkflowFeaturesUsageCommonModel) GetStandaloneAgentNode() (value *FeatureStatusCommonModel) {
	if w == nil {
		return
	}
	return w.StandaloneAgentNode
}

func (w *WorkflowFeaturesUsageCommonModel) GetPhoneNumberNode() (value *FeatureStatusCommonModel) {
	if w == nil {
		return
	}
	return w.PhoneNumberNode
}

func (w *WorkflowFeaturesUsageCommonModel) GetEndNode() (value *FeatureStatusCommonModel) {
	if w == nil {
		return
	}
	return w.EndNode
}

func (w *WorkflowFeaturesUsageCommonModel) GetOverrideAgentNode() (value *FeatureStatusCommonModel) {
	if w == nil {
		return
	}
	return w.OverrideAgentNode
}

func (w *WorkflowFeaturesUsageCommonModel) SetEnabled(enabled *bool) {
	w.Enabled = enabled
}

func (w *WorkflowFeaturesUsageCommonModel) SetToolNode(toolNode *FeatureStatusCommonModel) {
	w.ToolNode = toolNode
}

func (w *WorkflowFeaturesUsageCommonModel) SetStandaloneAgentNode(standaloneAgentNode *FeatureStatusCommonModel) {
	w.StandaloneAgentNode = standaloneAgentNode
}

func (w *WorkflowFeaturesUsageCommonModel) SetPhoneNumberNode(phoneNumberNode *FeatureStatusCommonModel) {
	w.PhoneNumberNode = phoneNumberNode
}

func (w *WorkflowFeaturesUsageCommonModel) SetEndNode(endNode *FeatureStatusCommonModel) {
	w.EndNode = endNode
}

func (w *WorkflowFeaturesUsageCommonModel) SetOverrideAgentNode(overrideAgentNode *FeatureStatusCommonModel) {
	w.OverrideAgentNode = overrideAgentNode
}

func (w *WorkflowFeaturesUsageCommonModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowFeaturesUsageCommonModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowFeaturesUsageCommonModel to string"
	}
	return string(jsonData)
}

func (w *WorkflowFeaturesUsageCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler WorkflowFeaturesUsageCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WorkflowFeaturesUsageCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowFeaturesUsageCommonModel) MarshalJSON() ([]byte, error) {
	type embed WorkflowFeaturesUsageCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
