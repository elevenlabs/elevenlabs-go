package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A template run and its outputs. Every output exists from the moment the
// run is created and reports its own `status`; the run's `status` rolls them
// up, and the run is finished once it is `completed` or `failed`.
type TemplateRunResponse struct {
	// The unique identifier of the run.
	Id string `json:"id" xml:"id" required:"true"`
	// The template this run executed, so a webhook consumer running several templates can tell their runs apart without keeping a run-to-template map.
	TemplateId string `json:"template_id" xml:"template_id" required:"true"`
	// The template version this run executed. Resolved when the run is created, so a run started with `latest` records the concrete version it ran.
	VersionId string `json:"version_id" xml:"version_id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The run's outputs, keyed by output port id. Each is a `TemplateOutput` discriminated on `type`, the `type` of its port's `content_schema`.
	Outputs         map[string]*TemplateOutput `json:"outputs" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TemplateRunResponse) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateRunResponse) GetTemplateId() (value string) {
	if t == nil {
		return
	}
	return t.TemplateId
}

func (t *TemplateRunResponse) GetVersionId() (value string) {
	if t == nil {
		return
	}
	return t.VersionId
}

func (t *TemplateRunResponse) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateRunResponse) GetOutputs() (value map[string]*TemplateOutput) {
	if t == nil {
		return
	}
	return t.Outputs
}

func (t *TemplateRunResponse) SetId(id string) {
	t.Id = id
}

func (t *TemplateRunResponse) SetTemplateId(templateId string) {
	t.TemplateId = templateId
}

func (t *TemplateRunResponse) SetVersionId(versionId string) {
	t.VersionId = versionId
}

func (t *TemplateRunResponse) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateRunResponse) SetOutputs(outputs map[string]*TemplateOutput) {
	t.Outputs = outputs
}

func (t *TemplateRunResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateRunResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateRunResponse to string"
	}
	return string(jsonData)
}

func (t *TemplateRunResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TemplateRunResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TemplateRunResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateRunResponse) MarshalJSON() ([]byte, error) {
	type embed TemplateRunResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
