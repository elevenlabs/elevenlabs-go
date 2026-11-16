package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FileInputConfig struct {
	// When enabled, users may attach images, PDFs, or text files in chat when the LLM supports multimodal input.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Number of most-recent files kept in memory during a conversation. Older files are summarized and their bytes freed.
	MaxFilesInMemory *int `json:"max_files_in_memory,omitempty" xml:"max_files_in_memory,omitempty" min:"1" max:"30"`
	// Total files a user can upload in one conversation. Uploads are billed per file. Use -1 for no limit, or a value >= max_files_in_memory.
	MaxFilesPerConversation *int `json:"max_files_per_conversation,omitempty" xml:"max_files_per_conversation,omitempty"`
	extraProperties         map[string]interface{}
}

func (f *FileInputConfig) GetEnabled() (value *bool) {
	if f == nil {
		return
	}
	return f.Enabled
}

func (f *FileInputConfig) GetMaxFilesInMemory() (value *int) {
	if f == nil {
		return
	}
	return f.MaxFilesInMemory
}

func (f *FileInputConfig) GetMaxFilesPerConversation() (value *int) {
	if f == nil {
		return
	}
	return f.MaxFilesPerConversation
}

func (f *FileInputConfig) SetEnabled(enabled *bool) {
	f.Enabled = enabled
}

func (f *FileInputConfig) SetMaxFilesInMemory(maxFilesInMemory *int) {
	f.MaxFilesInMemory = maxFilesInMemory
}

func (f *FileInputConfig) SetMaxFilesPerConversation(maxFilesPerConversation *int) {
	f.MaxFilesPerConversation = maxFilesPerConversation
}

func (f *FileInputConfig) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FileInputConfig) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FileInputConfig to string"
	}
	return string(jsonData)
}

func (f *FileInputConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler FileInputConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FileInputConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FileInputConfig) MarshalJSON() ([]byte, error) {
	type embed FileInputConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
