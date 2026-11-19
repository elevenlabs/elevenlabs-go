package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A template the caller can run, and its runnable versions.
type TemplateSummary struct {
	// Pass as `template_id` on `POST /v1/flows/templates/{template_id}/runs`.
	Id string `json:"id" xml:"id" required:"true"`
	// The template's name.
	Name string `json:"name" xml:"name" required:"true"`
	// The template's description, if it has one.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The published versions this caller can run, newest first. A version whose graph uses a model that is not available to you through the API is left out, as is one whose stored snapshot is gone; either way the list can be empty while the template still has versions the ElevenLabs app can run.
	Versions []*TemplateVersion `json:"versions" xml:"versions" required:"true"`
	// Whether this template has further published versions beyond the `versions_per_template` returned here. Fetch `GET /v1/flows/templates/{template_id}` with a larger `versions_per_template` to see more of them.
	HasMoreVersions bool `json:"has_more_versions" xml:"has_more_versions" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TemplateSummary) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateSummary) GetName() (value string) {
	if t == nil {
		return
	}
	return t.Name
}

func (t *TemplateSummary) GetDescription() (value *string) {
	if t == nil {
		return
	}
	return t.Description
}

func (t *TemplateSummary) GetVersions() (value []*TemplateVersion) {
	if t == nil {
		return
	}
	return t.Versions
}

func (t *TemplateSummary) GetHasMoreVersions() (value bool) {
	if t == nil {
		return
	}
	return t.HasMoreVersions
}

func (t *TemplateSummary) SetId(id string) {
	t.Id = id
}

func (t *TemplateSummary) SetName(name string) {
	t.Name = name
}

func (t *TemplateSummary) SetDescription(description *string) {
	t.Description = description
}

func (t *TemplateSummary) SetVersions(versions []*TemplateVersion) {
	t.Versions = versions
}

func (t *TemplateSummary) SetHasMoreVersions(hasMoreVersions bool) {
	t.HasMoreVersions = hasMoreVersions
}

func (t *TemplateSummary) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateSummary) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateSummary to string"
	}
	return string(jsonData)
}

func (t *TemplateSummary) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateSummary) MarshalJSON() ([]byte, error) {
	type embed TemplateSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
