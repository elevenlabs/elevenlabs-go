package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for backup LLM cascading. Can be disabled, use system defaults, or specify custom order.
type PromptAgentApiModelInputBackupLlmConfig struct {
	Preference string
	Default    *BackupLlmDefault
	Disabled   *BackupLlmDisabled
	Override   *BackupLlmOverride
}

func (p *PromptAgentApiModelInputBackupLlmConfig) GetDefault() *BackupLlmDefault {
	if p == nil {
		return nil
	}
	if p.Default == nil {
		return nil
	}
	return p.Default
}

func (p *PromptAgentApiModelInputBackupLlmConfig) GetDisabled() *BackupLlmDisabled {
	if p == nil {
		return nil
	}
	if p.Disabled == nil {
		return nil
	}
	return p.Disabled
}

func (p *PromptAgentApiModelInputBackupLlmConfig) GetOverride() *BackupLlmOverride {
	if p == nil {
		return nil
	}
	if p.Override == nil {
		return nil
	}
	return p.Override
}

func (p *PromptAgentApiModelInputBackupLlmConfig) GetPreference() string {
	if p == nil {
		return ""
	}
	return p.Preference
}

func (p *PromptAgentApiModelInputBackupLlmConfig) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Preference string `json:"preference"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Preference = unmarshaler.Preference
	if unmarshaler.Preference == "" {
		return fmt.Errorf("%T did not include discriminant preference", p)
	}
	switch unmarshaler.Preference {
	case "default":
		value := new(BackupLlmDefault)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Default = value
	case "disabled":
		value := new(BackupLlmDisabled)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Disabled = value
	case "override":
		value := new(BackupLlmOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Override = value
	}
	return nil
}

func (p PromptAgentApiModelInputBackupLlmConfig) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.Default != nil {
		return internal.MarshalJSONWithDiscriminant(p.Default, "preference", "default")
	}
	if p.Disabled != nil {
		return internal.MarshalJSONWithDiscriminant(p.Disabled, "preference", "disabled")
	}
	if p.Override != nil {
		return internal.MarshalJSONWithDiscriminant(p.Override, "preference", "override")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PromptAgentApiModelInputBackupLlmConfigVisitor interface {
	VisitDefault(*BackupLlmDefault) error
	VisitDisabled(*BackupLlmDisabled) error
	VisitOverride(*BackupLlmOverride) error
}

func (p *PromptAgentApiModelInputBackupLlmConfig) Accept(visitor PromptAgentApiModelInputBackupLlmConfigVisitor) error {
	if p.Default != nil {
		return visitor.VisitDefault(p.Default)
	}
	if p.Disabled != nil {
		return visitor.VisitDisabled(p.Disabled)
	}
	if p.Override != nil {
		return visitor.VisitOverride(p.Override)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PromptAgentApiModelInputBackupLlmConfig) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.Default != nil {
		fields = append(fields, "default")
	}
	if p.Disabled != nil {
		fields = append(fields, "disabled")
	}
	if p.Override != nil {
		fields = append(fields, "override")
	}
	if len(fields) == 0 {
		if p.Preference != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Preference)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Preference != "" {
		field := fields[0]
		if p.Preference != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Preference,
				p,
			)
		}
	}
	return nil
}
