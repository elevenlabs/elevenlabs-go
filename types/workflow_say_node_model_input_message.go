package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Message spoken by this node. Use `literal` for exact text or `prompt` to have the LLM generate the message from a description.
type WorkflowSayNodeModelInputMessage struct {
	Type    string
	Literal *SayNodeLiteralMessageInput
	Prompt  *SayNodePromptMessageInput
}

func (w *WorkflowSayNodeModelInputMessage) GetLiteral() *SayNodeLiteralMessageInput {
	if w == nil {
		return nil
	}
	if w.Literal == nil {
		return nil
	}
	return w.Literal
}

func (w *WorkflowSayNodeModelInputMessage) GetPrompt() *SayNodePromptMessageInput {
	if w == nil {
		return nil
	}
	if w.Prompt == nil {
		return nil
	}
	return w.Prompt
}

func (w *WorkflowSayNodeModelInputMessage) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowSayNodeModelInputMessage) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "literal":
		value := new(SayNodeLiteralMessageInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Literal = value
	case "prompt":
		value := new(SayNodePromptMessageInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Prompt = value
	}
	return nil
}

func (w WorkflowSayNodeModelInputMessage) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Literal != nil {
		return internal.MarshalJSONWithDiscriminant(w.Literal, "type", "literal")
	}
	if w.Prompt != nil {
		return internal.MarshalJSONWithDiscriminant(w.Prompt, "type", "prompt")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowSayNodeModelInputMessageVisitor interface {
	VisitLiteral(*SayNodeLiteralMessageInput) error
	VisitPrompt(*SayNodePromptMessageInput) error
}

func (w *WorkflowSayNodeModelInputMessage) Accept(visitor WorkflowSayNodeModelInputMessageVisitor) error {
	if w.Literal != nil {
		return visitor.VisitLiteral(w.Literal)
	}
	if w.Prompt != nil {
		return visitor.VisitPrompt(w.Prompt)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowSayNodeModelInputMessage) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Literal != nil {
		fields = append(fields, "literal")
	}
	if w.Prompt != nil {
		fields = append(fields, "prompt")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}
