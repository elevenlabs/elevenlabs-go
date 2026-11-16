package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PendingCancellationResponseModel struct {
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The timestamp of the cancellation.
	TimestampSeconds int `json:"timestamp_seconds" xml:"timestamp_seconds" required:"true"`
	extraProperties  map[string]interface{}
}

func (p *PendingCancellationResponseModel) GetKind() (value *string) {
	if p == nil {
		return
	}
	return p.Kind
}

func (p *PendingCancellationResponseModel) GetTimestampSeconds() (value int) {
	if p == nil {
		return
	}
	return p.TimestampSeconds
}

func (p *PendingCancellationResponseModel) SetKind(kind *string) {
	p.Kind = kind
}

func (p *PendingCancellationResponseModel) SetTimestampSeconds(timestampSeconds int) {
	p.TimestampSeconds = timestampSeconds
}

func (p *PendingCancellationResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PendingCancellationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PendingCancellationResponseModel to string"
	}
	return string(jsonData)
}

func (p *PendingCancellationResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PendingCancellationResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PendingCancellationResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PendingCancellationResponseModel) MarshalJSON() ([]byte, error) {
	type embed PendingCancellationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
