package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Error struct {
	Message         string  `json:"message" xml:"message" required:"true"`
	Error           string  `json:"error" xml:"error" required:"true"`
	Code            int     `json:"code" xml:"code" required:"true"`
	Param           *string `json:"param,omitempty" xml:"param,omitempty"`
	extraProperties map[string]interface{}
}

func (e *Error) GetMessage() (value string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *Error) GetError() (value string) {
	if e == nil {
		return
	}
	return e.Error
}

func (e *Error) GetCode() (value int) {
	if e == nil {
		return
	}
	return e.Code
}

func (e *Error) GetParam() (value *string) {
	if e == nil {
		return
	}
	return e.Param
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetError(error string) {
	e.Error = error
}

func (e *Error) SetCode(code int) {
	e.Code = code
}

func (e *Error) SetParam(param *string) {
	e.Param = param
}

func (e *Error) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e Error) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Error to string"
	}
	return string(jsonData)
}

func (e *Error) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias Error
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = Error(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *Error) MarshalJSON() ([]byte, error) {
	type embed Error
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
