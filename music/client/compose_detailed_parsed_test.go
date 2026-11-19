// Hand-written (see .fernignore): tests for ComposeDetailedParsed and its multipart parser. Mirrors
// tests/music_detailed.rs in the Rust SDK and tests/test_music_multipart.py in the Python SDK.

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

const testBoundary = "--boundary123"

func testMetadata() map[string]any {
	return map[string]any{
		"composition_plan": map[string]any{
			"positive_global_styles": []any{"pop"},
			"negative_global_styles": []any{},
			"sections":               []any{},
		},
		"song_metadata": map[string]any{
			"title":       "Test",
			"description": "A test song",
			"genres":      []any{"pop"},
			"languages":   []any{"en"},
			"is_explicit": false,
		},
		"words_timestamps": nil,
	}
}

// testAudio holds every byte value, so an off-by-one at either end of the audio part shows up.
func testAudio() []byte {
	audio := make([]byte, 1024)
	for i := range audio {
		audio[i] = byte(i)
	}
	return audio
}

func metadataJSON(t *testing.T) string {
	t.Helper()
	encoded, err := json.Marshal(testMetadata())
	if err != nil {
		t.Fatal(err)
	}
	return string(encoded)
}

func multipartBody(t *testing.T, lineBreak, jsonPart string, disposition bool) []byte {
	t.Helper()
	if jsonPart == "" {
		jsonPart = metadataJSON(t)
	}
	lb := lineBreak
	head := testBoundary + lb + "Content-Type: application/json" + lb + lb + jsonPart + lb +
		testBoundary + lb + "Content-Type: audio/mpeg" + lb
	if disposition {
		head += `Content-Disposition: attachment; filename="test_song.mp3"` + lb
	}
	head += lb
	body := []byte(head)
	body = append(body, testAudio()...)
	body = append(body, []byte(lb+testBoundary+"--"+lb)...)
	return body
}

type recordedRequest struct {
	method string
	path   string
	query  string
	header http.Header
	body   []byte
}

// serve starts a server that answers every request with status, the multipart Content-Type and
// body, and records the requests it receives.
func serve(t *testing.T, status int, header http.Header, body []byte) (*Client, *[]recordedRequest) {
	t.Helper()
	var requests []recordedRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received, _ := io.ReadAll(r.Body)
		requests = append(requests, recordedRequest{
			method: r.Method,
			path:   r.URL.Path,
			query:  r.URL.RawQuery,
			header: r.Header.Clone(),
			body:   received,
		})
		w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary123")
		for name, values := range header {
			w.Header()[name] = values
		}
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}))
	t.Cleanup(server.Close)
	apiKey := "test_api_key"
	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithApiKey(&apiKey),
		option.WithoutRetries(),
	))
	return client, &requests
}

func testRequest() *types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost {
	length := 10000
	prompt := "Epic orchestral music"
	return &types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost{
		OutputFormat:  types.ComposeDetailedMusicRequestOutputFormatMp344100128.Ptr(),
		Prompt:        &prompt,
		MusicLengthMs: &length,
	}
}

func compose(t *testing.T, body []byte) (*ComposeDetailedResponse, error) {
	t.Helper()
	client, _ := serve(t, http.StatusOK, nil, body)
	return client.ComposeDetailedParsed(context.Background(), testRequest())
}

func mustCompose(t *testing.T, body []byte) *ComposeDetailedResponse {
	t.Helper()
	song, err := compose(t, body)
	if err != nil {
		t.Fatalf("ComposeDetailedParsed: %v", err)
	}
	return song
}

func decodeErrorMessage(t *testing.T, body []byte) string {
	t.Helper()
	_, err := compose(t, body)
	var decodeErr *ComposeDetailedDecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("expected *ComposeDetailedDecodeError, got %T: %v", err, err)
	}
	if decodeErr.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want 200", decodeErr.StatusCode)
	}
	return decodeErr.Message
}

func assertSong(t *testing.T, song *ComposeDetailedResponse) {
	t.Helper()
	if !reflect.DeepEqual(song.JSON, testMetadata()) {
		t.Errorf("JSON = %#v, want %#v", song.JSON, testMetadata())
	}
	if !bytes.Equal(song.Audio, testAudio()) {
		t.Errorf("audio must be byte-for-byte, with the closing delimiter stripped: got %d bytes", len(song.Audio))
	}
}

func TestComposeDetailedParsedCRLF(t *testing.T) {
	song := mustCompose(t, multipartBody(t, "\r\n", "", true))
	assertSong(t, song)
	if song.Filename != "test_song.mp3" {
		t.Errorf("Filename = %q, want test_song.mp3", song.Filename)
	}
}

func TestComposeDetailedParsedLF(t *testing.T) {
	assertSong(t, mustCompose(t, multipartBody(t, "\n", "", true)))
}

func TestComposeDetailedParsedSendsTheGeneratedRequest(t *testing.T) {
	client, requests := serve(t, http.StatusOK, nil, multipartBody(t, "\r\n", "", true))
	if _, err := client.ComposeDetailedParsed(context.Background(), testRequest()); err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 1 {
		t.Fatalf("got %d requests, want 1", len(*requests))
	}
	got := (*requests)[0]
	if got.method != http.MethodPost || got.path != "/v1/music/detailed" {
		t.Errorf("request = %s %s, want POST /v1/music/detailed", got.method, got.path)
	}
	if got.query != "output_format=mp3_44100_128" {
		t.Errorf("query = %q, want output_format=mp3_44100_128", got.query)
	}
	if key := got.header.Get("xi-api-key"); key != "test_api_key" {
		t.Errorf("xi-api-key = %q", key)
	}
	if ct := got.header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q", ct)
	}
	var sent map[string]any
	if err := json.Unmarshal(got.body, &sent); err != nil {
		t.Fatalf("request body %q: %v", got.body, err)
	}
	want := map[string]any{"prompt": "Epic orchestral music", "music_length_ms": float64(10000)}
	if !reflect.DeepEqual(sent, want) {
		t.Errorf("request body = %v, want %v", sent, want)
	}
}

func TestComposeDetailedParsedReturnsTheSongIDHeader(t *testing.T) {
	client, _ := serve(t, http.StatusOK, http.Header{"Song-Id": {"song_abc"}}, multipartBody(t, "\r\n", "", true))
	song, err := client.ComposeDetailedParsed(context.Background(), testRequest())
	if err != nil {
		t.Fatal(err)
	}
	if song.SongID != "song_abc" {
		t.Errorf("SongID = %q, want song_abc", song.SongID)
	}
}

func TestComposeDetailedParsedSongIDIsEmptyWithoutTheHeader(t *testing.T) {
	if song := mustCompose(t, multipartBody(t, "\r\n", "", true)); song.SongID != "" {
		t.Errorf("SongID = %q, want empty", song.SongID)
	}
}

func TestComposeDetailedParsedDefaultsTheFilename(t *testing.T) {
	song := mustCompose(t, multipartBody(t, "\r\n", "", false))
	if song.Filename != DefaultComposeDetailedFilename {
		t.Errorf("Filename = %q, want %q", song.Filename, DefaultComposeDetailedFilename)
	}
	if !bytes.Equal(song.Audio, testAudio()) {
		t.Error("audio mismatch")
	}
}

func TestComposeDetailedParsedPrettyPrintedJSON(t *testing.T) {
	pretty, err := json.MarshalIndent(testMetadata(), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	jsonPart := strings.ReplaceAll(string(pretty), "\n", "\r\n")
	assertSong(t, mustCompose(t, multipartBody(t, "\r\n", jsonPart, true)))
}

func TestComposeDetailedParsedErrorsWithoutAJSONPart(t *testing.T) {
	body := []byte(testBoundary + "\r\nContent-Type: audio/mpeg\r\n\r\n")
	body = append(body, testAudio()...)
	body = append(body, []byte("\r\n"+testBoundary+"--\r\n")...)
	if message := decodeErrorMessage(t, body); !strings.Contains(message, "JSON part") {
		t.Errorf("message = %q", message)
	}
}

func TestComposeDetailedParsedErrorsWithoutAnAudioPart(t *testing.T) {
	body := []byte(testBoundary + "\r\nContent-Type: application/json\r\n\r\n{}\r\n" + testBoundary + "--\r\n")
	if message := decodeErrorMessage(t, body); !strings.Contains(message, "audio part") {
		t.Errorf("message = %q", message)
	}
}

func TestComposeDetailedParsedErrorsWhenTheJSONPartIsNotAnObject(t *testing.T) {
	body := multipartBody(t, "\r\n", "[1, 2]", true)
	if message := decodeErrorMessage(t, body); !strings.Contains(message, "not an object") {
		t.Errorf("message = %q", message)
	}
}

func TestComposeDetailedParsedAPIErrorsKeepTheGeneratedErrorHandling(t *testing.T) {
	client, _ := serve(t, http.StatusUnprocessableEntity, http.Header{"Content-Type": {"application/json"}}, []byte(`{"detail":[]}`))
	_, err := client.ComposeDetailedParsed(context.Background(), testRequest())
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusUnprocessableEntity || string(apiErr.Body) != `{"detail":[]}` {
		t.Errorf("APIError = %d %q", apiErr.StatusCode, apiErr.Body)
	}
}

func TestComposeDetailedParsedErrorsWhenTheBodyEndsMidAudio(t *testing.T) {
	body := multipartBody(t, "\r\n", "", true)
	// Cut inside the audio, so the closing delimiter never arrives.
	body = body[:len(body)-len("\r\n"+testBoundary+"--\r\n")-100]
	if message := decodeErrorMessage(t, body); !strings.Contains(message, "incomplete") {
		t.Errorf("message = %q", message)
	}
}

func TestComposeDetailedParsedErrorsOnAPartialClosingBoundary(t *testing.T) {
	body := multipartBody(t, "\r\n", "", true)
	body = body[:len(body)-len("3--\r\n")] // ends in "...\r\n--boundary12"
	if message := decodeErrorMessage(t, body); !strings.Contains(message, "incomplete") {
		t.Errorf("message = %q", message)
	}
}

func TestComposeDetailedParsedSkipsAPreambleUsingTheContentTypeBoundary(t *testing.T) {
	body := append([]byte("This is a multipart message.\r\n"), multipartBody(t, "\r\n", "", true)...)
	assertSong(t, mustCompose(t, body))
}

func TestComposeDetailedParsedTheDeadlineCoversAStalledBody(t *testing.T) {
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary123")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testBoundary + "\r\nContent-Type: application/json\r\n\r\n{}"))
		w.(http.Flusher).Flush()
		select {
		case <-release:
		case <-r.Context().Done():
		}
	}))
	t.Cleanup(server.Close)
	t.Cleanup(func() { close(release) })
	client := NewClient(core.NewRequestOptions(option.WithBaseURL(server.URL), option.WithoutRetries()))

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := client.ComposeDetailedParsed(ctx, testRequest())
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded, got %T: %v", err, err)
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Errorf("took %s to time out", elapsed)
	}
}

// Parser cases with no HTTP round trip.

func TestParseMultipartReadsTheBoundaryParameter(t *testing.T) {
	cases := map[string]string{
		`multipart/mixed; boundary="abc"`:           "abc",
		`multipart/mixed; boundary=----boundaryf00`: "----boundaryf00",
		`multipart/mixed; BOUNDARY=abc; charset=x`:  "abc",
		`application/json`:                          "",
		``:                                          "",
	}
	for contentType, want := range cases {
		if got := boundaryFrom(contentType); got != want {
			t.Errorf("boundaryFrom(%q) = %q, want %q", contentType, got, want)
		}
	}
}

func TestParseMultipartReadsAQuotedBoundaryFromTheResponse(t *testing.T) {
	client, _ := serve(t, http.StatusOK, http.Header{"Content-Type": {`multipart/mixed; boundary="boundary123"`}}, multipartBody(t, "\r\n", "", true))
	song, err := client.ComposeDetailedParsed(context.Background(), testRequest())
	if err != nil {
		t.Fatal(err)
	}
	assertSong(t, song)
}

func TestParseMultipartReadsAnUnquotedFilename(t *testing.T) {
	if got := filename("attachment; filename=song.mp3; size=1"); got != "song.mp3" {
		t.Errorf("filename = %q, want song.mp3", got)
	}
	if got := filename(""); got != DefaultComposeDetailedFilename {
		t.Errorf("filename = %q, want %q", got, DefaultComposeDetailedFilename)
	}
}

func TestParseMultipartFindsTheAudioPartBeforeTheJSONPart(t *testing.T) {
	body := []byte("--b\r\nContent-Type: audio/mpeg\r\n\r\n")
	body = append(body, testAudio()...)
	body = append(body, []byte("\r\n--b\r\nContent-Type: application/json\r\n\r\n{\"a\":1}\r\n--b--\r\n")...)
	parsed, err := parseMultipart(body, "")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.json["a"] != float64(1) || !bytes.Equal(parsed.audio, testAudio()) {
		t.Errorf("parsed = %v, %d bytes", parsed.json, len(parsed.audio))
	}
}

func TestParseMultipartRejectsABodyCutOffBeforeTheClosingBoundary(t *testing.T) {
	body := []byte("--b\r\nContent-Type: application/json\r\n\r\n{}\r\n--b\r\nContent-Type: audio/mpeg\r\n\r\n")
	body = append(body, testAudio()...)
	_, err := parseMultipart(body, "b")
	if err == nil || !strings.Contains(err.Error(), "incomplete") {
		t.Errorf("err = %v", err)
	}
	body = append(body, []byte("\r\n--")...)
	if _, err := parseMultipart(body, "b"); err == nil {
		t.Error("expected an error for a partial closing boundary")
	}
}

func TestParseMultipartSkipsAPreamble(t *testing.T) {
	body := []byte("preamble\r\n--b\r\nContent-Type: application/json\r\n\r\n{\"a\":1}\r\n--b\r\nContent-Type: audio/mpeg\r\n\r\n")
	body = append(body, testAudio()...)
	body = append(body, []byte("\r\n--b--\r\n")...)
	parsed, err := parseMultipart(body, "b")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.json["a"] != float64(1) || !bytes.Equal(parsed.audio, testAudio()) {
		t.Errorf("parsed = %v, %d bytes", parsed.json, len(parsed.audio))
	}
}

func TestParseMultipartSkipsTransportPaddingAfterEachDelimiter(t *testing.T) {
	body := []byte("--b \t\r\nContent-Type: application/json\r\n\r\n{\"a\":1}\r\n--b\t\r\nContent-Type: audio/mpeg\r\n\r\n")
	body = append(body, testAudio()...)
	body = append(body, []byte("\r\n--b--\r\n")...)
	parsed, err := parseMultipart(body, "b")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.json["a"] != float64(1) || !bytes.Equal(parsed.audio, testAudio()) {
		t.Errorf("parsed = %v, %d bytes", parsed.json, len(parsed.audio))
	}
}

func TestParseMultipartReadsAPartWithNoHeaders(t *testing.T) {
	body := []byte("--b\r\nContent-Type: application/json\r\n\r\n{\"a\":1}\r\n--b\r\n\r\n")
	body = append(body, testAudio()...)
	body = append(body, []byte("\r\n--b--\r\n")...)
	parsed, err := parseMultipart(body, "")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(parsed.audio, testAudio()) {
		t.Errorf("audio = %d bytes", len(parsed.audio))
	}
	if parsed.filename != DefaultComposeDetailedFilename {
		t.Errorf("filename = %q", parsed.filename)
	}
}

func TestParseMultipartRejectsABodyWithoutABoundary(t *testing.T) {
	if _, err := parseMultipart([]byte("not multipart"), ""); err == nil {
		t.Error("expected an error")
	}
	if _, err := parseMultipart(nil, ""); err == nil {
		t.Error("expected an error for an empty body")
	}
}
