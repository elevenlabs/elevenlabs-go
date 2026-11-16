package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PendingClipTask struct {
	Type            PendingClipTaskType `json:"type" xml:"type" required:"true"`
	Progress        *float64            `json:"progress,omitempty" xml:"progress,omitempty"`
	StartedAtMs     *int                `json:"started_at_ms,omitempty" xml:"started_at_ms,omitempty"`
	UpdatedAtMs     *int                `json:"updated_at_ms,omitempty" xml:"updated_at_ms,omitempty"`
	Metadata        map[string]any      `json:"metadata,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PendingClipTask) GetType() (value PendingClipTaskType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *PendingClipTask) GetProgress() (value *float64) {
	if p == nil {
		return
	}
	return p.Progress
}

func (p *PendingClipTask) GetStartedAtMs() (value *int) {
	if p == nil {
		return
	}
	return p.StartedAtMs
}

func (p *PendingClipTask) GetUpdatedAtMs() (value *int) {
	if p == nil {
		return
	}
	return p.UpdatedAtMs
}

func (p *PendingClipTask) GetMetadata() (value map[string]any) {
	if p == nil {
		return
	}
	return p.Metadata
}

func (p *PendingClipTask) SetType(type_ PendingClipTaskType) {
	p.Type = type_
}

func (p *PendingClipTask) SetProgress(progress *float64) {
	p.Progress = progress
}

func (p *PendingClipTask) SetStartedAtMs(startedAtMs *int) {
	p.StartedAtMs = startedAtMs
}

func (p *PendingClipTask) SetUpdatedAtMs(updatedAtMs *int) {
	p.UpdatedAtMs = updatedAtMs
}

func (p *PendingClipTask) SetMetadata(metadata map[string]any) {
	p.Metadata = metadata
}

func (p *PendingClipTask) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PendingClipTask) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PendingClipTask to string"
	}
	return string(jsonData)
}

func (p *PendingClipTask) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PendingClipTask
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PendingClipTask(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PendingClipTask) MarshalJSON() ([]byte, error) {
	type embed PendingClipTask
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
