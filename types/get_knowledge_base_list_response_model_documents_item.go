package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetKnowledgeBaseListResponseModelDocumentsItem struct {
	Type   string
	File   *GetKnowledgeBaseSummaryFileResponseModel
	Folder *GetKnowledgeBaseSummaryFolderResponseModel
	Text   *GetKnowledgeBaseSummaryTextResponseModel
	Url    *GetKnowledgeBaseSummaryUrlResponseModel
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) GetFile() *GetKnowledgeBaseSummaryFileResponseModel {
	if g == nil {
		return nil
	}
	if g.File == nil {
		return nil
	}
	return g.File
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) GetFolder() *GetKnowledgeBaseSummaryFolderResponseModel {
	if g == nil {
		return nil
	}
	if g.Folder == nil {
		return nil
	}
	return g.Folder
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) GetText() *GetKnowledgeBaseSummaryTextResponseModel {
	if g == nil {
		return nil
	}
	if g.Text == nil {
		return nil
	}
	return g.Text
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) GetUrl() *GetKnowledgeBaseSummaryUrlResponseModel {
	if g == nil {
		return nil
	}
	if g.Url == nil {
		return nil
	}
	return g.Url
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", g)
	}
	switch unmarshaler.Type {
	case "file":
		value := new(GetKnowledgeBaseSummaryFileResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.File = value
	case "folder":
		value := new(GetKnowledgeBaseSummaryFolderResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Folder = value
	case "text":
		value := new(GetKnowledgeBaseSummaryTextResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Text = value
	case "url":
		value := new(GetKnowledgeBaseSummaryUrlResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Url = value
	}
	return nil
}

func (g GetKnowledgeBaseListResponseModelDocumentsItem) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.File != nil {
		return internal.MarshalJSONWithDiscriminant(g.File, "type", "file")
	}
	if g.Folder != nil {
		return internal.MarshalJSONWithDiscriminant(g.Folder, "type", "folder")
	}
	if g.Text != nil {
		return internal.MarshalJSONWithDiscriminant(g.Text, "type", "text")
	}
	if g.Url != nil {
		return internal.MarshalJSONWithDiscriminant(g.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetKnowledgeBaseListResponseModelDocumentsItemVisitor interface {
	VisitFile(*GetKnowledgeBaseSummaryFileResponseModel) error
	VisitFolder(*GetKnowledgeBaseSummaryFolderResponseModel) error
	VisitText(*GetKnowledgeBaseSummaryTextResponseModel) error
	VisitUrl(*GetKnowledgeBaseSummaryUrlResponseModel) error
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) Accept(visitor GetKnowledgeBaseListResponseModelDocumentsItemVisitor) error {
	if g.File != nil {
		return visitor.VisitFile(g.File)
	}
	if g.Folder != nil {
		return visitor.VisitFolder(g.Folder)
	}
	if g.Text != nil {
		return visitor.VisitText(g.Text)
	}
	if g.Url != nil {
		return visitor.VisitUrl(g.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetKnowledgeBaseListResponseModelDocumentsItem) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.File != nil {
		fields = append(fields, "file")
	}
	if g.Folder != nil {
		fields = append(fields, "folder")
	}
	if g.Text != nil {
		fields = append(fields, "text")
	}
	if g.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if g.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Type)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Type != "" {
		field := fields[0]
		if g.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Type,
				g,
			)
		}
	}
	return nil
}
