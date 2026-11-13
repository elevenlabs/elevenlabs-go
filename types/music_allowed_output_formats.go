package types

import (
	"encoding/json"
	"fmt"
)

type MusicAllowedOutputFormats struct {
	AllowedOutputFormats   AllowedOutputFormats
	MusicOnlyOutputFormats MusicOnlyOutputFormats
	typ                    string
}

func (m *MusicAllowedOutputFormats) GetAllowedOutputFormats() AllowedOutputFormats {
	if m == nil {
		return ""
	}
	return m.AllowedOutputFormats
}

func (m *MusicAllowedOutputFormats) GetMusicOnlyOutputFormats() MusicOnlyOutputFormats {
	if m == nil {
		return ""
	}
	return m.MusicOnlyOutputFormats
}

func (m *MusicAllowedOutputFormats) UnmarshalJSON(data []byte) error {
	var valueAllowedOutputFormats AllowedOutputFormats
	if err := json.Unmarshal(data, &valueAllowedOutputFormats); err == nil {
		m.typ = "AllowedOutputFormats"
		m.AllowedOutputFormats = valueAllowedOutputFormats
		return nil
	}
	var valueMusicOnlyOutputFormats MusicOnlyOutputFormats
	if err := json.Unmarshal(data, &valueMusicOnlyOutputFormats); err == nil {
		m.typ = "MusicOnlyOutputFormats"
		m.MusicOnlyOutputFormats = valueMusicOnlyOutputFormats
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MusicAllowedOutputFormats) MarshalJSON() ([]byte, error) {
	if m.typ == "AllowedOutputFormats" || m.AllowedOutputFormats != "" {
		return json.Marshal(m.AllowedOutputFormats)
	}
	if m.typ == "MusicOnlyOutputFormats" || m.MusicOnlyOutputFormats != "" {
		return json.Marshal(m.MusicOnlyOutputFormats)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MusicAllowedOutputFormatsVisitor interface {
	VisitAllowedOutputFormats(AllowedOutputFormats) error
	VisitMusicOnlyOutputFormats(MusicOnlyOutputFormats) error
}

func (m *MusicAllowedOutputFormats) Accept(visitor MusicAllowedOutputFormatsVisitor) error {
	if m.typ == "AllowedOutputFormats" || m.AllowedOutputFormats != "" {
		return visitor.VisitAllowedOutputFormats(m.AllowedOutputFormats)
	}
	if m.typ == "MusicOnlyOutputFormats" || m.MusicOnlyOutputFormats != "" {
		return visitor.VisitMusicOnlyOutputFormats(m.MusicOnlyOutputFormats)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
