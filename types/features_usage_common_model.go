package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FeaturesUsageCommonModel struct {
	LanguageDetection          *FeatureStatusCommonModel         `json:"language_detection,omitempty" xml:"language_detection,omitempty"`
	TransferToAgent            *FeatureStatusCommonModel         `json:"transfer_to_agent,omitempty" xml:"transfer_to_agent,omitempty"`
	TransferToNumber           *FeatureStatusCommonModel         `json:"transfer_to_number,omitempty" xml:"transfer_to_number,omitempty"`
	Multivoice                 *FeatureStatusCommonModel         `json:"multivoice,omitempty" xml:"multivoice,omitempty"`
	DtmfTones                  *FeatureStatusCommonModel         `json:"dtmf_tones,omitempty" xml:"dtmf_tones,omitempty"`
	ExternalMcpServers         *FeatureStatusCommonModel         `json:"external_mcp_servers,omitempty" xml:"external_mcp_servers,omitempty"`
	PiiZrmWorkspace            *bool                             `json:"pii_zrm_workspace,omitempty" xml:"pii_zrm_workspace,omitempty"`
	PiiZrmAgent                *bool                             `json:"pii_zrm_agent,omitempty" xml:"pii_zrm_agent,omitempty"`
	ToolDynamicVariableUpdates *FeatureStatusCommonModel         `json:"tool_dynamic_variable_updates,omitempty" xml:"tool_dynamic_variable_updates,omitempty"`
	IsLivekit                  *bool                             `json:"is_livekit,omitempty" xml:"is_livekit,omitempty"`
	VoicemailDetection         *FeatureStatusCommonModel         `json:"voicemail_detection,omitempty" xml:"voicemail_detection,omitempty"`
	DtmfInput                  *FeatureStatusCommonModel         `json:"dtmf_input,omitempty" xml:"dtmf_input,omitempty"`
	Workflow                   *WorkflowFeaturesUsageCommonModel `json:"workflow,omitempty" xml:"workflow,omitempty"`
	AgentTesting               *TestsFeatureUsageCommonModel     `json:"agent_testing,omitempty" xml:"agent_testing,omitempty"`
	Versioning                 *FeatureStatusCommonModel         `json:"versioning,omitempty" xml:"versioning,omitempty"`
	FileInput                  *FeatureStatusCommonModel         `json:"file_input,omitempty" xml:"file_input,omitempty"`
	FreeformProcedure          *FeatureStatusCommonModel         `json:"freeform_procedure,omitempty" xml:"freeform_procedure,omitempty"`
	StructuredProcedure        *FeatureStatusCommonModel         `json:"structured_procedure,omitempty" xml:"structured_procedure,omitempty"`
	extraProperties            map[string]interface{}
}

func (f *FeaturesUsageCommonModel) GetLanguageDetection() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.LanguageDetection
}

func (f *FeaturesUsageCommonModel) GetTransferToAgent() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.TransferToAgent
}

func (f *FeaturesUsageCommonModel) GetTransferToNumber() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.TransferToNumber
}

func (f *FeaturesUsageCommonModel) GetMultivoice() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.Multivoice
}

func (f *FeaturesUsageCommonModel) GetDtmfTones() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.DtmfTones
}

func (f *FeaturesUsageCommonModel) GetExternalMcpServers() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.ExternalMcpServers
}

func (f *FeaturesUsageCommonModel) GetPiiZrmWorkspace() (value *bool) {
	if f == nil {
		return
	}
	return f.PiiZrmWorkspace
}

func (f *FeaturesUsageCommonModel) GetPiiZrmAgent() (value *bool) {
	if f == nil {
		return
	}
	return f.PiiZrmAgent
}

func (f *FeaturesUsageCommonModel) GetToolDynamicVariableUpdates() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.ToolDynamicVariableUpdates
}

func (f *FeaturesUsageCommonModel) GetIsLivekit() (value *bool) {
	if f == nil {
		return
	}
	return f.IsLivekit
}

func (f *FeaturesUsageCommonModel) GetVoicemailDetection() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.VoicemailDetection
}

func (f *FeaturesUsageCommonModel) GetDtmfInput() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.DtmfInput
}

func (f *FeaturesUsageCommonModel) GetWorkflow() (value *WorkflowFeaturesUsageCommonModel) {
	if f == nil {
		return
	}
	return f.Workflow
}

func (f *FeaturesUsageCommonModel) GetAgentTesting() (value *TestsFeatureUsageCommonModel) {
	if f == nil {
		return
	}
	return f.AgentTesting
}

func (f *FeaturesUsageCommonModel) GetVersioning() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.Versioning
}

func (f *FeaturesUsageCommonModel) GetFileInput() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.FileInput
}

func (f *FeaturesUsageCommonModel) GetFreeformProcedure() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.FreeformProcedure
}

func (f *FeaturesUsageCommonModel) GetStructuredProcedure() (value *FeatureStatusCommonModel) {
	if f == nil {
		return
	}
	return f.StructuredProcedure
}

func (f *FeaturesUsageCommonModel) SetLanguageDetection(languageDetection *FeatureStatusCommonModel) {
	f.LanguageDetection = languageDetection
}

func (f *FeaturesUsageCommonModel) SetTransferToAgent(transferToAgent *FeatureStatusCommonModel) {
	f.TransferToAgent = transferToAgent
}

func (f *FeaturesUsageCommonModel) SetTransferToNumber(transferToNumber *FeatureStatusCommonModel) {
	f.TransferToNumber = transferToNumber
}

func (f *FeaturesUsageCommonModel) SetMultivoice(multivoice *FeatureStatusCommonModel) {
	f.Multivoice = multivoice
}

func (f *FeaturesUsageCommonModel) SetDtmfTones(dtmfTones *FeatureStatusCommonModel) {
	f.DtmfTones = dtmfTones
}

func (f *FeaturesUsageCommonModel) SetExternalMcpServers(externalMcpServers *FeatureStatusCommonModel) {
	f.ExternalMcpServers = externalMcpServers
}

func (f *FeaturesUsageCommonModel) SetPiiZrmWorkspace(piiZrmWorkspace *bool) {
	f.PiiZrmWorkspace = piiZrmWorkspace
}

func (f *FeaturesUsageCommonModel) SetPiiZrmAgent(piiZrmAgent *bool) {
	f.PiiZrmAgent = piiZrmAgent
}

func (f *FeaturesUsageCommonModel) SetToolDynamicVariableUpdates(toolDynamicVariableUpdates *FeatureStatusCommonModel) {
	f.ToolDynamicVariableUpdates = toolDynamicVariableUpdates
}

func (f *FeaturesUsageCommonModel) SetIsLivekit(isLivekit *bool) {
	f.IsLivekit = isLivekit
}

func (f *FeaturesUsageCommonModel) SetVoicemailDetection(voicemailDetection *FeatureStatusCommonModel) {
	f.VoicemailDetection = voicemailDetection
}

func (f *FeaturesUsageCommonModel) SetDtmfInput(dtmfInput *FeatureStatusCommonModel) {
	f.DtmfInput = dtmfInput
}

func (f *FeaturesUsageCommonModel) SetWorkflow(workflow *WorkflowFeaturesUsageCommonModel) {
	f.Workflow = workflow
}

func (f *FeaturesUsageCommonModel) SetAgentTesting(agentTesting *TestsFeatureUsageCommonModel) {
	f.AgentTesting = agentTesting
}

func (f *FeaturesUsageCommonModel) SetVersioning(versioning *FeatureStatusCommonModel) {
	f.Versioning = versioning
}

func (f *FeaturesUsageCommonModel) SetFileInput(fileInput *FeatureStatusCommonModel) {
	f.FileInput = fileInput
}

func (f *FeaturesUsageCommonModel) SetFreeformProcedure(freeformProcedure *FeatureStatusCommonModel) {
	f.FreeformProcedure = freeformProcedure
}

func (f *FeaturesUsageCommonModel) SetStructuredProcedure(structuredProcedure *FeatureStatusCommonModel) {
	f.StructuredProcedure = structuredProcedure
}

func (f *FeaturesUsageCommonModel) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FeaturesUsageCommonModel) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FeaturesUsageCommonModel to string"
	}
	return string(jsonData)
}

func (f *FeaturesUsageCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler FeaturesUsageCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FeaturesUsageCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FeaturesUsageCommonModel) MarshalJSON() ([]byte, error) {
	type embed FeaturesUsageCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
