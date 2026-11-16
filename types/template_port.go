package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One input or output port of a published template version.
type TemplatePort struct {
	// The port id. Input ids are the keys of the `inputs` map on `POST /v1/flows/templates/{template_id}/runs`; output ids are the keys of the `outputs` map on the run response.
	Id              string         `json:"id" xml:"id" required:"true"`
	ContentSchema   *ContentSchema `json:"content_schema" xml:"content_schema" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TemplatePort) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplatePort) GetContentSchema() (value *ContentSchema) {
	if t == nil {
		return
	}
	return t.ContentSchema
}

func (t *TemplatePort) SetId(id string) {
	t.Id = id
}

func (t *TemplatePort) SetContentSchema(contentSchema *ContentSchema) {
	t.ContentSchema = contentSchema
}

func (t *TemplatePort) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplatePort) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplatePort to string"
	}
	return string(jsonData)
}

func (t *TemplatePort) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TemplatePort
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TemplatePort(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplatePort) MarshalJSON() ([]byte, error) {
	type embed TemplatePort
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
