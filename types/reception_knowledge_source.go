package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A file or web page the receptionists learn from.
type ReceptionKnowledgeSource struct {
	// Type of knowledge source.
	SourceType *KnowledgeSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// The page or site URL for 'url' sources.
	SourceUrl *string `json:"source_url,omitempty" xml:"source_url,omitempty"`
	// When the URL was last read, as a Unix timestamp in seconds.
	LastScrapedAtUnix *int `json:"last_scraped_at_unix,omitempty" xml:"last_scraped_at_unix,omitempty"`
	// The uploaded file's name for 'file' sources; empty for URLs.
	OriginalFilename *string `json:"original_filename,omitempty" xml:"original_filename,omitempty"`
	// The uploaded file's MIME type; empty for URLs.
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The uploaded file's size in bytes; 0 for URLs.
	FileSizeBytes *int `json:"file_size_bytes,omitempty" xml:"file_size_bytes,omitempty"`
	// Processing status for knowledge sources.
	//
	// - `processing`: the file or site is being read.
	// - `ready`: done; the results are available.
	// - `error`: it failed and can be retried.
	//
	// See `progress` for detail.
	Status *KnowledgeSourceStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Why processing failed, when status is 'error'.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// Machine-readable cause of the failure, when status is 'error'.
	FailureReason *IngestFailure `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// Whether the whole site was crawled rather than the single page at source_url.
	CrawlEntireSite *bool `json:"crawl_entire_site,omitempty" xml:"crawl_entire_site,omitempty"`
	// How many pages were read for 'url' sources.
	ScrapedPagesCount *int `json:"scraped_pages_count,omitempty" xml:"scraped_pages_count,omitempty"`
	// A short plain-text preview of the content.
	ContentPreview *string `json:"content_preview,omitempty" xml:"content_preview,omitempty"`
	// When processing started, as a Unix timestamp in seconds.
	ProcessingStartedAtUnix *int `json:"processing_started_at_unix,omitempty" xml:"processing_started_at_unix,omitempty"`
	// The knowledge source's ID.
	Id string `json:"id" xml:"id" required:"true"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	// Unix timestamp (seconds since epoch)
	UpdatedAtUnix   *int `json:"updated_at_unix,omitempty" xml:"updated_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionKnowledgeSource) GetSourceType() (value *KnowledgeSourceType) {
	if r == nil {
		return
	}
	return r.SourceType
}

func (r *ReceptionKnowledgeSource) GetSourceUrl() (value *string) {
	if r == nil {
		return
	}
	return r.SourceUrl
}

func (r *ReceptionKnowledgeSource) GetLastScrapedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.LastScrapedAtUnix
}

func (r *ReceptionKnowledgeSource) GetOriginalFilename() (value *string) {
	if r == nil {
		return
	}
	return r.OriginalFilename
}

func (r *ReceptionKnowledgeSource) GetContentType() (value *string) {
	if r == nil {
		return
	}
	return r.ContentType
}

func (r *ReceptionKnowledgeSource) GetFileSizeBytes() (value *int) {
	if r == nil {
		return
	}
	return r.FileSizeBytes
}

func (r *ReceptionKnowledgeSource) GetStatus() (value *KnowledgeSourceStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *ReceptionKnowledgeSource) GetErrorMessage() (value *string) {
	if r == nil {
		return
	}
	return r.ErrorMessage
}

func (r *ReceptionKnowledgeSource) GetFailureReason() (value *IngestFailure) {
	if r == nil {
		return
	}
	return r.FailureReason
}

func (r *ReceptionKnowledgeSource) GetCrawlEntireSite() (value *bool) {
	if r == nil {
		return
	}
	return r.CrawlEntireSite
}

func (r *ReceptionKnowledgeSource) GetScrapedPagesCount() (value *int) {
	if r == nil {
		return
	}
	return r.ScrapedPagesCount
}

func (r *ReceptionKnowledgeSource) GetContentPreview() (value *string) {
	if r == nil {
		return
	}
	return r.ContentPreview
}

func (r *ReceptionKnowledgeSource) GetProcessingStartedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.ProcessingStartedAtUnix
}

func (r *ReceptionKnowledgeSource) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionKnowledgeSource) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionKnowledgeSource) GetUpdatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.UpdatedAtUnix
}

func (r *ReceptionKnowledgeSource) SetSourceType(sourceType *KnowledgeSourceType) {
	r.SourceType = sourceType
}

func (r *ReceptionKnowledgeSource) SetSourceUrl(sourceUrl *string) {
	r.SourceUrl = sourceUrl
}

func (r *ReceptionKnowledgeSource) SetLastScrapedAtUnix(lastScrapedAtUnix *int) {
	r.LastScrapedAtUnix = lastScrapedAtUnix
}

func (r *ReceptionKnowledgeSource) SetOriginalFilename(originalFilename *string) {
	r.OriginalFilename = originalFilename
}

func (r *ReceptionKnowledgeSource) SetContentType(contentType *string) {
	r.ContentType = contentType
}

func (r *ReceptionKnowledgeSource) SetFileSizeBytes(fileSizeBytes *int) {
	r.FileSizeBytes = fileSizeBytes
}

func (r *ReceptionKnowledgeSource) SetStatus(status *KnowledgeSourceStatus) {
	r.Status = status
}

func (r *ReceptionKnowledgeSource) SetErrorMessage(errorMessage *string) {
	r.ErrorMessage = errorMessage
}

func (r *ReceptionKnowledgeSource) SetFailureReason(failureReason *IngestFailure) {
	r.FailureReason = failureReason
}

func (r *ReceptionKnowledgeSource) SetCrawlEntireSite(crawlEntireSite *bool) {
	r.CrawlEntireSite = crawlEntireSite
}

func (r *ReceptionKnowledgeSource) SetScrapedPagesCount(scrapedPagesCount *int) {
	r.ScrapedPagesCount = scrapedPagesCount
}

func (r *ReceptionKnowledgeSource) SetContentPreview(contentPreview *string) {
	r.ContentPreview = contentPreview
}

func (r *ReceptionKnowledgeSource) SetProcessingStartedAtUnix(processingStartedAtUnix *int) {
	r.ProcessingStartedAtUnix = processingStartedAtUnix
}

func (r *ReceptionKnowledgeSource) SetId(id string) {
	r.Id = id
}

func (r *ReceptionKnowledgeSource) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionKnowledgeSource) SetUpdatedAtUnix(updatedAtUnix *int) {
	r.UpdatedAtUnix = updatedAtUnix
}

func (r *ReceptionKnowledgeSource) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionKnowledgeSource) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionKnowledgeSource to string"
	}
	return string(jsonData)
}

func (r *ReceptionKnowledgeSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionKnowledgeSource) MarshalJSON() ([]byte, error) {
	type embed ReceptionKnowledgeSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
