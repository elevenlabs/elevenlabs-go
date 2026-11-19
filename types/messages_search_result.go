package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// transcript_index: index of the message in the conversation transcript
// chunk_text: text of the transcript; transcript messages if very long could have several chunks.
// chunk_highlights: chunk_text split into matched/unmatched segments for highlighting.
// Only populated for keyword/text search, not semantic search.
// score: similarity score of the message to the search query
type MessagesSearchResult struct {
	ConversationId                string                    `json:"conversation_id" xml:"conversation_id" required:"true"`
	AgentId                       string                    `json:"agent_id" xml:"agent_id" required:"true"`
	AgentName                     *string                   `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
	TranscriptIndex               int                       `json:"transcript_index" xml:"transcript_index" required:"true"`
	ChunkText                     string                    `json:"chunk_text" xml:"chunk_text" required:"true"`
	ChunkHighlights               []*SearchHighlightSegment `json:"chunk_highlights,omitempty" xml:"chunk_highlights,omitempty"`
	Score                         float64                   `json:"score" xml:"score" required:"true"`
	ConversationStartTimeUnixSecs int                       `json:"conversation_start_time_unix_secs" xml:"conversation_start_time_unix_secs" required:"true"`
	extraProperties               map[string]interface{}
}

func (m *MessagesSearchResult) GetConversationId() (value string) {
	if m == nil {
		return
	}
	return m.ConversationId
}

func (m *MessagesSearchResult) GetAgentId() (value string) {
	if m == nil {
		return
	}
	return m.AgentId
}

func (m *MessagesSearchResult) GetAgentName() (value *string) {
	if m == nil {
		return
	}
	return m.AgentName
}

func (m *MessagesSearchResult) GetTranscriptIndex() (value int) {
	if m == nil {
		return
	}
	return m.TranscriptIndex
}

func (m *MessagesSearchResult) GetChunkText() (value string) {
	if m == nil {
		return
	}
	return m.ChunkText
}

func (m *MessagesSearchResult) GetChunkHighlights() (value []*SearchHighlightSegment) {
	if m == nil {
		return
	}
	return m.ChunkHighlights
}

func (m *MessagesSearchResult) GetScore() (value float64) {
	if m == nil {
		return
	}
	return m.Score
}

func (m *MessagesSearchResult) GetConversationStartTimeUnixSecs() (value int) {
	if m == nil {
		return
	}
	return m.ConversationStartTimeUnixSecs
}

func (m *MessagesSearchResult) SetConversationId(conversationId string) {
	m.ConversationId = conversationId
}

func (m *MessagesSearchResult) SetAgentId(agentId string) {
	m.AgentId = agentId
}

func (m *MessagesSearchResult) SetAgentName(agentName *string) {
	m.AgentName = agentName
}

func (m *MessagesSearchResult) SetTranscriptIndex(transcriptIndex int) {
	m.TranscriptIndex = transcriptIndex
}

func (m *MessagesSearchResult) SetChunkText(chunkText string) {
	m.ChunkText = chunkText
}

func (m *MessagesSearchResult) SetChunkHighlights(chunkHighlights []*SearchHighlightSegment) {
	m.ChunkHighlights = chunkHighlights
}

func (m *MessagesSearchResult) SetScore(score float64) {
	m.Score = score
}

func (m *MessagesSearchResult) SetConversationStartTimeUnixSecs(conversationStartTimeUnixSecs int) {
	m.ConversationStartTimeUnixSecs = conversationStartTimeUnixSecs
}

func (m *MessagesSearchResult) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MessagesSearchResult) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MessagesSearchResult to string"
	}
	return string(jsonData)
}

func (m *MessagesSearchResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MessagesSearchResult) MarshalJSON() ([]byte, error) {
	type embed MessagesSearchResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
