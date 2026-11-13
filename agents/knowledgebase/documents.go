package knowledgebase

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
	big "math/big"
)

type BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost struct {
	// URL to a page of documentation that the agent will have access to in order to interact with users.
	Url string `json:"url" xml:"url" required:"true"`
	// A custom, human-readable name for the document.
	Name *string `json:"name,omitempty" xml:"name,omitempty" minLength:"1"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	// Whether to enable auto-sync for this URL document.
	EnableAutoSync *bool `json:"enable_auto_sync,omitempty" xml:"enable_auto_sync,omitempty"`
	// Whether to automatically remove the document if the URL becomes unavailable. Only applicable when auto-sync is enabled.
	AutoRemove *bool `json:"auto_remove,omitempty" xml:"auto_remove,omitempty"`
	// Minimum frequency (in days) at which the document is refreshed. The actual interval may be shorter, never longer. Defaults to 7, tightened to the parent folder's frequency if that is stricter. Only applicable when auto-sync is enabled.
	MinimumFrequencyDays *int `json:"minimum_frequency_days,omitempty" xml:"minimum_frequency_days,omitempty" min:"1" max:"180"`
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetUrl() (value string) {
	if b == nil {
		return
	}
	return b.Url
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetEnableAutoSync() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableAutoSync
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetAutoRemove() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoRemove
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetMinimumFrequencyDays() (value *int) {
	if b == nil {
		return
	}
	return b.MinimumFrequencyDays
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetUrl(url string) {
	b.Url = url
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetName(name *string) {
	b.Name = name
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetEnableAutoSync(enableAutoSync *bool) {
	b.EnableAutoSync = enableAutoSync
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetAutoRemove(autoRemove *bool) {
	b.AutoRemove = autoRemove
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) SetMinimumFrequencyDays(minimumFrequencyDays *int) {
	b.MinimumFrequencyDays = minimumFrequencyDays
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost struct {
	// Documentation that the agent will have access to in order to interact with users.
	File io.Reader `json:"file" xml:"file" required:"true"`
	// A custom, human-readable name for the document.
	Name *string `json:"name,omitempty" xml:"name,omitempty" minLength:"1"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) SetName(name *string) {
	b.Name = name
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateFileDocumentV1ConvaiKnowledgeBaseFilePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost struct {
	// Text content to be added to the knowledge base.
	Text string `json:"text" xml:"text" required:"true"`
	// A custom, human-readable name for the document.
	Name *string `json:"name,omitempty" xml:"name,omitempty" minLength:"1"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) SetText(text string) {
	b.Text = text
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) SetName(name *string) {
	b.Name = name
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost struct {
	// A custom, human-readable name for the document.
	Name string `json:"name" xml:"name" required:"true" minLength:"1"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	// Whether to enable auto-sync for this URL document.
	EnableAutoSync *bool `json:"enable_auto_sync,omitempty" xml:"enable_auto_sync,omitempty"`
	// Whether to automatically remove the document if the URL becomes unavailable. Only applicable when auto-sync is enabled.
	AutoRemove *bool `json:"auto_remove,omitempty" xml:"auto_remove,omitempty"`
	// Minimum frequency (in days) at which the underlying eligible documents are refreshed. The actual interval may be shorter, never longer. Defaults to 7, tightened to the parent folder's frequency if that is stricter. Only applicable when auto-sync is enabled.
	MinimumFrequencyDays *int `json:"minimum_frequency_days,omitempty" xml:"minimum_frequency_days,omitempty" min:"1" max:"180"`
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetEnableAutoSync() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableAutoSync
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetAutoRemove() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoRemove
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetMinimumFrequencyDays() (value *int) {
	if b == nil {
		return
	}
	return b.MinimumFrequencyDays
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) SetEnableAutoSync(enableAutoSync *bool) {
	b.EnableAutoSync = enableAutoSync
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) SetAutoRemove(autoRemove *bool) {
	b.AutoRemove = autoRemove
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) SetMinimumFrequencyDays(minimumFrequencyDays *int) {
	b.MinimumFrequencyDays = minimumFrequencyDays
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type GetDocumentsResponse struct {
	Type   string
	File   *types.GetKnowledgeBaseFileResponseModel
	Folder *types.GetKnowledgeBaseFolderResponseModel
	Text   *types.GetKnowledgeBaseTextResponseModel
	Url    *types.GetKnowledgeBaseUrlResponseModel
}

func (g *GetDocumentsResponse) GetFile() *types.GetKnowledgeBaseFileResponseModel {
	if g == nil {
		return nil
	}
	if g.File == nil {
		return nil
	}
	return g.File
}

func (g *GetDocumentsResponse) GetFolder() *types.GetKnowledgeBaseFolderResponseModel {
	if g == nil {
		return nil
	}
	if g.Folder == nil {
		return nil
	}
	return g.Folder
}

func (g *GetDocumentsResponse) GetText() *types.GetKnowledgeBaseTextResponseModel {
	if g == nil {
		return nil
	}
	if g.Text == nil {
		return nil
	}
	return g.Text
}

func (g *GetDocumentsResponse) GetUrl() *types.GetKnowledgeBaseUrlResponseModel {
	if g == nil {
		return nil
	}
	if g.Url == nil {
		return nil
	}
	return g.Url
}

func (g *GetDocumentsResponse) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetDocumentsResponse) UnmarshalJSON(data []byte) error {
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
		value := new(types.GetKnowledgeBaseFileResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.File = value
	case "folder":
		value := new(types.GetKnowledgeBaseFolderResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Folder = value
	case "text":
		value := new(types.GetKnowledgeBaseTextResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Text = value
	case "url":
		value := new(types.GetKnowledgeBaseUrlResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Url = value
	}
	return nil
}

func (g GetDocumentsResponse) MarshalJSON() ([]byte, error) {
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

type GetDocumentsResponseVisitor interface {
	VisitFile(*types.GetKnowledgeBaseFileResponseModel) error
	VisitFolder(*types.GetKnowledgeBaseFolderResponseModel) error
	VisitText(*types.GetKnowledgeBaseTextResponseModel) error
	VisitUrl(*types.GetKnowledgeBaseUrlResponseModel) error
}

func (g *GetDocumentsResponse) Accept(visitor GetDocumentsResponseVisitor) error {
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

func (g *GetDocumentsResponse) validate() error {
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

type UpdateDocumentsResponse struct {
	Type   string
	File   *types.GetKnowledgeBaseFileResponseModel
	Folder *types.GetKnowledgeBaseFolderResponseModel
	Text   *types.GetKnowledgeBaseTextResponseModel
	Url    *types.GetKnowledgeBaseUrlResponseModel
}

func (u *UpdateDocumentsResponse) GetFile() *types.GetKnowledgeBaseFileResponseModel {
	if u == nil {
		return nil
	}
	if u.File == nil {
		return nil
	}
	return u.File
}

func (u *UpdateDocumentsResponse) GetFolder() *types.GetKnowledgeBaseFolderResponseModel {
	if u == nil {
		return nil
	}
	if u.Folder == nil {
		return nil
	}
	return u.Folder
}

func (u *UpdateDocumentsResponse) GetText() *types.GetKnowledgeBaseTextResponseModel {
	if u == nil {
		return nil
	}
	if u.Text == nil {
		return nil
	}
	return u.Text
}

func (u *UpdateDocumentsResponse) GetUrl() *types.GetKnowledgeBaseUrlResponseModel {
	if u == nil {
		return nil
	}
	if u.Url == nil {
		return nil
	}
	return u.Url
}

func (u *UpdateDocumentsResponse) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UpdateDocumentsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "file":
		value := new(types.GetKnowledgeBaseFileResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.File = value
	case "folder":
		value := new(types.GetKnowledgeBaseFolderResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Folder = value
	case "text":
		value := new(types.GetKnowledgeBaseTextResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Text = value
	case "url":
		value := new(types.GetKnowledgeBaseUrlResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Url = value
	}
	return nil
}

func (u UpdateDocumentsResponse) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.File != nil {
		return internal.MarshalJSONWithDiscriminant(u.File, "type", "file")
	}
	if u.Folder != nil {
		return internal.MarshalJSONWithDiscriminant(u.Folder, "type", "folder")
	}
	if u.Text != nil {
		return internal.MarshalJSONWithDiscriminant(u.Text, "type", "text")
	}
	if u.Url != nil {
		return internal.MarshalJSONWithDiscriminant(u.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateDocumentsResponseVisitor interface {
	VisitFile(*types.GetKnowledgeBaseFileResponseModel) error
	VisitFolder(*types.GetKnowledgeBaseFolderResponseModel) error
	VisitText(*types.GetKnowledgeBaseTextResponseModel) error
	VisitUrl(*types.GetKnowledgeBaseUrlResponseModel) error
}

func (u *UpdateDocumentsResponse) Accept(visitor UpdateDocumentsResponseVisitor) error {
	if u.File != nil {
		return visitor.VisitFile(u.File)
	}
	if u.Folder != nil {
		return visitor.VisitFolder(u.Folder)
	}
	if u.Text != nil {
		return visitor.VisitText(u.Text)
	}
	if u.Url != nil {
		return visitor.VisitUrl(u.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateDocumentsResponse) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.File != nil {
		fields = append(fields, "file")
	}
	if u.Folder != nil {
		fields = append(fields, "folder")
	}
	if u.Text != nil {
		fields = append(fields, "text")
	}
	if u.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}

type BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch struct {
	// The id of a document from the knowledge base. This is returned on document addition.
	DocumentationId string `json:"-" url:"-"`
	// A custom, human-readable name for the document.
	Name *string `json:"name,omitempty" xml:"name,omitempty" minLength:"1"`
	// Updated content for the document. Only supported for text documents, URL documents with auto-sync disabled, and file documents that are not synced from an external source.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Whether the document or folder is refreshed automatically. Only applicable to URL documents, folders, and externally-synced files. Leave unset to keep the current value.
	EnableAutoSync *bool `json:"enable_auto_sync,omitempty" xml:"enable_auto_sync,omitempty"`
	// Whether to automatically remove the document once its source is no longer available. Only applicable when auto-sync is enabled. Leave unset to keep the current value.
	AutoRemove *bool `json:"auto_remove,omitempty" xml:"auto_remove,omitempty"`
	// Minimum frequency (in days) at which the document is refreshed. The actual interval may be shorter, never longer. Only applicable when auto-sync is enabled. Leave unset to keep the current value.
	MinimumFrequencyDays *int `json:"minimum_frequency_days,omitempty" xml:"minimum_frequency_days,omitempty" min:"1" max:"180"`
	// Folders only. When true, items already in the folder are only tightened to stay at least as strict as the folder, and keep any setting that is stricter than the folder's. When false, they are set to exactly the folder's auto-sync settings. Ignored for externally synced folders, whose contents always follow the folder's schedule.
	TightenOnly *bool `json:"tighten_only,omitempty" xml:"tighten_only,omitempty"`
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetDocumentationId() (value string) {
	if b == nil {
		return
	}
	return b.DocumentationId
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetContent() (value *string) {
	if b == nil {
		return
	}
	return b.Content
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetEnableAutoSync() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableAutoSync
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetAutoRemove() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoRemove
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetMinimumFrequencyDays() (value *int) {
	if b == nil {
		return
	}
	return b.MinimumFrequencyDays
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetTightenOnly() (value *bool) {
	if b == nil {
		return
	}
	return b.TightenOnly
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetDocumentationId(documentationId string) {
	b.DocumentationId = documentationId
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetName(name *string) {
	b.Name = name
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetContent(content *string) {
	b.Content = content
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetEnableAutoSync(enableAutoSync *bool) {
	b.EnableAutoSync = enableAutoSync
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetAutoRemove(autoRemove *bool) {
	b.AutoRemove = autoRemove
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetMinimumFrequencyDays(minimumFrequencyDays *int) {
	b.MinimumFrequencyDays = minimumFrequencyDays
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) SetTightenOnly(tightenOnly *bool) {
	b.TightenOnly = tightenOnly
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor *string `json:"-" url:"cursor,omitempty" explode:"true" serializationStyle:"form" queryParam:"cursor"`
	// How many documents to return at maximum. Can not exceed 100, defaults to 30.
	PageSize *int `json:"-" url:"page_size,omitempty" explode:"true" serializationStyle:"form" min:"1" max:"100" queryParam:"page_size"`
	// Type of dependent agents to return.
	DependentType *types.KnowledgeBaseDependentType `json:"-" url:"dependent_type,omitempty" explode:"true" serializationStyle:"form" queryParam:"dependent_type"`
	// The ids of documents or folders from the knowledge base.
	DocumentIds []string `json:"document_ids" xml:"document_ids" required:"true" minItems:"1" maxItems:"20" uniqueItems:"true"`
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) GetCursor() (value *string) {
	if b == nil {
		return
	}
	return b.Cursor
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) GetPageSize() (value *int) {
	if b == nil {
		return
	}
	return b.PageSize
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) GetDependentType() (value *types.KnowledgeBaseDependentType) {
	if b == nil {
		return
	}
	return b.DependentType
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) GetDocumentIds() (value []string) {
	if b == nil {
		return
	}
	return b.DocumentIds
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) SetCursor(cursor *string) {
	b.Cursor = cursor
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) SetPageSize(pageSize *int) {
	b.PageSize = pageSize
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) SetDependentType(dependentType *types.KnowledgeBaseDependentType) {
	b.DependentType = dependentType
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) SetDocumentIds(documentIds []string) {
	b.DocumentIds = documentIds
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost to string"
	}
	return string(jsonData)
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost) MarshalJSON() ([]byte, error) {
	type embed BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost struct {
	// The id of a document from the knowledge base. This is returned on document addition.
	DocumentId string `json:"-" url:"-"`
	// The folder to move the entities to. If not set, the entities will be moved to the root folder.
	MoveTo *string `json:"move_to,omitempty" xml:"move_to,omitempty"`
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) GetDocumentId() (value string) {
	if b == nil {
		return
	}
	return b.DocumentId
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) GetMoveTo() (value *string) {
	if b == nil {
		return
	}
	return b.MoveTo
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) SetDocumentId(documentId string) {
	b.DocumentId = documentId
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) SetMoveTo(moveTo *string) {
	b.MoveTo = moveTo
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost to string"
	}
	return string(jsonData)
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost) MarshalJSON() ([]byte, error) {
	type embed BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost struct {
	// The ids of documents or folders from the knowledge base.
	DocumentIds []string `json:"document_ids" xml:"document_ids" required:"true" minItems:"1" maxItems:"20" uniqueItems:"true"`
	// The folder to move the entities to. If not set, the entities will be moved to the root folder.
	MoveTo *string `json:"move_to,omitempty" xml:"move_to,omitempty"`
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) GetDocumentIds() (value []string) {
	if b == nil {
		return
	}
	return b.DocumentIds
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) GetMoveTo() (value *string) {
	if b == nil {
		return
	}
	return b.MoveTo
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) SetDocumentIds(documentIds []string) {
	b.DocumentIds = documentIds
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) SetMoveTo(moveTo *string) {
	b.MoveTo = moveTo
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost to string"
	}
	return string(jsonData)
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost) MarshalJSON() ([]byte, error) {
	type embed BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BulkDeleteDocumentsResponse struct {
	Status  string
	Failure *types.BatchFailureResponseModel
	Success *types.KnowledgeBaseBulkDeleteSuccessfulResponseModel
}

func (b *BulkDeleteDocumentsResponse) GetFailure() *types.BatchFailureResponseModel {
	if b == nil {
		return nil
	}
	if b.Failure == nil {
		return nil
	}
	return b.Failure
}

func (b *BulkDeleteDocumentsResponse) GetSuccess() *types.KnowledgeBaseBulkDeleteSuccessfulResponseModel {
	if b == nil {
		return nil
	}
	if b.Success == nil {
		return nil
	}
	return b.Success
}

func (b *BulkDeleteDocumentsResponse) GetStatus() string {
	if b == nil {
		return ""
	}
	return b.Status
}

func (b *BulkDeleteDocumentsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	b.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", b)
	}
	switch unmarshaler.Status {
	case "failure":
		value := new(types.BatchFailureResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Failure = value
	case "success":
		value := new(types.KnowledgeBaseBulkDeleteSuccessfulResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Success = value
	}
	return nil
}

func (b BulkDeleteDocumentsResponse) MarshalJSON() ([]byte, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	if b.Failure != nil {
		return internal.MarshalJSONWithDiscriminant(b.Failure, "status", "failure")
	}
	if b.Success != nil {
		return internal.MarshalJSONWithDiscriminant(b.Success, "status", "success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", b)
}

type BulkDeleteDocumentsResponseVisitor interface {
	VisitFailure(*types.BatchFailureResponseModel) error
	VisitSuccess(*types.KnowledgeBaseBulkDeleteSuccessfulResponseModel) error
}

func (b *BulkDeleteDocumentsResponse) Accept(visitor BulkDeleteDocumentsResponseVisitor) error {
	if b.Failure != nil {
		return visitor.VisitFailure(b.Failure)
	}
	if b.Success != nil {
		return visitor.VisitSuccess(b.Success)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", b)
}

func (b *BulkDeleteDocumentsResponse) validate() error {
	if b == nil {
		return fmt.Errorf("type %T is nil", b)
	}
	var fields []string
	if b.Failure != nil {
		fields = append(fields, "failure")
	}
	if b.Success != nil {
		fields = append(fields, "success")
	}
	if len(fields) == 0 {
		if b.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", b, b.Status)
		}
		return fmt.Errorf("type %T is empty", b)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", b, fields)
	}
	if b.Status != "" {
		field := fields[0]
		if b.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				b,
				b.Status,
				b,
			)
		}
	}
	return nil
}

type BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost struct {
	// The ids of documents or folders from the knowledge base.
	DocumentIds []string `json:"document_ids" xml:"document_ids" required:"true" minItems:"1" maxItems:"20" uniqueItems:"true"`
	// If set to true, documents or folders will be deleted regardless of whether they are used by any agents and will be removed from the dependent agents. For non-empty folders, this will also delete all child documents and folders.
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) GetDocumentIds() (value []string) {
	if b == nil {
		return
	}
	return b.DocumentIds
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) GetForce() (value *bool) {
	if b == nil {
		return
	}
	return b.Force
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) SetDocumentIds(documentIds []string) {
	b.DocumentIds = documentIds
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) SetForce(force *bool) {
	b.Force = force
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost to string"
	}
	return string(jsonData)
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) UnmarshalJSON(data []byte) error {
	type unmarshaler BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost(value)
	return nil
}

func (b *BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost) MarshalJSON() ([]byte, error) {
	type embed BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	getDocumentsRequestFieldAgentId = big.NewInt(1 << 0)
)

type GetDocumentsRequest struct {
	AgentId *string `json:"-" url:"agent_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetDocumentsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDocumentsRequest) SetAgentId(agentId *string) {
	g.AgentId = agentId
	g.require(getDocumentsRequestFieldAgentId)
}

var (
	deleteDocumentsRequestFieldForce = big.NewInt(1 << 0)
)

type DeleteDocumentsRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteDocumentsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteDocumentsRequest) SetForce(force *bool) {
	d.Force = force
	d.require(deleteDocumentsRequestFieldForce)
}

var (
	getAgentsDocumentsRequestFieldDependentType = big.NewInt(1 << 0)
	getAgentsDocumentsRequestFieldPageSize      = big.NewInt(1 << 1)
	getAgentsDocumentsRequestFieldCursor        = big.NewInt(1 << 2)
)

type GetAgentsDocumentsRequest struct {
	DependentType *types.KnowledgeBaseDependentType `json:"-" url:"dependent_type,omitempty"`
	PageSize      *int                              `json:"-" url:"page_size,omitempty"`
	Cursor        *string                           `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAgentsDocumentsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetDependentType sets the DependentType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAgentsDocumentsRequest) SetDependentType(dependentType *types.KnowledgeBaseDependentType) {
	g.DependentType = dependentType
	g.require(getAgentsDocumentsRequestFieldDependentType)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAgentsDocumentsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getAgentsDocumentsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAgentsDocumentsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getAgentsDocumentsRequestFieldCursor)
}
