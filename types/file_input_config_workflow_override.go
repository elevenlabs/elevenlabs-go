package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FileInputConfigWorkflowOverride struct {
	// When enabled, users may attach images, PDFs, or text files in chat when the LLM supports multimodal input.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Number of most-recent files kept in memory during a conversation. Older files are summarized and their bytes freed.
	MaxFilesInMemory *int `json:"max_files_in_memory,omitempty" xml:"max_files_in_memory,omitempty"`
	// Total files a user can upload in one conversation. Uploads are billed per file. Use -1 for no limit, or a value >= max_files_in_memory.
	MaxFilesPerConversation *int `json:"max_files_per_conversation,omitempty" xml:"max_files_per_conversation,omitempty"`
	extraProperties         map[string]interface{}
}

func (f *FileInputConfigWorkflowOverride) GetEnabled() (value *bool) {
	if f == nil {
		return
	}
	return f.Enabled
}

func (f *FileInputConfigWorkflowOverride) GetMaxFilesInMemory() (value *int) {
	if f == nil {
		return
	}
	return f.MaxFilesInMemory
}

func (f *FileInputConfigWorkflowOverride) GetMaxFilesPerConversation() (value *int) {
	if f == nil {
		return
	}
	return f.MaxFilesPerConversation
}

func (f *FileInputConfigWorkflowOverride) SetEnabled(enabled *bool) {
	f.Enabled = enabled
}

func (f *FileInputConfigWorkflowOverride) SetMaxFilesInMemory(maxFilesInMemory *int) {
	f.MaxFilesInMemory = maxFilesInMemory
}

func (f *FileInputConfigWorkflowOverride) SetMaxFilesPerConversation(maxFilesPerConversation *int) {
	f.MaxFilesPerConversation = maxFilesPerConversation
}

func (f *FileInputConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FileInputConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FileInputConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (f *FileInputConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, f); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FileInputConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed FileInputConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
