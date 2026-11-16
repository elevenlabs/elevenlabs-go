package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ChapterSnapshotsResponse struct {
	// List of chapter snapshots.
	Snapshots       []*ChapterSnapshotResponse `json:"snapshots" xml:"snapshots" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterSnapshotsResponse) GetSnapshots() (value []*ChapterSnapshotResponse) {
	if c == nil {
		return
	}
	return c.Snapshots
}

func (c *ChapterSnapshotsResponse) SetSnapshots(snapshots []*ChapterSnapshotResponse) {
	c.Snapshots = snapshots
}

func (c *ChapterSnapshotsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterSnapshotsResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterSnapshotsResponse to string"
	}
	return string(jsonData)
}

func (c *ChapterSnapshotsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ChapterSnapshotsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ChapterSnapshotsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterSnapshotsResponse) MarshalJSON() ([]byte, error) {
	type embed ChapterSnapshotsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
