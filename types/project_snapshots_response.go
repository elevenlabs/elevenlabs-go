package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ProjectSnapshotsResponse struct {
	// List of project snapshots.
	Snapshots       []*ProjectSnapshotResponse `json:"snapshots" xml:"snapshots" required:"true"`
	extraProperties map[string]interface{}
}

func (p *ProjectSnapshotsResponse) GetSnapshots() (value []*ProjectSnapshotResponse) {
	if p == nil {
		return
	}
	return p.Snapshots
}

func (p *ProjectSnapshotsResponse) SetSnapshots(snapshots []*ProjectSnapshotResponse) {
	p.Snapshots = snapshots
}

func (p *ProjectSnapshotsResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectSnapshotsResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectSnapshotsResponse to string"
	}
	return string(jsonData)
}

func (p *ProjectSnapshotsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ProjectSnapshotsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ProjectSnapshotsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectSnapshotsResponse) MarshalJSON() ([]byte, error) {
	type embed ProjectSnapshotsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
