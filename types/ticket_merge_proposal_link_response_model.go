package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TicketMergeProposalLinkResponseModel struct {
	MergeProposalId  string  `json:"merge_proposal_id" xml:"merge_proposal_id" required:"true"`
	LinkedByUserId   *string `json:"linked_by_user_id" xml:"linked_by_user_id" required:"true"`
	LinkedAtUnixSecs int     `json:"linked_at_unix_secs" xml:"linked_at_unix_secs" required:"true"`
	extraProperties  map[string]interface{}
}

func (t *TicketMergeProposalLinkResponseModel) GetMergeProposalId() (value string) {
	if t == nil {
		return
	}
	return t.MergeProposalId
}

func (t *TicketMergeProposalLinkResponseModel) GetLinkedByUserId() (value *string) {
	if t == nil {
		return
	}
	return t.LinkedByUserId
}

func (t *TicketMergeProposalLinkResponseModel) GetLinkedAtUnixSecs() (value int) {
	if t == nil {
		return
	}
	return t.LinkedAtUnixSecs
}

func (t *TicketMergeProposalLinkResponseModel) SetMergeProposalId(mergeProposalId string) {
	t.MergeProposalId = mergeProposalId
}

func (t *TicketMergeProposalLinkResponseModel) SetLinkedByUserId(linkedByUserId *string) {
	t.LinkedByUserId = linkedByUserId
}

func (t *TicketMergeProposalLinkResponseModel) SetLinkedAtUnixSecs(linkedAtUnixSecs int) {
	t.LinkedAtUnixSecs = linkedAtUnixSecs
}

func (t *TicketMergeProposalLinkResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TicketMergeProposalLinkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TicketMergeProposalLinkResponseModel to string"
	}
	return string(jsonData)
}

func (t *TicketMergeProposalLinkResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TicketMergeProposalLinkResponseModel) MarshalJSON() ([]byte, error) {
	type embed TicketMergeProposalLinkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
