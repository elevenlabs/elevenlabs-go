package types

import (
	"encoding/json"
	"fmt"
)

// Whether to generate and return the composition plan for the uploaded song. Pass a model id (`music_v1`, `music_v2` or `music_v2_5`) to control which composition plan format is returned. Passing `true`/`false` is deprecated; `true` defaults to the `music_v1` plan format. Enabling this will increase the latency.
type BodyUploadMusicV1MusicUploadPostExtractCompositionPlan struct {
	Boolean      bool
	MusicModelId MusicModelId
	typ          string
}

func (bodyUploadMusicV1MusicUploadPostExtractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) GetBoolean() bool {
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan == nil {
		return false
	}
	return bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean
}

func (bodyUploadMusicV1MusicUploadPostExtractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) GetMusicModelId() MusicModelId {
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan == nil {
		return ""
	}
	return bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId
}

func (bodyUploadMusicV1MusicUploadPostExtractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ = "Boolean"
		bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean = valueBoolean
		return nil
	}
	var valueMusicModelId MusicModelId
	if err := json.Unmarshal(data, &valueMusicModelId); err == nil {
		bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ = "MusicModelId"
		bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId = valueMusicModelId
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, bodyUploadMusicV1MusicUploadPostExtractCompositionPlan)
}

func (bodyUploadMusicV1MusicUploadPostExtractCompositionPlan BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) MarshalJSON() ([]byte, error) {
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ == "Boolean" || bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean != false {
		return json.Marshal(bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean)
	}
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ == "MusicModelId" || bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId != "" {
		return json.Marshal(bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", bodyUploadMusicV1MusicUploadPostExtractCompositionPlan)
}

type BodyUploadMusicV1MusicUploadPostExtractCompositionPlanVisitor interface {
	VisitBoolean(bool) error
	VisitMusicModelId(MusicModelId) error
}

func (bodyUploadMusicV1MusicUploadPostExtractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) Accept(visitor BodyUploadMusicV1MusicUploadPostExtractCompositionPlanVisitor) error {
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ == "Boolean" || bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean != false {
		return visitor.VisitBoolean(bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.Boolean)
	}
	if bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.typ == "MusicModelId" || bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId != "" {
		return visitor.VisitMusicModelId(bodyUploadMusicV1MusicUploadPostExtractCompositionPlan.MusicModelId)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", bodyUploadMusicV1MusicUploadPostExtractCompositionPlan)
}
