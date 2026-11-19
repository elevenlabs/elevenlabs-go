package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferRuleInput struct {
	// At most one transfer configuration per telephony provider. A provider without an entry cannot use this transfer rule.
	ProviderConfigs []*TransferRuleInputProviderConfigsItem `json:"provider_configs" xml:"provider_configs" required:"true"`
	// When the agent should pick this transfer. Supports dynamic variables.
	Condition       string `json:"condition" xml:"condition" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferRuleInput) GetProviderConfigs() (value []*TransferRuleInputProviderConfigsItem) {
	if t == nil {
		return
	}
	return t.ProviderConfigs
}

func (t *TransferRuleInput) GetCondition() (value string) {
	if t == nil {
		return
	}
	return t.Condition
}

func (t *TransferRuleInput) SetProviderConfigs(providerConfigs []*TransferRuleInputProviderConfigsItem) {
	t.ProviderConfigs = providerConfigs
}

func (t *TransferRuleInput) SetCondition(condition string) {
	t.Condition = condition
}

func (t *TransferRuleInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferRuleInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferRuleInput to string"
	}
	return string(jsonData)
}

func (t *TransferRuleInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferRuleInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferRuleInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferRuleInput) MarshalJSON() ([]byte, error) {
	type embed TransferRuleInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
