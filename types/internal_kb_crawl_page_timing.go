package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Per-document timing data for one crawled page.
type InternalKbCrawlPageTiming struct {
	Url             string `json:"url" xml:"url" required:"true"`
	DocId           string `json:"doc_id" xml:"doc_id" required:"true"`
	ReadMs          *int   `json:"read_ms,omitempty" xml:"read_ms,omitempty"`
	IndexMs         *int   `json:"index_ms,omitempty" xml:"index_ms,omitempty"`
	HtmlLen         *int   `json:"html_len,omitempty" xml:"html_len,omitempty"`
	Skipped         *bool  `json:"skipped,omitempty" xml:"skipped,omitempty"`
	Error           *bool  `json:"error,omitempty" xml:"error,omitempty"`
	extraProperties map[string]interface{}
}

func (i *InternalKbCrawlPageTiming) GetUrl() (value string) {
	if i == nil {
		return
	}
	return i.Url
}

func (i *InternalKbCrawlPageTiming) GetDocId() (value string) {
	if i == nil {
		return
	}
	return i.DocId
}

func (i *InternalKbCrawlPageTiming) GetReadMs() (value *int) {
	if i == nil {
		return
	}
	return i.ReadMs
}

func (i *InternalKbCrawlPageTiming) GetIndexMs() (value *int) {
	if i == nil {
		return
	}
	return i.IndexMs
}

func (i *InternalKbCrawlPageTiming) GetHtmlLen() (value *int) {
	if i == nil {
		return
	}
	return i.HtmlLen
}

func (i *InternalKbCrawlPageTiming) GetSkipped() (value *bool) {
	if i == nil {
		return
	}
	return i.Skipped
}

func (i *InternalKbCrawlPageTiming) GetError() (value *bool) {
	if i == nil {
		return
	}
	return i.Error
}

func (i *InternalKbCrawlPageTiming) SetUrl(url string) {
	i.Url = url
}

func (i *InternalKbCrawlPageTiming) SetDocId(docId string) {
	i.DocId = docId
}

func (i *InternalKbCrawlPageTiming) SetReadMs(readMs *int) {
	i.ReadMs = readMs
}

func (i *InternalKbCrawlPageTiming) SetIndexMs(indexMs *int) {
	i.IndexMs = indexMs
}

func (i *InternalKbCrawlPageTiming) SetHtmlLen(htmlLen *int) {
	i.HtmlLen = htmlLen
}

func (i *InternalKbCrawlPageTiming) SetSkipped(skipped *bool) {
	i.Skipped = skipped
}

func (i *InternalKbCrawlPageTiming) SetError(error *bool) {
	i.Error = error
}

func (i *InternalKbCrawlPageTiming) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InternalKbCrawlPageTiming) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InternalKbCrawlPageTiming to string"
	}
	return string(jsonData)
}

func (i *InternalKbCrawlPageTiming) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InternalKbCrawlPageTiming) MarshalJSON() ([]byte, error) {
	type embed InternalKbCrawlPageTiming
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
