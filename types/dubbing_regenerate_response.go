package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The accepted re-dub: what it covers and what it cost.
type DubbingRegenerateResponse struct {
	// The segments this re-dub re-synthesizes: those with edits to apply.
	RegeneratedSegmentIds []string `json:"regenerated_segment_ids" xml:"regenerated_segment_ids" required:"true"`
	// Seconds of audio this re-dub covers — the edited regions only, never the whole target. `charged_seconds` is the part of it that was billed.
	RegeneratedSeconds float64 `json:"regenerated_seconds" xml:"regenerated_seconds" required:"true"`
	// Seconds actually billed, after the free-regeneration allowance. Zero when the re-dub cost nothing — either the allowance covered all of it, or the project's included generation did.
	ChargedSeconds float64 `json:"charged_seconds" xml:"charged_seconds" required:"true"`
	// Free-regeneration seconds left for this language target after this re-dub. The allowance is the source's own duration.
	FreeRegenerationSecondsRemaining float64 `json:"free_regeneration_seconds_remaining" xml:"free_regeneration_seconds_remaining" required:"true"`
	extraProperties                  map[string]interface{}
}

func (d *DubbingRegenerateResponse) GetRegeneratedSegmentIds() (value []string) {
	if d == nil {
		return
	}
	return d.RegeneratedSegmentIds
}

func (d *DubbingRegenerateResponse) GetRegeneratedSeconds() (value float64) {
	if d == nil {
		return
	}
	return d.RegeneratedSeconds
}

func (d *DubbingRegenerateResponse) GetChargedSeconds() (value float64) {
	if d == nil {
		return
	}
	return d.ChargedSeconds
}

func (d *DubbingRegenerateResponse) GetFreeRegenerationSecondsRemaining() (value float64) {
	if d == nil {
		return
	}
	return d.FreeRegenerationSecondsRemaining
}

func (d *DubbingRegenerateResponse) SetRegeneratedSegmentIds(regeneratedSegmentIds []string) {
	d.RegeneratedSegmentIds = regeneratedSegmentIds
}

func (d *DubbingRegenerateResponse) SetRegeneratedSeconds(regeneratedSeconds float64) {
	d.RegeneratedSeconds = regeneratedSeconds
}

func (d *DubbingRegenerateResponse) SetChargedSeconds(chargedSeconds float64) {
	d.ChargedSeconds = chargedSeconds
}

func (d *DubbingRegenerateResponse) SetFreeRegenerationSecondsRemaining(freeRegenerationSecondsRemaining float64) {
	d.FreeRegenerationSecondsRemaining = freeRegenerationSecondsRemaining
}

func (d *DubbingRegenerateResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingRegenerateResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingRegenerateResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingRegenerateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingRegenerateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingRegenerateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingRegenerateResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingRegenerateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
