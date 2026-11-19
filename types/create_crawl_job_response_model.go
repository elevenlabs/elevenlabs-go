package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateCrawlJobResponseModel struct {
	Id           string    `json:"id" xml:"id" required:"true"`
	Type         CrawlType `json:"type" xml:"type" required:"true"`
	RootFolderId string    `json:"root_folder_id" xml:"root_folder_id" required:"true"`
	Status       string    `json:"status" xml:"status" required:"true"`
	CreatedAt    int       `json:"created_at" xml:"created_at" required:"true"`
	// The folder path segments leading to the root folder, from root to parent folder.
	FolderPath      []*KnowledgeBaseFolderPathSegmentSummaryResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CreateCrawlJobResponseModel) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *CreateCrawlJobResponseModel) GetType() (value CrawlType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CreateCrawlJobResponseModel) GetRootFolderId() (value string) {
	if c == nil {
		return
	}
	return c.RootFolderId
}

func (c *CreateCrawlJobResponseModel) GetStatus() (value string) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *CreateCrawlJobResponseModel) GetCreatedAt() (value int) {
	if c == nil {
		return
	}
	return c.CreatedAt
}

func (c *CreateCrawlJobResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if c == nil {
		return
	}
	return c.FolderPath
}

func (c *CreateCrawlJobResponseModel) SetId(id string) {
	c.Id = id
}

func (c *CreateCrawlJobResponseModel) SetType(type_ CrawlType) {
	c.Type = type_
}

func (c *CreateCrawlJobResponseModel) SetRootFolderId(rootFolderId string) {
	c.RootFolderId = rootFolderId
}

func (c *CreateCrawlJobResponseModel) SetStatus(status string) {
	c.Status = status
}

func (c *CreateCrawlJobResponseModel) SetCreatedAt(createdAt int) {
	c.CreatedAt = createdAt
}

func (c *CreateCrawlJobResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	c.FolderPath = folderPath
}

func (c *CreateCrawlJobResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateCrawlJobResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateCrawlJobResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreateCrawlJobResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateCrawlJobResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateCrawlJobResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateCrawlJobResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreateCrawlJobResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
