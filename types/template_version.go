package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A published snapshot of a template, and the ports it runs with.
type TemplateVersion struct {
	// Pass as `version_id` on `POST /v1/flows/templates/{template_id}/runs` to pin a run to this snapshot.
	VersionId string `json:"version_id" xml:"version_id" required:"true"`
	// When this version was published, as a Unix timestamp in seconds.
	PublishedAtUnix int `json:"published_at_unix" xml:"published_at_unix" required:"true"`
	// Whether this is the version a run gets when `version_id` is omitted or set to `latest`.
	IsLatest bool `json:"is_latest" xml:"is_latest" required:"true"`
	// The inputs this version accepts, in canvas order. Every input is required on a run.
	Inputs []*TemplatePort `json:"inputs" xml:"inputs" required:"true"`
	// The outputs this version produces, in canvas order.
	Outputs         []*TemplatePort `json:"outputs" xml:"outputs" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TemplateVersion) GetVersionId() (value string) {
	if t == nil {
		return
	}
	return t.VersionId
}

func (t *TemplateVersion) GetPublishedAtUnix() (value int) {
	if t == nil {
		return
	}
	return t.PublishedAtUnix
}

func (t *TemplateVersion) GetIsLatest() (value bool) {
	if t == nil {
		return
	}
	return t.IsLatest
}

func (t *TemplateVersion) GetInputs() (value []*TemplatePort) {
	if t == nil {
		return
	}
	return t.Inputs
}

func (t *TemplateVersion) GetOutputs() (value []*TemplatePort) {
	if t == nil {
		return
	}
	return t.Outputs
}

func (t *TemplateVersion) SetVersionId(versionId string) {
	t.VersionId = versionId
}

func (t *TemplateVersion) SetPublishedAtUnix(publishedAtUnix int) {
	t.PublishedAtUnix = publishedAtUnix
}

func (t *TemplateVersion) SetIsLatest(isLatest bool) {
	t.IsLatest = isLatest
}

func (t *TemplateVersion) SetInputs(inputs []*TemplatePort) {
	t.Inputs = inputs
}

func (t *TemplateVersion) SetOutputs(outputs []*TemplatePort) {
	t.Outputs = outputs
}

func (t *TemplateVersion) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateVersion) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateVersion to string"
	}
	return string(jsonData)
}

func (t *TemplateVersion) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TemplateVersion
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TemplateVersion(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateVersion) MarshalJSON() ([]byte, error) {
	type embed TemplateVersion
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
