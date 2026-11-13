package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost struct {
	RenderType RenderType `json:"render_type" xml:"render_type" required:"true"`
	// Whether to normalize the volume of the rendered audio.
	NormalizeVolume *bool `json:"normalize_volume,omitempty" xml:"normalize_volume,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) GetRenderType() (value RenderType) {
	if r == nil {
		return
	}
	return r.RenderType
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) GetNormalizeVolume() (value *bool) {
	if r == nil {
		return
	}
	return r.NormalizeVolume
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) SetRenderType(renderType RenderType) {
	r.RenderType = renderType
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) SetNormalizeVolume(normalizeVolume *bool) {
	r.NormalizeVolume = normalizeVolume
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost to string"
	}
	return string(jsonData)
}

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) UnmarshalJSON(data []byte) error {
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

func (r *RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost) MarshalJSON() ([]byte, error) {
	type embed RenderAudioOrVideoForTheGivenLanguageV1DubbingResourceDubbingIdRenderLanguagePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
