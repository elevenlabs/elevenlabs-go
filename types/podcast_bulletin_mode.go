package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastBulletinMode struct {
	// The type of podcast to create.
	type_           string
	Bulletin        *PodcastBulletinModeData `json:"bulletin" xml:"bulletin" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastBulletinMode) GetBulletin() (value *PodcastBulletinModeData) {
	if p == nil {
		return
	}
	return p.Bulletin
}

func (p *PodcastBulletinMode) SetBulletin(bulletin *PodcastBulletinModeData) {
	p.Bulletin = bulletin
}

func (p *PodcastBulletinMode) Type() string {
	return p.type_
}

func (p *PodcastBulletinMode) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastBulletinMode) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastBulletinMode to string"
	}
	return string(jsonData)
}

func (p *PodcastBulletinMode) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastBulletinMode
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastBulletinMode(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "bulletin" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "bulletin", *literals.Type)
	}
	p.type_ = "bulletin"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastBulletinMode) MarshalJSON() ([]byte, error) {
	type embed PodcastBulletinMode
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "bulletin",
	}
	return json.Marshal(marshaler)
}
