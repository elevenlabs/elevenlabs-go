package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MergeProposalComment struct {
	Id              string `json:"id" xml:"id" required:"true"`
	UserId          string `json:"user_id" xml:"user_id" required:"true"`
	Body            string `json:"body" xml:"body" required:"true"`
	CreatedAt       int    `json:"created_at" xml:"created_at" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MergeProposalComment) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MergeProposalComment) GetUserId() (value string) {
	if m == nil {
		return
	}
	return m.UserId
}

func (m *MergeProposalComment) GetBody() (value string) {
	if m == nil {
		return
	}
	return m.Body
}

func (m *MergeProposalComment) GetCreatedAt() (value int) {
	if m == nil {
		return
	}
	return m.CreatedAt
}

func (m *MergeProposalComment) SetId(id string) {
	m.Id = id
}

func (m *MergeProposalComment) SetUserId(userId string) {
	m.UserId = userId
}

func (m *MergeProposalComment) SetBody(body string) {
	m.Body = body
}

func (m *MergeProposalComment) SetCreatedAt(createdAt int) {
	m.CreatedAt = createdAt
}

func (m *MergeProposalComment) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MergeProposalComment) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergeProposalComment to string"
	}
	return string(jsonData)
}

func (m *MergeProposalComment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MergeProposalComment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MergeProposalComment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MergeProposalComment) MarshalJSON() ([]byte, error) {
	type embed MergeProposalComment
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
