package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SayNodePromptMessageOutput struct {
	type_ string
	// Description of the message the LLM should generate. It is appended to the agent's system prompt while this node is active.
	Prompt          string `json:"prompt" xml:"prompt" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SayNodePromptMessageOutput) GetPrompt() (value string) {
	if s == nil {
		return
	}
	return s.Prompt
}

func (s *SayNodePromptMessageOutput) SetPrompt(prompt string) {
	s.Prompt = prompt
}

func (s *SayNodePromptMessageOutput) Type() string {
	return s.type_
}

func (s *SayNodePromptMessageOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SayNodePromptMessageOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SayNodePromptMessageOutput to string"
	}
	return string(jsonData)
}

func (s *SayNodePromptMessageOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SayNodePromptMessageOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SayNodePromptMessageOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "prompt" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "prompt", *literals.Type)
	}
	s.type_ = "prompt"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SayNodePromptMessageOutput) MarshalJSON() ([]byte, error) {
	type embed SayNodePromptMessageOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "prompt",
	}
	return json.Marshal(marshaler)
}
