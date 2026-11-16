package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A time-bound promotion on a service/variant or on add-ons.
type ReceptionServicePromotion struct {
	// Stable ID of this promotion.
	Id   *string     `json:"id,omitempty" xml:"id,omitempty"`
	Name StrippedStr `json:"name" xml:"name" required:"true"`
	// What a promotion discounts: the service price or add-on prices.
	Target *PromotionTarget `json:"target,omitempty" xml:"target,omitempty"`
	// When set (target=service), the promotion applies only to this variant.
	VariantId *string `json:"variant_id,omitempty" xml:"variant_id,omitempty"`
	// When set (target=add_on), the promotion applies only to this add-on.
	AddOnId *string `json:"add_on_id,omitempty" xml:"add_on_id,omitempty"`
	// The final price when this promotion is applied to a matching line item.
	PromotionalPrice float64 `json:"promotional_price" xml:"promotional_price" required:"true" min:"0"`
	// Days when the promotion is active (0=Monday, 6=Sunday); empty means every day.
	DaysOfWeek []int `json:"days_of_week,omitempty" xml:"days_of_week,omitempty"`
	// First day the promotion applies, inclusive.
	StartDate string `json:"start_date" xml:"start_date" required:"true"`
	// Last day the promotion applies, inclusive; on or after `start_date`.
	EndDate string `json:"end_date" xml:"end_date" required:"true"`
	// Whether the promotion applies; inactive promotions are ignored.
	IsActive        *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionServicePromotion) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionServicePromotion) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionServicePromotion) GetTarget() (value *PromotionTarget) {
	if r == nil {
		return
	}
	return r.Target
}

func (r *ReceptionServicePromotion) GetVariantId() (value *string) {
	if r == nil {
		return
	}
	return r.VariantId
}

func (r *ReceptionServicePromotion) GetAddOnId() (value *string) {
	if r == nil {
		return
	}
	return r.AddOnId
}

func (r *ReceptionServicePromotion) GetPromotionalPrice() (value float64) {
	if r == nil {
		return
	}
	return r.PromotionalPrice
}

func (r *ReceptionServicePromotion) GetDaysOfWeek() (value []int) {
	if r == nil {
		return
	}
	return r.DaysOfWeek
}

func (r *ReceptionServicePromotion) GetStartDate() (value string) {
	if r == nil {
		return
	}
	return r.StartDate
}

func (r *ReceptionServicePromotion) GetEndDate() (value string) {
	if r == nil {
		return
	}
	return r.EndDate
}

func (r *ReceptionServicePromotion) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionServicePromotion) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionServicePromotion) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionServicePromotion) SetTarget(target *PromotionTarget) {
	r.Target = target
}

func (r *ReceptionServicePromotion) SetVariantId(variantId *string) {
	r.VariantId = variantId
}

func (r *ReceptionServicePromotion) SetAddOnId(addOnId *string) {
	r.AddOnId = addOnId
}

func (r *ReceptionServicePromotion) SetPromotionalPrice(promotionalPrice float64) {
	r.PromotionalPrice = promotionalPrice
}

func (r *ReceptionServicePromotion) SetDaysOfWeek(daysOfWeek []int) {
	r.DaysOfWeek = daysOfWeek
}

func (r *ReceptionServicePromotion) SetStartDate(startDate string) {
	r.StartDate = startDate
}

func (r *ReceptionServicePromotion) SetEndDate(endDate string) {
	r.EndDate = endDate
}

func (r *ReceptionServicePromotion) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionServicePromotion) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionServicePromotion) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionServicePromotion to string"
	}
	return string(jsonData)
}

func (r *ReceptionServicePromotion) UnmarshalJSON(data []byte) error {
	data = internal.ParseDateFields(data, "start_date", "end_date")
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

func (r *ReceptionServicePromotion) MarshalJSON() ([]byte, error) {
	type embed ReceptionServicePromotion
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return internal.FormatDateFields("start_date", "end_date")(json.Marshal(marshaler))
}
