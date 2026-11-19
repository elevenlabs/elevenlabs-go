package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PaginatedResultAgentMergeProposalResponse struct {
	Results         []*AgentMergeProposalResponse `json:"results" xml:"results" required:"true"`
	NextCursor      *string                       `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         *bool                         `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PaginatedResultAgentMergeProposalResponse) GetResults() (value []*AgentMergeProposalResponse) {
	if p == nil {
		return
	}
	return p.Results
}

func (p *PaginatedResultAgentMergeProposalResponse) GetNextCursor() (value *string) {
	if p == nil {
		return
	}
	return p.NextCursor
}

func (p *PaginatedResultAgentMergeProposalResponse) GetHasMore() (value *bool) {
	if p == nil {
		return
	}
	return p.HasMore
}

func (p *PaginatedResultAgentMergeProposalResponse) SetResults(results []*AgentMergeProposalResponse) {
	p.Results = results
}

func (p *PaginatedResultAgentMergeProposalResponse) SetNextCursor(nextCursor *string) {
	p.NextCursor = nextCursor
}

func (p *PaginatedResultAgentMergeProposalResponse) SetHasMore(hasMore *bool) {
	p.HasMore = hasMore
}

func (p *PaginatedResultAgentMergeProposalResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PaginatedResultAgentMergeProposalResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PaginatedResultAgentMergeProposalResponse to string"
	}
	return string(jsonData)
}

func (p *PaginatedResultAgentMergeProposalResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PaginatedResultAgentMergeProposalResponse) MarshalJSON() ([]byte, error) {
	type embed PaginatedResultAgentMergeProposalResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
