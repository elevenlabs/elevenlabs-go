package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterContentBlockTtsNodeResponseModel struct {
	type_             string
	ProjectVoiceRefId string  `json:"project_voice_ref_id" xml:"project_voice_ref_id" required:"true"`
	Text              string  `json:"text" xml:"text" required:"true"`
	VoiceId           *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ChapterContentBlockTtsNodeResponseModel) GetProjectVoiceRefId() (value string) {
	if c == nil {
		return
	}
	return c.ProjectVoiceRefId
}

func (c *ChapterContentBlockTtsNodeResponseModel) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *ChapterContentBlockTtsNodeResponseModel) GetVoiceId() (value *string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ChapterContentBlockTtsNodeResponseModel) SetProjectVoiceRefId(projectVoiceRefId string) {
	c.ProjectVoiceRefId = projectVoiceRefId
}

func (c *ChapterContentBlockTtsNodeResponseModel) SetText(text string) {
	c.Text = text
}

func (c *ChapterContentBlockTtsNodeResponseModel) SetVoiceId(voiceId *string) {
	c.VoiceId = voiceId
}

func (c *ChapterContentBlockTtsNodeResponseModel) Type() string {
	return c.type_
}

func (c *ChapterContentBlockTtsNodeResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentBlockTtsNodeResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentBlockTtsNodeResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentBlockTtsNodeResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterContentBlockTtsNodeResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterContentBlockTtsNodeResponseModel(tmp)
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

func (c *ChapterContentBlockTtsNodeResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentBlockTtsNodeResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "tts_node",
	}
	return json.Marshal(marshaler)
}
