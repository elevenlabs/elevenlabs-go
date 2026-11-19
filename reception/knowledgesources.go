package reception

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
	big "math/big"
)

type BodySmbUploadKnowledgeSource struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// The document the receptionist should learn from.
	File io.Reader `json:"file" xml:"file" required:"true"`
}

func (b *BodySmbUploadKnowledgeSource) GetXiWorkspaceId() (value string) {
	if b == nil {
		return
	}
	return b.XiWorkspaceId
}

func (b *BodySmbUploadKnowledgeSource) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodySmbUploadKnowledgeSource) SetXiWorkspaceId(xiWorkspaceId string) {
	b.XiWorkspaceId = xiWorkspaceId
}

func (b *BodySmbUploadKnowledgeSource) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodySmbUploadKnowledgeSource) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySmbUploadKnowledgeSource) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySmbUploadKnowledgeSource to string"
	}
	return string(jsonData)
}

func (b *BodySmbUploadKnowledgeSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodySmbUploadKnowledgeSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodySmbUploadKnowledgeSource(tmp)
	return nil
}

func (b *BodySmbUploadKnowledgeSource) MarshalJSON() ([]byte, error) {
	type embed BodySmbUploadKnowledgeSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// A web page or site to learn from.
type AddUrlRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// The HTTPS page or site to read.
	Url string `json:"url" xml:"url" required:"true" maxLength:"2048" minLength:"1"`
	// Crawl the whole site instead of only the page at url.
	CrawlEntireSite *bool `json:"crawl_entire_site,omitempty" xml:"crawl_entire_site,omitempty"`
}

func (a *AddUrlRequest) GetXiWorkspaceId() (value string) {
	if a == nil {
		return
	}
	return a.XiWorkspaceId
}

func (a *AddUrlRequest) GetUrl() (value string) {
	if a == nil {
		return
	}
	return a.Url
}

func (a *AddUrlRequest) GetCrawlEntireSite() (value *bool) {
	if a == nil {
		return
	}
	return a.CrawlEntireSite
}

func (a *AddUrlRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	a.XiWorkspaceId = xiWorkspaceId
}

func (a *AddUrlRequest) SetUrl(url string) {
	a.Url = url
}

func (a *AddUrlRequest) SetCrawlEntireSite(crawlEntireSite *bool) {
	a.CrawlEntireSite = crawlEntireSite
}

func (a *AddUrlRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddUrlRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddUrlRequest to string"
	}
	return string(jsonData)
}

func (a *AddUrlRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddUrlRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddUrlRequest(tmp)
	return nil
}

func (a *AddUrlRequest) MarshalJSON() ([]byte, error) {
	type embed AddUrlRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

var (
	listKnowledgeSourcesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListKnowledgeSourcesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListKnowledgeSourcesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeSourcesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listKnowledgeSourcesRequestFieldXiWorkspaceId)
}

var (
	getKnowledgeSourcesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type GetKnowledgeSourcesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetKnowledgeSourcesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetKnowledgeSourcesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	g.XiWorkspaceId = xiWorkspaceId
	g.require(getKnowledgeSourcesRequestFieldXiWorkspaceId)
}

var (
	deleteKnowledgeSourcesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type DeleteKnowledgeSourcesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteKnowledgeSourcesRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteKnowledgeSourcesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	d.XiWorkspaceId = xiWorkspaceId
	d.require(deleteKnowledgeSourcesRequestFieldXiWorkspaceId)
}
