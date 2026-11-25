package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The policy the campaign took when it was created, laid out like a submitted `policy`.
type CampaignPolicySnapshot struct {
	Timezone       string                 `json:"timezone" xml:"timezone" required:"true"`
	Retry          *RetryPolicy           `json:"retry" xml:"retry" required:"true"`
	CallingPeriods []*CallingPeriodPolicy `json:"calling_periods" xml:"calling_periods" required:"true"`
	DateOverrides  []*DateOverridePolicy  `json:"date_overrides,omitempty" xml:"date_overrides,omitempty"`
	DataCollection *DataCollectionPolicy  `json:"data_collection" xml:"data_collection" required:"true"`
	// Which `campaign__` dynamic variables follow-up calls get.
	FollowupContext *FollowupContextPolicy `json:"followup_context" xml:"followup_context" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CampaignPolicySnapshot) GetTimezone() (value string) {
	if c == nil {
		return
	}
	return c.Timezone
}

func (c *CampaignPolicySnapshot) GetRetry() (value *RetryPolicy) {
	if c == nil {
		return
	}
	return c.Retry
}

func (c *CampaignPolicySnapshot) GetCallingPeriods() (value []*CallingPeriodPolicy) {
	if c == nil {
		return
	}
	return c.CallingPeriods
}

func (c *CampaignPolicySnapshot) GetDateOverrides() (value []*DateOverridePolicy) {
	if c == nil {
		return
	}
	return c.DateOverrides
}

func (c *CampaignPolicySnapshot) GetDataCollection() (value *DataCollectionPolicy) {
	if c == nil {
		return
	}
	return c.DataCollection
}

func (c *CampaignPolicySnapshot) GetFollowupContext() (value *FollowupContextPolicy) {
	if c == nil {
		return
	}
	return c.FollowupContext
}

func (c *CampaignPolicySnapshot) SetTimezone(timezone string) {
	c.Timezone = timezone
}

func (c *CampaignPolicySnapshot) SetRetry(retry *RetryPolicy) {
	c.Retry = retry
}

func (c *CampaignPolicySnapshot) SetCallingPeriods(callingPeriods []*CallingPeriodPolicy) {
	c.CallingPeriods = callingPeriods
}

func (c *CampaignPolicySnapshot) SetDateOverrides(dateOverrides []*DateOverridePolicy) {
	c.DateOverrides = dateOverrides
}

func (c *CampaignPolicySnapshot) SetDataCollection(dataCollection *DataCollectionPolicy) {
	c.DataCollection = dataCollection
}

func (c *CampaignPolicySnapshot) SetFollowupContext(followupContext *FollowupContextPolicy) {
	c.FollowupContext = followupContext
}

func (c *CampaignPolicySnapshot) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CampaignPolicySnapshot) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CampaignPolicySnapshot to string"
	}
	return string(jsonData)
}

func (c *CampaignPolicySnapshot) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CampaignPolicySnapshot
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CampaignPolicySnapshot(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CampaignPolicySnapshot) MarshalJSON() ([]byte, error) {
	type embed CampaignPolicySnapshot
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
