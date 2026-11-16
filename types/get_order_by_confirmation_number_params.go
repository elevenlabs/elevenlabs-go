package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetOrderByConfirmationNumberParams struct {
	extraProperties map[string]interface{}
}

func (g *GetOrderByConfirmationNumberParams) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetOrderByConfirmationNumberParams) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetOrderByConfirmationNumberParams to string"
	}
	return string(jsonData)
}

func (g *GetOrderByConfirmationNumberParams) UnmarshalJSON(data []byte) error {
	type unmarshaler GetOrderByConfirmationNumberParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetOrderByConfirmationNumberParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "smb_tool_type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetOrderByConfirmationNumberParams) MarshalJSON() ([]byte, error) {
	type embed GetOrderByConfirmationNumberParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
