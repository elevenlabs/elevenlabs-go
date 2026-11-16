package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterContentParagraphTtsNodeInputModel struct {
	type_           string
	Text            string `json:"text" xml:"text" required:"true"`
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterContentParagraphTtsNodeInputModel) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *ChapterContentParagraphTtsNodeInputModel) GetVoiceId() (value string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ChapterContentParagraphTtsNodeInputModel) SetText(text string) {
	c.Text = text
}

func (c *ChapterContentParagraphTtsNodeInputModel) SetVoiceId(voiceId string) {
	c.VoiceId = voiceId
}

func (c *ChapterContentParagraphTtsNodeInputModel) Type() string {
	return c.type_
}

func (c *ChapterContentParagraphTtsNodeInputModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentParagraphTtsNodeInputModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentParagraphTtsNodeInputModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentParagraphTtsNodeInputModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterContentParagraphTtsNodeInputModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterContentParagraphTtsNodeInputModel(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "tts_node" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "tts_node", *literals.Type)
	}
	c.type_ = "tts_node"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterContentParagraphTtsNodeInputModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentParagraphTtsNodeInputModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "tts_node",
	}
	return json.Marshal(marshaler)
}
