package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SayNodePromptMessageInput struct {
	// Description of the message the LLM should generate. It is appended to the agent's system prompt while this node is active.
	Prompt          string `json:"prompt" xml:"prompt" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SayNodePromptMessageInput) GetPrompt() (value string) {
	if s == nil {
		return
	}
	return s.Prompt
}

func (s *SayNodePromptMessageInput) SetPrompt(prompt string) {
	s.Prompt = prompt
}

func (s *SayNodePromptMessageInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SayNodePromptMessageInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SayNodePromptMessageInput to string"
	}
	return string(jsonData)
}

func (s *SayNodePromptMessageInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SayNodePromptMessageInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SayNodePromptMessageInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SayNodePromptMessageInput) MarshalJSON() ([]byte, error) {
	type embed SayNodePromptMessageInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
