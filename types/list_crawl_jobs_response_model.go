package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ListCrawlJobsResponseModel struct {
	CrawlJobs       []*GetCrawlJobResponseModel `json:"crawl_jobs" xml:"crawl_jobs" required:"true"`
	NextCursor      *string                     `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListCrawlJobsResponseModel) GetCrawlJobs() (value []*GetCrawlJobResponseModel) {
	if l == nil {
		return
	}
	return l.CrawlJobs
}

func (l *ListCrawlJobsResponseModel) GetNextCursor() (value *string) {
	if l == nil {
		return
	}
	return l.NextCursor
}

func (l *ListCrawlJobsResponseModel) SetCrawlJobs(crawlJobs []*GetCrawlJobResponseModel) {
	l.CrawlJobs = crawlJobs
}

func (l *ListCrawlJobsResponseModel) SetNextCursor(nextCursor *string) {
	l.NextCursor = nextCursor
}

func (l *ListCrawlJobsResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListCrawlJobsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListCrawlJobsResponseModel to string"
	}
	return string(jsonData)
}

func (l *ListCrawlJobsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListCrawlJobsResponseModel) MarshalJSON() ([]byte, error) {
	type embed ListCrawlJobsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
