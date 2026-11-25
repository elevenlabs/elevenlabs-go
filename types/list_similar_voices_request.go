package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type ListSimilarVoicesRequest struct {
	// Filter voices by public owner ID
	OwnerId *string `json:"-" url:"owner_id,omitempty" explode:"true" serializationStyle:"form" queryParam:"owner_id"`
	// Include/exclude voices that are live moderated
	IncludeLiveModerated *bool `json:"-" url:"include_live_moderated,omitempty" explode:"true" serializationStyle:"form" queryParam:"include_live_moderated"`
	// Include/exclude voices with custom rates
	IncludeCustomRates *bool `json:"-" url:"include_custom_rates,omitempty" explode:"true" serializationStyle:"form" queryParam:"include_custom_rates"`
	// Filter voices with a minimum notice period of the given number of days.
	MinNoticePeriodDays *int `json:"-" url:"min_notice_period_days,omitempty" explode:"true" serializationStyle:"form" queryParam:"min_notice_period_days"`
	// Use-case used for filtering
	UseCases []*string `json:"-" url:"use_cases,omitempty" explode:"true" serializationStyle:"form" queryParam:"use_cases"`
	// Age used for filtering
	Age *string `json:"-" url:"age,omitempty" explode:"true" serializationStyle:"form" queryParam:"age"`
	// Gender used for filtering
	Gender *string `json:"-" url:"gender,omitempty" explode:"true" serializationStyle:"form" queryParam:"gender"`
	// Voice category used for filtering
	Category  *FindSimilarVoicesVoicesRequestCategory `json:"-" url:"category,omitempty" explode:"true" serializationStyle:"form" queryParam:"category"`
	AudioFile io.Reader                               `json:"audio_file,omitempty" xml:"audio_file,omitempty"`
	// Threshold for voice similarity between provided sample and library voices. Values range from 0 to 2. The smaller the value the more similar voices will be returned.
	SimilarityThreshold *float64 `json:"similarity_threshold,omitempty" xml:"similarity_threshold,omitempty"`
	// Number of most similar voices to return. If similarity_threshold is provided, less than this number of voices may be returned. Values range from 1 to 100.
	TopK *int `json:"top_k,omitempty" xml:"top_k,omitempty"`
}

func (l *ListSimilarVoicesRequest) GetOwnerId() (value *string) {
	if l == nil {
		return
	}
	return l.OwnerId
}

func (l *ListSimilarVoicesRequest) GetIncludeLiveModerated() (value *bool) {
	if l == nil {
		return
	}
	return l.IncludeLiveModerated
}

func (l *ListSimilarVoicesRequest) GetIncludeCustomRates() (value *bool) {
	if l == nil {
		return
	}
	return l.IncludeCustomRates
}

func (l *ListSimilarVoicesRequest) GetMinNoticePeriodDays() (value *int) {
	if l == nil {
		return
	}
	return l.MinNoticePeriodDays
}

func (l *ListSimilarVoicesRequest) GetUseCases() (value []*string) {
	if l == nil {
		return
	}
	return l.UseCases
}

func (l *ListSimilarVoicesRequest) GetAge() (value *string) {
	if l == nil {
		return
	}
	return l.Age
}

func (l *ListSimilarVoicesRequest) GetGender() (value *string) {
	if l == nil {
		return
	}
	return l.Gender
}

func (l *ListSimilarVoicesRequest) GetCategory() (value *FindSimilarVoicesVoicesRequestCategory) {
	if l == nil {
		return
	}
	return l.Category
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

func (l *ListSimilarVoicesRequest) SetOwnerId(ownerId *string) {
	l.OwnerId = ownerId
}

func (l *ListSimilarVoicesRequest) SetIncludeLiveModerated(includeLiveModerated *bool) {
	l.IncludeLiveModerated = includeLiveModerated
}

func (l *ListSimilarVoicesRequest) SetIncludeCustomRates(includeCustomRates *bool) {
	l.IncludeCustomRates = includeCustomRates
}

func (l *ListSimilarVoicesRequest) SetMinNoticePeriodDays(minNoticePeriodDays *int) {
	l.MinNoticePeriodDays = minNoticePeriodDays
}

func (l *ListSimilarVoicesRequest) SetUseCases(useCases []*string) {
	l.UseCases = useCases
}

func (l *ListSimilarVoicesRequest) SetAge(age *string) {
	l.Age = age
}

func (l *ListSimilarVoicesRequest) SetGender(gender *string) {
	l.Gender = gender
}

func (l *ListSimilarVoicesRequest) SetCategory(category *FindSimilarVoicesVoicesRequestCategory) {
	l.Category = category
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
