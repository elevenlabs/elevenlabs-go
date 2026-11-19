# ElevenLabs Go Library

![LOGO](.github/assets/logo.png)

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://buildwithfern.com/?utm_source=fern-elevenlabs/elevenlabs-go/readme)
[![Discord](https://badgen.net/badge/black/ElevenLabs/icon?icon=discord&label)](https://discord.gg/elevenlabs)
[![Twitter](https://badgen.net/badge/black/elevenlabsio/icon?icon=twitter&label)](https://twitter.com/elevenlabsio)
[![Go Reference](https://pkg.go.dev/badge/github.com/elevenlabs/elevenlabs-go.svg)](https://pkg.go.dev/github.com/elevenlabs/elevenlabs-go)

The official Go SDK for [ElevenLabs](https://elevenlabs.io/). ElevenLabs brings the most compelling, rich and lifelike voices to creators and developers in just a few lines of code.

## 📖 API & Docs

Check out the [HTTP API documentation](https://elevenlabs.io/docs/api-reference). A full reference for this library is available [here](./reference.md).

## Requirements

This SDK requires Go `1.22` or later.

## Install

```sh
go get github.com/elevenlabs/elevenlabs-go
```

## Usage

### Main Models

1. **Eleven v4** (`eleven_v4`)
    - Highest quality and most expressive speech, with reliable voice cloning across long generations
    - 90+ languages with consistent voice quality across language switches
    - Up to 10,000 characters (~10 minutes of audio) per request
    - Used via the [Text to Dialogue API](https://elevenlabs.io/docs/api-reference/text-to-dialogue/convert)

2. **Eleven v4 Turbo** (`eleven_v4_turbo`)
    - State-of-the-art real-time speech, with a median latency of ~100ms
    - 90+ languages
    - Ideal for voice agents and other latency-sensitive use cases

3. **Eleven Multilingual v2** (`eleven_multilingual_v2`)
    - Excels in stability, language diversity, and accent accuracy
    - The most stable choice for long-form Text to Speech
    - Supports 29 languages

4. **Eleven Flash v2.5** (`eleven_flash_v2_5`)
    - Ultra-low latency and lowest cost
    - Supports 32 languages

For more detailed information about these models and others, visit the [ElevenLabs Models documentation](https://elevenlabs.io/docs/models).

Instantiate and use the client with the following:

```go
package main

import (
	"context"
	"io"
	"log"
	"os"

	elevenlabs "github.com/elevenlabs/elevenlabs-go"
	elevenlabsclient "github.com/elevenlabs/elevenlabs-go/client"
	"github.com/elevenlabs/elevenlabs-go/types"
)

func main() {
	// Reads the API key from the ELEVENLABS_API_KEY environment variable.
	// Pass it explicitly with option.WithApiKey(elevenlabs.String("YOUR_API_KEY")) if you prefer.
	client := elevenlabsclient.NewClient()

	audio, err := client.TextToSpeech.Convert(
		context.Background(),
		"JBFqnCBsd6RMkjVDRZzb",
		&types.BodyTextToSpeechFull{
			Text:         "The first move is what sets everything in motion.",
			ModelId:      elevenlabs.String("eleven_multilingual_v2"),
			OutputFormat: types.ConvertTextToSpeechRequestOutputFormatMp344100128.Ptr(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("output.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// The response body streams as it arrives. io.Copy writes it to the file until the end.
	if _, err := io.Copy(file, audio); err != nil {
		log.Fatal(err)
	}
}
```

`NewClient` falls back to `ELEVENLABS_API_KEY` only when no key is passed. To set the key in code,
use `option.WithApiKey` from `github.com/elevenlabs/elevenlabs-go/option`:

```go
client := elevenlabsclient.NewClient(
	option.WithApiKey(elevenlabs.String("YOUR_API_KEY")),
)
```

Optional fields are pointers. `elevenlabs.String`, `elevenlabs.Int`, `elevenlabs.Bool` and the
other helpers in the root package take a value and return a pointer to it, and enum constants in the
`types` package have a `Ptr()` method.

### Text to Dialogue (Eleven v4)

Eleven v4 and Eleven v4 Turbo are served through the Text to Dialogue API, which turns a list of speaker turns into a single expressive audio stream.

```go
audio, err := client.TextToDialogue.Convert(ctx, &types.BodyTextToDialogueMultiVoiceV1TextToDialoguePost{
	ModelId:      elevenlabs.String("eleven_v4"),
	OutputFormat: elevenlabs.String("mp3_44100_128"),
	Inputs: []*types.DialogueInput{
		{Text: "[giggling] Knock knock", VoiceId: "JBFqnCBsd6RMkjVDRZzb"},
		{Text: "[curious] Who is there?", VoiceId: "Aw4FAjKCGjjNkVhN1Xmq"},
	},
})
if err != nil {
	return err
}

file, err := os.Create("dialogue.mp3")
if err != nil {
	return err
}
defer file.Close()

_, err = io.Copy(file, audio)
return err
```

### Music with metadata

`client.Music.ComposeDetailedParsed` composes a song and returns its metadata and audio file as
separate fields. It takes the same arguments as `ComposeDetailed`. The API answers with a
`multipart/mixed` body (a JSON part, then the audio), which the generated `ComposeDetailed` can't
decode.

```go
song, err := client.Music.ComposeDetailedParsed(ctx, &types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost{
	Prompt:        elevenlabs.String("Epic orchestral music with dramatic strings"),
	MusicLengthMs: elevenlabs.Int(60000),
})
if err != nil {
	return err
}

metadata, _ := song.JSON["song_metadata"].(map[string]any)
fmt.Println(metadata["title"])
return os.WriteFile(song.Filename, song.Audio, 0o644)
```

`song.JSON` holds the metadata part as the API sent it (`composition_plan`, `song_metadata`,
`words_timestamps` and, with `WithWaveformVisual`, `waveform_visual`). `song.Filename` is the name
the API suggests for the audio, or `generated_music.mp3`, and `song.SongID` is the `song-id`
response header, sent when `StoreForInpainting` is set. The whole body is read before the method
returns, so a context deadline or HTTP client `Timeout` covers the download. An error status
returns a `*core.APIError`, as `ComposeDetailed` does; a body that isn't a complete multipart body
with a JSON part and an audio part returns a `*ComposeDetailedDecodeError` from the `music/client`
package.

## Voices

List all your available voices with `Search`.

```go
response, err := client.Voices.Search(ctx, &types.SearchVoicesRequest{})
if err != nil {
	return err
}
for _, voice := range response.Voices {
	if name := voice.GetName(); name != nil {
		fmt.Println(*name)
	}
}
```

For information about the structure of the voices output, please refer to the [official ElevenLabs API documentation for Get Voices](https://elevenlabs.io/docs/api-reference/get-voices).

## Streaming

Stream audio in real-time, as it's being generated. `Convert` and `Stream` both return an
`io.Reader` over the live response body: each `Read` returns the bytes that have arrived so far, so
playback or a file write can start before the server has finished.

```go
audio, err := client.TextToSpeech.Stream(
	ctx,
	"JBFqnCBsd6RMkjVDRZzb",
	&types.StreamTextToSpeechRequest{
		Text:         "This is a... streaming voice.",
		ModelId:      elevenlabs.String("eleven_multilingual_v2"),
		OutputFormat: types.StreamTextToSpeechRequestOutputFormatMp344100128.Ptr(),
	},
)
if err != nil {
	return err
}

chunk := make([]byte, 32*1024)
for {
	n, err := audio.Read(chunk)
	if n > 0 {
		// Process each chunk as it arrives.
		os.Stdout.Write(chunk[:n])
	}
	if err == io.EOF {
		break
	}
	if err != nil {
		return err
	}
}
```

The reader can be read once. It releases the connection when it reaches `io.EOF` or a read error.
If you stop reading early, close it (it also implements `io.Closer`) or cancel the request's context:

```go
if closer, ok := audio.(io.Closer); ok {
	closer.Close()
}
```

## Flows

Flows image, video and text-to-speech generations, and flows template runs, are asynchronous:
`Create` returns straight away with a `pending` status, and `Get` reports progress until the work
is `completed` or `failed`. `CreateAndWait` does both. It polls `Get` at the pace the server asks
for in its `Retry-After` header, never faster than `MinPollInterval` (3 seconds by default), and
returns the final generation:

```go
generation, err := client.Flows.Image.CreateAndWait(
	ctx,
	&types.ImageGenerationRequest{
		ModelId:   "gpt-image-2",
		GptImage2: &types.GptImage2Request{Prompt: "A lighthouse at dawn, watercolor"},
	},
	&flows.WaitOptions{WaitTimeout: 5 * time.Minute},
)
var timeoutErr *flows.WaitTimeoutError
if errors.As(err, &timeoutErr) {
	// Still running on the server: resume later with client.Flows.Image.Get(ctx, timeoutErr.ID).
	return err
}
if err != nil {
	return err
}
if completed := generation.GetCompleted(); completed != nil {
	fmt.Println(completed.ContentUrl)
} else if failed := generation.GetFailed(); failed != nil {
	fmt.Println("Generation failed:", failed.FailureReason, failed.ErrorMessage)
}
```

`client.Flows.Video.CreateAndWait` and `client.Flows.TextToSpeech.CreateAndWait` work the same
way. For a template run, pass the template id; the run comes back with its `Status` and `Outputs`:

```go
run, err := client.Flows.Templates.Runs.CreateAndWait(
	ctx,
	"TEMPLATE_ID",
	&templates.TemplateRunCreateRequest{
		Inputs: map[string]*types.TemplateRunInput{"prompt": {String: "A fox in the snow"}},
	},
	nil, // the default wait options: no timeout, 3-second minimum poll interval
)
```

A `failed` generation or run is returned with a `nil` error, so check its status. The error is a
`*flows.WaitTimeoutError` (carrying `ID`, and `TemplateID` for a run) when `WaitTimeout` runs out,
wraps `ctx.Err()` when the context is cancelled, and is the API error when the create call or a
status check fails. Request options apply to the create call and to every status check, except
`option.WithBodyProperties`, which only the create call gets.

`CreateAndWait` is meant for scripts and quickstarts. In production, set `Webhook` on the request
and handle the webhook event instead of holding a process open.

## Environments

This SDK allows you to configure different environments for API requests. It defaults to the production environment.

```go
client := elevenlabsclient.NewClient(
	option.WithBaseURL(elevenlabs.Environments.ProductionUs),
)
```

Available environments:
- `elevenlabs.Environments.Production`
- `elevenlabs.Environments.ProductionUs`
- `elevenlabs.Environments.ProductionEu`
- `elevenlabs.Environments.ProductionIndia`
- `elevenlabs.Environments.ProductionSingapore`

`WithBaseURL` also accepts any URL string, such as a staging server or a local mock.

## Errors

When the API returns a non-success status code (4xx or 5xx), the method returns a `*core.APIError`
(from `github.com/elevenlabs/elevenlabs-go/core`) as its error. The HTTP status code is in
`StatusCode`, the response headers in `Header` and the raw response body in `Body`. Use `errors.As`
to reach it, and compare `StatusCode` to handle a specific status:

```go
_, err := client.TextToSpeech.Convert(ctx, "JBFqnCBsd6RMkjVDRZzb", request)

var apiErr *core.APIError
if errors.As(err, &apiErr) {
	fmt.Println("Status code:", apiErr.StatusCode)
	fmt.Println("Response body:", string(apiErr.Body))
	if apiErr.StatusCode == http.StatusNotFound {
		// Handle a 404 specifically.
	}
} else if err != nil {
	// The request never got a response: a network failure, or a cancelled or expired context.
	fmt.Println("Request failed:", err)
}
```

## Webhooks

ElevenLabs signs each webhook it sends with your webhook secret. `client.Webhooks.ConstructEvent`
verifies the `ElevenLabs-Signature` header and returns the decoded event as a `map[string]any`. Pass
the raw request body exactly as received: re-encoding a decoded body changes the bytes and breaks
verification.

```go
http.HandleFunc("/webhooks/elevenlabs", func(w http.ResponseWriter, r *http.Request) {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusBadRequest)
		return
	}

	event, err := client.Webhooks.ConstructEvent(
		rawBody,
		r.Header.Get("ElevenLabs-Signature"),
		os.Getenv("ELEVENLABS_WEBHOOK_SECRET"),
	)
	if err != nil {
		// Missing or invalid signature, or a timestamp older than 30 minutes.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(event["type"])
	w.WriteHeader(http.StatusOK)
})
```

`webhooks.ConstructEvent` (package `github.com/elevenlabs/elevenlabs-go/webhooks`) does the same
without a client. Each failure is a sentinel error you can match with `errors.Is`, such as
`webhooks.ErrSignatureMismatch` or `webhooks.ErrTimestampOutsideTolerance`.

## WebSockets

`client.TextToSpeech.Realtime(...)` opens a connection and returns a `*texttospeech.TextToSpeechSocket`.
Register handlers with its `On...` methods, call `Start` to begin delivering messages, send with its
typed `Send...` methods (or `SendText` and `SendBinary` for raw frames), and call `Close` when done.
Handlers run one at a time, in arrival order; messages that arrive before `Start` are held, not
dropped.

```go
socket, err := client.TextToSpeech.Realtime(
	ctx,
	&texttospeech.TextToSpeechConnectRequest{
		VoiceId:      "JBFqnCBsd6RMkjVDRZzb",
		ModelId:      elevenlabs.String("eleven_flash_v2_5"),
		OutputFormat: elevenlabs.String("mp3_44100_128"),
	},
	// Stay closed once the server finishes, instead of reconnecting.
	option.WithoutWebSocketReconnect(),
)
if err != nil {
	return err
}
defer socket.Close()

file, err := os.Create("realtime.mp3")
if err != nil {
	return err
}
defer file.Close()

done := make(chan struct{})
socket.OnReceiveMessage(func(message *types.ReceiveMessage) {
	if output := message.GetAudioOutput(); output != nil {
		chunk, err := base64.StdEncoding.DecodeString(output.GetAudio())
		if err == nil {
			file.Write(chunk)
		}
	}
})
socket.OnError(func(err error) {
	log.Println("socket error:", err)
})
socket.OnDisconnected(func(reason core.DisconnectReason) {
	close(done)
})
socket.Start()

for _, message := range []*types.SendMessage{
	{InitializeConnection: &types.InitializeConnection{}},
	{SendText: &types.SendText{Text: "Hello from the Go SDK. "}},
	{CloseConnection: &types.CloseConnection{}},
} {
	if err := socket.SendSendMessage(ctx, message); err != nil {
		return err
	}
}

// The server sends the remaining audio, then closes the connection.
<-done
```

The SDK has one socket per channel:

- `client.Dubbing.Live` returns a `*DubbingSocket` (package `github.com/elevenlabs/elevenlabs-go/dubbing/client`)
- `client.TextToSpeech.Realtime` returns a `*texttospeech.TextToSpeechSocket`
- `client.TextToDialogue.Realtime` returns a `*texttodialogue.TextToDialogueSocket`
- `client.SpeechToText.Realtime` returns a `*SpeechToTextSocket` (package `github.com/elevenlabs/elevenlabs-go/speechtotext/client`)
- `client.SpeechEngine.Upstream` returns a `*speechengine.SpeechEngineSocket`
- `client.TextToDialogueMultiContext.Realtime` returns a `*texttodialoguemulticontext.TextToDialogueMultiContextSocket`
- `client.TextToSpeechMultiContext.Realtime` returns a `*texttospeechmulticontext.TextToSpeechMultiContextSocket`
- `client.Translation.Realtime` returns a `*translation.TranslationSocket`

Every socket has the same lifecycle methods: `OnConnected`, `OnDisconnected`, `OnError`, `Start`,
`ReadyState`, `SendText`, `SendBinary` and `Close`.

### Reconnection

By default a socket reconnects whenever the connection is lost, including when the server closes it
normally, and sends the handshake (credentials and connect parameters) again each time. Sends made
while it is reconnecting wait for the new connection, bounded by their `ctx`. The default backoff
starts at 1 second, grows by 1.3x per attempt up to 10 seconds, and keeps retrying until `Close` is
called; each handshake attempt times out after 10 seconds.

Pass `option.WithoutWebSocketReconnect()` for a one-shot session that should stay closed, as in the
example above, or tune the backoff with `option.WithWebSocketReconnect`. Zero fields keep their
defaults:

```go
socket, err := client.TextToSpeech.Realtime(
	ctx,
	request,
	option.WithWebSocketReconnect(core.ReconnectOptions{
		MinDelay:    500 * time.Millisecond,
		MaxDelay:    5 * time.Second,
		MaxAttempts: 5,
	}),
)
```

## Advanced

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as
long as the request is deemed retryable and the number of attempts has not grown larger than the
configured limit (default: 2 attempts, that is the first request and one retry).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses) (Internal Server Error)

Between attempts the SDK waits as long as the `Retry-After` or `X-RateLimit-Reset` response header
asks, and otherwise backs off exponentially from 1 second, with jitter. No wait exceeds 60 seconds.
A request that fails without a response, such as a refused connection, is not retried.

Set the limit on the client with `option.WithMaxAttempts`, or turn retries off for one call with
`option.WithoutRetries` (`WithMaxAttempts(0)` falls back to the default):

```go
client := elevenlabsclient.NewClient(
	option.WithMaxAttempts(3),
)

// Send this one request exactly once.
client.TextToSpeech.Convert(ctx, "JBFqnCBsd6RMkjVDRZzb", request, option.WithoutRetries())
```

### Timeouts

The SDK sets no timeout of its own: without a custom HTTP client it uses Go's `http.DefaultClient`,
which waits indefinitely. Bound a call with a context deadline, or set `Timeout` on an
`*http.Client` that you pass to the client. Both limits also cover reading the response body, so
leave room for long audio.

```go
client := elevenlabsclient.NewClient(
	option.WithHTTPClient(&http.Client{Timeout: 60 * time.Second}),
)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client.TextToSpeech.Convert(ctx, "JBFqnCBsd6RMkjVDRZzb", request)
```

### Custom HTTP Client

This SDK is built to work with any instance of `*http.Client` (or any type with a
`Do(*http.Request) (*http.Response, error)` method). By default, if no client is provided, the SDK
uses `http.DefaultClient`. Pass your own, for example to set a proxy or transport, through
`option.WithHTTPClient`:

```go
httpClient := &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	},
}

client := elevenlabsclient.NewClient(
	option.WithHTTPClient(httpClient),
)
```

### Per-request Options

Every method accepts variadic `opts ...option.RequestOption` to override behavior for a single call —
a different base URL, additional headers, additional query parameters, or additional body
properties. Passed to `NewClient` instead, the base URL, headers, API key, HTTP client and retry
options apply to every call; query parameters and body properties take effect only per call.

```go
client.TextToSpeech.Convert(
	ctx,
	"JBFqnCBsd6RMkjVDRZzb",
	request,
	option.WithBaseURL("https://staging.example.com"),
	option.WithHTTPHeader(http.Header{"X-Custom-Header": []string{"value"}}),
	option.WithQueryParameters(url.Values{"trace": []string{"true"}}),
	option.WithBodyProperties(map[string]interface{}{"extra_field": "value"}),
)
```

### Raw Response

Each resource client that makes HTTP calls, such as `client.TextToSpeech`, has a `WithRawResponse`
field with the same methods. They return a `*core.Response` holding the status code and response
headers alongside the body:

```go
response, err := client.TextToSpeech.WithRawResponse.Convert(ctx, "JBFqnCBsd6RMkjVDRZzb", request)
if err != nil {
	return err
}
fmt.Println(response.StatusCode)
fmt.Println(response.Header.Get("Content-Type"))
audio := response.Body // the io.Reader that Convert returns
```

## Languages Supported

Explore [all models & languages](https://elevenlabs.io/docs/models).

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
