package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CharacterResponseModel struct {
	CharacterId         string                          `json:"character_id" xml:"character_id" required:"true"`
	Name                string                          `json:"name" xml:"name" required:"true"`
	Metadata            *CharacterMetadataResponseModel `json:"metadata" xml:"metadata" required:"true"`
	RecommendedVoiceIds []string                        `json:"recommended_voice_ids,omitempty" xml:"recommended_voice_ids,omitempty"`
	extraProperties     map[string]interface{}
}

func (c *CharacterResponseModel) GetCharacterId() (value string) {
	if c == nil {
		return
	}
	return c.CharacterId
}

func (c *CharacterResponseModel) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CharacterResponseModel) GetMetadata() (value *CharacterMetadataResponseModel) {
	if c == nil {
		return
	}
	return c.Metadata
}

func (c *CharacterResponseModel) GetRecommendedVoiceIds() (value []string) {
	if c == nil {
		return
	}
	return c.RecommendedVoiceIds
}

func (c *CharacterResponseModel) SetCharacterId(characterId string) {
	c.CharacterId = characterId
}

func (c *CharacterResponseModel) SetName(name string) {
	c.Name = name
}

func (c *CharacterResponseModel) SetMetadata(metadata *CharacterMetadataResponseModel) {
	c.Metadata = metadata
}

func (c *CharacterResponseModel) SetRecommendedVoiceIds(recommendedVoiceIds []string) {
	c.RecommendedVoiceIds = recommendedVoiceIds
}

func (c *CharacterResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CharacterResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CharacterResponseModel to string"
	}
	return string(jsonData)
}

func (c *CharacterResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CharacterResponseModel) MarshalJSON() ([]byte, error) {
	type embed CharacterResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
