package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Render struct {
	Id              string                 `json:"id" xml:"id" required:"true"`
	Version         int                    `json:"version" xml:"version" required:"true"`
	Language        *string                `json:"language" xml:"language" required:"true"`
	Type            *RenderType            `json:"type" xml:"type" required:"true"`
	MediaRef        *DubbingMediaReference `json:"media_ref" xml:"media_ref" required:"true"`
	Status          RenderStatus           `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (r *Render) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *Render) GetVersion() (value int) {
	if r == nil {
		return
	}
	return r.Version
}

func (r *Render) GetLanguage() (value *string) {
	if r == nil {
		return
	}
	return r.Language
}

func (r *Render) GetType() (value *RenderType) {
	if r == nil {
		return
	}
	return r.Type
}

func (r *Render) GetMediaRef() (value *DubbingMediaReference) {
	if r == nil {
		return
	}
	return r.MediaRef
}

func (r *Render) GetStatus() (value RenderStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *Render) SetId(id string) {
	r.Id = id
}

func (r *Render) SetVersion(version int) {
	r.Version = version
}

func (r *Render) SetLanguage(language *string) {
	r.Language = language
}

func (r *Render) SetType(type_ *RenderType) {
	r.Type = type_
}

func (r *Render) SetMediaRef(mediaRef *DubbingMediaReference) {
	r.MediaRef = mediaRef
}

func (r *Render) SetStatus(status RenderStatus) {
	r.Status = status
}

func (r *Render) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r Render) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: Render to string"
	}
	return string(jsonData)
}

func (r *Render) UnmarshalJSON(data []byte) error {
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

func (r *Render) MarshalJSON() ([]byte, error) {
	type embed Render
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
