package types

import (
	"encoding/json"
	"fmt"
)

type MusicOutputFormat struct {
	MusicAllowedOutputFormats *MusicAllowedOutputFormats
	AutoStringLiteral         string
	typ                       string
}

func (m *MusicOutputFormat) GetMusicAllowedOutputFormats() *MusicAllowedOutputFormats {
	if m == nil {
		return nil
	}
	if m.MusicAllowedOutputFormats == nil {
		return nil
	}
	return m.MusicAllowedOutputFormats
}

func (m *MusicOutputFormat) GetAutoStringLiteral() string {
	if m == nil {
		return ""
	}
	return m.AutoStringLiteral
}

func (m *MusicOutputFormat) UnmarshalJSON(data []byte) error {
	var valueMusicAllowedOutputFormats *MusicAllowedOutputFormats
	if err := json.Unmarshal(data, &valueMusicAllowedOutputFormats); err == nil {
		m.typ = "MusicAllowedOutputFormats"
		m.MusicAllowedOutputFormats = valueMusicAllowedOutputFormats
		return nil
	}
	var valueAutoStringLiteral string
	if err := json.Unmarshal(data, &valueAutoStringLiteral); err == nil {
		m.typ = "AutoStringLiteral"
		m.AutoStringLiteral = valueAutoStringLiteral
		if m.AutoStringLiteral != "auto" {
			return fmt.Errorf("unexpected value for literal on type %T; expected auto got %v", m, m.AutoStringLiteral)
		}
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MusicOutputFormat) MarshalJSON() ([]byte, error) {
	if m.typ == "MusicAllowedOutputFormats" || m.MusicAllowedOutputFormats != nil {
		return json.Marshal(m.MusicAllowedOutputFormats)
	}
	if m.typ == "AutoStringLiteral" || m.AutoStringLiteral != "" {
		return json.Marshal(m.AutoStringLiteral)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MusicOutputFormatVisitor interface {
	VisitMusicAllowedOutputFormats(*MusicAllowedOutputFormats) error
	VisitAutoStringLiteral(string) error
}

func (m *MusicOutputFormat) Accept(visitor MusicOutputFormatVisitor) error {
	if m.typ == "MusicAllowedOutputFormats" || m.MusicAllowedOutputFormats != nil {
		return visitor.VisitMusicAllowedOutputFormats(m.MusicAllowedOutputFormats)
	}
	if m.typ == "AutoStringLiteral" || m.AutoStringLiteral != "" {
		return visitor.VisitAutoStringLiteral(m.AutoStringLiteral)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

func NewMusicOutputFormatWithAutoStringLiteral() *MusicOutputFormat {
	return &MusicOutputFormat{typ: "AutoStringLiteral", AutoStringLiteral: "auto"}
}
