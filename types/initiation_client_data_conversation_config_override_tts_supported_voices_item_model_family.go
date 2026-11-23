package types

import (
	"fmt"
)

type InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily string

const (
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyTurbo            InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "turbo"
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyFlash            InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "flash"
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyMultilingual     InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "multilingual"
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV3Conversational InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "v3_conversational"
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV4               InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "v4"
	InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV4Turbo          InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily = "v4_turbo"
)

func NewInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyFromString(s string) (InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily, error) {
	switch s {
	case "turbo":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyTurbo, nil
	case "flash":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyFlash, nil
	case "multilingual":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyMultilingual, nil
	case "v3_conversational":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV3Conversational, nil
	case "v4":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV4, nil
	case "v4_turbo":
		return InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamilyV4Turbo, nil
	}
	var t InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily) Ptr() *InitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily {
	return &e
}
