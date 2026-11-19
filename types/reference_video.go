package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReferenceVideo struct {
	GenerationId    *string            `json:"generation_id,omitempty" xml:"generation_id,omitempty"`
	ContentAssetId  *string            `json:"content_asset_id,omitempty" xml:"content_asset_id,omitempty"`
	TemplateNodeId  *string            `json:"template_node_id,omitempty" xml:"template_node_id,omitempty"`
	StudioClip      *StudioClipLocator `json:"studio_clip,omitempty" xml:"studio_clip,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReferenceVideo) GetGenerationId() (value *string) {
	if r == nil {
		return
	}
	return r.GenerationId
}

func (r *ReferenceVideo) GetContentAssetId() (value *string) {
	if r == nil {
		return
	}
	return r.ContentAssetId
}

func (r *ReferenceVideo) GetTemplateNodeId() (value *string) {
	if r == nil {
		return
	}
	return r.TemplateNodeId
}

func (r *ReferenceVideo) GetStudioClip() (value *StudioClipLocator) {
	if r == nil {
		return
	}
	return r.StudioClip
}

func (r *ReferenceVideo) SetGenerationId(generationId *string) {
	r.GenerationId = generationId
}

func (r *ReferenceVideo) SetContentAssetId(contentAssetId *string) {
	r.ContentAssetId = contentAssetId
}

func (r *ReferenceVideo) SetTemplateNodeId(templateNodeId *string) {
	r.TemplateNodeId = templateNodeId
}

func (r *ReferenceVideo) SetStudioClip(studioClip *StudioClipLocator) {
	r.StudioClip = studioClip
}

func (r *ReferenceVideo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReferenceVideo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReferenceVideo to string"
	}
	return string(jsonData)
}

func (r *ReferenceVideo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReferenceVideo) MarshalJSON() ([]byte, error) {
	type embed ReferenceVideo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
