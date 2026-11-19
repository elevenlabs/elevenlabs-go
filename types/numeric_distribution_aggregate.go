package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type NumericDistributionAggregate struct {
	Count           *int     `json:"count,omitempty" xml:"count,omitempty"`
	Sum             *float64 `json:"sum,omitempty" xml:"sum,omitempty"`
	Min             *float64 `json:"min,omitempty" xml:"min,omitempty"`
	Max             *float64 `json:"max,omitempty" xml:"max,omitempty"`
	extraProperties map[string]interface{}
}

func (n *NumericDistributionAggregate) GetCount() (value *int) {
	if n == nil {
		return
	}
	return n.Count
}

func (n *NumericDistributionAggregate) GetSum() (value *float64) {
	if n == nil {
		return
	}
	return n.Sum
}

func (n *NumericDistributionAggregate) GetMin() (value *float64) {
	if n == nil {
		return
	}
	return n.Min
}

func (n *NumericDistributionAggregate) GetMax() (value *float64) {
	if n == nil {
		return
	}
	return n.Max
}

func (n *NumericDistributionAggregate) SetCount(count *int) {
	n.Count = count
}

func (n *NumericDistributionAggregate) SetSum(sum *float64) {
	n.Sum = sum
}

func (n *NumericDistributionAggregate) SetMin(min *float64) {
	n.Min = min
}

func (n *NumericDistributionAggregate) SetMax(max *float64) {
	n.Max = max
}

func (n *NumericDistributionAggregate) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n NumericDistributionAggregate) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: NumericDistributionAggregate to string"
	}
	return string(jsonData)
}

func (n *NumericDistributionAggregate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, n); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	return nil
}

func (n *NumericDistributionAggregate) MarshalJSON() ([]byte, error) {
	type embed NumericDistributionAggregate
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	return json.Marshal(marshaler)
}
