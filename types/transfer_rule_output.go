package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferRuleOutput struct {
	// At most one transfer configuration per telephony provider. A provider without an entry cannot use this transfer rule.
	ProviderConfigs []*TransferRuleOutputProviderConfigsItem `json:"provider_configs" xml:"provider_configs" required:"true"`
	// When the agent should pick this transfer. Supports dynamic variables.
	Condition       string `json:"condition" xml:"condition" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferRuleOutput) GetProviderConfigs() (value []*TransferRuleOutputProviderConfigsItem) {
	if t == nil {
		return
	}
	return t.ProviderConfigs
}

func (t *TransferRuleOutput) GetCondition() (value string) {
	if t == nil {
		return
	}
	return t.Condition
}

func (t *TransferRuleOutput) SetProviderConfigs(providerConfigs []*TransferRuleOutputProviderConfigsItem) {
	t.ProviderConfigs = providerConfigs
}

func (t *TransferRuleOutput) SetCondition(condition string) {
	t.Condition = condition
}

func (t *TransferRuleOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferRuleOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferRuleOutput to string"
	}
	return string(jsonData)
}

func (t *TransferRuleOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferRuleOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferRuleOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferRuleOutput) MarshalJSON() ([]byte, error) {
	type embed TransferRuleOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
