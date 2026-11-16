package client

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	assets "github.com/elevenlabs/elevenlabs-go/reception/assets"
	billing "github.com/elevenlabs/elevenlabs-go/reception/billing"
	businessprofile "github.com/elevenlabs/elevenlabs-go/reception/businessprofile"
	holidays "github.com/elevenlabs/elevenlabs-go/reception/holidays"
	knowledgesources "github.com/elevenlabs/elevenlabs-go/reception/knowledgesources"
	locations "github.com/elevenlabs/elevenlabs-go/reception/locations"
	phonenumbers "github.com/elevenlabs/elevenlabs-go/reception/phonenumbers"
	productsclient "github.com/elevenlabs/elevenlabs-go/reception/products/client"
	receptionistsclient "github.com/elevenlabs/elevenlabs-go/reception/receptionists/client"
	services "github.com/elevenlabs/elevenlabs-go/reception/services"
	staff "github.com/elevenlabs/elevenlabs-go/reception/staff"
	subscriptions "github.com/elevenlabs/elevenlabs-go/reception/subscriptions"
	teamclient "github.com/elevenlabs/elevenlabs-go/reception/team/client"
	tiers "github.com/elevenlabs/elevenlabs-go/reception/tiers"
	workspaces "github.com/elevenlabs/elevenlabs-go/reception/workspaces"
)

type Client struct {
	Services         *services.Client
	Assets           *assets.Client
	Products         *productsclient.Client
	Staff            *staff.Client
	Receptionists    *receptionistsclient.Client
	PhoneNumbers     *phonenumbers.Client
	Holidays         *holidays.Client
	Locations        *locations.Client
	BusinessProfile  *businessprofile.Client
	KnowledgeSources *knowledgesources.Client
	Workspaces       *workspaces.Client
	Subscriptions    *subscriptions.Client
	Tiers            *tiers.Client
	Billing          *billing.Client
	Team             *teamclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		Services:         services.NewClient(options),
		Assets:           assets.NewClient(options),
		Products:         productsclient.NewClient(options),
		Staff:            staff.NewClient(options),
		Receptionists:    receptionistsclient.NewClient(options),
		PhoneNumbers:     phonenumbers.NewClient(options),
		Holidays:         holidays.NewClient(options),
		Locations:        locations.NewClient(options),
		BusinessProfile:  businessprofile.NewClient(options),
		KnowledgeSources: knowledgesources.NewClient(options),
		Workspaces:       workspaces.NewClient(options),
		Subscriptions:    subscriptions.NewClient(options),
		Tiers:            tiers.NewClient(options),
		Billing:          billing.NewClient(options),
		Team:             teamclient.NewClient(options),
		options:          options,
		baseURL:          options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
