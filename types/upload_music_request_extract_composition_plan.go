package types

import (
	"encoding/json"
	"fmt"
)

// Whether to generate and return the composition plan for the uploaded song. Pass a model id (`music_v1`, `music_v2` or `music_v2_5`) to control which composition plan format is returned. Passing `true`/`false` is deprecated; `true` defaults to the `music_v1` plan format. Enabling this will increase the latency.
type UploadMusicRequestExtractCompositionPlan struct {
	Boolean      bool
	MusicModelId MusicModelId
	typ          string
}

func (u *UploadMusicRequestExtractCompositionPlan) GetBoolean() bool {
	if u == nil {
		return false
	}
	return u.Boolean
}

func (u *UploadMusicRequestExtractCompositionPlan) GetMusicModelId() MusicModelId {
	if u == nil {
		return ""
	}
	return u.MusicModelId
}

func (u *UploadMusicRequestExtractCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		u.typ = "Boolean"
		u.Boolean = valueBoolean
		return nil
	}
	var valueMusicModelId MusicModelId
	if err := json.Unmarshal(data, &valueMusicModelId); err == nil {
		u.typ = "MusicModelId"
		u.MusicModelId = valueMusicModelId
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UploadMusicRequestExtractCompositionPlan) MarshalJSON() ([]byte, error) {
	if u.typ == "Boolean" || u.Boolean != false {
		return json.Marshal(u.Boolean)
	}
	if u.typ == "MusicModelId" || u.MusicModelId != "" {
		return json.Marshal(u.MusicModelId)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UploadMusicRequestExtractCompositionPlanVisitor interface {
	VisitBoolean(bool) error
	VisitMusicModelId(MusicModelId) error
}

func (u *UploadMusicRequestExtractCompositionPlan) Accept(visitor UploadMusicRequestExtractCompositionPlanVisitor) error {
	if u.typ == "Boolean" || u.Boolean != false {
		return visitor.VisitBoolean(u.Boolean)
	}
	if u.typ == "MusicModelId" || u.MusicModelId != "" {
		return visitor.VisitMusicModelId(u.MusicModelId)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}
