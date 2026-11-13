package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type ListSimilarVoicesRequest struct {
	AudioFile io.Reader `json:"audio_file,omitempty" xml:"audio_file,omitempty"`
	// Threshold for voice similarity between provided sample and library voices. Values range from 0 to 2. The smaller the value the more similar voices will be returned.
	SimilarityThreshold *float64 `json:"similarity_threshold,omitempty" xml:"similarity_threshold,omitempty"`
	// Number of most similar voices to return. If similarity_threshold is provided, less than this number of voices may be returned. Values range from 1 to 100.
	TopK *int `json:"top_k,omitempty" xml:"top_k,omitempty"`
}

func (l *ListSimilarVoicesRequest) GetAudioFile() (value io.Reader) {
	if l == nil {
		return
	}
	return l.AudioFile
}

func (l *ListSimilarVoicesRequest) GetSimilarityThreshold() (value *float64) {
	if l == nil {
		return
	}
	return l.SimilarityThreshold
}

func (l *ListSimilarVoicesRequest) GetTopK() (value *int) {
	if l == nil {
		return
	}
	return l.TopK
}

func (l *ListSimilarVoicesRequest) SetAudioFile(audioFile io.Reader) {
	l.AudioFile = audioFile
}

func (l *ListSimilarVoicesRequest) SetSimilarityThreshold(similarityThreshold *float64) {
	l.SimilarityThreshold = similarityThreshold
}

func (l *ListSimilarVoicesRequest) SetTopK(topK *int) {
	l.TopK = topK
}

func (l *ListSimilarVoicesRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (l ListSimilarVoicesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListSimilarVoicesRequest to string"
	}
	return string(jsonData)
}

func (l *ListSimilarVoicesRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}

func (l *ListSimilarVoicesRequest) MarshalJSON() ([]byte, error) {
	type embed ListSimilarVoicesRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
