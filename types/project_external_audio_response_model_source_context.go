package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ProjectExternalAudioResponseModelSourceContext struct {
	SourceType       string
	MusicExploreSong *MusicExploreSongSourceContext
	Sfx              *SfxSourceContext
	Song             *SongSourceContext
}

func (p *ProjectExternalAudioResponseModelSourceContext) GetMusicExploreSong() *MusicExploreSongSourceContext {
	if p == nil {
		return nil
	}
	if p.MusicExploreSong == nil {
		return nil
	}
	return p.MusicExploreSong
}

func (p *ProjectExternalAudioResponseModelSourceContext) GetSfx() *SfxSourceContext {
	if p == nil {
		return nil
	}
	if p.Sfx == nil {
		return nil
	}
	return p.Sfx
}

func (p *ProjectExternalAudioResponseModelSourceContext) GetSong() *SongSourceContext {
	if p == nil {
		return nil
	}
	if p.Song == nil {
		return nil
	}
	return p.Song
}

func (p *ProjectExternalAudioResponseModelSourceContext) GetSourceType() string {
	if p == nil {
		return ""
	}
	return p.SourceType
}

func (p *ProjectExternalAudioResponseModelSourceContext) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		SourceType string `json:"source_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.SourceType = unmarshaler.SourceType
	if unmarshaler.SourceType == "" {
		return fmt.Errorf("%T did not include discriminant source_type", p)
	}
	switch unmarshaler.SourceType {
	case "music_explore_song":
		value := new(MusicExploreSongSourceContext)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.MusicExploreSong = value
	case "sfx":
		value := new(SfxSourceContext)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Sfx = value
	case "song":
		value := new(SongSourceContext)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Song = value
	}
	return nil
}

func (p ProjectExternalAudioResponseModelSourceContext) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.MusicExploreSong != nil {
		return internal.MarshalJSONWithDiscriminant(p.MusicExploreSong, "source_type", "music_explore_song")
	}
	if p.Sfx != nil {
		return internal.MarshalJSONWithDiscriminant(p.Sfx, "source_type", "sfx")
	}
	if p.Song != nil {
		return internal.MarshalJSONWithDiscriminant(p.Song, "source_type", "song")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type ProjectExternalAudioResponseModelSourceContextVisitor interface {
	VisitMusicExploreSong(*MusicExploreSongSourceContext) error
	VisitSfx(*SfxSourceContext) error
	VisitSong(*SongSourceContext) error
}

func (p *ProjectExternalAudioResponseModelSourceContext) Accept(visitor ProjectExternalAudioResponseModelSourceContextVisitor) error {
	if p.MusicExploreSong != nil {
		return visitor.VisitMusicExploreSong(p.MusicExploreSong)
	}
	if p.Sfx != nil {
		return visitor.VisitSfx(p.Sfx)
	}
	if p.Song != nil {
		return visitor.VisitSong(p.Song)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *ProjectExternalAudioResponseModelSourceContext) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.MusicExploreSong != nil {
		fields = append(fields, "music_explore_song")
	}
	if p.Sfx != nil {
		fields = append(fields, "sfx")
	}
	if p.Song != nil {
		fields = append(fields, "song")
	}
	if len(fields) == 0 {
		if p.SourceType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.SourceType)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.SourceType != "" {
		field := fields[0]
		if p.SourceType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.SourceType,
				p,
			)
		}
	}
	return nil
}
