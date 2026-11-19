package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ListPhoneNumbersResponseItem struct {
	Provider string
	Exotel   *types.GetPhoneNumberExotelResponseModel
	SipTrunk *types.GetPhoneNumberSipTrunkResponseModel
	Twilio   *types.GetPhoneNumberTwilioResponseModel
}

func (l *ListPhoneNumbersResponseItem) GetExotel() *types.GetPhoneNumberExotelResponseModel {
	if l == nil {
		return nil
	}
	if l.Exotel == nil {
		return nil
	}
	return l.Exotel
}

func (l *ListPhoneNumbersResponseItem) GetSipTrunk() *types.GetPhoneNumberSipTrunkResponseModel {
	if l == nil {
		return nil
	}
	if l.SipTrunk == nil {
		return nil
	}
	return l.SipTrunk
}

func (l *ListPhoneNumbersResponseItem) GetTwilio() *types.GetPhoneNumberTwilioResponseModel {
	if l == nil {
		return nil
	}
	if l.Twilio == nil {
		return nil
	}
	return l.Twilio
}

func (l *ListPhoneNumbersResponseItem) GetProvider() string {
	if l == nil {
		return ""
	}
	return l.Provider
}

func (l *ListPhoneNumbersResponseItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	l.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", l)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(types.GetPhoneNumberExotelResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Exotel = value
	case "sip_trunk":
		value := new(types.GetPhoneNumberSipTrunkResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.SipTrunk = value
	case "twilio":
		value := new(types.GetPhoneNumberTwilioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Twilio = value
	}
	return nil
}

func (l ListPhoneNumbersResponseItem) MarshalJSON() ([]byte, error) {
	if err := l.validate(); err != nil {
		return nil, err
	}
	if l.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(l.Exotel, "provider", "exotel")
	}
	if l.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(l.SipTrunk, "provider", "sip_trunk")
	}
	if l.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(l.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", l)
}

type ListPhoneNumbersResponseItemVisitor interface {
	VisitExotel(*types.GetPhoneNumberExotelResponseModel) error
	VisitSipTrunk(*types.GetPhoneNumberSipTrunkResponseModel) error
	VisitTwilio(*types.GetPhoneNumberTwilioResponseModel) error
}

func (l *ListPhoneNumbersResponseItem) Accept(visitor ListPhoneNumbersResponseItemVisitor) error {
	if l.Exotel != nil {
		return visitor.VisitExotel(l.Exotel)
	}
	if l.SipTrunk != nil {
		return visitor.VisitSipTrunk(l.SipTrunk)
	}
	if l.Twilio != nil {
		return visitor.VisitTwilio(l.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", l)
}

func (l *ListPhoneNumbersResponseItem) validate() error {
	if l == nil {
		return fmt.Errorf("type %T is nil", l)
	}
	var fields []string
	if l.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if l.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if l.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if l.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", l, l.Provider)
		}
		return fmt.Errorf("type %T is empty", l)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", l, fields)
	}
	if l.Provider != "" {
		field := fields[0]
		if l.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				l,
				l.Provider,
				l,
			)
		}
	}
	return nil
}

// Create Phone Request Information
type CreatePhoneNumbersRequestBody struct {
	Provider string
	Exotel   *types.CreateExotelPhoneNumberRequest
	SipTrunk *types.CreateSipTrunkPhoneNumberRequest
	Twilio   *types.CreateTwilioPhoneNumberRequest
}

func (c *CreatePhoneNumbersRequestBody) GetExotel() *types.CreateExotelPhoneNumberRequest {
	if c == nil {
		return nil
	}
	if c.Exotel == nil {
		return nil
	}
	return c.Exotel
}

func (c *CreatePhoneNumbersRequestBody) GetSipTrunk() *types.CreateSipTrunkPhoneNumberRequest {
	if c == nil {
		return nil
	}
	if c.SipTrunk == nil {
		return nil
	}
	return c.SipTrunk
}

func (c *CreatePhoneNumbersRequestBody) GetTwilio() *types.CreateTwilioPhoneNumberRequest {
	if c == nil {
		return nil
	}
	if c.Twilio == nil {
		return nil
	}
	return c.Twilio
}

func (c *CreatePhoneNumbersRequestBody) GetProvider() string {
	if c == nil {
		return ""
	}
	return c.Provider
}

func (c *CreatePhoneNumbersRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", c)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(types.CreateExotelPhoneNumberRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Exotel = value
	case "sip_trunk":
		value := new(types.CreateSipTrunkPhoneNumberRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.SipTrunk = value
	case "twilio":
		value := new(types.CreateTwilioPhoneNumberRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Twilio = value
	}
	return nil
}

func (c CreatePhoneNumbersRequestBody) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(c.Exotel, "provider", "exotel")
	}
	if c.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(c.SipTrunk, "provider", "sip_trunk")
	}
	if c.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(c.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreatePhoneNumbersRequestBodyVisitor interface {
	VisitExotel(*types.CreateExotelPhoneNumberRequest) error
	VisitSipTrunk(*types.CreateSipTrunkPhoneNumberRequest) error
	VisitTwilio(*types.CreateTwilioPhoneNumberRequest) error
}

func (c *CreatePhoneNumbersRequestBody) Accept(visitor CreatePhoneNumbersRequestBodyVisitor) error {
	if c.Exotel != nil {
		return visitor.VisitExotel(c.Exotel)
	}
	if c.SipTrunk != nil {
		return visitor.VisitSipTrunk(c.SipTrunk)
	}
	if c.Twilio != nil {
		return visitor.VisitTwilio(c.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreatePhoneNumbersRequestBody) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if c.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if c.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if c.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Provider)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Provider != "" {
		field := fields[0]
		if c.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Provider,
				c,
			)
		}
	}
	return nil
}

type GetPhoneNumbersResponse struct {
	Provider string
	Exotel   *types.GetPhoneNumberExotelResponseModel
	SipTrunk *types.GetPhoneNumberSipTrunkResponseModel
	Twilio   *types.GetPhoneNumberTwilioResponseModel
}

func (g *GetPhoneNumbersResponse) GetExotel() *types.GetPhoneNumberExotelResponseModel {
	if g == nil {
		return nil
	}
	if g.Exotel == nil {
		return nil
	}
	return g.Exotel
}

func (g *GetPhoneNumbersResponse) GetSipTrunk() *types.GetPhoneNumberSipTrunkResponseModel {
	if g == nil {
		return nil
	}
	if g.SipTrunk == nil {
		return nil
	}
	return g.SipTrunk
}

func (g *GetPhoneNumbersResponse) GetTwilio() *types.GetPhoneNumberTwilioResponseModel {
	if g == nil {
		return nil
	}
	if g.Twilio == nil {
		return nil
	}
	return g.Twilio
}

func (g *GetPhoneNumbersResponse) GetProvider() string {
	if g == nil {
		return ""
	}
	return g.Provider
}

func (g *GetPhoneNumbersResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", g)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(types.GetPhoneNumberExotelResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Exotel = value
	case "sip_trunk":
		value := new(types.GetPhoneNumberSipTrunkResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.SipTrunk = value
	case "twilio":
		value := new(types.GetPhoneNumberTwilioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Twilio = value
	}
	return nil
}

func (g GetPhoneNumbersResponse) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(g.Exotel, "provider", "exotel")
	}
	if g.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(g.SipTrunk, "provider", "sip_trunk")
	}
	if g.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(g.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetPhoneNumbersResponseVisitor interface {
	VisitExotel(*types.GetPhoneNumberExotelResponseModel) error
	VisitSipTrunk(*types.GetPhoneNumberSipTrunkResponseModel) error
	VisitTwilio(*types.GetPhoneNumberTwilioResponseModel) error
}

func (g *GetPhoneNumbersResponse) Accept(visitor GetPhoneNumbersResponseVisitor) error {
	if g.Exotel != nil {
		return visitor.VisitExotel(g.Exotel)
	}
	if g.SipTrunk != nil {
		return visitor.VisitSipTrunk(g.SipTrunk)
	}
	if g.Twilio != nil {
		return visitor.VisitTwilio(g.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetPhoneNumbersResponse) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if g.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if g.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if g.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Provider)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Provider != "" {
		field := fields[0]
		if g.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Provider,
				g,
			)
		}
	}
	return nil
}

type UpdatePhoneNumbersResponse struct {
	Provider string
	Exotel   *types.GetPhoneNumberExotelResponseModel
	SipTrunk *types.GetPhoneNumberSipTrunkResponseModel
	Twilio   *types.GetPhoneNumberTwilioResponseModel
}

func (u *UpdatePhoneNumbersResponse) GetExotel() *types.GetPhoneNumberExotelResponseModel {
	if u == nil {
		return nil
	}
	if u.Exotel == nil {
		return nil
	}
	return u.Exotel
}

func (u *UpdatePhoneNumbersResponse) GetSipTrunk() *types.GetPhoneNumberSipTrunkResponseModel {
	if u == nil {
		return nil
	}
	if u.SipTrunk == nil {
		return nil
	}
	return u.SipTrunk
}

func (u *UpdatePhoneNumbersResponse) GetTwilio() *types.GetPhoneNumberTwilioResponseModel {
	if u == nil {
		return nil
	}
	if u.Twilio == nil {
		return nil
	}
	return u.Twilio
}

func (u *UpdatePhoneNumbersResponse) GetProvider() string {
	if u == nil {
		return ""
	}
	return u.Provider
}

func (u *UpdatePhoneNumbersResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Provider string `json:"provider"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Provider = unmarshaler.Provider
	if unmarshaler.Provider == "" {
		return fmt.Errorf("%T did not include discriminant provider", u)
	}
	switch unmarshaler.Provider {
	case "exotel":
		value := new(types.GetPhoneNumberExotelResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Exotel = value
	case "sip_trunk":
		value := new(types.GetPhoneNumberSipTrunkResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.SipTrunk = value
	case "twilio":
		value := new(types.GetPhoneNumberTwilioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Twilio = value
	}
	return nil
}

func (u UpdatePhoneNumbersResponse) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Exotel != nil {
		return internal.MarshalJSONWithDiscriminant(u.Exotel, "provider", "exotel")
	}
	if u.SipTrunk != nil {
		return internal.MarshalJSONWithDiscriminant(u.SipTrunk, "provider", "sip_trunk")
	}
	if u.Twilio != nil {
		return internal.MarshalJSONWithDiscriminant(u.Twilio, "provider", "twilio")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdatePhoneNumbersResponseVisitor interface {
	VisitExotel(*types.GetPhoneNumberExotelResponseModel) error
	VisitSipTrunk(*types.GetPhoneNumberSipTrunkResponseModel) error
	VisitTwilio(*types.GetPhoneNumberTwilioResponseModel) error
}

func (u *UpdatePhoneNumbersResponse) Accept(visitor UpdatePhoneNumbersResponseVisitor) error {
	if u.Exotel != nil {
		return visitor.VisitExotel(u.Exotel)
	}
	if u.SipTrunk != nil {
		return visitor.VisitSipTrunk(u.SipTrunk)
	}
	if u.Twilio != nil {
		return visitor.VisitTwilio(u.Twilio)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdatePhoneNumbersResponse) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Exotel != nil {
		fields = append(fields, "exotel")
	}
	if u.SipTrunk != nil {
		fields = append(fields, "sip_trunk")
	}
	if u.Twilio != nil {
		fields = append(fields, "twilio")
	}
	if len(fields) == 0 {
		if u.Provider != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Provider)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Provider != "" {
		field := fields[0]
		if u.Provider != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Provider,
				u,
			)
		}
	}
	return nil
}

type UpdatePhoneNumberRequest struct {
	// The phone number ID. This is returned when a phone number is imported.
	PhoneNumberId       string                                    `json:"-" url:"-"`
	AgentId             *string                                   `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Label               *string                                   `json:"label,omitempty" xml:"label,omitempty"`
	InboundTrunkConfig  *types.InboundSipTrunkConfigRequestModel  `json:"inbound_trunk_config,omitempty" xml:"inbound_trunk_config,omitempty"`
	OutboundTrunkConfig *types.OutboundSipTrunkConfigRequestModel `json:"outbound_trunk_config,omitempty" xml:"outbound_trunk_config,omitempty"`
	LivekitStack        *types.LivekitStackType                   `json:"livekit_stack,omitempty" xml:"livekit_stack,omitempty"`
	StoreSipMessages    *bool                                     `json:"store_sip_messages,omitempty" xml:"store_sip_messages,omitempty"`
	// Environment to use for resolving environment variables on calls to this number.
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// Agent branch to use for calls to this number.
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
}

func (u *UpdatePhoneNumberRequest) GetPhoneNumberId() (value string) {
	if u == nil {
		return
	}
	return u.PhoneNumberId
}

func (u *UpdatePhoneNumberRequest) GetAgentId() (value *string) {
	if u == nil {
		return
	}
	return u.AgentId
}

func (u *UpdatePhoneNumberRequest) GetLabel() (value *string) {
	if u == nil {
		return
	}
	return u.Label
}

func (u *UpdatePhoneNumberRequest) GetInboundTrunkConfig() (value *types.InboundSipTrunkConfigRequestModel) {
	if u == nil {
		return
	}
	return u.InboundTrunkConfig
}

func (u *UpdatePhoneNumberRequest) GetOutboundTrunkConfig() (value *types.OutboundSipTrunkConfigRequestModel) {
	if u == nil {
		return
	}
	return u.OutboundTrunkConfig
}

func (u *UpdatePhoneNumberRequest) GetLivekitStack() (value *types.LivekitStackType) {
	if u == nil {
		return
	}
	return u.LivekitStack
}

func (u *UpdatePhoneNumberRequest) GetStoreSipMessages() (value *bool) {
	if u == nil {
		return
	}
	return u.StoreSipMessages
}

func (u *UpdatePhoneNumberRequest) GetEnvironment() (value *string) {
	if u == nil {
		return
	}
	return u.Environment
}

func (u *UpdatePhoneNumberRequest) GetBranchId() (value *string) {
	if u == nil {
		return
	}
	return u.BranchId
}

func (u *UpdatePhoneNumberRequest) SetPhoneNumberId(phoneNumberId string) {
	u.PhoneNumberId = phoneNumberId
}

func (u *UpdatePhoneNumberRequest) SetAgentId(agentId *string) {
	u.AgentId = agentId
}

func (u *UpdatePhoneNumberRequest) SetLabel(label *string) {
	u.Label = label
}

func (u *UpdatePhoneNumberRequest) SetInboundTrunkConfig(inboundTrunkConfig *types.InboundSipTrunkConfigRequestModel) {
	u.InboundTrunkConfig = inboundTrunkConfig
}

func (u *UpdatePhoneNumberRequest) SetOutboundTrunkConfig(outboundTrunkConfig *types.OutboundSipTrunkConfigRequestModel) {
	u.OutboundTrunkConfig = outboundTrunkConfig
}

func (u *UpdatePhoneNumberRequest) SetLivekitStack(livekitStack *types.LivekitStackType) {
	u.LivekitStack = livekitStack
}

func (u *UpdatePhoneNumberRequest) SetStoreSipMessages(storeSipMessages *bool) {
	u.StoreSipMessages = storeSipMessages
}

func (u *UpdatePhoneNumberRequest) SetEnvironment(environment *string) {
	u.Environment = environment
}

func (u *UpdatePhoneNumberRequest) SetBranchId(branchId *string) {
	u.BranchId = branchId
}

func (u *UpdatePhoneNumberRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdatePhoneNumberRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdatePhoneNumberRequest to string"
	}
	return string(jsonData)
}

func (u *UpdatePhoneNumberRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdatePhoneNumberRequest) MarshalJSON() ([]byte, error) {
	type embed UpdatePhoneNumberRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

var (
	listPhoneNumbersRequestFieldProvider = big.NewInt(1 << 0)
	listPhoneNumbersRequestFieldAgentId  = big.NewInt(1 << 1)
	listPhoneNumbersRequestFieldBranchId = big.NewInt(1 << 2)
)

type ListPhoneNumbersRequest struct {
	Provider *types.TelephonyProvider `json:"-" url:"provider,omitempty"`
	AgentId  *string                  `json:"-" url:"agent_id,omitempty"`
	BranchId *string                  `json:"-" url:"branch_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListPhoneNumbersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetProvider sets the Provider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListPhoneNumbersRequest) SetProvider(provider *types.TelephonyProvider) {
	l.Provider = provider
	l.require(listPhoneNumbersRequestFieldProvider)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListPhoneNumbersRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listPhoneNumbersRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListPhoneNumbersRequest) SetBranchId(branchId *string) {
	l.BranchId = branchId
	l.require(listPhoneNumbersRequestFieldBranchId)
}

var (
	listV2PhoneNumbersRequestFieldPageSize         = big.NewInt(1 << 0)
	listV2PhoneNumbersRequestFieldSearch           = big.NewInt(1 << 1)
	listV2PhoneNumbersRequestFieldLabel            = big.NewInt(1 << 2)
	listV2PhoneNumbersRequestFieldPhoneNumber      = big.NewInt(1 << 3)
	listV2PhoneNumbersRequestFieldProvider         = big.NewInt(1 << 4)
	listV2PhoneNumbersRequestFieldSupportsOutbound = big.NewInt(1 << 5)
	listV2PhoneNumbersRequestFieldAgentId          = big.NewInt(1 << 6)
	listV2PhoneNumbersRequestFieldBranchId         = big.NewInt(1 << 7)
	listV2PhoneNumbersRequestFieldSortBy           = big.NewInt(1 << 8)
	listV2PhoneNumbersRequestFieldSortDirection    = big.NewInt(1 << 9)
	listV2PhoneNumbersRequestFieldCursor           = big.NewInt(1 << 10)
)

type ListV2PhoneNumbersRequest struct {
	PageSize         *int                     `json:"-" url:"page_size,omitempty"`
	Search           *string                  `json:"-" url:"search,omitempty"`
	Label            *string                  `json:"-" url:"label,omitempty"`
	PhoneNumber      *string                  `json:"-" url:"phone_number,omitempty"`
	Provider         *types.TelephonyProvider `json:"-" url:"provider,omitempty"`
	SupportsOutbound *bool                    `json:"-" url:"supports_outbound,omitempty"`
	AgentId          *string                  `json:"-" url:"agent_id,omitempty"`
	BranchId         *string                  `json:"-" url:"branch_id,omitempty"`
	SortBy           *types.PhoneNumberSortBy `json:"-" url:"sort_by,omitempty"`
	SortDirection    *types.SortDirection     `json:"-" url:"sort_direction,omitempty"`
	Cursor           *string                  `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListV2PhoneNumbersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listV2PhoneNumbersRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listV2PhoneNumbersRequestFieldSearch)
}

// SetLabel sets the Label field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetLabel(label *string) {
	l.Label = label
	l.require(listV2PhoneNumbersRequestFieldLabel)
}

// SetPhoneNumber sets the PhoneNumber field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetPhoneNumber(phoneNumber *string) {
	l.PhoneNumber = phoneNumber
	l.require(listV2PhoneNumbersRequestFieldPhoneNumber)
}

// SetProvider sets the Provider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetProvider(provider *types.TelephonyProvider) {
	l.Provider = provider
	l.require(listV2PhoneNumbersRequestFieldProvider)
}

// SetSupportsOutbound sets the SupportsOutbound field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetSupportsOutbound(supportsOutbound *bool) {
	l.SupportsOutbound = supportsOutbound
	l.require(listV2PhoneNumbersRequestFieldSupportsOutbound)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listV2PhoneNumbersRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetBranchId(branchId *string) {
	l.BranchId = branchId
	l.require(listV2PhoneNumbersRequestFieldBranchId)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetSortBy(sortBy *types.PhoneNumberSortBy) {
	l.SortBy = sortBy
	l.require(listV2PhoneNumbersRequestFieldSortBy)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listV2PhoneNumbersRequestFieldSortDirection)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListV2PhoneNumbersRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listV2PhoneNumbersRequestFieldCursor)
}

var (
	getSipMessagesPhoneNumbersRequestFieldPageSize = big.NewInt(1 << 0)
	getSipMessagesPhoneNumbersRequestFieldCursor   = big.NewInt(1 << 1)
)

type GetSipMessagesPhoneNumbersRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSipMessagesPhoneNumbersRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSipMessagesPhoneNumbersRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getSipMessagesPhoneNumbersRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSipMessagesPhoneNumbersRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getSipMessagesPhoneNumbersRequestFieldCursor)
}
