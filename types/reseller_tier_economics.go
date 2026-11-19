package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A tier's prices and the limits shown on its tier card.
type ResellerTierEconomics struct {
	// What one month is sold for.
	MonthlyPriceUsd float64 `json:"monthly_price_usd" xml:"monthly_price_usd" required:"true"`
	// What a year up front is sold for. Null for the trial.
	AnnualPriceUsd *float64 `json:"annual_price_usd" xml:"annual_price_usd" required:"true"`
	// Phone minutes included each month. A web minute draws half a minute, a chat message a tenth.
	IncludedMinutesPerMonth float64 `json:"included_minutes_per_month" xml:"included_minutes_per_month" required:"true"`
	// Charged per minute past the included ones. Null for the trial.
	OveragePricePerMinuteUsd *float64 `json:"overage_price_per_minute_usd" xml:"overage_price_per_minute_usd" required:"true"`
	// Phone numbers a workspace on this tier can hold at once.
	MaxPhoneNumbers int `json:"max_phone_numbers" xml:"max_phone_numbers" required:"true"`
	// Calls a workspace on this tier can take at the same time.
	ConcurrentCalls int `json:"concurrent_calls" xml:"concurrent_calls" required:"true"`
	// Business locations a workspace on this tier can have active.
	MaxActiveLocations int `json:"max_active_locations" xml:"max_active_locations" required:"true"`
	// Receptionists a workspace on this tier can have active.
	MaxCustomerFacingAgents int `json:"max_customer_facing_agents" xml:"max_customer_facing_agents" required:"true"`
	// Team members a workspace on this tier can have, counting pending invites.
	MaxTeamMembers int `json:"max_team_members" xml:"max_team_members" required:"true"`
	// Messages the owner can send the Assistant each day.
	AssistantDailyMessages int `json:"assistant_daily_messages" xml:"assistant_daily_messages" required:"true"`
	extraProperties        map[string]interface{}
}

func (r *ResellerTierEconomics) GetMonthlyPriceUsd() (value float64) {
	if r == nil {
		return
	}
	return r.MonthlyPriceUsd
}

func (r *ResellerTierEconomics) GetAnnualPriceUsd() (value *float64) {
	if r == nil {
		return
	}
	return r.AnnualPriceUsd
}

func (r *ResellerTierEconomics) GetIncludedMinutesPerMonth() (value float64) {
	if r == nil {
		return
	}
	return r.IncludedMinutesPerMonth
}

func (r *ResellerTierEconomics) GetOveragePricePerMinuteUsd() (value *float64) {
	if r == nil {
		return
	}
	return r.OveragePricePerMinuteUsd
}

func (r *ResellerTierEconomics) GetMaxPhoneNumbers() (value int) {
	if r == nil {
		return
	}
	return r.MaxPhoneNumbers
}

func (r *ResellerTierEconomics) GetConcurrentCalls() (value int) {
	if r == nil {
		return
	}
	return r.ConcurrentCalls
}

func (r *ResellerTierEconomics) GetMaxActiveLocations() (value int) {
	if r == nil {
		return
	}
	return r.MaxActiveLocations
}

func (r *ResellerTierEconomics) GetMaxCustomerFacingAgents() (value int) {
	if r == nil {
		return
	}
	return r.MaxCustomerFacingAgents
}

func (r *ResellerTierEconomics) GetMaxTeamMembers() (value int) {
	if r == nil {
		return
	}
	return r.MaxTeamMembers
}

func (r *ResellerTierEconomics) GetAssistantDailyMessages() (value int) {
	if r == nil {
		return
	}
	return r.AssistantDailyMessages
}

func (r *ResellerTierEconomics) SetMonthlyPriceUsd(monthlyPriceUsd float64) {
	r.MonthlyPriceUsd = monthlyPriceUsd
}

func (r *ResellerTierEconomics) SetAnnualPriceUsd(annualPriceUsd *float64) {
	r.AnnualPriceUsd = annualPriceUsd
}

func (r *ResellerTierEconomics) SetIncludedMinutesPerMonth(includedMinutesPerMonth float64) {
	r.IncludedMinutesPerMonth = includedMinutesPerMonth
}

func (r *ResellerTierEconomics) SetOveragePricePerMinuteUsd(overagePricePerMinuteUsd *float64) {
	r.OveragePricePerMinuteUsd = overagePricePerMinuteUsd
}

func (r *ResellerTierEconomics) SetMaxPhoneNumbers(maxPhoneNumbers int) {
	r.MaxPhoneNumbers = maxPhoneNumbers
}

func (r *ResellerTierEconomics) SetConcurrentCalls(concurrentCalls int) {
	r.ConcurrentCalls = concurrentCalls
}

func (r *ResellerTierEconomics) SetMaxActiveLocations(maxActiveLocations int) {
	r.MaxActiveLocations = maxActiveLocations
}

func (r *ResellerTierEconomics) SetMaxCustomerFacingAgents(maxCustomerFacingAgents int) {
	r.MaxCustomerFacingAgents = maxCustomerFacingAgents
}

func (r *ResellerTierEconomics) SetMaxTeamMembers(maxTeamMembers int) {
	r.MaxTeamMembers = maxTeamMembers
}

func (r *ResellerTierEconomics) SetAssistantDailyMessages(assistantDailyMessages int) {
	r.AssistantDailyMessages = assistantDailyMessages
}

func (r *ResellerTierEconomics) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerTierEconomics) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerTierEconomics to string"
	}
	return string(jsonData)
}

func (r *ResellerTierEconomics) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerTierEconomics) MarshalJSON() ([]byte, error) {
	type embed ResellerTierEconomics
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
