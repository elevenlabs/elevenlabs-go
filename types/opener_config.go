package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenerConfig struct {
	Llm             Llm `json:"llm" xml:"llm" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OpenerConfig) GetLlm() (value Llm) {
	if o == nil {
		return
	}
	return o.Llm
}

func (o *OpenerConfig) SetLlm(llm Llm) {
	o.Llm = llm
}

func (o *OpenerConfig) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenerConfig) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenerConfig to string"
	}
	return string(jsonData)
}

func (o *OpenerConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OpenerConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OpenerConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenerConfig) MarshalJSON() ([]byte, error) {
	type embed OpenerConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
