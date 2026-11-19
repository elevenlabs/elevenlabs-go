package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetCrawlJobResponseModel struct {
	Type    *CrawlType `json:"type,omitempty" xml:"type,omitempty"`
	SeedUrl string     `json:"seed_url" xml:"seed_url" required:"true"`
	Pattern *string    `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// URLs matching any of these patterns are skipped, even if they match `pattern`.
	ExcludePatterns []string `json:"exclude_patterns,omitempty" xml:"exclude_patterns,omitempty"`
	MaxPages        int      `json:"max_pages" xml:"max_pages" required:"true"`
	// Whether to automatically discover and enqueue additional pages found while crawling.
	AutoDiscover    *bool        `json:"auto_discover,omitempty" xml:"auto_discover,omitempty"`
	Status          *CrawlStatus `json:"status,omitempty" xml:"status,omitempty"`
	PagesIdentified *int         `json:"pages_identified,omitempty" xml:"pages_identified,omitempty"`
	PagesScraped    *int         `json:"pages_scraped,omitempty" xml:"pages_scraped,omitempty"`
	PagesSkipped    *int         `json:"pages_skipped,omitempty" xml:"pages_skipped,omitempty"`
	PagesFailed     *int         `json:"pages_failed,omitempty" xml:"pages_failed,omitempty"`
	RootFolderId    string       `json:"root_folder_id" xml:"root_folder_id" required:"true"`
	UpdatedAt       int          `json:"updated_at" xml:"updated_at" required:"true"`
	Id              string       `json:"id" xml:"id" required:"true"`
	CreatedAt       int          `json:"created_at" xml:"created_at" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetCrawlJobResponseModel) GetType() (value *CrawlType) {
	if g == nil {
		return
	}
	return g.Type
}

func (g *GetCrawlJobResponseModel) GetSeedUrl() (value string) {
	if g == nil {
		return
	}
	return g.SeedUrl
}

func (g *GetCrawlJobResponseModel) GetPattern() (value *string) {
	if g == nil {
		return
	}
	return g.Pattern
}

func (g *GetCrawlJobResponseModel) GetExcludePatterns() (value []string) {
	if g == nil {
		return
	}
	return g.ExcludePatterns
}

func (g *GetCrawlJobResponseModel) GetMaxPages() (value int) {
	if g == nil {
		return
	}
	return g.MaxPages
}

func (g *GetCrawlJobResponseModel) GetAutoDiscover() (value *bool) {
	if g == nil {
		return
	}
	return g.AutoDiscover
}

func (g *GetCrawlJobResponseModel) GetStatus() (value *CrawlStatus) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetCrawlJobResponseModel) GetPagesIdentified() (value *int) {
	if g == nil {
		return
	}
	return g.PagesIdentified
}

func (g *GetCrawlJobResponseModel) GetPagesScraped() (value *int) {
	if g == nil {
		return
	}
	return g.PagesScraped
}

func (g *GetCrawlJobResponseModel) GetPagesSkipped() (value *int) {
	if g == nil {
		return
	}
	return g.PagesSkipped
}

func (g *GetCrawlJobResponseModel) GetPagesFailed() (value *int) {
	if g == nil {
		return
	}
	return g.PagesFailed
}

func (g *GetCrawlJobResponseModel) GetRootFolderId() (value string) {
	if g == nil {
		return
	}
	return g.RootFolderId
}

func (g *GetCrawlJobResponseModel) GetUpdatedAt() (value int) {
	if g == nil {
		return
	}
	return g.UpdatedAt
}

func (g *GetCrawlJobResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetCrawlJobResponseModel) GetCreatedAt() (value int) {
	if g == nil {
		return
	}
	return g.CreatedAt
}

func (g *GetCrawlJobResponseModel) SetType(type_ *CrawlType) {
	g.Type = type_
}

func (g *GetCrawlJobResponseModel) SetSeedUrl(seedUrl string) {
	g.SeedUrl = seedUrl
}

func (g *GetCrawlJobResponseModel) SetPattern(pattern *string) {
	g.Pattern = pattern
}

func (g *GetCrawlJobResponseModel) SetExcludePatterns(excludePatterns []string) {
	g.ExcludePatterns = excludePatterns
}

func (g *GetCrawlJobResponseModel) SetMaxPages(maxPages int) {
	g.MaxPages = maxPages
}

func (g *GetCrawlJobResponseModel) SetAutoDiscover(autoDiscover *bool) {
	g.AutoDiscover = autoDiscover
}

func (g *GetCrawlJobResponseModel) SetStatus(status *CrawlStatus) {
	g.Status = status
}

func (g *GetCrawlJobResponseModel) SetPagesIdentified(pagesIdentified *int) {
	g.PagesIdentified = pagesIdentified
}

func (g *GetCrawlJobResponseModel) SetPagesScraped(pagesScraped *int) {
	g.PagesScraped = pagesScraped
}

func (g *GetCrawlJobResponseModel) SetPagesSkipped(pagesSkipped *int) {
	g.PagesSkipped = pagesSkipped
}

func (g *GetCrawlJobResponseModel) SetPagesFailed(pagesFailed *int) {
	g.PagesFailed = pagesFailed
}

func (g *GetCrawlJobResponseModel) SetRootFolderId(rootFolderId string) {
	g.RootFolderId = rootFolderId
}

func (g *GetCrawlJobResponseModel) SetUpdatedAt(updatedAt int) {
	g.UpdatedAt = updatedAt
}

func (g *GetCrawlJobResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetCrawlJobResponseModel) SetCreatedAt(createdAt int) {
	g.CreatedAt = createdAt
}

func (g *GetCrawlJobResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetCrawlJobResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetCrawlJobResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetCrawlJobResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetCrawlJobResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetCrawlJobResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
