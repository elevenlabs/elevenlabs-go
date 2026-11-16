package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryRulesResponseModel struct {
	// The ID of the pronunciation dictionary.
	Id string `json:"id" xml:"id" required:"true"`
	// The version ID of the pronunciation dictionary.
	VersionId string `json:"version_id" xml:"version_id" required:"true"`
	// The number of rules in the version of the pronunciation dictionary.
	VersionRulesNum int `json:"version_rules_num" xml:"version_rules_num" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryRulesResponseModel) GetId() (value string) {
	if p == nil {
		return
	}
	return p.Id
}

func (p *PronunciationDictionaryRulesResponseModel) GetVersionId() (value string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *PronunciationDictionaryRulesResponseModel) GetVersionRulesNum() (value int) {
	if p == nil {
		return
	}
	return p.VersionRulesNum
}

func (p *PronunciationDictionaryRulesResponseModel) SetId(id string) {
	p.Id = id
}

func (p *PronunciationDictionaryRulesResponseModel) SetVersionId(versionId string) {
	p.VersionId = versionId
}

func (p *PronunciationDictionaryRulesResponseModel) SetVersionRulesNum(versionRulesNum int) {
	p.VersionRulesNum = versionRulesNum
}

func (p *PronunciationDictionaryRulesResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryRulesResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryRulesResponseModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryRulesResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionaryRulesResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionaryRulesResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PronunciationDictionaryRulesResponseModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryRulesResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
