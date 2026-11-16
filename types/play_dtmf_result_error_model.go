package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PlayDtmfResultErrorModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Error           string  `json:"error" xml:"error" required:"true"`
	Details         *string `json:"details,omitempty" xml:"details,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PlayDtmfResultErrorModel) GetStatus() (value *string) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *PlayDtmfResultErrorModel) GetError() (value string) {
	if p == nil {
		return
	}
	return p.Error
}

func (p *PlayDtmfResultErrorModel) GetDetails() (value *string) {
	if p == nil {
		return
	}
	return p.Details
}

func (p *PlayDtmfResultErrorModel) SetStatus(status *string) {
	p.Status = status
}

func (p *PlayDtmfResultErrorModel) SetError(error string) {
	p.Error = error
}

func (p *PlayDtmfResultErrorModel) SetDetails(details *string) {
	p.Details = details
}

func (p *PlayDtmfResultErrorModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlayDtmfResultErrorModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlayDtmfResultErrorModel to string"
	}
	return string(jsonData)
}

func (p *PlayDtmfResultErrorModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "result_type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlayDtmfResultErrorModel) MarshalJSON() ([]byte, error) {
	type embed PlayDtmfResultErrorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
