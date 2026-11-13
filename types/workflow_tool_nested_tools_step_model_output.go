package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowToolNestedToolsStepModelOutput struct {
	StepLatencySecs float64                                                   `json:"step_latency_secs" xml:"step_latency_secs" required:"true"`
	NodeId          string                                                    `json:"node_id" xml:"node_id" required:"true"`
	Requests        []*ConversationHistoryTranscriptToolCallCommonModelOutput `json:"requests" xml:"requests" required:"true"`
	Results         []*WorkflowToolNestedToolsStepModelOutputResultsItem      `json:"results" xml:"results" required:"true"`
	IsSuccessful    bool                                                      `json:"is_successful" xml:"is_successful" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetStepLatencySecs() (value float64) {
	if w == nil {
		return
	}
	return w.StepLatencySecs
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetNodeId() (value string) {
	if w == nil {
		return
	}
	return w.NodeId
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetRequests() (value []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	if w == nil {
		return
	}
	return w.Requests
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetResults() (value []*WorkflowToolNestedToolsStepModelOutputResultsItem) {
	if w == nil {
		return
	}
	return w.Results
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetIsSuccessful() (value bool) {
	if w == nil {
		return
	}
	return w.IsSuccessful
}

func (w *WorkflowToolNestedToolsStepModelOutput) SetStepLatencySecs(stepLatencySecs float64) {
	w.StepLatencySecs = stepLatencySecs
}

func (w *WorkflowToolNestedToolsStepModelOutput) SetNodeId(nodeId string) {
	w.NodeId = nodeId
}

func (w *WorkflowToolNestedToolsStepModelOutput) SetRequests(requests []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	w.Requests = requests
}

func (w *WorkflowToolNestedToolsStepModelOutput) SetResults(results []*WorkflowToolNestedToolsStepModelOutputResultsItem) {
	w.Results = results
}

func (w *WorkflowToolNestedToolsStepModelOutput) SetIsSuccessful(isSuccessful bool) {
	w.IsSuccessful = isSuccessful
}

func (w *WorkflowToolNestedToolsStepModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowToolNestedToolsStepModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowToolNestedToolsStepModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowToolNestedToolsStepModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowToolNestedToolsStepModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowToolNestedToolsStepModelOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowToolNestedToolsStepModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowToolNestedToolsStepModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
