package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreatePronunciationDictionaryResponseModel struct {
	// The status of the create pronunciation dictionary request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePronunciationDictionaryResponseModel) GetStatus() (value string) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *CreatePronunciationDictionaryResponseModel) SetStatus(status string) {
	c.Status = status
}

func (c *CreatePronunciationDictionaryResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePronunciationDictionaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePronunciationDictionaryResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreatePronunciationDictionaryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePronunciationDictionaryResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePronunciationDictionaryResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePronunciationDictionaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreatePronunciationDictionaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
