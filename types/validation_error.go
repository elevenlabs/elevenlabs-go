package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ValidationError struct {
	Loc             []ValidationErrorLocItem `json:"loc" xml:"loc" required:"true"`
	Msg             string                   `json:"msg" xml:"msg" required:"true"`
	Type            string                   `json:"type" xml:"type" required:"true"`
	extraProperties map[string]interface{}
}

func (v *ValidationError) GetLoc() (value []ValidationErrorLocItem) {
	if v == nil {
		return
	}
	return v.Loc
}

func (v *ValidationError) GetMsg() (value string) {
	if v == nil {
		return
	}
	return v.Msg
}

func (v *ValidationError) GetType() (value string) {
	if v == nil {
		return
	}
	return v.Type
}

func (v *ValidationError) SetLoc(loc []ValidationErrorLocItem) {
	v.Loc = loc
}

func (v *ValidationError) SetMsg(msg string) {
	v.Msg = msg
}

func (v *ValidationError) SetType(type_ string) {
	v.Type = type_
}

func (v *ValidationError) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v ValidationError) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: ValidationError to string"
	}
	return string(jsonData)
}

func (v *ValidationError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias ValidationError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = ValidationError(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *ValidationError) MarshalJSON() ([]byte, error) {
	type embed ValidationError
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
