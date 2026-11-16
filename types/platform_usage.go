package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Per-category breakdown of “platform_charge“ (the analogue of “llm_usage“).
type PlatformUsage struct {
	CategoryUsage   map[string]*PlatformCategoryUsage `json:"category_usage,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PlatformUsage) GetCategoryUsage() (value map[string]*PlatformCategoryUsage) {
	if p == nil {
		return
	}
	return p.CategoryUsage
}

func (p *PlatformUsage) SetCategoryUsage(categoryUsage map[string]*PlatformCategoryUsage) {
	p.CategoryUsage = categoryUsage
}

func (p *PlatformUsage) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlatformUsage) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlatformUsage to string"
	}
	return string(jsonData)
}

func (p *PlatformUsage) UnmarshalJSON(data []byte) error {
	type unmarshaler PlatformUsage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PlatformUsage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlatformUsage) MarshalJSON() ([]byte, error) {
	type embed PlatformUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
