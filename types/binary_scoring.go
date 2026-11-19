package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BinaryScoring struct {
	extraProperties map[string]interface{}
}

func (b *BinaryScoring) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BinaryScoring) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BinaryScoring to string"
	}
	return string(jsonData)
}

func (b *BinaryScoring) UnmarshalJSON(data []byte) error {
	type unmarshaler BinaryScoring
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BinaryScoring(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "mode")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BinaryScoring) MarshalJSON() ([]byte, error) {
	type embed BinaryScoring
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
