package chapters

import "encoding/json"

type AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost struct {
	// Whether to convert the audio to mpeg format.
	ConvertToMpeg *bool `json:"convert_to_mpeg,omitempty" xml:"convert_to_mpeg,omitempty"`
}

func (a *AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) GetConvertToMpeg() (value *bool) {
	if a == nil {
		return
	}
	return a.ConvertToMpeg
}

func (a *AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) SetConvertToMpeg(convertToMpeg *bool) {
	a.ConvertToMpeg = convertToMpeg
}

func (a *AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost to string"
	}
	return string(jsonData)
}

func (a *AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost(value)
	return nil
}

func (a *AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost) MarshalJSON() ([]byte, error) {
	type embed AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
