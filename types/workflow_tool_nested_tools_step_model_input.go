package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolNestedToolsStepModelInput struct {
	StepLatencySecs float64                                                  `json:"step_latency_secs" xml:"step_latency_secs" required:"true"`
	NodeId          string                                                   `json:"node_id" xml:"node_id" required:"true"`
	Requests        []*ConversationHistoryTranscriptToolCallCommonModelInput `json:"requests" xml:"requests" required:"true"`
	Results         []*WorkflowToolNestedToolsStepModelInputResultsItem      `json:"results" xml:"results" required:"true"`
	IsSuccessful    bool                                                     `json:"is_successful" xml:"is_successful" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolNestedToolsStepModelInput) GetStepLatencySecs() (value float64) {
	if w == nil {
		return
	}
	return w.StepLatencySecs
}

func (w *WorkflowToolNestedToolsStepModelInput) GetNodeId() (value string) {
	if w == nil {
		return
	}
	return w.NodeId
}

func (w *WorkflowToolNestedToolsStepModelInput) GetRequests() (value []*ConversationHistoryTranscriptToolCallCommonModelInput) {
	if w == nil {
		return
	}
	return w.Requests
}

func (w *WorkflowToolNestedToolsStepModelInput) GetResults() (value []*WorkflowToolNestedToolsStepModelInputResultsItem) {
	if w == nil {
		return
	}
	return w.Results
}

func (w *WorkflowToolNestedToolsStepModelInput) GetIsSuccessful() (value bool) {
	if w == nil {
		return
	}
	return w.IsSuccessful
}

func (w *WorkflowToolNestedToolsStepModelInput) SetStepLatencySecs(stepLatencySecs float64) {
	w.StepLatencySecs = stepLatencySecs
}

func (w *WorkflowToolNestedToolsStepModelInput) SetNodeId(nodeId string) {
	w.NodeId = nodeId
}

func (w *WorkflowToolNestedToolsStepModelInput) SetRequests(requests []*ConversationHistoryTranscriptToolCallCommonModelInput) {
	w.Requests = requests
}

func (w *WorkflowToolNestedToolsStepModelInput) SetResults(results []*WorkflowToolNestedToolsStepModelInputResultsItem) {
	w.Results = results
}

func (w *WorkflowToolNestedToolsStepModelInput) SetIsSuccessful(isSuccessful bool) {
	w.IsSuccessful = isSuccessful
}

func (w *WorkflowToolNestedToolsStepModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolNestedToolsStepModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolNestedToolsStepModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolNestedToolsStepModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowToolNestedToolsStepModelInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowToolNestedToolsStepModelInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolNestedToolsStepModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolNestedToolsStepModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
