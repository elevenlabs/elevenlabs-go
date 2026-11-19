package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoSegment struct {
	StartMs         int      `json:"start_ms" xml:"start_ms" required:"true"`
	EndMs           int      `json:"end_ms" xml:"end_ms" required:"true"`
	Description     string   `json:"description" xml:"description" required:"true"`
	Subjects        []string `json:"subjects,omitempty" xml:"subjects,omitempty"`
	ShotType        *string  `json:"shot_type,omitempty" xml:"shot_type,omitempty"`
	CameraMovement  *string  `json:"camera_movement,omitempty" xml:"camera_movement,omitempty"`
	TransitionIn    *string  `json:"transition_in,omitempty" xml:"transition_in,omitempty"`
	HasSpeech       *bool    `json:"has_speech,omitempty" xml:"has_speech,omitempty"`
	HasMusic        *bool    `json:"has_music,omitempty" xml:"has_music,omitempty"`
	Pacing          *string  `json:"pacing,omitempty" xml:"pacing,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VideoSegment) GetStartMs() (value int) {
	if v == nil {
		return
	}
	return v.StartMs
}

func (v *VideoSegment) GetEndMs() (value int) {
	if v == nil {
		return
	}
	return v.EndMs
}

func (v *VideoSegment) GetDescription() (value string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VideoSegment) GetSubjects() (value []string) {
	if v == nil {
		return
	}
	return v.Subjects
}

func (v *VideoSegment) GetShotType() (value *string) {
	if v == nil {
		return
	}
	return v.ShotType
}

func (v *VideoSegment) GetCameraMovement() (value *string) {
	if v == nil {
		return
	}
	return v.CameraMovement
}

func (v *VideoSegment) GetTransitionIn() (value *string) {
	if v == nil {
		return
	}
	return v.TransitionIn
}

func (v *VideoSegment) GetHasSpeech() (value *bool) {
	if v == nil {
		return
	}
	return v.HasSpeech
}

func (v *VideoSegment) GetHasMusic() (value *bool) {
	if v == nil {
		return
	}
	return v.HasMusic
}

func (v *VideoSegment) GetPacing() (value *string) {
	if v == nil {
		return
	}
	return v.Pacing
}

func (v *VideoSegment) SetStartMs(startMs int) {
	v.StartMs = startMs
}

func (v *VideoSegment) SetEndMs(endMs int) {
	v.EndMs = endMs
}

func (v *VideoSegment) SetDescription(description string) {
	v.Description = description
}

func (v *VideoSegment) SetSubjects(subjects []string) {
	v.Subjects = subjects
}

func (v *VideoSegment) SetShotType(shotType *string) {
	v.ShotType = shotType
}

func (v *VideoSegment) SetCameraMovement(cameraMovement *string) {
	v.CameraMovement = cameraMovement
}

func (v *VideoSegment) SetTransitionIn(transitionIn *string) {
	v.TransitionIn = transitionIn
}

func (v *VideoSegment) SetHasSpeech(hasSpeech *bool) {
	v.HasSpeech = hasSpeech
}

func (v *VideoSegment) SetHasMusic(hasMusic *bool) {
	v.HasMusic = hasMusic
}

func (v *VideoSegment) SetPacing(pacing *string) {
	v.Pacing = pacing
}

func (v *VideoSegment) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoSegment) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoSegment to string"
	}
	return string(jsonData)
}

func (v *VideoSegment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoSegment) MarshalJSON() ([]byte, error) {
	type embed VideoSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
