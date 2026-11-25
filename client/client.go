package client

import (
	"github.com/elevenlabs/elevenlabs-go/agentbuilder"
	"github.com/elevenlabs/elevenlabs-go/assets"
	"github.com/elevenlabs/elevenlabs-go/audioeffects"
	"github.com/elevenlabs/elevenlabs-go/audioisolation"
	"github.com/elevenlabs/elevenlabs-go/audionative"
	"github.com/elevenlabs/elevenlabs-go/contentprovenance"
	"github.com/elevenlabs/elevenlabs-go/core"
	dubbingclient "github.com/elevenlabs/elevenlabs-go/dubbing/client"
	"github.com/elevenlabs/elevenlabs-go/environments"
	"github.com/elevenlabs/elevenlabs-go/environmentvariables"
	flowsclient "github.com/elevenlabs/elevenlabs-go/flows/client"
	"github.com/elevenlabs/elevenlabs-go/forcedalignment"
	"github.com/elevenlabs/elevenlabs-go/history"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/models"
	musicclient "github.com/elevenlabs/elevenlabs-go/music/client"
	"github.com/elevenlabs/elevenlabs-go/option"
	productionsclient "github.com/elevenlabs/elevenlabs-go/productions/client"
	pronunciationdictionariesclient "github.com/elevenlabs/elevenlabs-go/pronunciationdictionaries/client"
	receptionclient "github.com/elevenlabs/elevenlabs-go/reception/client"
	"github.com/elevenlabs/elevenlabs-go/samples"
	serviceaccountsclient "github.com/elevenlabs/elevenlabs-go/serviceaccounts/client"
	"github.com/elevenlabs/elevenlabs-go/speechengine"
	"github.com/elevenlabs/elevenlabs-go/speechtospeech"
	speechtotextclient "github.com/elevenlabs/elevenlabs-go/speechtotext/client"
	studioclient "github.com/elevenlabs/elevenlabs-go/studio/client"
	"github.com/elevenlabs/elevenlabs-go/subtitling"
	"github.com/elevenlabs/elevenlabs-go/texttodialogue"
	texttodialoguemulticontext "github.com/elevenlabs/elevenlabs-go/texttodialoguemulticontext"
	"github.com/elevenlabs/elevenlabs-go/texttosoundeffects"
	"github.com/elevenlabs/elevenlabs-go/texttospeech"
	texttospeechmulticontext "github.com/elevenlabs/elevenlabs-go/texttospeechmulticontext"
	texttovoiceclient "github.com/elevenlabs/elevenlabs-go/texttovoice/client"
	tokensclient "github.com/elevenlabs/elevenlabs-go/tokens/client"
	translation "github.com/elevenlabs/elevenlabs-go/translation"
	userclient "github.com/elevenlabs/elevenlabs-go/user/client"
	"github.com/elevenlabs/elevenlabs-go/voices/client"
	"github.com/elevenlabs/elevenlabs-go/webhooks"
	workspaceclient "github.com/elevenlabs/elevenlabs-go/workspace/client"
	workspacesclient "github.com/elevenlabs/elevenlabs-go/workspaces/client"
	"os"
)

type Client struct {
	History                    *history.Client
	TextToSoundEffects         *texttosoundeffects.Client
	AudioIsolation             *audioisolation.Client
	AudioEffects               *audioeffects.Client
	Samples                    *samples.Client
	Voices                     *client.Client
	TextToSpeech               *texttospeech.Client
	TextToDialogue             *texttodialogue.Client
	SpeechToSpeech             *speechtospeech.Client
	TextToVoice                *texttovoiceclient.Client
	User                       *userclient.Client
	Studio                     *studioclient.Client
	Dubbing                    *dubbingclient.Client
	Models                     *models.Client
	AudioNative                *audionative.Client
	ContentProvenance          *contentprovenance.Client
	PronunciationDictionaries  *pronunciationdictionariesclient.Client
	Workspaces                 *workspacesclient.Client
	Workspace                  *workspaceclient.Client
	ServiceAccounts            *serviceaccountsclient.Client
	Webhooks                   *webhooks.Client
	Music                      *musicclient.Client
	SpeechToText               *speechtotextclient.Client
	Tokens                     *tokensclient.Client
	ForcedAlignment            *forcedalignment.Client
	Agents                     *agentbuilder.AgentsClient
	SpeechEngine               *speechengine.SpeechEngineClient
	EnvironmentVariables       *environmentvariables.Client
	Environments               *environments.Client
	Productions                *productionsclient.Client
	Flows                      *flowsclient.Client
	Assets                     *assets.Client
	Reception                  *receptionclient.Client
	TextToDialogueMultiContext *texttodialoguemulticontext.Client
	TextToSpeechMultiContext   *texttospeechmulticontext.Client
	Subtitling                 *subtitling.Client
	Translation                *translation.Client
	options                    *core.RequestOptions
	baseURL                    string
	caller                     *internal.Caller
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.ApiKey == nil {
		if value := os.Getenv("ELEVENLABS_API_KEY"); value != "" {
			options.ApiKey = &value
		}
	}
	return &Client{
		History:                    history.NewClient(options),
		TextToSoundEffects:         texttosoundeffects.NewClient(options),
		AudioIsolation:             audioisolation.NewClient(options),
		AudioEffects:               audioeffects.NewClient(options),
		Samples:                    samples.NewClient(options),
		Voices:                     client.NewClient(options),
		TextToSpeech:               texttospeech.NewClient(options),
		TextToDialogue:             texttodialogue.NewClient(options),
		SpeechToSpeech:             speechtospeech.NewClient(options),
		TextToVoice:                texttovoiceclient.NewClient(options),
		User:                       userclient.NewClient(options),
		Studio:                     studioclient.NewClient(options),
		Dubbing:                    dubbingclient.NewClient(options),
		Models:                     models.NewClient(options),
		AudioNative:                audionative.NewClient(options),
		ContentProvenance:          contentprovenance.NewClient(options),
		PronunciationDictionaries:  pronunciationdictionariesclient.NewClient(options),
		Workspaces:                 workspacesclient.NewClient(options),
		Workspace:                  workspaceclient.NewClient(options),
		ServiceAccounts:            serviceaccountsclient.NewClient(options),
		Webhooks:                   webhooks.NewClient(options),
		Music:                      musicclient.NewClient(options),
		SpeechToText:               speechtotextclient.NewClient(options),
		Tokens:                     tokensclient.NewClient(options),
		ForcedAlignment:            forcedalignment.NewClient(options),
		Agents:                     agentbuilder.NewAgentsClient(options),
		SpeechEngine:               speechengine.NewSpeechEngineClient(options),
		EnvironmentVariables:       environmentvariables.NewClient(options),
		Environments:               environments.NewClient(options),
		Productions:                productionsclient.NewClient(options),
		Flows:                      flowsclient.NewClient(options),
		Assets:                     assets.NewClient(options),
		Reception:                  receptionclient.NewClient(options),
		TextToDialogueMultiContext: texttodialoguemulticontext.NewClient(options),
		TextToSpeechMultiContext:   texttospeechmulticontext.NewClient(options),
		Subtitling:                 subtitling.NewClient(options),
		Translation:                translation.NewClient(options),
		options:                    options,
		baseURL:                    options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
