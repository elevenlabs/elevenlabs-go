package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseSummaryBatchSuccessfulResponseModelData struct {
	Type   string
	File   *GetKnowledgeBaseSummaryFileResponseModel
	Folder *GetKnowledgeBaseSummaryFolderResponseModel
	Text   *GetKnowledgeBaseSummaryTextResponseModel
	Url    *GetKnowledgeBaseSummaryUrlResponseModel
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) GetFile() *GetKnowledgeBaseSummaryFileResponseModel {
	if k == nil {
		return nil
	}
	if k.File == nil {
		return nil
	}
	return k.File
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) GetFolder() *GetKnowledgeBaseSummaryFolderResponseModel {
	if k == nil {
		return nil
	}
	if k.Folder == nil {
		return nil
	}
	return k.Folder
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) GetText() *GetKnowledgeBaseSummaryTextResponseModel {
	if k == nil {
		return nil
	}
	if k.Text == nil {
		return nil
	}
	return k.Text
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) GetUrl() *GetKnowledgeBaseSummaryUrlResponseModel {
	if k == nil {
		return nil
	}
	if k.Url == nil {
		return nil
	}
	return k.Url
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) GetType() string {
	if k == nil {
		return ""
	}
	return k.Type
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	k.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", k)
	}
	switch unmarshaler.Type {
	case "file":
		value := new(GetKnowledgeBaseSummaryFileResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		k.File = value
	case "folder":
		value := new(GetKnowledgeBaseSummaryFolderResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		k.Folder = value
	case "text":
		value := new(GetKnowledgeBaseSummaryTextResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		k.Text = value
	case "url":
		value := new(GetKnowledgeBaseSummaryUrlResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		k.Url = value
	}
	return nil
}

func (k KnowledgeBaseSummaryBatchSuccessfulResponseModelData) MarshalJSON() ([]byte, error) {
	if err := k.validate(); err != nil {
		return nil, err
	}
	if k.File != nil {
		return internal.MarshalJSONWithDiscriminant(k.File, "type", "file")
	}
	if k.Folder != nil {
		return internal.MarshalJSONWithDiscriminant(k.Folder, "type", "folder")
	}
	if k.Text != nil {
		return internal.MarshalJSONWithDiscriminant(k.Text, "type", "text")
	}
	if k.Url != nil {
		return internal.MarshalJSONWithDiscriminant(k.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", k)
}

type KnowledgeBaseSummaryBatchSuccessfulResponseModelDataVisitor interface {
	VisitFile(*GetKnowledgeBaseSummaryFileResponseModel) error
	VisitFolder(*GetKnowledgeBaseSummaryFolderResponseModel) error
	VisitText(*GetKnowledgeBaseSummaryTextResponseModel) error
	VisitUrl(*GetKnowledgeBaseSummaryUrlResponseModel) error
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) Accept(visitor KnowledgeBaseSummaryBatchSuccessfulResponseModelDataVisitor) error {
	if k.File != nil {
		return visitor.VisitFile(k.File)
	}
	if k.Folder != nil {
		return visitor.VisitFolder(k.Folder)
	}
	if k.Text != nil {
		return visitor.VisitText(k.Text)
	}
	if k.Url != nil {
		return visitor.VisitUrl(k.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", k)
}

func (k *KnowledgeBaseSummaryBatchSuccessfulResponseModelData) validate() error {
	if k == nil {
		return fmt.Errorf("type %T is nil", k)
	}
	var fields []string
	if k.File != nil {
		fields = append(fields, "file")
	}
	if k.Folder != nil {
		fields = append(fields, "folder")
	}
	if k.Text != nil {
		fields = append(fields, "text")
	}
	if k.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if k.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", k, k.Type)
		}
		return fmt.Errorf("type %T is empty", k)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", k, fields)
	}
	if k.Type != "" {
		field := fields[0]
		if k.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				k,
				k.Type,
				k,
			)
		}
	}
	return nil
}
