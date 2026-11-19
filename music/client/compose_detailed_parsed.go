// Hand-written (see .fernignore): music compose-detailed, parsed.
//
// The API answers POST /v1/music/detailed with a multipart/mixed body (a JSON metadata part, then
// the audio). The generated ComposeDetailed JSON-decodes that body into
// types.ComposeDetailedMusicResponse and always fails, so ComposeDetailedParsed sends the same
// request through the same caller (base URL, headers, credentials, retries, HTTP client and error
// handling) and splits the parts instead. A Go method can't be redefined with a different return
// type, so this is a separate method next to the generated one. Mirrors music.compose_detailed in
// the Python SDK.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// DefaultComposeDetailedFilename is the file name used when the audio part doesn't suggest one.
const DefaultComposeDetailedFilename = "generated_music.mp3"

// ComposeDetailedResponse is a composed song from ComposeDetailedParsed: the metadata the API
// returns with it, the audio file, and its name.
type ComposeDetailedResponse struct {
	// JSON is the JSON metadata part as the API sent it (snake_case keys): composition_plan,
	// song_metadata (title, description, genres, languages, is_explicit), words_timestamps (null
	// unless WithTimestamps is set) and, with WithWaveformVisual, waveform_visual.
	JSON map[string]any
	// Audio is the audio file's bytes, exactly as sent, in the requested output format.
	Audio []byte
	// Filename is the file name the API suggests for the audio, or DefaultComposeDetailedFilename.
	Filename string
	// SongID is the song-id response header, sent when StoreForInpainting is set; empty otherwise.
	SongID string
}

// ComposeDetailedDecodeError is returned by ComposeDetailedParsed when the API answers with a
// success status but the body isn't a complete multipart body with a JSON part and an audio part.
// Error statuses are returned as the generated ComposeDetailed returns them (a *core.APIError).
type ComposeDetailedDecodeError struct {
	StatusCode int
	Message    string
}

func (e *ComposeDetailedDecodeError) Error() string {
	return e.Message
}

// ComposeDetailedParsed composes a song from a prompt or a composition plan and returns its
// metadata and audio as separate fields.
//
// It takes the same arguments as the generated ComposeDetailed and sends the same request. The whole
// body is read before it returns, so a context deadline or the HTTP client's Timeout covers the
// download as well as the request.
func (c *Client) ComposeDetailedParsed(ctx context.Context, request *types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost, opts ...option.RequestOption) (*ComposeDetailedResponse, error) {
	// Built exactly as the generated RawClient.ComposeDetailed builds it.
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/music/detailed"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost{}
	}
	// Where the generated method decodes JSON, buffer the bytes (the caller copies the body into an
	// io.Writer before it returns) and split the parts.
	var body bytes.Buffer
	raw, err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &body,
		},
	)
	if err != nil {
		return nil, err
	}
	parsed, err := parseMultipart(body.Bytes(), boundaryFrom(raw.Header.Get("Content-Type")))
	if err != nil {
		return nil, &ComposeDetailedDecodeError{StatusCode: raw.StatusCode, Message: err.Error()}
	}
	return &ComposeDetailedResponse{
		JSON:     parsed.json,
		Audio:    parsed.audio,
		Filename: parsed.filename,
		SongID:   raw.Header.Get("song-id"),
	}, nil
}

type parsedMultipart struct {
	json     map[string]any
	audio    []byte
	filename string
}

type multipartPart struct {
	headers   map[string]string
	body      []byte
	separated bool
}

// boundaryFrom returns the boundary parameter of a Content-Type header, unquoted, or "".
func boundaryFrom(contentType string) string {
	return parameter(contentType, "boundary=")
}

// parseMultipart splits the multipart/mixed body into its JSON part and its audio part. It works on
// bytes throughout, so the audio is returned exactly as sent: the line break before each delimiter
// and the closing --boundary-- are stripped, never left on the end of the audio. Mirrors
// parseMultipart in the TypeScript SDK.
//
// boundary is the boundary parameter of the response's Content-Type; without it the delimiter is
// read from the body's first line. A body that ends before its closing --boundary-- (truncated) is
// an error.
func parseMultipart(body []byte, boundary string) (*parsedMultipart, error) {
	delimiter, err := delimiterFor(body, boundary)
	if err != nil {
		return nil, err
	}

	// The first delimiter starts the body or a line: anything before it is a preamble, which is
	// ignored.
	pos := 0
	if !bytes.HasPrefix(body, delimiter) {
		found := bytes.Index(body, append([]byte("\n"), delimiter...))
		if found < 0 {
			return nil, errors.New("Could not find the multipart boundary in the music response")
		}
		pos = found + 1
	}
	// The body uses one line ending throughout; read it from the end of the first delimiter line,
	// past any transport padding (spaces or tabs, which MIME allows after a delimiter).
	lineBreak := []byte("\n")
	if bytes.HasPrefix(body[skipPadding(body, pos+len(delimiter)):], []byte("\r\n")) {
		lineBreak = []byte("\r\n")
	}
	breakThenDelimiter := append(append([]byte{}, lineBreak...), delimiter...)

	var parts []multipartPart
	for {
		after := pos + len(delimiter)
		if bytes.HasPrefix(body[after:], []byte("--")) {
			break // Closing delimiter.
		}
		after = skipPadding(body, after)
		start := after
		if bytes.HasPrefix(body[after:], lineBreak) {
			start += len(lineBreak)
		}
		// Every part ends at a delimiter, and the body at --boundary--; the API always sends it. A
		// body that stops first was cut off, so its last part (the audio, usually) is incomplete.
		next := bytes.Index(body[start:], breakThenDelimiter)
		if next < 0 {
			return nil, errors.New("The music response ended before its closing multipart boundary; it is incomplete")
		}
		next += start
		parts = append(parts, splitPart(body[start:next:next], lineBreak))
		pos = next + len(lineBreak)
	}

	var jsonPart, audioPart *multipartPart
	for i := range parts {
		part := &parts[i]
		isJSON := strings.HasPrefix(asciiLower(part.headers["content-type"]), "application/json")
		if jsonPart == nil && isJSON {
			jsonPart = part
		} else if audioPart == nil && !isJSON {
			audioPart = part
		}
	}

	if jsonPart == nil {
		return nil, errors.New("Could not find the JSON part in the music response")
	}
	var value any
	if err := json.Unmarshal(trimASCII(jsonPart.body), &value); err != nil {
		return nil, fmt.Errorf("Could not parse the JSON part of the music response: %v", err)
	}
	object, ok := value.(map[string]any)
	if !ok {
		return nil, errors.New("The JSON part of the music response is not an object")
	}
	if audioPart == nil || !audioPart.separated {
		return nil, errors.New("Could not find the audio part in the music response")
	}

	return &parsedMultipart{
		json:     object,
		audio:    audioPart.body,
		filename: filename(audioPart.headers["content-disposition"]),
	}, nil
}

func delimiterFor(body []byte, boundary string) ([]byte, error) {
	if boundary != "" {
		return []byte("--" + boundary), nil
	}
	firstLine := body
	if i := bytes.IndexByte(body, '\n'); i >= 0 {
		firstLine = body[:i]
	}
	delimiter := trimASCII(firstLine)
	if !bytes.HasPrefix(delimiter, []byte("--")) || len(delimiter) < 3 {
		return nil, errors.New("Could not find the multipart boundary in the music response")
	}
	return delimiter, nil
}

// skipPadding returns the offset past the spaces and tabs at offset, which are MIME transport
// padding after a delimiter.
func skipPadding(body []byte, offset int) int {
	for offset < len(body) && (body[offset] == ' ' || body[offset] == '\t') {
		offset++
	}
	return offset
}

func splitPart(part []byte, lineBreak []byte) multipartPart {
	// A part with no headers starts with the blank line that ends them.
	if bytes.HasPrefix(part, lineBreak) {
		return multipartPart{headers: map[string]string{}, body: part[len(lineBreak):], separated: true}
	}
	separator := bytes.Index(part, append(append([]byte{}, lineBreak...), lineBreak...))
	if separator < 0 {
		return multipartPart{headers: partHeaders(part, lineBreak)}
	}
	return multipartPart{
		headers:   partHeaders(part[:separator], lineBreak),
		body:      part[separator+2*len(lineBreak):],
		separated: true,
	}
}

func partHeaders(block []byte, lineBreak []byte) map[string]string {
	headers := map[string]string{}
	for _, line := range strings.Split(string(block), string(lineBreak)) {
		name, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		headers[asciiLower(strings.TrimSpace(name))] = strings.TrimSpace(value)
	}
	return headers
}

func filename(disposition string) string {
	if name := parameter(disposition, "filename="); name != "" {
		return name
	}
	return DefaultComposeDetailedFilename
}

// parameter returns the value after key (lowercase, ending in "=") in a header value, unquoted, or
// "". A quoted value ends at the next quote; a bare one at a semicolon or whitespace.
func parameter(header, key string) string {
	// ASCII lowercasing keeps byte offsets, so an index into it indexes the original too.
	i := strings.Index(asciiLower(header), key)
	if i < 0 {
		return ""
	}
	rest := header[i+len(key):]
	if quoted, ok := strings.CutPrefix(rest, `"`); ok {
		value, _, _ := strings.Cut(quoted, `"`)
		return value
	}
	if end := strings.IndexFunc(rest, func(r rune) bool {
		return r == ';' || r == ' ' || r == '\t' || r == '\r' || r == '\n' || r == '\f' || r == '\v'
	}); end >= 0 {
		return rest[:end]
	}
	return rest
}

func asciiLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if 'A' <= c && c <= 'Z' {
			b[i] = c + ('a' - 'A')
		}
	}
	return string(b)
}

func trimASCII(b []byte) []byte {
	return bytes.Trim(b, " \t\n\r\f")
}
