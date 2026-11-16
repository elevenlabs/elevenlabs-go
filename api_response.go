package elevenlabs

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/clients/rest/httptransport"
	"net/http"
)

// ApiResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type ApiResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata ApiResponseMetadata
}

// ApiResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type ApiResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewApiResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewApiResponse[T any](resp *httptransport.Response[T]) *ApiResponse[T] {
	return &ApiResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: ApiResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *ApiResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r ApiResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ApiResponse to string"
	}
	return string(jsonData)
}
