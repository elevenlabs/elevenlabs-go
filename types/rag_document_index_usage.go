package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagDocumentIndexUsage struct {
	UsedBytes       int `json:"used_bytes" xml:"used_bytes" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagDocumentIndexUsage) GetUsedBytes() (value int) {
	if r == nil {
		return
	}
	return r.UsedBytes
}

func (r *RagDocumentIndexUsage) SetUsedBytes(usedBytes int) {
	r.UsedBytes = usedBytes
}

func (r *RagDocumentIndexUsage) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagDocumentIndexUsage) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagDocumentIndexUsage to string"
	}
	return string(jsonData)
}

func (r *RagDocumentIndexUsage) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagDocumentIndexUsage
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagDocumentIndexUsage(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagDocumentIndexUsage) MarshalJSON() ([]byte, error) {
	type embed RagDocumentIndexUsage
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
