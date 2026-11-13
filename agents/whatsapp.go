package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost struct {
	WhatsappPhoneNumberId                             string                                              `json:"whatsapp_phone_number_id" xml:"whatsapp_phone_number_id" required:"true"`
	WhatsappUserId                                    string                                              `json:"whatsapp_user_id" xml:"whatsapp_user_id" required:"true"`
	WhatsappCallPermissionRequestTemplateName         string                                              `json:"whatsapp_call_permission_request_template_name" xml:"whatsapp_call_permission_request_template_name" required:"true"`
	WhatsappCallPermissionRequestTemplateLanguageCode string                                              `json:"whatsapp_call_permission_request_template_language_code" xml:"whatsapp_call_permission_request_template_language_code" required:"true"`
	AgentId                                           string                                              `json:"agent_id" xml:"agent_id" required:"true"`
	ConversationInitiationClientData                  *types.ConversationInitiationClientDataRequestInput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetWhatsappPhoneNumberId() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappPhoneNumberId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetWhatsappUserId() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappUserId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetWhatsappCallPermissionRequestTemplateName() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappCallPermissionRequestTemplateName
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetWhatsappCallPermissionRequestTemplateLanguageCode() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappCallPermissionRequestTemplateLanguageCode
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetWhatsappPhoneNumberId(whatsappPhoneNumberId string) {
	b.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetWhatsappUserId(whatsappUserId string) {
	b.WhatsappUserId = whatsappUserId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetWhatsappCallPermissionRequestTemplateName(whatsappCallPermissionRequestTemplateName string) {
	b.WhatsappCallPermissionRequestTemplateName = whatsappCallPermissionRequestTemplateName
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetWhatsappCallPermissionRequestTemplateLanguageCode(whatsappCallPermissionRequestTemplateLanguageCode string) {
	b.WhatsappCallPermissionRequestTemplateLanguageCode = whatsappCallPermissionRequestTemplateLanguageCode
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost to string"
	}
	return string(jsonData)
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost) MarshalJSON() ([]byte, error) {
	type embed BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost struct {
	WhatsappPhoneNumberId            string                                                                                   `json:"whatsapp_phone_number_id" xml:"whatsapp_phone_number_id" required:"true"`
	WhatsappUserId                   string                                                                                   `json:"whatsapp_user_id" xml:"whatsapp_user_id" required:"true"`
	TemplateName                     string                                                                                   `json:"template_name" xml:"template_name" required:"true"`
	TemplateLanguageCode             string                                                                                   `json:"template_language_code" xml:"template_language_code" required:"true"`
	TemplateParams                   []*SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem `json:"template_params" xml:"template_params" required:"true"`
	AgentId                          string                                                                                   `json:"agent_id" xml:"agent_id" required:"true"`
	ConversationInitiationClientData *types.ConversationInitiationClientDataRequestInput                                      `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetWhatsappPhoneNumberId() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappPhoneNumberId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetWhatsappUserId() (value string) {
	if b == nil {
		return
	}
	return b.WhatsappUserId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetTemplateName() (value string) {
	if b == nil {
		return
	}
	return b.TemplateName
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetTemplateLanguageCode() (value string) {
	if b == nil {
		return
	}
	return b.TemplateLanguageCode
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetTemplateParams() (value []*SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) {
	if b == nil {
		return
	}
	return b.TemplateParams
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetConversationInitiationClientData() (value *types.ConversationInitiationClientDataRequestInput) {
	if b == nil {
		return
	}
	return b.ConversationInitiationClientData
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetWhatsappPhoneNumberId(whatsappPhoneNumberId string) {
	b.WhatsappPhoneNumberId = whatsappPhoneNumberId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetWhatsappUserId(whatsappUserId string) {
	b.WhatsappUserId = whatsappUserId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetTemplateName(templateName string) {
	b.TemplateName = templateName
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetTemplateLanguageCode(templateLanguageCode string) {
	b.TemplateLanguageCode = templateLanguageCode
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetTemplateParams(templateParams []*SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) {
	b.TemplateParams = templateParams
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) SetConversationInitiationClientData(conversationInitiationClientData *types.ConversationInitiationClientDataRequestInput) {
	b.ConversationInitiationClientData = conversationInitiationClientData
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost to string"
	}
	return string(jsonData)
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost) MarshalJSON() ([]byte, error) {
	type embed BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem struct {
	Type   string
	Body   *types.WhatsAppTemplateBodyComponentParams
	Button *types.WhatsAppTemplateButtonComponentParams
	Header *types.WhatsAppTemplateHeaderComponentParams
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) GetBody() *types.WhatsAppTemplateBodyComponentParams {
	if s == nil {
		return nil
	}
	if s.Body == nil {
		return nil
	}
	return s.Body
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) GetButton() *types.WhatsAppTemplateButtonComponentParams {
	if s == nil {
		return nil
	}
	if s.Button == nil {
		return nil
	}
	return s.Button
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) GetHeader() *types.WhatsAppTemplateHeaderComponentParams {
	if s == nil {
		return nil
	}
	if s.Header == nil {
		return nil
	}
	return s.Header
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) GetType() string {
	if s == nil {
		return ""
	}
	return s.Type
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", s)
	}
	switch unmarshaler.Type {
	case "body":
		value := new(types.WhatsAppTemplateBodyComponentParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Body = value
	case "button":
		value := new(types.WhatsAppTemplateButtonComponentParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Button = value
	case "header":
		value := new(types.WhatsAppTemplateHeaderComponentParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Header = value
	}
	return nil
}

func (s SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Body != nil {
		return internal.MarshalJSONWithDiscriminant(s.Body, "type", "body")
	}
	if s.Button != nil {
		return internal.MarshalJSONWithDiscriminant(s.Button, "type", "button")
	}
	if s.Header != nil {
		return internal.MarshalJSONWithDiscriminant(s.Header, "type", "header")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItemVisitor interface {
	VisitBody(*types.WhatsAppTemplateBodyComponentParams) error
	VisitButton(*types.WhatsAppTemplateButtonComponentParams) error
	VisitHeader(*types.WhatsAppTemplateHeaderComponentParams) error
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) Accept(visitor SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItemVisitor) error {
	if s.Body != nil {
		return visitor.VisitBody(s.Body)
	}
	if s.Button != nil {
		return visitor.VisitButton(s.Button)
	}
	if s.Header != nil {
		return visitor.VisitHeader(s.Header)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePostTemplateParamsItem) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Body != nil {
		fields = append(fields, "body")
	}
	if s.Button != nil {
		fields = append(fields, "button")
	}
	if s.Header != nil {
		fields = append(fields, "header")
	}
	if len(fields) == 0 {
		if s.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Type)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Type != "" {
		field := fields[0]
		if s.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Type,
				s,
			)
		}
	}
	return nil
}
