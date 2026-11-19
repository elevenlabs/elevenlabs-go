package types

import (
	"fmt"
)

// Entity types for the API configuration.
//
// This enum contains all valid entity type configurations that users can specify:
// - Parent types (e.g., "name", "financial_id") that expand to all subtypes
// - Specific subtypes using dot notation (e.g., "name.full_name")
// - Standalone terminal types (e.g., "email_address")
//
// When converted for service use, parent types expand to all their terminal subtypes.
type ConfigEntityType string

const (
	ConfigEntityTypeName                                                ConfigEntityType = "name"
	ConfigEntityTypeNameNameGiven                                       ConfigEntityType = "name.name_given"
	ConfigEntityTypeNameNameFamily                                      ConfigEntityType = "name.name_family"
	ConfigEntityTypeNameNameOther                                       ConfigEntityType = "name.name_other"
	ConfigEntityTypeEmailAddress                                        ConfigEntityType = "email_address"
	ConfigEntityTypeContactNumber                                       ConfigEntityType = "contact_number"
	ConfigEntityTypeDob                                                 ConfigEntityType = "dob"
	ConfigEntityTypeAge                                                 ConfigEntityType = "age"
	ConfigEntityTypeReligiousBelief                                     ConfigEntityType = "religious_belief"
	ConfigEntityTypePoliticalOpinion                                    ConfigEntityType = "political_opinion"
	ConfigEntityTypeSexualOrientation                                   ConfigEntityType = "sexual_orientation"
	ConfigEntityTypeEthnicityRace                                       ConfigEntityType = "ethnicity_race"
	ConfigEntityTypeMaritalStatus                                       ConfigEntityType = "marital_status"
	ConfigEntityTypeOccupation                                          ConfigEntityType = "occupation"
	ConfigEntityTypePhysicalAttribute                                   ConfigEntityType = "physical_attribute"
	ConfigEntityTypeLanguage                                            ConfigEntityType = "language"
	ConfigEntityTypeCriminalRecord                                      ConfigEntityType = "criminal_record"
	ConfigEntityTypeUsername                                            ConfigEntityType = "username"
	ConfigEntityTypePassword                                            ConfigEntityType = "password"
	ConfigEntityTypeUrl                                                 ConfigEntityType = "url"
	ConfigEntityTypeOrganization                                        ConfigEntityType = "organization"
	ConfigEntityTypeFinancialId                                         ConfigEntityType = "financial_id"
	ConfigEntityTypeFinancialIdPaymentCard                              ConfigEntityType = "financial_id.payment_card"
	ConfigEntityTypeFinancialIdPaymentCardPaymentCardNumber             ConfigEntityType = "financial_id.payment_card.payment_card_number"
	ConfigEntityTypeFinancialIdPaymentCardPaymentCardExpirationDate     ConfigEntityType = "financial_id.payment_card.payment_card_expiration_date"
	ConfigEntityTypeFinancialIdPaymentCardPaymentCardCvv                ConfigEntityType = "financial_id.payment_card.payment_card_cvv"
	ConfigEntityTypeFinancialIdBankAccount                              ConfigEntityType = "financial_id.bank_account"
	ConfigEntityTypeFinancialIdBankAccountBankAccountNumber             ConfigEntityType = "financial_id.bank_account.bank_account_number"
	ConfigEntityTypeFinancialIdBankAccountBankRoutingNumber             ConfigEntityType = "financial_id.bank_account.bank_routing_number"
	ConfigEntityTypeFinancialIdBankAccountSwiftBicCode                  ConfigEntityType = "financial_id.bank_account.swift_bic_code"
	ConfigEntityTypeFinancialIdFinancialIdOther                         ConfigEntityType = "financial_id.financial_id_other"
	ConfigEntityTypeLocation                                            ConfigEntityType = "location"
	ConfigEntityTypeLocationLocationAddress                             ConfigEntityType = "location.location_address"
	ConfigEntityTypeLocationLocationCity                                ConfigEntityType = "location.location_city"
	ConfigEntityTypeLocationLocationPostalCode                          ConfigEntityType = "location.location_postal_code"
	ConfigEntityTypeLocationLocationCoordinate                          ConfigEntityType = "location.location_coordinate"
	ConfigEntityTypeLocationLocationState                               ConfigEntityType = "location.location_state"
	ConfigEntityTypeLocationLocationCountry                             ConfigEntityType = "location.location_country"
	ConfigEntityTypeLocationLocationOther                               ConfigEntityType = "location.location_other"
	ConfigEntityTypeDate                                                ConfigEntityType = "date"
	ConfigEntityTypeDateInterval                                        ConfigEntityType = "date_interval"
	ConfigEntityTypeUniqueId                                            ConfigEntityType = "unique_id"
	ConfigEntityTypeUniqueIdGovernmentIssuedId                          ConfigEntityType = "unique_id.government_issued_id"
	ConfigEntityTypeUniqueIdAccountNumber                               ConfigEntityType = "unique_id.account_number"
	ConfigEntityTypeUniqueIdVehicleId                                   ConfigEntityType = "unique_id.vehicle_id"
	ConfigEntityTypeUniqueIdHealthcareNumber                            ConfigEntityType = "unique_id.healthcare_number"
	ConfigEntityTypeUniqueIdHealthcareNumberMedicalRecordNumber         ConfigEntityType = "unique_id.healthcare_number.medical_record_number"
	ConfigEntityTypeUniqueIdHealthcareNumberHealthPlanBeneficiaryNumber ConfigEntityType = "unique_id.healthcare_number.health_plan_beneficiary_number"
	ConfigEntityTypeUniqueIdDeviceId                                    ConfigEntityType = "unique_id.device_id"
	ConfigEntityTypeUniqueIdUniqueIdOther                               ConfigEntityType = "unique_id.unique_id_other"
	ConfigEntityTypeMedical                                             ConfigEntityType = "medical"
	ConfigEntityTypeMedicalMedicalCondition                             ConfigEntityType = "medical.medical_condition"
	ConfigEntityTypeMedicalMedication                                   ConfigEntityType = "medical.medication"
	ConfigEntityTypeMedicalMedicalProcedure                             ConfigEntityType = "medical.medical_procedure"
	ConfigEntityTypeMedicalMedicalMeasurement                           ConfigEntityType = "medical.medical_measurement"
	ConfigEntityTypeMedicalMedicalOther                                 ConfigEntityType = "medical.medical_other"
)

func NewConfigEntityTypeFromString(s string) (ConfigEntityType, error) {
	switch s {
	case "name":
		return ConfigEntityTypeName, nil
	case "name.name_given":
		return ConfigEntityTypeNameNameGiven, nil
	case "name.name_family":
		return ConfigEntityTypeNameNameFamily, nil
	case "name.name_other":
		return ConfigEntityTypeNameNameOther, nil
	case "email_address":
		return ConfigEntityTypeEmailAddress, nil
	case "contact_number":
		return ConfigEntityTypeContactNumber, nil
	case "dob":
		return ConfigEntityTypeDob, nil
	case "age":
		return ConfigEntityTypeAge, nil
	case "religious_belief":
		return ConfigEntityTypeReligiousBelief, nil
	case "political_opinion":
		return ConfigEntityTypePoliticalOpinion, nil
	case "sexual_orientation":
		return ConfigEntityTypeSexualOrientation, nil
	case "ethnicity_race":
		return ConfigEntityTypeEthnicityRace, nil
	case "marital_status":
		return ConfigEntityTypeMaritalStatus, nil
	case "occupation":
		return ConfigEntityTypeOccupation, nil
	case "physical_attribute":
		return ConfigEntityTypePhysicalAttribute, nil
	case "language":
		return ConfigEntityTypeLanguage, nil
	case "criminal_record":
		return ConfigEntityTypeCriminalRecord, nil
	case "username":
		return ConfigEntityTypeUsername, nil
	case "password":
		return ConfigEntityTypePassword, nil
	case "url":
		return ConfigEntityTypeUrl, nil
	case "organization":
		return ConfigEntityTypeOrganization, nil
	case "financial_id":
		return ConfigEntityTypeFinancialId, nil
	case "financial_id.payment_card":
		return ConfigEntityTypeFinancialIdPaymentCard, nil
	case "financial_id.payment_card.payment_card_number":
		return ConfigEntityTypeFinancialIdPaymentCardPaymentCardNumber, nil
	case "financial_id.payment_card.payment_card_expiration_date":
		return ConfigEntityTypeFinancialIdPaymentCardPaymentCardExpirationDate, nil
	case "financial_id.payment_card.payment_card_cvv":
		return ConfigEntityTypeFinancialIdPaymentCardPaymentCardCvv, nil
	case "financial_id.bank_account":
		return ConfigEntityTypeFinancialIdBankAccount, nil
	case "financial_id.bank_account.bank_account_number":
		return ConfigEntityTypeFinancialIdBankAccountBankAccountNumber, nil
	case "financial_id.bank_account.bank_routing_number":
		return ConfigEntityTypeFinancialIdBankAccountBankRoutingNumber, nil
	case "financial_id.bank_account.swift_bic_code":
		return ConfigEntityTypeFinancialIdBankAccountSwiftBicCode, nil
	case "financial_id.financial_id_other":
		return ConfigEntityTypeFinancialIdFinancialIdOther, nil
	case "location":
		return ConfigEntityTypeLocation, nil
	case "location.location_address":
		return ConfigEntityTypeLocationLocationAddress, nil
	case "location.location_city":
		return ConfigEntityTypeLocationLocationCity, nil
	case "location.location_postal_code":
		return ConfigEntityTypeLocationLocationPostalCode, nil
	case "location.location_coordinate":
		return ConfigEntityTypeLocationLocationCoordinate, nil
	case "location.location_state":
		return ConfigEntityTypeLocationLocationState, nil
	case "location.location_country":
		return ConfigEntityTypeLocationLocationCountry, nil
	case "location.location_other":
		return ConfigEntityTypeLocationLocationOther, nil
	case "date":
		return ConfigEntityTypeDate, nil
	case "date_interval":
		return ConfigEntityTypeDateInterval, nil
	case "unique_id":
		return ConfigEntityTypeUniqueId, nil
	case "unique_id.government_issued_id":
		return ConfigEntityTypeUniqueIdGovernmentIssuedId, nil
	case "unique_id.account_number":
		return ConfigEntityTypeUniqueIdAccountNumber, nil
	case "unique_id.vehicle_id":
		return ConfigEntityTypeUniqueIdVehicleId, nil
	case "unique_id.healthcare_number":
		return ConfigEntityTypeUniqueIdHealthcareNumber, nil
	case "unique_id.healthcare_number.medical_record_number":
		return ConfigEntityTypeUniqueIdHealthcareNumberMedicalRecordNumber, nil
	case "unique_id.healthcare_number.health_plan_beneficiary_number":
		return ConfigEntityTypeUniqueIdHealthcareNumberHealthPlanBeneficiaryNumber, nil
	case "unique_id.device_id":
		return ConfigEntityTypeUniqueIdDeviceId, nil
	case "unique_id.unique_id_other":
		return ConfigEntityTypeUniqueIdUniqueIdOther, nil
	case "medical":
		return ConfigEntityTypeMedical, nil
	case "medical.medical_condition":
		return ConfigEntityTypeMedicalMedicalCondition, nil
	case "medical.medication":
		return ConfigEntityTypeMedicalMedication, nil
	case "medical.medical_procedure":
		return ConfigEntityTypeMedicalMedicalProcedure, nil
	case "medical.medical_measurement":
		return ConfigEntityTypeMedicalMedicalMeasurement, nil
	case "medical.medical_other":
		return ConfigEntityTypeMedicalMedicalOther, nil
	}
	var t ConfigEntityType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConfigEntityType) Ptr() *ConfigEntityType {
	return &e
}
