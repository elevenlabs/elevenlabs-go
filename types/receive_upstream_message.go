package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ReceiveUpstreamMessage struct {
	Type           string
	Close          *Close
	Error          *SpeechEngineError
	Init           *Init
	Ping           *Ping
	UserTranscript *UserTranscript
}

func (r *ReceiveUpstreamMessage) GetClose() *Close {
	if r == nil {
		return nil
	}
	if r.Close == nil {
		return nil
	}
	return r.Close
}

func (r *ReceiveUpstreamMessage) GetError() *SpeechEngineError {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveUpstreamMessage) GetInit() *Init {
	if r == nil {
		return nil
	}
	if r.Init == nil {
		return nil
	}
	return r.Init
}

func (r *ReceiveUpstreamMessage) GetPing() *Ping {
	if r == nil {
		return nil
	}
	if r.Ping == nil {
		return nil
	}
	return r.Ping
}

func (r *ReceiveUpstreamMessage) GetUserTranscript() *UserTranscript {
	if r == nil {
		return nil
	}
	if r.UserTranscript == nil {
		return nil
	}
	return r.UserTranscript
}

func (r *ReceiveUpstreamMessage) GetType() string {
	if r == nil {
		return ""
	}
	return r.Type
}

func (r *ReceiveUpstreamMessage) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", r)
	}
	switch unmarshaler.Type {
	case "close":
		value := new(Close)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Close = value
	case "error":
		value := new(SpeechEngineError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Error = value
	case "init":
		value := new(Init)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Init = value
	case "ping":
		value := new(Ping)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Ping = value
	case "user_transcript":
		value := new(UserTranscript)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.UserTranscript = value
	}
	return nil
}

func (r ReceiveUpstreamMessage) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.Close != nil {
		return internal.MarshalJSONWithDiscriminant(r.Close, "type", "close")
	}
	if r.Error != nil {
		return internal.MarshalJSONWithDiscriminant(r.Error, "type", "error")
	}
	if r.Init != nil {
		return internal.MarshalJSONWithDiscriminant(r.Init, "type", "init")
	}
	if r.Ping != nil {
		return internal.MarshalJSONWithDiscriminant(r.Ping, "type", "ping")
	}
	if r.UserTranscript != nil {
		return internal.MarshalJSONWithDiscriminant(r.UserTranscript, "type", "user_transcript")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type ReceiveUpstreamMessageVisitor interface {
	VisitClose(*Close) error
	VisitError(*SpeechEngineError) error
	VisitInit(*Init) error
	VisitPing(*Ping) error
	VisitUserTranscript(*UserTranscript) error
}

func (r *ReceiveUpstreamMessage) Accept(visitor ReceiveUpstreamMessageVisitor) error {
	if r.Close != nil {
		return visitor.VisitClose(r.Close)
	}
	if r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	if r.Init != nil {
		return visitor.VisitInit(r.Init)
	}
	if r.Ping != nil {
		return visitor.VisitPing(r.Ping)
	}
	if r.UserTranscript != nil {
		return visitor.VisitUserTranscript(r.UserTranscript)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *ReceiveUpstreamMessage) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.Close != nil {
		fields = append(fields, "close")
	}
	if r.Error != nil {
		fields = append(fields, "error")
	}
	if r.Init != nil {
		fields = append(fields, "init")
	}
	if r.Ping != nil {
		fields = append(fields, "ping")
	}
	if r.UserTranscript != nil {
		fields = append(fields, "user_transcript")
	}
	if len(fields) == 0 {
		if r.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", r, r.Type)
		}
		return fmt.Errorf("type %T is empty", r)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", r, fields)
	}
	if r.Type != "" {
		field := fields[0]
		if r.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				r,
				r.Type,
				r,
			)
		}
	}
	return nil
}
