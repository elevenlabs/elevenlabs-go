package conversations

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyUploadFileV1ConvaiConversationsConversationIdFilesPost struct {
	ConversationId string `json:"-" url:"-"`
	// Image, PDF, or text file to upload
	File io.Reader `json:"file" xml:"file" required:"true"`
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) GetConversationId() (value string) {
	if b == nil {
		return
	}
	return b.ConversationId
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) SetConversationId(conversationId string) {
	b.ConversationId = conversationId
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUploadFileV1ConvaiConversationsConversationIdFilesPost to string"
	}
	return string(jsonData)
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUploadFileV1ConvaiConversationsConversationIdFilesPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUploadFileV1ConvaiConversationsConversationIdFilesPost(tmp)
	return nil
}

func (b *BodyUploadFileV1ConvaiConversationsConversationIdFilesPost) MarshalJSON() ([]byte, error) {
	type embed BodyUploadFileV1ConvaiConversationsConversationIdFilesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
