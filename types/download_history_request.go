package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DownloadHistoryRequest struct {
	// A list of history items to download, you can get IDs of history items and other metadata using the GET https://api.elevenlabs.io/v1/history endpoint.
	HistoryItemIds []string `json:"history_item_ids" xml:"history_item_ids" required:"true"`
	// Output format to transcode the audio file, can be wav or default.
	OutputFormat *string `json:"output_format,omitempty" xml:"output_format,omitempty"`
}

func (d *DownloadHistoryRequest) GetHistoryItemIds() (value []string) {
	if d == nil {
		return
	}
	return d.HistoryItemIds
}

func (d *DownloadHistoryRequest) GetOutputFormat() (value *string) {
	if d == nil {
		return
	}
	return d.OutputFormat
}

func (d *DownloadHistoryRequest) SetHistoryItemIds(historyItemIds []string) {
	d.HistoryItemIds = historyItemIds
}

func (d *DownloadHistoryRequest) SetOutputFormat(outputFormat *string) {
	d.OutputFormat = outputFormat
}

func (d *DownloadHistoryRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DownloadHistoryRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DownloadHistoryRequest to string"
	}
	return string(jsonData)
}

func (d *DownloadHistoryRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}

func (d *DownloadHistoryRequest) MarshalJSON() ([]byte, error) {
	type embed DownloadHistoryRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
