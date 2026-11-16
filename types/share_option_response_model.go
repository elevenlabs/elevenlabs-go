package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ShareOptionResponseModel struct {
	// The name of the principal.
	Name string `json:"name" xml:"name" required:"true"`
	// The ID of the principal.
	Id string `json:"id" xml:"id" required:"true"`
	// The type of the principal: user, group, or service account (under 'key').
	Type            ShareOptionResponseModelType `json:"type" xml:"type" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ShareOptionResponseModel) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *ShareOptionResponseModel) GetId() (value string) {
	if s == nil {
		return
	}
	return s.Id
}

func (s *ShareOptionResponseModel) GetType() (value ShareOptionResponseModelType) {
	if s == nil {
		return
	}
	return s.Type
}

func (s *ShareOptionResponseModel) SetName(name string) {
	s.Name = name
}

func (s *ShareOptionResponseModel) SetId(id string) {
	s.Id = id
}

func (s *ShareOptionResponseModel) SetType(type_ ShareOptionResponseModelType) {
	s.Type = type_
}

func (s *ShareOptionResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ShareOptionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ShareOptionResponseModel to string"
	}
	return string(jsonData)
}

func (s *ShareOptionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ShareOptionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ShareOptionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ShareOptionResponseModel) MarshalJSON() ([]byte, error) {
	type embed ShareOptionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
