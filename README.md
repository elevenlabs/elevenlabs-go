# Api Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=Api%2FGo)

This is the documentation for the ElevenLabs API. You can use this API to use our service programmatically, this is done by using your API key. You can find your API key in the dashboard at https://elevenlabs.io/app/settings/api-keys.

## Table of Contents

- [Reference](#reference)
- [Installation](#installation)
- [Usage](#usage)
- [Environments](#environments)
- [Errors](#errors)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Custom Client](#custom-client)
  - [Timeouts](#timeouts)
  - [WebSockets](#websockets)
- [Contributing](#contributing)

## Reference

A full reference for this library is available [here](./reference.md).

## Installation

This SDK requires Go 1.22 or later.

```sh
go get github.com/elevenlabs/elevenlabs-go
```

## Usage

```go
import (
	context "context"
	elevenlabs "github.com/elevenlabs/elevenlabs-go"
	elevenlabsclient "github.com/elevenlabs/elevenlabs-go/client"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

client := elevenlabsclient.NewClient()
client.History.List(
	context.TODO(),
	&types.ListHistoryRequest{
		PageSize: elevenlabs.Int(
			1,
		),
		StartAfterHistoryItemId: elevenlabs.String(
			"start_after_history_item_id",
		),
		VoiceId: elevenlabs.String(
			"voice_id",
		),
		ModelId: elevenlabs.String(
			"model_id",
		),
		DateBeforeUnix: elevenlabs.Int(
			1,
		),
		DateAfterUnix: elevenlabs.Int(
			1,
		),
		SortDirection: types.ListHistoryRequestSortDirectionAsc.Ptr(),
		Search: elevenlabs.String(
			"search",
		),
		Source: types.ListHistoryRequestSourceTts.Ptr(),
	},
)
```

## Environments

This SDK allows you to configure a different environment for API requests by passing the
`WithBaseURL` request option to the client.

```go
import (
  "os"
  "github.com/elevenlabs/elevenlabs-go/client"
  "github.com/elevenlabs/elevenlabs-go/option"
  "github.com/elevenlabs/elevenlabs-go"
)

client := client.NewClient(
  option.WithApiKey(elevenlabs.String(os.Getenv("ELEVENLABS_API_KEY"))),
  option.WithBaseURL(elevenlabs.Environments.Production),
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

A non-2xx response returns an `error` wrapping a `*core.APIError`, which carries the status code
and raw response body:

```go
var apiErr *core.APIError
if errors.As(err, &apiErr) {
  fmt.Println(apiErr.StatusCode)
}
```

The `core` package also exposes typed errors for endpoints that declare them, each
embedding `*core.APIError` so the status code is still reachable through it.

## Request Options

Every method accepts variadic `opts ...option.RequestOption` to override behavior for a single
call — a different base URL, additional headers, or a longer timeout:

```go
response, err := client.SomeResource.SomeMethod(
  ctx,
  request,
  option.WithHTTPHeader(httpHeader),
  option.WithMaxAttempts(1),
)
```

## Advanced

### Custom Client

This SDK is built to work with any instance of `*http.Client`. By default, if no client is
provided, the SDK constructs one. You can pass your own through `option.WithHTTPClient`:

```go
customClient := &http.Client{}

client := client.NewClient(
  option.WithHTTPClient(customClient),
)
```

### Timeouts

```go
import (
  "os"
  "net/http"
  "time"
  "github.com/elevenlabs/elevenlabs-go/client"
  "github.com/elevenlabs/elevenlabs-go/option"
  "github.com/elevenlabs/elevenlabs-go"
)

client := client.NewClient(
  option.WithApiKey(elevenlabs.String(os.Getenv("ELEVENLABS_API_KEY"))),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 10 * time.Second,
    },
  ),
)
```

### WebSockets

`client.TextToSpeech.Realtime()` opens the connection and returns a `*TextToSpeechSocket`. Register handlers
with its `On...` methods, call `Start()` to begin delivering messages, send with its typed `Send...` methods, and call
`Close()` when done.

```go
socket, err := client.TextToSpeech.Realtime(ctx, nil)
if err != nil {
  return err
}
defer socket.Close()

socket.OnReceiveMessage(func(message *types.ReceiveMessage) {
  // Handle each message.
})
socket.Start()

var payload types.SendMessage
if err := socket.SendSendMessage(ctx, &payload); err != nil {
  return err
}
```

The SDK has one socket per channel:

- `client.TextToSpeech.Realtime()` returns a `*TextToSpeechSocket`
- `client.TextToDialogue.Realtime()` returns a `*TextToDialogueSocket`
- `client.Dubbing.Live()` returns a `*DubbingSocket`
- `client.SpeechToText.Realtime()` returns a `*SpeechToTextSocket`
- `client.SpeechEngine.Upstream()` returns a `*SpeechEngineSocket`
- `client.TextToDialogueMultiContext.Realtime()` returns a `*TextToDialogueMultiContextSocket`
- `client.TextToSpeechMultiContext.Realtime()` returns a `*TextToSpeechMultiContextSocket`
- `client.Translation.Realtime()` returns a `*TranslationSocket`

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
