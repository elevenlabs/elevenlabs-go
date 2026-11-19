package knowledgebase

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	"io"
)

type UpdateFileDocumentResponse struct {
	Type   string
	File   *types.GetKnowledgeBaseFileResponseModel
	Folder *types.GetKnowledgeBaseFolderResponseModel
	Text   *types.GetKnowledgeBaseTextResponseModel
	Url    *types.GetKnowledgeBaseUrlResponseModel
}

func (u *UpdateFileDocumentResponse) GetFile() *types.GetKnowledgeBaseFileResponseModel {
	if u == nil {
		return nil
	}
	if u.File == nil {
		return nil
	}
	return u.File
}

func (u *UpdateFileDocumentResponse) GetFolder() *types.GetKnowledgeBaseFolderResponseModel {
	if u == nil {
		return nil
	}
	if u.Folder == nil {
		return nil
	}
	return u.Folder
}

func (u *UpdateFileDocumentResponse) GetText() *types.GetKnowledgeBaseTextResponseModel {
	if u == nil {
		return nil
	}
	if u.Text == nil {
		return nil
	}
	return u.Text
}

func (u *UpdateFileDocumentResponse) GetUrl() *types.GetKnowledgeBaseUrlResponseModel {
	if u == nil {
		return nil
	}
	if u.Url == nil {
		return nil
	}
	return u.Url
}

func (u *UpdateFileDocumentResponse) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UpdateFileDocumentResponse) UnmarshalJSON(data []byte) error {
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

func (u UpdateFileDocumentResponse) MarshalJSON() ([]byte, error) {
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

type UpdateFileDocumentResponseVisitor interface {
	VisitFile(*types.GetKnowledgeBaseFileResponseModel) error
	VisitFolder(*types.GetKnowledgeBaseFolderResponseModel) error
	VisitText(*types.GetKnowledgeBaseTextResponseModel) error
	VisitUrl(*types.GetKnowledgeBaseUrlResponseModel) error
}

func (u *UpdateFileDocumentResponse) Accept(visitor UpdateFileDocumentResponseVisitor) error {
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

func (u *UpdateFileDocumentResponse) validate() error {
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

type BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch struct {
	// The id of a document from the knowledge base. This is returned on document addition.
	DocumentationId string `json:"-" url:"-"`
	// Documentation that the agent will have access to in order to interact with users.
	File io.Reader `json:"file" xml:"file" required:"true"`
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) GetDocumentationId() (value string) {
	if b == nil {
		return
	}
	return b.DocumentationId
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) SetDocumentationId(documentationId string) {
	b.DocumentationId = documentationId
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch(tmp)
	return nil
}

func (b *BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type RefreshDocumentResponse struct {
	Type   string
	File   *types.GetKnowledgeBaseFileResponseModel
	Folder *types.GetKnowledgeBaseFolderResponseModel
	Text   *types.GetKnowledgeBaseTextResponseModel
	Url    *types.GetKnowledgeBaseUrlResponseModel
}

func (r *RefreshDocumentResponse) GetFile() *types.GetKnowledgeBaseFileResponseModel {
	if r == nil {
		return nil
	}
	if r.File == nil {
		return nil
	}
	return r.File
}

func (r *RefreshDocumentResponse) GetFolder() *types.GetKnowledgeBaseFolderResponseModel {
	if r == nil {
		return nil
	}
	if r.Folder == nil {
		return nil
	}
	return r.Folder
}

func (r *RefreshDocumentResponse) GetText() *types.GetKnowledgeBaseTextResponseModel {
	if r == nil {
		return nil
	}
	if r.Text == nil {
		return nil
	}
	return r.Text
}

func (r *RefreshDocumentResponse) GetUrl() *types.GetKnowledgeBaseUrlResponseModel {
	if r == nil {
		return nil
	}
	if r.Url == nil {
		return nil
	}
	return r.Url
}

func (r *RefreshDocumentResponse) GetType() string {
	if r == nil {
		return ""
	}
	return r.Type
}

func (r *RefreshDocumentResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", r)
	}
	switch unmarshaler.Type {
	case "file":
		value := new(types.GetKnowledgeBaseFileResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.File = value
	case "folder":
		value := new(types.GetKnowledgeBaseFolderResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Folder = value
	case "text":
		value := new(types.GetKnowledgeBaseTextResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Text = value
	case "url":
		value := new(types.GetKnowledgeBaseUrlResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Url = value
	}
	return nil
}

func (r RefreshDocumentResponse) MarshalJSON() ([]byte, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if r.File != nil {
		return internal.MarshalJSONWithDiscriminant(r.File, "type", "file")
	}
	if r.Folder != nil {
		return internal.MarshalJSONWithDiscriminant(r.Folder, "type", "folder")
	}
	if r.Text != nil {
		return internal.MarshalJSONWithDiscriminant(r.Text, "type", "text")
	}
	if r.Url != nil {
		return internal.MarshalJSONWithDiscriminant(r.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", r)
}

type RefreshDocumentResponseVisitor interface {
	VisitFile(*types.GetKnowledgeBaseFileResponseModel) error
	VisitFolder(*types.GetKnowledgeBaseFolderResponseModel) error
	VisitText(*types.GetKnowledgeBaseTextResponseModel) error
	VisitUrl(*types.GetKnowledgeBaseUrlResponseModel) error
}

func (r *RefreshDocumentResponse) Accept(visitor RefreshDocumentResponseVisitor) error {
	if r.File != nil {
		return visitor.VisitFile(r.File)
	}
	if r.Folder != nil {
		return visitor.VisitFolder(r.Folder)
	}
	if r.Text != nil {
		return visitor.VisitText(r.Text)
	}
	if r.Url != nil {
		return visitor.VisitUrl(r.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", r)
}

func (r *RefreshDocumentResponse) validate() error {
	if r == nil {
		return fmt.Errorf("type %T is nil", r)
	}
	var fields []string
	if r.File != nil {
		fields = append(fields, "file")
	}
	if r.Folder != nil {
		fields = append(fields, "folder")
	}
	if r.Text != nil {
		fields = append(fields, "text")
	}
	if r.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if r.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", r, r.Type)
		}
		return fmt.Errorf("type %T is empty", r)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", r, fields)
	}
	if r.Type != "" {
		field := fields[0]
		if r.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				r,
				r.Type,
				r,
			)
		}
	}
	return nil
}

type RagIndexRequestModel struct {
	// The id of a document from the knowledge base. This is returned on document addition.
	DocumentationId string                   `json:"-" url:"-"`
	Model           types.EmbeddingModelEnum `json:"model" xml:"model" required:"true"`
}

func (r *RagIndexRequestModel) GetDocumentationId() (value string) {
	if r == nil {
		return
	}
	return r.DocumentationId
}

func (r *RagIndexRequestModel) GetModel() (value types.EmbeddingModelEnum) {
	if r == nil {
		return
	}
	return r.Model
}

func (r *RagIndexRequestModel) SetDocumentationId(documentationId string) {
	r.DocumentationId = documentationId
}

func (r *RagIndexRequestModel) SetModel(model types.EmbeddingModelEnum) {
	r.Model = model
}

func (r *RagIndexRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r RagIndexRequestModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagIndexRequestModel to string"
	}
	return string(jsonData)
}

func (r *RagIndexRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagIndexRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagIndexRequestModel(tmp)
	return nil
}

func (r *RagIndexRequestModel) MarshalJSON() ([]byte, error) {
	type embed RagIndexRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
