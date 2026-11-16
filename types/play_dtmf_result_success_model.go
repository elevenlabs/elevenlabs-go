package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PlayDtmfResultSuccessModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	DtmfTones       string  `json:"dtmf_tones" xml:"dtmf_tones" required:"true"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PlayDtmfResultSuccessModel) GetStatus() (value *string) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *PlayDtmfResultSuccessModel) GetDtmfTones() (value string) {
	if p == nil {
		return
	}
	return p.DtmfTones
}

func (p *PlayDtmfResultSuccessModel) GetReason() (value *string) {
	if p == nil {
		return
	}
	return p.Reason
}

func (p *PlayDtmfResultSuccessModel) SetStatus(status *string) {
	p.Status = status
}

func (p *PlayDtmfResultSuccessModel) SetDtmfTones(dtmfTones string) {
	p.DtmfTones = dtmfTones
}

func (p *PlayDtmfResultSuccessModel) SetReason(reason *string) {
	p.Reason = reason
}

func (p *PlayDtmfResultSuccessModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlayDtmfResultSuccessModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlayDtmfResultSuccessModel to string"
	}
	return string(jsonData)
}

func (p *PlayDtmfResultSuccessModel) UnmarshalJSON(data []byte) error {
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

func (p *PlayDtmfResultSuccessModel) MarshalJSON() ([]byte, error) {
	type embed PlayDtmfResultSuccessModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
