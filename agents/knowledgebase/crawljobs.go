package knowledgebase

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost struct {
	// URL to a page of documentation that the agent will have access to in order to interact with users.
	Url string `json:"url" xml:"url" required:"true"`
	// Deprecated - this field is a no-op and will be removed in a future version.
	MaxDepth *int `json:"max_depth,omitempty" xml:"max_depth,omitempty"`
	// Maximum number of pages to crawl (1-10,000), defaults to 1000.
	MaxPages *int `json:"max_pages,omitempty" xml:"max_pages,omitempty" min:"1" max:"10000"`
	// If set, only URLs that match this pattern are included.
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty" maxLength:"2048"`
	// If set, URLs that match any of these patterns are skipped, even if they match `pattern`.
	ExcludePatterns []string `json:"exclude_patterns,omitempty" xml:"exclude_patterns,omitempty" maxItems:"50"`
	// List of URLs to crawl from sitemap (optional, overrides automatic URL discovery).
	SitemapUrls []string `json:"sitemap_urls,omitempty" xml:"sitemap_urls,omitempty" maxItems:"10000"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	// Whether to enable auto-sync for this URL document.
	EnableAutoSync *bool `json:"enable_auto_sync,omitempty" xml:"enable_auto_sync,omitempty"`
	// Whether to automatically remove the document if the URL becomes unavailable. Only applicable when auto-sync is enabled.
	AutoRemove *bool `json:"auto_remove,omitempty" xml:"auto_remove,omitempty"`
	// Automatically discover and add new pages during auto-sync: pages linked from already-crawled pages for website crawls, or new sitemap entries matching `pattern` and `exclude_patterns` for sitemap crawls. Requires enable_auto_sync=true.
	AutoDiscover *bool `json:"auto_discover,omitempty" xml:"auto_discover,omitempty"`
	// Minimum frequency (in days) at which the underlying eligible documents are refreshed. The actual interval may be shorter, never longer. Defaults to 7, tightened to the parent folder's frequency if that is stricter. Only applicable when auto-sync is enabled.
	MinimumFrequencyDays *int `json:"minimum_frequency_days,omitempty" xml:"minimum_frequency_days,omitempty" min:"1" max:"180"`
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetUrl() (value string) {
	if b == nil {
		return
	}
	return b.Url
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetMaxDepth() (value *int) {
	if b == nil {
		return
	}
	return b.MaxDepth
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetMaxPages() (value *int) {
	if b == nil {
		return
	}
	return b.MaxPages
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetPattern() (value *string) {
	if b == nil {
		return
	}
	return b.Pattern
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetExcludePatterns() (value []string) {
	if b == nil {
		return
	}
	return b.ExcludePatterns
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetSitemapUrls() (value []string) {
	if b == nil {
		return
	}
	return b.SitemapUrls
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetEnableAutoSync() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableAutoSync
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetAutoRemove() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoRemove
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetAutoDiscover() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoDiscover
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetMinimumFrequencyDays() (value *int) {
	if b == nil {
		return
	}
	return b.MinimumFrequencyDays
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetUrl(url string) {
	b.Url = url
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetMaxDepth(maxDepth *int) {
	b.MaxDepth = maxDepth
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetMaxPages(maxPages *int) {
	b.MaxPages = maxPages
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetPattern(pattern *string) {
	b.Pattern = pattern
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetExcludePatterns(excludePatterns []string) {
	b.ExcludePatterns = excludePatterns
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetSitemapUrls(sitemapUrls []string) {
	b.SitemapUrls = sitemapUrls
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetEnableAutoSync(enableAutoSync *bool) {
	b.EnableAutoSync = enableAutoSync
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetAutoRemove(autoRemove *bool) {
	b.AutoRemove = autoRemove
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetAutoDiscover(autoDiscover *bool) {
	b.AutoDiscover = autoDiscover
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) SetMinimumFrequencyDays(minimumFrequencyDays *int) {
	b.MinimumFrequencyDays = minimumFrequencyDays
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listCrawlJobsRequestFieldIncludeJobIds = big.NewInt(1 << 0)
	listCrawlJobsRequestFieldPageSize      = big.NewInt(1 << 1)
	listCrawlJobsRequestFieldCursor        = big.NewInt(1 << 2)
)

type ListCrawlJobsRequest struct {
	IncludeJobIds []*string `json:"-" url:"include_job_ids,omitempty"`
	PageSize      *int      `json:"-" url:"page_size,omitempty"`
	Cursor        *string   `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListCrawlJobsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetIncludeJobIds sets the IncludeJobIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCrawlJobsRequest) SetIncludeJobIds(includeJobIds []*string) {
	l.IncludeJobIds = includeJobIds
	l.require(listCrawlJobsRequestFieldIncludeJobIds)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCrawlJobsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listCrawlJobsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCrawlJobsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listCrawlJobsRequestFieldCursor)
}
