package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UnitTestToolCallParameter struct {
	Eval            *UnitTestToolCallParameterEval `json:"eval" xml:"eval" required:"true"`
	Path            string                         `json:"path" xml:"path" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UnitTestToolCallParameter) GetEval() (value *UnitTestToolCallParameterEval) {
	if u == nil {
		return
	}
	return u.Eval
}

func (u *UnitTestToolCallParameter) GetPath() (value string) {
	if u == nil {
		return
	}
	return u.Path
}

func (u *UnitTestToolCallParameter) SetEval(eval *UnitTestToolCallParameterEval) {
	u.Eval = eval
}

func (u *UnitTestToolCallParameter) SetPath(path string) {
	u.Path = path
}

func (u *UnitTestToolCallParameter) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UnitTestToolCallParameter) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnitTestToolCallParameter to string"
	}
	return string(jsonData)
}

func (u *UnitTestToolCallParameter) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UnitTestToolCallParameter
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UnitTestToolCallParameter(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UnitTestToolCallParameter) MarshalJSON() ([]byte, error) {
	type embed UnitTestToolCallParameter
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
