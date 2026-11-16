package projects

import "encoding/json"

type StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost struct {
	// Whether to convert the audio to mpeg format.
	ConvertToMpeg *bool `json:"convert_to_mpeg,omitempty" xml:"convert_to_mpeg,omitempty"`
}

func (s *StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetConvertToMpeg() (value *bool) {
	if s == nil {
		return
	}
	return s.ConvertToMpeg
}

func (s *StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) SetConvertToMpeg(convertToMpeg *bool) {
	s.ConvertToMpeg = convertToMpeg
}

func (s *StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (s StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost to string"
	}
	return string(jsonData)
}

func (s *StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) UnmarshalJSON(data []byte) error {
	type unmarshaler StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost(value)
	return nil
}

func (s *StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost) MarshalJSON() ([]byte, error) {
	type embed StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
