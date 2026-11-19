package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceBatchCallsResponse struct {
	BatchCalls []*BatchCallResponse `json:"batch_calls" xml:"batch_calls" required:"true"`
	// The next document, used to paginate through the batch calls
	NextDoc *string `json:"next_doc,omitempty" xml:"next_doc,omitempty"`
	// Whether there are more batch calls to paginate through
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceBatchCallsResponse) GetBatchCalls() (value []*BatchCallResponse) {
	if w == nil {
		return
	}
	return w.BatchCalls
}

func (w *WorkspaceBatchCallsResponse) GetNextDoc() (value *string) {
	if w == nil {
		return
	}
	return w.NextDoc
}

func (w *WorkspaceBatchCallsResponse) GetHasMore() (value *bool) {
	if w == nil {
		return
	}
	return w.HasMore
}

func (w *WorkspaceBatchCallsResponse) SetBatchCalls(batchCalls []*BatchCallResponse) {
	w.BatchCalls = batchCalls
}

func (w *WorkspaceBatchCallsResponse) SetNextDoc(nextDoc *string) {
	w.NextDoc = nextDoc
}

func (w *WorkspaceBatchCallsResponse) SetHasMore(hasMore *bool) {
	w.HasMore = hasMore
}

func (w *WorkspaceBatchCallsResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceBatchCallsResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceBatchCallsResponse to string"
	}
	return string(jsonData)
}

func (w *WorkspaceBatchCallsResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceBatchCallsResponse) MarshalJSON() ([]byte, error) {
	type embed WorkspaceBatchCallsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
