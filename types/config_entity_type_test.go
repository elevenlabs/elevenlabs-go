package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConfigEntityTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConfigEntityType
		wire    string
	}
	cases := []enumCase{
		{variant: ConfigEntityTypeName, wire: "\"name\""},
		{variant: ConfigEntityTypeNameNameGiven, wire: "\"name.name_given\""},
		{variant: ConfigEntityTypeNameNameFamily, wire: "\"name.name_family\""},
		{variant: ConfigEntityTypeNameNameOther, wire: "\"name.name_other\""},
		{variant: ConfigEntityTypeEmailAddress, wire: "\"email_address\""},
		{variant: ConfigEntityTypeContactNumber, wire: "\"contact_number\""},
		{variant: ConfigEntityTypeDob, wire: "\"dob\""},
		{variant: ConfigEntityTypeAge, wire: "\"age\""},
		{variant: ConfigEntityTypeReligiousBelief, wire: "\"religious_belief\""},
		{variant: ConfigEntityTypePoliticalOpinion, wire: "\"political_opinion\""},
		{variant: ConfigEntityTypeSexualOrientation, wire: "\"sexual_orientation\""},
		{variant: ConfigEntityTypeEthnicityRace, wire: "\"ethnicity_race\""},
		{variant: ConfigEntityTypeMaritalStatus, wire: "\"marital_status\""},
		{variant: ConfigEntityTypeOccupation, wire: "\"occupation\""},
		{variant: ConfigEntityTypePhysicalAttribute, wire: "\"physical_attribute\""},
		{variant: ConfigEntityTypeLanguage, wire: "\"language\""},
		{variant: ConfigEntityTypeCriminalRecord, wire: "\"criminal_record\""},
		{variant: ConfigEntityTypeUsername, wire: "\"username\""},
		{variant: ConfigEntityTypePassword, wire: "\"password\""},
		{variant: ConfigEntityTypeUrl, wire: "\"url\""},
		{variant: ConfigEntityTypeOrganization, wire: "\"organization\""},
		{variant: ConfigEntityTypeFinancialId, wire: "\"financial_id\""},
		{variant: ConfigEntityTypeFinancialIdPaymentCard, wire: "\"financial_id.payment_card\""},
		{variant: ConfigEntityTypeFinancialIdPaymentCardPaymentCardNumber, wire: "\"financial_id.payment_card.payment_card_number\""},
		{variant: ConfigEntityTypeFinancialIdPaymentCardPaymentCardExpirationDate, wire: "\"financial_id.payment_card.payment_card_expiration_date\""},
		{variant: ConfigEntityTypeFinancialIdPaymentCardPaymentCardCvv, wire: "\"financial_id.payment_card.payment_card_cvv\""},
		{variant: ConfigEntityTypeFinancialIdBankAccount, wire: "\"financial_id.bank_account\""},
		{variant: ConfigEntityTypeFinancialIdBankAccountBankAccountNumber, wire: "\"financial_id.bank_account.bank_account_number\""},
		{variant: ConfigEntityTypeFinancialIdBankAccountBankRoutingNumber, wire: "\"financial_id.bank_account.bank_routing_number\""},
		{variant: ConfigEntityTypeFinancialIdBankAccountSwiftBicCode, wire: "\"financial_id.bank_account.swift_bic_code\""},
		{variant: ConfigEntityTypeFinancialIdFinancialIdOther, wire: "\"financial_id.financial_id_other\""},
		{variant: ConfigEntityTypeLocation, wire: "\"location\""},
		{variant: ConfigEntityTypeLocationLocationAddress, wire: "\"location.location_address\""},
		{variant: ConfigEntityTypeLocationLocationCity, wire: "\"location.location_city\""},
		{variant: ConfigEntityTypeLocationLocationPostalCode, wire: "\"location.location_postal_code\""},
		{variant: ConfigEntityTypeLocationLocationCoordinate, wire: "\"location.location_coordinate\""},
		{variant: ConfigEntityTypeLocationLocationState, wire: "\"location.location_state\""},
		{variant: ConfigEntityTypeLocationLocationCountry, wire: "\"location.location_country\""},
		{variant: ConfigEntityTypeLocationLocationOther, wire: "\"location.location_other\""},
		{variant: ConfigEntityTypeDate, wire: "\"date\""},
		{variant: ConfigEntityTypeDateInterval, wire: "\"date_interval\""},
		{variant: ConfigEntityTypeUniqueId, wire: "\"unique_id\""},
		{variant: ConfigEntityTypeUniqueIdGovernmentIssuedId, wire: "\"unique_id.government_issued_id\""},
		{variant: ConfigEntityTypeUniqueIdAccountNumber, wire: "\"unique_id.account_number\""},
		{variant: ConfigEntityTypeUniqueIdVehicleId, wire: "\"unique_id.vehicle_id\""},
		{variant: ConfigEntityTypeUniqueIdHealthcareNumber, wire: "\"unique_id.healthcare_number\""},
		{variant: ConfigEntityTypeUniqueIdHealthcareNumberMedicalRecordNumber, wire: "\"unique_id.healthcare_number.medical_record_number\""},
		{variant: ConfigEntityTypeUniqueIdHealthcareNumberHealthPlanBeneficiaryNumber, wire: "\"unique_id.healthcare_number.health_plan_beneficiary_number\""},
		{variant: ConfigEntityTypeUniqueIdDeviceId, wire: "\"unique_id.device_id\""},
		{variant: ConfigEntityTypeUniqueIdUniqueIdOther, wire: "\"unique_id.unique_id_other\""},
		{variant: ConfigEntityTypeMedical, wire: "\"medical\""},
		{variant: ConfigEntityTypeMedicalMedicalCondition, wire: "\"medical.medical_condition\""},
		{variant: ConfigEntityTypeMedicalMedication, wire: "\"medical.medication\""},
		{variant: ConfigEntityTypeMedicalMedicalProcedure, wire: "\"medical.medical_procedure\""},
		{variant: ConfigEntityTypeMedicalMedicalMeasurement, wire: "\"medical.medical_measurement\""},
		{variant: ConfigEntityTypeMedicalMedicalOther, wire: "\"medical.medical_other\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConfigEntityType
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
