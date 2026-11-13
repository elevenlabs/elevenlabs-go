package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationVoiceRewardModel struct {
	VoiceId         string  `json:"voice_id" xml:"voice_id" required:"true"`
	RewardUsdCents  float64 `json:"reward_usd_cents" xml:"reward_usd_cents" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationVoiceRewardModel) GetVoiceId() (value string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ConversationVoiceRewardModel) GetRewardUsdCents() (value float64) {
	if c == nil {
		return
	}
	return c.RewardUsdCents
}

func (c *ConversationVoiceRewardModel) SetVoiceId(voiceId string) {
	c.VoiceId = voiceId
}

func (c *ConversationVoiceRewardModel) SetRewardUsdCents(rewardUsdCents float64) {
	c.RewardUsdCents = rewardUsdCents
}

func (c *ConversationVoiceRewardModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationVoiceRewardModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationVoiceRewardModel to string"
	}
	return string(jsonData)
}

func (c *ConversationVoiceRewardModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationVoiceRewardModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationVoiceRewardModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationVoiceRewardModel) MarshalJSON() ([]byte, error) {
	type embed ConversationVoiceRewardModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
