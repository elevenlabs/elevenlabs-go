package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SmbToolConfigParams struct {
	SmbToolType                        string
	CancelCalendarEvent                *CancelCalendarEventParams
	CancelGroupSessionForAll           *CancelGroupSessionForAllParams
	CancelGroupSessionRegistration     *CancelGroupSessionRegistrationParams
	CancelOrder                        *CancelOrderParams
	CheckServiceAvailability           *CheckServiceAvailabilityParams
	CreateAgentProcedure               *CreateAgentProcedureParams
	CreateAgentRule                    *CreateAgentRuleParams
	CreateAsset                        *CreateAssetParams
	CreateClient                       *CreateClientParams
	CreateClientAppointment            *CreateClientAppointmentParams
	CreateClientInteraction            *CreateClientInteractionParams
	CreateHoliday                      *CreateHolidayParams
	CreateLocation                     *CreateLocationParams
	CreateOrder                        *CreateOrderParams
	CreateProduct                      *CreateProductParams
	CreateProductQuoteRequest          *CreateProductQuoteRequestParams
	CreateService                      *CreateServiceParams
	CreateServiceQuoteRequest          *CreateServiceQuoteRequestParams
	CreateStaff                        *CreateStaffParams
	CreateTransferRule                 *CreateTransferRuleParams
	DeleteAgentProcedure               *DeleteAgentProcedureParams
	DeleteAgentRule                    *DeleteAgentRuleParams
	DeleteAsset                        *DeleteAssetParams
	DeleteCalendarEvent                *DeleteCalendarEventParams
	DeleteClient                       *DeleteClientParams
	DeleteClientInteraction            *DeleteClientInteractionParams
	DeleteGroupSession                 *DeleteGroupSessionParams
	DeleteHoliday                      *DeleteHolidayParams
	DeleteLocation                     *DeleteLocationParams
	DeleteProduct                      *DeleteProductParams
	DeleteService                      *DeleteServiceParams
	DeleteStaff                        *DeleteStaffParams
	DeleteTransferRule                 *DeleteTransferRuleParams
	GetAgentConfig                     *GetAgentConfigParams
	GetAnalyticsSummary                *GetAnalyticsSummaryParams
	GetAppointmentByConfirmationNumber *GetAppointmentByConfirmationNumberParams
	GetBookingPageSettings             *GetBookingPageSettingsParams
	GetBookingSlugStatus               *GetBookingSlugStatusParams
	GetClientAppointments              *GetClientAppointmentsParams
	GetClientByPhone                   *GetClientByPhoneParams
	GetClientOrders                    *GetClientOrdersParams
	GetOrderByConfirmationNumber       *GetOrderByConfirmationNumberParams
	GetSchedule                        *GetScheduleParams
	GetWorkspacePlan                   *GetWorkspacePlanParams
	LeaveMessage                       *LeaveMessageParams
	ListAgentProcedures                *ListAgentProceduresParams
	ListAgentReferences                *ListAgentReferencesParams
	ListAgentRules                     *ListAgentRulesParams
	ListAssets                         *ListAssetsParams
	ListCalendarEvents                 *ListCalendarEventsParams
	ListClientInteractions             *ListClientInteractionsParams
	ListClients                        *ListClientsParams
	ListCustomerFacingAgents           *ListCustomerFacingAgentsParams
	ListGroupSessions                  *ListGroupSessionsParams
	ListHolidays                       *ListHolidaysParams
	ListLocations                      *ListLocationsParams
	ListProducts                       *ListProductsParams
	ListServices                       *ListServicesParams
	ListStaff                          *ListStaffParams
	ListTransferRules                  *ListTransferRulesParams
	OptInSmsReminder                   *OptInSmsReminderParams
	OptOutSmsReminder                  *OptOutSmsReminderParams
	OwnerGetClientAppointments         *OwnerGetClientAppointmentsParams
	OwnerListAssets                    *OwnerListAssetsParams
	OwnerListLocations                 *OwnerListLocationsParams
	OwnerListProducts                  *OwnerListProductsParams
	OwnerListServices                  *OwnerListServicesParams
	OwnerListStaff                     *OwnerListStaffParams
	RegisterForGroupSession            *RegisterForGroupSessionParams
	ReportKnowledgeGap                 *ReportKnowledgeGapParams
	RestoreCalendarEvent               *RestoreCalendarEventParams
	ScheduleGroupSession               *ScheduleGroupSessionParams
	SearchClients                      *SearchClientsParams
	SendCustomEmail                    *SendCustomEmailParams
	SetBookingSlug                     *SetBookingSlugParams
	SubmitBusinessInfo                 *SubmitBusinessInfoParams
	TriggerUserVerification            *TriggerUserVerificationParams
	UpdateAgentConfig                  *UpdateAgentConfigParams
	UpdateAgentProcedure               *UpdateAgentProcedureParams
	UpdateAgentRule                    *UpdateAgentRuleParams
	UpdateAsset                        *UpdateAssetParams
	UpdateBookingPageAppearance        *UpdateBookingPageAppearanceParams
	UpdateBookingPageSettings          *UpdateBookingPageSettingsParams
	UpdateBusinessInfo                 *UpdateBusinessInfoParams
	UpdateCalendarEvent                *UpdateCalendarEventParams
	UpdateClient                       *UpdateClientParams
	UpdateGroupSessionSeats            *UpdateGroupSessionSeatsParams
	UpdateHoliday                      *UpdateHolidayParams
	UpdateLocation                     *UpdateLocationParams
	UpdateOrder                        *UpdateOrderParams
	UpdateProduct                      *UpdateProductParams
	UpdateService                      *UpdateServiceParams
	UpdateStaff                        *UpdateStaffParams
	UpdateTransferRule                 *UpdateTransferRuleParams
	ValidateUserVerificationCode       *ValidateUserVerificationCodeParams
}

func (s *SmbToolConfigParams) GetCancelCalendarEvent() *CancelCalendarEventParams {
	if s == nil {
		return nil
	}
	if s.CancelCalendarEvent == nil {
		return nil
	}
	return s.CancelCalendarEvent
}

func (s *SmbToolConfigParams) GetCancelGroupSessionForAll() *CancelGroupSessionForAllParams {
	if s == nil {
		return nil
	}
	if s.CancelGroupSessionForAll == nil {
		return nil
	}
	return s.CancelGroupSessionForAll
}

func (s *SmbToolConfigParams) GetCancelGroupSessionRegistration() *CancelGroupSessionRegistrationParams {
	if s == nil {
		return nil
	}
	if s.CancelGroupSessionRegistration == nil {
		return nil
	}
	return s.CancelGroupSessionRegistration
}

func (s *SmbToolConfigParams) GetCancelOrder() *CancelOrderParams {
	if s == nil {
		return nil
	}
	if s.CancelOrder == nil {
		return nil
	}
	return s.CancelOrder
}

func (s *SmbToolConfigParams) GetCheckServiceAvailability() *CheckServiceAvailabilityParams {
	if s == nil {
		return nil
	}
	if s.CheckServiceAvailability == nil {
		return nil
	}
	return s.CheckServiceAvailability
}

func (s *SmbToolConfigParams) GetCreateAgentProcedure() *CreateAgentProcedureParams {
	if s == nil {
		return nil
	}
	if s.CreateAgentProcedure == nil {
		return nil
	}
	return s.CreateAgentProcedure
}

func (s *SmbToolConfigParams) GetCreateAgentRule() *CreateAgentRuleParams {
	if s == nil {
		return nil
	}
	if s.CreateAgentRule == nil {
		return nil
	}
	return s.CreateAgentRule
}

func (s *SmbToolConfigParams) GetCreateAsset() *CreateAssetParams {
	if s == nil {
		return nil
	}
	if s.CreateAsset == nil {
		return nil
	}
	return s.CreateAsset
}

func (s *SmbToolConfigParams) GetCreateClient() *CreateClientParams {
	if s == nil {
		return nil
	}
	if s.CreateClient == nil {
		return nil
	}
	return s.CreateClient
}

func (s *SmbToolConfigParams) GetCreateClientAppointment() *CreateClientAppointmentParams {
	if s == nil {
		return nil
	}
	if s.CreateClientAppointment == nil {
		return nil
	}
	return s.CreateClientAppointment
}

func (s *SmbToolConfigParams) GetCreateClientInteraction() *CreateClientInteractionParams {
	if s == nil {
		return nil
	}
	if s.CreateClientInteraction == nil {
		return nil
	}
	return s.CreateClientInteraction
}

func (s *SmbToolConfigParams) GetCreateHoliday() *CreateHolidayParams {
	if s == nil {
		return nil
	}
	if s.CreateHoliday == nil {
		return nil
	}
	return s.CreateHoliday
}

func (s *SmbToolConfigParams) GetCreateLocation() *CreateLocationParams {
	if s == nil {
		return nil
	}
	if s.CreateLocation == nil {
		return nil
	}
	return s.CreateLocation
}

func (s *SmbToolConfigParams) GetCreateOrder() *CreateOrderParams {
	if s == nil {
		return nil
	}
	if s.CreateOrder == nil {
		return nil
	}
	return s.CreateOrder
}

func (s *SmbToolConfigParams) GetCreateProduct() *CreateProductParams {
	if s == nil {
		return nil
	}
	if s.CreateProduct == nil {
		return nil
	}
	return s.CreateProduct
}

func (s *SmbToolConfigParams) GetCreateProductQuoteRequest() *CreateProductQuoteRequestParams {
	if s == nil {
		return nil
	}
	if s.CreateProductQuoteRequest == nil {
		return nil
	}
	return s.CreateProductQuoteRequest
}

func (s *SmbToolConfigParams) GetCreateService() *CreateServiceParams {
	if s == nil {
		return nil
	}
	if s.CreateService == nil {
		return nil
	}
	return s.CreateService
}

func (s *SmbToolConfigParams) GetCreateServiceQuoteRequest() *CreateServiceQuoteRequestParams {
	if s == nil {
		return nil
	}
	if s.CreateServiceQuoteRequest == nil {
		return nil
	}
	return s.CreateServiceQuoteRequest
}

func (s *SmbToolConfigParams) GetCreateStaff() *CreateStaffParams {
	if s == nil {
		return nil
	}
	if s.CreateStaff == nil {
		return nil
	}
	return s.CreateStaff
}

func (s *SmbToolConfigParams) GetCreateTransferRule() *CreateTransferRuleParams {
	if s == nil {
		return nil
	}
	if s.CreateTransferRule == nil {
		return nil
	}
	return s.CreateTransferRule
}

func (s *SmbToolConfigParams) GetDeleteAgentProcedure() *DeleteAgentProcedureParams {
	if s == nil {
		return nil
	}
	if s.DeleteAgentProcedure == nil {
		return nil
	}
	return s.DeleteAgentProcedure
}

func (s *SmbToolConfigParams) GetDeleteAgentRule() *DeleteAgentRuleParams {
	if s == nil {
		return nil
	}
	if s.DeleteAgentRule == nil {
		return nil
	}
	return s.DeleteAgentRule
}

func (s *SmbToolConfigParams) GetDeleteAsset() *DeleteAssetParams {
	if s == nil {
		return nil
	}
	if s.DeleteAsset == nil {
		return nil
	}
	return s.DeleteAsset
}

func (s *SmbToolConfigParams) GetDeleteCalendarEvent() *DeleteCalendarEventParams {
	if s == nil {
		return nil
	}
	if s.DeleteCalendarEvent == nil {
		return nil
	}
	return s.DeleteCalendarEvent
}

func (s *SmbToolConfigParams) GetDeleteClient() *DeleteClientParams {
	if s == nil {
		return nil
	}
	if s.DeleteClient == nil {
		return nil
	}
	return s.DeleteClient
}

func (s *SmbToolConfigParams) GetDeleteClientInteraction() *DeleteClientInteractionParams {
	if s == nil {
		return nil
	}
	if s.DeleteClientInteraction == nil {
		return nil
	}
	return s.DeleteClientInteraction
}

func (s *SmbToolConfigParams) GetDeleteGroupSession() *DeleteGroupSessionParams {
	if s == nil {
		return nil
	}
	if s.DeleteGroupSession == nil {
		return nil
	}
	return s.DeleteGroupSession
}

func (s *SmbToolConfigParams) GetDeleteHoliday() *DeleteHolidayParams {
	if s == nil {
		return nil
	}
	if s.DeleteHoliday == nil {
		return nil
	}
	return s.DeleteHoliday
}

func (s *SmbToolConfigParams) GetDeleteLocation() *DeleteLocationParams {
	if s == nil {
		return nil
	}
	if s.DeleteLocation == nil {
		return nil
	}
	return s.DeleteLocation
}

func (s *SmbToolConfigParams) GetDeleteProduct() *DeleteProductParams {
	if s == nil {
		return nil
	}
	if s.DeleteProduct == nil {
		return nil
	}
	return s.DeleteProduct
}

func (s *SmbToolConfigParams) GetDeleteService() *DeleteServiceParams {
	if s == nil {
		return nil
	}
	if s.DeleteService == nil {
		return nil
	}
	return s.DeleteService
}

func (s *SmbToolConfigParams) GetDeleteStaff() *DeleteStaffParams {
	if s == nil {
		return nil
	}
	if s.DeleteStaff == nil {
		return nil
	}
	return s.DeleteStaff
}

func (s *SmbToolConfigParams) GetDeleteTransferRule() *DeleteTransferRuleParams {
	if s == nil {
		return nil
	}
	if s.DeleteTransferRule == nil {
		return nil
	}
	return s.DeleteTransferRule
}

func (s *SmbToolConfigParams) GetGetAgentConfig() *GetAgentConfigParams {
	if s == nil {
		return nil
	}
	if s.GetAgentConfig == nil {
		return nil
	}
	return s.GetAgentConfig
}

func (s *SmbToolConfigParams) GetGetAnalyticsSummary() *GetAnalyticsSummaryParams {
	if s == nil {
		return nil
	}
	if s.GetAnalyticsSummary == nil {
		return nil
	}
	return s.GetAnalyticsSummary
}

func (s *SmbToolConfigParams) GetGetAppointmentByConfirmationNumber() *GetAppointmentByConfirmationNumberParams {
	if s == nil {
		return nil
	}
	if s.GetAppointmentByConfirmationNumber == nil {
		return nil
	}
	return s.GetAppointmentByConfirmationNumber
}

func (s *SmbToolConfigParams) GetGetBookingPageSettings() *GetBookingPageSettingsParams {
	if s == nil {
		return nil
	}
	if s.GetBookingPageSettings == nil {
		return nil
	}
	return s.GetBookingPageSettings
}

func (s *SmbToolConfigParams) GetGetBookingSlugStatus() *GetBookingSlugStatusParams {
	if s == nil {
		return nil
	}
	if s.GetBookingSlugStatus == nil {
		return nil
	}
	return s.GetBookingSlugStatus
}

func (s *SmbToolConfigParams) GetGetClientAppointments() *GetClientAppointmentsParams {
	if s == nil {
		return nil
	}
	if s.GetClientAppointments == nil {
		return nil
	}
	return s.GetClientAppointments
}

func (s *SmbToolConfigParams) GetGetClientByPhone() *GetClientByPhoneParams {
	if s == nil {
		return nil
	}
	if s.GetClientByPhone == nil {
		return nil
	}
	return s.GetClientByPhone
}

func (s *SmbToolConfigParams) GetGetClientOrders() *GetClientOrdersParams {
	if s == nil {
		return nil
	}
	if s.GetClientOrders == nil {
		return nil
	}
	return s.GetClientOrders
}

func (s *SmbToolConfigParams) GetGetOrderByConfirmationNumber() *GetOrderByConfirmationNumberParams {
	if s == nil {
		return nil
	}
	if s.GetOrderByConfirmationNumber == nil {
		return nil
	}
	return s.GetOrderByConfirmationNumber
}

func (s *SmbToolConfigParams) GetGetSchedule() *GetScheduleParams {
	if s == nil {
		return nil
	}
	if s.GetSchedule == nil {
		return nil
	}
	return s.GetSchedule
}

func (s *SmbToolConfigParams) GetGetWorkspacePlan() *GetWorkspacePlanParams {
	if s == nil {
		return nil
	}
	if s.GetWorkspacePlan == nil {
		return nil
	}
	return s.GetWorkspacePlan
}

func (s *SmbToolConfigParams) GetLeaveMessage() *LeaveMessageParams {
	if s == nil {
		return nil
	}
	if s.LeaveMessage == nil {
		return nil
	}
	return s.LeaveMessage
}

func (s *SmbToolConfigParams) GetListAgentProcedures() *ListAgentProceduresParams {
	if s == nil {
		return nil
	}
	if s.ListAgentProcedures == nil {
		return nil
	}
	return s.ListAgentProcedures
}

func (s *SmbToolConfigParams) GetListAgentReferences() *ListAgentReferencesParams {
	if s == nil {
		return nil
	}
	if s.ListAgentReferences == nil {
		return nil
	}
	return s.ListAgentReferences
}

func (s *SmbToolConfigParams) GetListAgentRules() *ListAgentRulesParams {
	if s == nil {
		return nil
	}
	if s.ListAgentRules == nil {
		return nil
	}
	return s.ListAgentRules
}

func (s *SmbToolConfigParams) GetListAssets() *ListAssetsParams {
	if s == nil {
		return nil
	}
	if s.ListAssets == nil {
		return nil
	}
	return s.ListAssets
}

func (s *SmbToolConfigParams) GetListCalendarEvents() *ListCalendarEventsParams {
	if s == nil {
		return nil
	}
	if s.ListCalendarEvents == nil {
		return nil
	}
	return s.ListCalendarEvents
}

func (s *SmbToolConfigParams) GetListClientInteractions() *ListClientInteractionsParams {
	if s == nil {
		return nil
	}
	if s.ListClientInteractions == nil {
		return nil
	}
	return s.ListClientInteractions
}

func (s *SmbToolConfigParams) GetListClients() *ListClientsParams {
	if s == nil {
		return nil
	}
	if s.ListClients == nil {
		return nil
	}
	return s.ListClients
}

func (s *SmbToolConfigParams) GetListCustomerFacingAgents() *ListCustomerFacingAgentsParams {
	if s == nil {
		return nil
	}
	if s.ListCustomerFacingAgents == nil {
		return nil
	}
	return s.ListCustomerFacingAgents
}

func (s *SmbToolConfigParams) GetListGroupSessions() *ListGroupSessionsParams {
	if s == nil {
		return nil
	}
	if s.ListGroupSessions == nil {
		return nil
	}
	return s.ListGroupSessions
}

func (s *SmbToolConfigParams) GetListHolidays() *ListHolidaysParams {
	if s == nil {
		return nil
	}
	if s.ListHolidays == nil {
		return nil
	}
	return s.ListHolidays
}

func (s *SmbToolConfigParams) GetListLocations() *ListLocationsParams {
	if s == nil {
		return nil
	}
	if s.ListLocations == nil {
		return nil
	}
	return s.ListLocations
}

func (s *SmbToolConfigParams) GetListProducts() *ListProductsParams {
	if s == nil {
		return nil
	}
	if s.ListProducts == nil {
		return nil
	}
	return s.ListProducts
}

func (s *SmbToolConfigParams) GetListServices() *ListServicesParams {
	if s == nil {
		return nil
	}
	if s.ListServices == nil {
		return nil
	}
	return s.ListServices
}

func (s *SmbToolConfigParams) GetListStaff() *ListStaffParams {
	if s == nil {
		return nil
	}
	if s.ListStaff == nil {
		return nil
	}
	return s.ListStaff
}

func (s *SmbToolConfigParams) GetListTransferRules() *ListTransferRulesParams {
	if s == nil {
		return nil
	}
	if s.ListTransferRules == nil {
		return nil
	}
	return s.ListTransferRules
}

func (s *SmbToolConfigParams) GetOptInSmsReminder() *OptInSmsReminderParams {
	if s == nil {
		return nil
	}
	if s.OptInSmsReminder == nil {
		return nil
	}
	return s.OptInSmsReminder
}

func (s *SmbToolConfigParams) GetOptOutSmsReminder() *OptOutSmsReminderParams {
	if s == nil {
		return nil
	}
	if s.OptOutSmsReminder == nil {
		return nil
	}
	return s.OptOutSmsReminder
}

func (s *SmbToolConfigParams) GetOwnerGetClientAppointments() *OwnerGetClientAppointmentsParams {
	if s == nil {
		return nil
	}
	if s.OwnerGetClientAppointments == nil {
		return nil
	}
	return s.OwnerGetClientAppointments
}

func (s *SmbToolConfigParams) GetOwnerListAssets() *OwnerListAssetsParams {
	if s == nil {
		return nil
	}
	if s.OwnerListAssets == nil {
		return nil
	}
	return s.OwnerListAssets
}

func (s *SmbToolConfigParams) GetOwnerListLocations() *OwnerListLocationsParams {
	if s == nil {
		return nil
	}
	if s.OwnerListLocations == nil {
		return nil
	}
	return s.OwnerListLocations
}

func (s *SmbToolConfigParams) GetOwnerListProducts() *OwnerListProductsParams {
	if s == nil {
		return nil
	}
	if s.OwnerListProducts == nil {
		return nil
	}
	return s.OwnerListProducts
}

func (s *SmbToolConfigParams) GetOwnerListServices() *OwnerListServicesParams {
	if s == nil {
		return nil
	}
	if s.OwnerListServices == nil {
		return nil
	}
	return s.OwnerListServices
}

func (s *SmbToolConfigParams) GetOwnerListStaff() *OwnerListStaffParams {
	if s == nil {
		return nil
	}
	if s.OwnerListStaff == nil {
		return nil
	}
	return s.OwnerListStaff
}

func (s *SmbToolConfigParams) GetRegisterForGroupSession() *RegisterForGroupSessionParams {
	if s == nil {
		return nil
	}
	if s.RegisterForGroupSession == nil {
		return nil
	}
	return s.RegisterForGroupSession
}

func (s *SmbToolConfigParams) GetReportKnowledgeGap() *ReportKnowledgeGapParams {
	if s == nil {
		return nil
	}
	if s.ReportKnowledgeGap == nil {
		return nil
	}
	return s.ReportKnowledgeGap
}

func (s *SmbToolConfigParams) GetRestoreCalendarEvent() *RestoreCalendarEventParams {
	if s == nil {
		return nil
	}
	if s.RestoreCalendarEvent == nil {
		return nil
	}
	return s.RestoreCalendarEvent
}

func (s *SmbToolConfigParams) GetScheduleGroupSession() *ScheduleGroupSessionParams {
	if s == nil {
		return nil
	}
	if s.ScheduleGroupSession == nil {
		return nil
	}
	return s.ScheduleGroupSession
}

func (s *SmbToolConfigParams) GetSearchClients() *SearchClientsParams {
	if s == nil {
		return nil
	}
	if s.SearchClients == nil {
		return nil
	}
	return s.SearchClients
}

func (s *SmbToolConfigParams) GetSendCustomEmail() *SendCustomEmailParams {
	if s == nil {
		return nil
	}
	if s.SendCustomEmail == nil {
		return nil
	}
	return s.SendCustomEmail
}

func (s *SmbToolConfigParams) GetSetBookingSlug() *SetBookingSlugParams {
	if s == nil {
		return nil
	}
	if s.SetBookingSlug == nil {
		return nil
	}
	return s.SetBookingSlug
}

func (s *SmbToolConfigParams) GetSubmitBusinessInfo() *SubmitBusinessInfoParams {
	if s == nil {
		return nil
	}
	if s.SubmitBusinessInfo == nil {
		return nil
	}
	return s.SubmitBusinessInfo
}

func (s *SmbToolConfigParams) GetTriggerUserVerification() *TriggerUserVerificationParams {
	if s == nil {
		return nil
	}
	if s.TriggerUserVerification == nil {
		return nil
	}
	return s.TriggerUserVerification
}

func (s *SmbToolConfigParams) GetUpdateAgentConfig() *UpdateAgentConfigParams {
	if s == nil {
		return nil
	}
	if s.UpdateAgentConfig == nil {
		return nil
	}
	return s.UpdateAgentConfig
}

func (s *SmbToolConfigParams) GetUpdateAgentProcedure() *UpdateAgentProcedureParams {
	if s == nil {
		return nil
	}
	if s.UpdateAgentProcedure == nil {
		return nil
	}
	return s.UpdateAgentProcedure
}

func (s *SmbToolConfigParams) GetUpdateAgentRule() *UpdateAgentRuleParams {
	if s == nil {
		return nil
	}
	if s.UpdateAgentRule == nil {
		return nil
	}
	return s.UpdateAgentRule
}

func (s *SmbToolConfigParams) GetUpdateAsset() *UpdateAssetParams {
	if s == nil {
		return nil
	}
	if s.UpdateAsset == nil {
		return nil
	}
	return s.UpdateAsset
}

func (s *SmbToolConfigParams) GetUpdateBookingPageAppearance() *UpdateBookingPageAppearanceParams {
	if s == nil {
		return nil
	}
	if s.UpdateBookingPageAppearance == nil {
		return nil
	}
	return s.UpdateBookingPageAppearance
}

func (s *SmbToolConfigParams) GetUpdateBookingPageSettings() *UpdateBookingPageSettingsParams {
	if s == nil {
		return nil
	}
	if s.UpdateBookingPageSettings == nil {
		return nil
	}
	return s.UpdateBookingPageSettings
}

func (s *SmbToolConfigParams) GetUpdateBusinessInfo() *UpdateBusinessInfoParams {
	if s == nil {
		return nil
	}
	if s.UpdateBusinessInfo == nil {
		return nil
	}
	return s.UpdateBusinessInfo
}

func (s *SmbToolConfigParams) GetUpdateCalendarEvent() *UpdateCalendarEventParams {
	if s == nil {
		return nil
	}
	if s.UpdateCalendarEvent == nil {
		return nil
	}
	return s.UpdateCalendarEvent
}

func (s *SmbToolConfigParams) GetUpdateClient() *UpdateClientParams {
	if s == nil {
		return nil
	}
	if s.UpdateClient == nil {
		return nil
	}
	return s.UpdateClient
}

func (s *SmbToolConfigParams) GetUpdateGroupSessionSeats() *UpdateGroupSessionSeatsParams {
	if s == nil {
		return nil
	}
	if s.UpdateGroupSessionSeats == nil {
		return nil
	}
	return s.UpdateGroupSessionSeats
}

func (s *SmbToolConfigParams) GetUpdateHoliday() *UpdateHolidayParams {
	if s == nil {
		return nil
	}
	if s.UpdateHoliday == nil {
		return nil
	}
	return s.UpdateHoliday
}

func (s *SmbToolConfigParams) GetUpdateLocation() *UpdateLocationParams {
	if s == nil {
		return nil
	}
	if s.UpdateLocation == nil {
		return nil
	}
	return s.UpdateLocation
}

func (s *SmbToolConfigParams) GetUpdateOrder() *UpdateOrderParams {
	if s == nil {
		return nil
	}
	if s.UpdateOrder == nil {
		return nil
	}
	return s.UpdateOrder
}

func (s *SmbToolConfigParams) GetUpdateProduct() *UpdateProductParams {
	if s == nil {
		return nil
	}
	if s.UpdateProduct == nil {
		return nil
	}
	return s.UpdateProduct
}

func (s *SmbToolConfigParams) GetUpdateService() *UpdateServiceParams {
	if s == nil {
		return nil
	}
	if s.UpdateService == nil {
		return nil
	}
	return s.UpdateService
}

func (s *SmbToolConfigParams) GetUpdateStaff() *UpdateStaffParams {
	if s == nil {
		return nil
	}
	if s.UpdateStaff == nil {
		return nil
	}
	return s.UpdateStaff
}

func (s *SmbToolConfigParams) GetUpdateTransferRule() *UpdateTransferRuleParams {
	if s == nil {
		return nil
	}
	if s.UpdateTransferRule == nil {
		return nil
	}
	return s.UpdateTransferRule
}

func (s *SmbToolConfigParams) GetValidateUserVerificationCode() *ValidateUserVerificationCodeParams {
	if s == nil {
		return nil
	}
	if s.ValidateUserVerificationCode == nil {
		return nil
	}
	return s.ValidateUserVerificationCode
}

func (s *SmbToolConfigParams) GetSmbToolType() string {
	if s == nil {
		return ""
	}
	return s.SmbToolType
}

func (s *SmbToolConfigParams) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		SmbToolType string `json:"smb_tool_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.SmbToolType = unmarshaler.SmbToolType
	if unmarshaler.SmbToolType == "" {
		return fmt.Errorf("%T did not include discriminant smb_tool_type", s)
	}
	switch unmarshaler.SmbToolType {
	case "cancel_calendar_event":
		value := new(CancelCalendarEventParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CancelCalendarEvent = value
	case "cancel_group_session_for_all":
		value := new(CancelGroupSessionForAllParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CancelGroupSessionForAll = value
	case "cancel_group_session_registration":
		value := new(CancelGroupSessionRegistrationParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CancelGroupSessionRegistration = value
	case "cancel_order":
		value := new(CancelOrderParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CancelOrder = value
	case "check_service_availability":
		value := new(CheckServiceAvailabilityParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CheckServiceAvailability = value
	case "create_agent_procedure":
		value := new(CreateAgentProcedureParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateAgentProcedure = value
	case "create_agent_rule":
		value := new(CreateAgentRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateAgentRule = value
	case "create_asset":
		value := new(CreateAssetParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateAsset = value
	case "create_client":
		value := new(CreateClientParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateClient = value
	case "create_client_appointment":
		value := new(CreateClientAppointmentParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateClientAppointment = value
	case "create_client_interaction":
		value := new(CreateClientInteractionParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateClientInteraction = value
	case "create_holiday":
		value := new(CreateHolidayParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateHoliday = value
	case "create_location":
		value := new(CreateLocationParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateLocation = value
	case "create_order":
		value := new(CreateOrderParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateOrder = value
	case "create_product":
		value := new(CreateProductParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateProduct = value
	case "create_product_quote_request":
		value := new(CreateProductQuoteRequestParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateProductQuoteRequest = value
	case "create_service":
		value := new(CreateServiceParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateService = value
	case "create_service_quote_request":
		value := new(CreateServiceQuoteRequestParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateServiceQuoteRequest = value
	case "create_staff":
		value := new(CreateStaffParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateStaff = value
	case "create_transfer_rule":
		value := new(CreateTransferRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.CreateTransferRule = value
	case "delete_agent_procedure":
		value := new(DeleteAgentProcedureParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteAgentProcedure = value
	case "delete_agent_rule":
		value := new(DeleteAgentRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteAgentRule = value
	case "delete_asset":
		value := new(DeleteAssetParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteAsset = value
	case "delete_calendar_event":
		value := new(DeleteCalendarEventParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteCalendarEvent = value
	case "delete_client":
		value := new(DeleteClientParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteClient = value
	case "delete_client_interaction":
		value := new(DeleteClientInteractionParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteClientInteraction = value
	case "delete_group_session":
		value := new(DeleteGroupSessionParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteGroupSession = value
	case "delete_holiday":
		value := new(DeleteHolidayParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteHoliday = value
	case "delete_location":
		value := new(DeleteLocationParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteLocation = value
	case "delete_product":
		value := new(DeleteProductParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteProduct = value
	case "delete_service":
		value := new(DeleteServiceParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteService = value
	case "delete_staff":
		value := new(DeleteStaffParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteStaff = value
	case "delete_transfer_rule":
		value := new(DeleteTransferRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.DeleteTransferRule = value
	case "get_agent_config":
		value := new(GetAgentConfigParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetAgentConfig = value
	case "get_analytics_summary":
		value := new(GetAnalyticsSummaryParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetAnalyticsSummary = value
	case "get_appointment_by_confirmation_number":
		value := new(GetAppointmentByConfirmationNumberParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetAppointmentByConfirmationNumber = value
	case "get_booking_page_settings":
		value := new(GetBookingPageSettingsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetBookingPageSettings = value
	case "get_booking_slug_status":
		value := new(GetBookingSlugStatusParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetBookingSlugStatus = value
	case "get_client_appointments":
		value := new(GetClientAppointmentsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetClientAppointments = value
	case "get_client_by_phone":
		value := new(GetClientByPhoneParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetClientByPhone = value
	case "get_client_orders":
		value := new(GetClientOrdersParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetClientOrders = value
	case "get_order_by_confirmation_number":
		value := new(GetOrderByConfirmationNumberParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetOrderByConfirmationNumber = value
	case "get_schedule":
		value := new(GetScheduleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetSchedule = value
	case "get_workspace_plan":
		value := new(GetWorkspacePlanParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.GetWorkspacePlan = value
	case "leave_message":
		value := new(LeaveMessageParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.LeaveMessage = value
	case "list_agent_procedures":
		value := new(ListAgentProceduresParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListAgentProcedures = value
	case "list_agent_references":
		value := new(ListAgentReferencesParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListAgentReferences = value
	case "list_agent_rules":
		value := new(ListAgentRulesParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListAgentRules = value
	case "list_assets":
		value := new(ListAssetsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListAssets = value
	case "list_calendar_events":
		value := new(ListCalendarEventsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListCalendarEvents = value
	case "list_client_interactions":
		value := new(ListClientInteractionsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListClientInteractions = value
	case "list_clients":
		value := new(ListClientsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListClients = value
	case "list_customer_facing_agents":
		value := new(ListCustomerFacingAgentsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListCustomerFacingAgents = value
	case "list_group_sessions":
		value := new(ListGroupSessionsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListGroupSessions = value
	case "list_holidays":
		value := new(ListHolidaysParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListHolidays = value
	case "list_locations":
		value := new(ListLocationsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListLocations = value
	case "list_products":
		value := new(ListProductsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListProducts = value
	case "list_services":
		value := new(ListServicesParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListServices = value
	case "list_staff":
		value := new(ListStaffParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListStaff = value
	case "list_transfer_rules":
		value := new(ListTransferRulesParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ListTransferRules = value
	case "opt_in_sms_reminder":
		value := new(OptInSmsReminderParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OptInSmsReminder = value
	case "opt_out_sms_reminder":
		value := new(OptOutSmsReminderParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OptOutSmsReminder = value
	case "owner_get_client_appointments":
		value := new(OwnerGetClientAppointmentsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerGetClientAppointments = value
	case "owner_list_assets":
		value := new(OwnerListAssetsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerListAssets = value
	case "owner_list_locations":
		value := new(OwnerListLocationsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerListLocations = value
	case "owner_list_products":
		value := new(OwnerListProductsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerListProducts = value
	case "owner_list_services":
		value := new(OwnerListServicesParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerListServices = value
	case "owner_list_staff":
		value := new(OwnerListStaffParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.OwnerListStaff = value
	case "register_for_group_session":
		value := new(RegisterForGroupSessionParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.RegisterForGroupSession = value
	case "report_knowledge_gap":
		value := new(ReportKnowledgeGapParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ReportKnowledgeGap = value
	case "restore_calendar_event":
		value := new(RestoreCalendarEventParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.RestoreCalendarEvent = value
	case "schedule_group_session":
		value := new(ScheduleGroupSessionParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ScheduleGroupSession = value
	case "search_clients":
		value := new(SearchClientsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SearchClients = value
	case "send_custom_email":
		value := new(SendCustomEmailParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SendCustomEmail = value
	case "set_booking_slug":
		value := new(SetBookingSlugParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SetBookingSlug = value
	case "submit_business_info":
		value := new(SubmitBusinessInfoParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.SubmitBusinessInfo = value
	case "trigger_user_verification":
		value := new(TriggerUserVerificationParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.TriggerUserVerification = value
	case "update_agent_config":
		value := new(UpdateAgentConfigParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateAgentConfig = value
	case "update_agent_procedure":
		value := new(UpdateAgentProcedureParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateAgentProcedure = value
	case "update_agent_rule":
		value := new(UpdateAgentRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateAgentRule = value
	case "update_asset":
		value := new(UpdateAssetParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateAsset = value
	case "update_booking_page_appearance":
		value := new(UpdateBookingPageAppearanceParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateBookingPageAppearance = value
	case "update_booking_page_settings":
		value := new(UpdateBookingPageSettingsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateBookingPageSettings = value
	case "update_business_info":
		value := new(UpdateBusinessInfoParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateBusinessInfo = value
	case "update_calendar_event":
		value := new(UpdateCalendarEventParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateCalendarEvent = value
	case "update_client":
		value := new(UpdateClientParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateClient = value
	case "update_group_session_seats":
		value := new(UpdateGroupSessionSeatsParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateGroupSessionSeats = value
	case "update_holiday":
		value := new(UpdateHolidayParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateHoliday = value
	case "update_location":
		value := new(UpdateLocationParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateLocation = value
	case "update_order":
		value := new(UpdateOrderParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateOrder = value
	case "update_product":
		value := new(UpdateProductParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateProduct = value
	case "update_service":
		value := new(UpdateServiceParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateService = value
	case "update_staff":
		value := new(UpdateStaffParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateStaff = value
	case "update_transfer_rule":
		value := new(UpdateTransferRuleParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.UpdateTransferRule = value
	case "validate_user_verification_code":
		value := new(ValidateUserVerificationCodeParams)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.ValidateUserVerificationCode = value
	}
	return nil
}

func (s SmbToolConfigParams) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.CancelCalendarEvent != nil {
		return internal.MarshalJSONWithDiscriminant(s.CancelCalendarEvent, "smb_tool_type", "cancel_calendar_event")
	}
	if s.CancelGroupSessionForAll != nil {
		return internal.MarshalJSONWithDiscriminant(s.CancelGroupSessionForAll, "smb_tool_type", "cancel_group_session_for_all")
	}
	if s.CancelGroupSessionRegistration != nil {
		return internal.MarshalJSONWithDiscriminant(s.CancelGroupSessionRegistration, "smb_tool_type", "cancel_group_session_registration")
	}
	if s.CancelOrder != nil {
		return internal.MarshalJSONWithDiscriminant(s.CancelOrder, "smb_tool_type", "cancel_order")
	}
	if s.CheckServiceAvailability != nil {
		return internal.MarshalJSONWithDiscriminant(s.CheckServiceAvailability, "smb_tool_type", "check_service_availability")
	}
	if s.CreateAgentProcedure != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateAgentProcedure, "smb_tool_type", "create_agent_procedure")
	}
	if s.CreateAgentRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateAgentRule, "smb_tool_type", "create_agent_rule")
	}
	if s.CreateAsset != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateAsset, "smb_tool_type", "create_asset")
	}
	if s.CreateClient != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateClient, "smb_tool_type", "create_client")
	}
	if s.CreateClientAppointment != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateClientAppointment, "smb_tool_type", "create_client_appointment")
	}
	if s.CreateClientInteraction != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateClientInteraction, "smb_tool_type", "create_client_interaction")
	}
	if s.CreateHoliday != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateHoliday, "smb_tool_type", "create_holiday")
	}
	if s.CreateLocation != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateLocation, "smb_tool_type", "create_location")
	}
	if s.CreateOrder != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateOrder, "smb_tool_type", "create_order")
	}
	if s.CreateProduct != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateProduct, "smb_tool_type", "create_product")
	}
	if s.CreateProductQuoteRequest != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateProductQuoteRequest, "smb_tool_type", "create_product_quote_request")
	}
	if s.CreateService != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateService, "smb_tool_type", "create_service")
	}
	if s.CreateServiceQuoteRequest != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateServiceQuoteRequest, "smb_tool_type", "create_service_quote_request")
	}
	if s.CreateStaff != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateStaff, "smb_tool_type", "create_staff")
	}
	if s.CreateTransferRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.CreateTransferRule, "smb_tool_type", "create_transfer_rule")
	}
	if s.DeleteAgentProcedure != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteAgentProcedure, "smb_tool_type", "delete_agent_procedure")
	}
	if s.DeleteAgentRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteAgentRule, "smb_tool_type", "delete_agent_rule")
	}
	if s.DeleteAsset != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteAsset, "smb_tool_type", "delete_asset")
	}
	if s.DeleteCalendarEvent != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteCalendarEvent, "smb_tool_type", "delete_calendar_event")
	}
	if s.DeleteClient != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteClient, "smb_tool_type", "delete_client")
	}
	if s.DeleteClientInteraction != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteClientInteraction, "smb_tool_type", "delete_client_interaction")
	}
	if s.DeleteGroupSession != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteGroupSession, "smb_tool_type", "delete_group_session")
	}
	if s.DeleteHoliday != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteHoliday, "smb_tool_type", "delete_holiday")
	}
	if s.DeleteLocation != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteLocation, "smb_tool_type", "delete_location")
	}
	if s.DeleteProduct != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteProduct, "smb_tool_type", "delete_product")
	}
	if s.DeleteService != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteService, "smb_tool_type", "delete_service")
	}
	if s.DeleteStaff != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteStaff, "smb_tool_type", "delete_staff")
	}
	if s.DeleteTransferRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.DeleteTransferRule, "smb_tool_type", "delete_transfer_rule")
	}
	if s.GetAgentConfig != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetAgentConfig, "smb_tool_type", "get_agent_config")
	}
	if s.GetAnalyticsSummary != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetAnalyticsSummary, "smb_tool_type", "get_analytics_summary")
	}
	if s.GetAppointmentByConfirmationNumber != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetAppointmentByConfirmationNumber, "smb_tool_type", "get_appointment_by_confirmation_number")
	}
	if s.GetBookingPageSettings != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetBookingPageSettings, "smb_tool_type", "get_booking_page_settings")
	}
	if s.GetBookingSlugStatus != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetBookingSlugStatus, "smb_tool_type", "get_booking_slug_status")
	}
	if s.GetClientAppointments != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetClientAppointments, "smb_tool_type", "get_client_appointments")
	}
	if s.GetClientByPhone != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetClientByPhone, "smb_tool_type", "get_client_by_phone")
	}
	if s.GetClientOrders != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetClientOrders, "smb_tool_type", "get_client_orders")
	}
	if s.GetOrderByConfirmationNumber != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetOrderByConfirmationNumber, "smb_tool_type", "get_order_by_confirmation_number")
	}
	if s.GetSchedule != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetSchedule, "smb_tool_type", "get_schedule")
	}
	if s.GetWorkspacePlan != nil {
		return internal.MarshalJSONWithDiscriminant(s.GetWorkspacePlan, "smb_tool_type", "get_workspace_plan")
	}
	if s.LeaveMessage != nil {
		return internal.MarshalJSONWithDiscriminant(s.LeaveMessage, "smb_tool_type", "leave_message")
	}
	if s.ListAgentProcedures != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListAgentProcedures, "smb_tool_type", "list_agent_procedures")
	}
	if s.ListAgentReferences != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListAgentReferences, "smb_tool_type", "list_agent_references")
	}
	if s.ListAgentRules != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListAgentRules, "smb_tool_type", "list_agent_rules")
	}
	if s.ListAssets != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListAssets, "smb_tool_type", "list_assets")
	}
	if s.ListCalendarEvents != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListCalendarEvents, "smb_tool_type", "list_calendar_events")
	}
	if s.ListClientInteractions != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListClientInteractions, "smb_tool_type", "list_client_interactions")
	}
	if s.ListClients != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListClients, "smb_tool_type", "list_clients")
	}
	if s.ListCustomerFacingAgents != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListCustomerFacingAgents, "smb_tool_type", "list_customer_facing_agents")
	}
	if s.ListGroupSessions != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListGroupSessions, "smb_tool_type", "list_group_sessions")
	}
	if s.ListHolidays != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListHolidays, "smb_tool_type", "list_holidays")
	}
	if s.ListLocations != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListLocations, "smb_tool_type", "list_locations")
	}
	if s.ListProducts != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListProducts, "smb_tool_type", "list_products")
	}
	if s.ListServices != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListServices, "smb_tool_type", "list_services")
	}
	if s.ListStaff != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListStaff, "smb_tool_type", "list_staff")
	}
	if s.ListTransferRules != nil {
		return internal.MarshalJSONWithDiscriminant(s.ListTransferRules, "smb_tool_type", "list_transfer_rules")
	}
	if s.OptInSmsReminder != nil {
		return internal.MarshalJSONWithDiscriminant(s.OptInSmsReminder, "smb_tool_type", "opt_in_sms_reminder")
	}
	if s.OptOutSmsReminder != nil {
		return internal.MarshalJSONWithDiscriminant(s.OptOutSmsReminder, "smb_tool_type", "opt_out_sms_reminder")
	}
	if s.OwnerGetClientAppointments != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerGetClientAppointments, "smb_tool_type", "owner_get_client_appointments")
	}
	if s.OwnerListAssets != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerListAssets, "smb_tool_type", "owner_list_assets")
	}
	if s.OwnerListLocations != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerListLocations, "smb_tool_type", "owner_list_locations")
	}
	if s.OwnerListProducts != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerListProducts, "smb_tool_type", "owner_list_products")
	}
	if s.OwnerListServices != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerListServices, "smb_tool_type", "owner_list_services")
	}
	if s.OwnerListStaff != nil {
		return internal.MarshalJSONWithDiscriminant(s.OwnerListStaff, "smb_tool_type", "owner_list_staff")
	}
	if s.RegisterForGroupSession != nil {
		return internal.MarshalJSONWithDiscriminant(s.RegisterForGroupSession, "smb_tool_type", "register_for_group_session")
	}
	if s.ReportKnowledgeGap != nil {
		return internal.MarshalJSONWithDiscriminant(s.ReportKnowledgeGap, "smb_tool_type", "report_knowledge_gap")
	}
	if s.RestoreCalendarEvent != nil {
		return internal.MarshalJSONWithDiscriminant(s.RestoreCalendarEvent, "smb_tool_type", "restore_calendar_event")
	}
	if s.ScheduleGroupSession != nil {
		return internal.MarshalJSONWithDiscriminant(s.ScheduleGroupSession, "smb_tool_type", "schedule_group_session")
	}
	if s.SearchClients != nil {
		return internal.MarshalJSONWithDiscriminant(s.SearchClients, "smb_tool_type", "search_clients")
	}
	if s.SendCustomEmail != nil {
		return internal.MarshalJSONWithDiscriminant(s.SendCustomEmail, "smb_tool_type", "send_custom_email")
	}
	if s.SetBookingSlug != nil {
		return internal.MarshalJSONWithDiscriminant(s.SetBookingSlug, "smb_tool_type", "set_booking_slug")
	}
	if s.SubmitBusinessInfo != nil {
		return internal.MarshalJSONWithDiscriminant(s.SubmitBusinessInfo, "smb_tool_type", "submit_business_info")
	}
	if s.TriggerUserVerification != nil {
		return internal.MarshalJSONWithDiscriminant(s.TriggerUserVerification, "smb_tool_type", "trigger_user_verification")
	}
	if s.UpdateAgentConfig != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateAgentConfig, "smb_tool_type", "update_agent_config")
	}
	if s.UpdateAgentProcedure != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateAgentProcedure, "smb_tool_type", "update_agent_procedure")
	}
	if s.UpdateAgentRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateAgentRule, "smb_tool_type", "update_agent_rule")
	}
	if s.UpdateAsset != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateAsset, "smb_tool_type", "update_asset")
	}
	if s.UpdateBookingPageAppearance != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateBookingPageAppearance, "smb_tool_type", "update_booking_page_appearance")
	}
	if s.UpdateBookingPageSettings != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateBookingPageSettings, "smb_tool_type", "update_booking_page_settings")
	}
	if s.UpdateBusinessInfo != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateBusinessInfo, "smb_tool_type", "update_business_info")
	}
	if s.UpdateCalendarEvent != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateCalendarEvent, "smb_tool_type", "update_calendar_event")
	}
	if s.UpdateClient != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateClient, "smb_tool_type", "update_client")
	}
	if s.UpdateGroupSessionSeats != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateGroupSessionSeats, "smb_tool_type", "update_group_session_seats")
	}
	if s.UpdateHoliday != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateHoliday, "smb_tool_type", "update_holiday")
	}
	if s.UpdateLocation != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateLocation, "smb_tool_type", "update_location")
	}
	if s.UpdateOrder != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateOrder, "smb_tool_type", "update_order")
	}
	if s.UpdateProduct != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateProduct, "smb_tool_type", "update_product")
	}
	if s.UpdateService != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateService, "smb_tool_type", "update_service")
	}
	if s.UpdateStaff != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateStaff, "smb_tool_type", "update_staff")
	}
	if s.UpdateTransferRule != nil {
		return internal.MarshalJSONWithDiscriminant(s.UpdateTransferRule, "smb_tool_type", "update_transfer_rule")
	}
	if s.ValidateUserVerificationCode != nil {
		return internal.MarshalJSONWithDiscriminant(s.ValidateUserVerificationCode, "smb_tool_type", "validate_user_verification_code")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SmbToolConfigParamsVisitor interface {
	VisitCancelCalendarEvent(*CancelCalendarEventParams) error
	VisitCancelGroupSessionForAll(*CancelGroupSessionForAllParams) error
	VisitCancelGroupSessionRegistration(*CancelGroupSessionRegistrationParams) error
	VisitCancelOrder(*CancelOrderParams) error
	VisitCheckServiceAvailability(*CheckServiceAvailabilityParams) error
	VisitCreateAgentProcedure(*CreateAgentProcedureParams) error
	VisitCreateAgentRule(*CreateAgentRuleParams) error
	VisitCreateAsset(*CreateAssetParams) error
	VisitCreateClient(*CreateClientParams) error
	VisitCreateClientAppointment(*CreateClientAppointmentParams) error
	VisitCreateClientInteraction(*CreateClientInteractionParams) error
	VisitCreateHoliday(*CreateHolidayParams) error
	VisitCreateLocation(*CreateLocationParams) error
	VisitCreateOrder(*CreateOrderParams) error
	VisitCreateProduct(*CreateProductParams) error
	VisitCreateProductQuoteRequest(*CreateProductQuoteRequestParams) error
	VisitCreateService(*CreateServiceParams) error
	VisitCreateServiceQuoteRequest(*CreateServiceQuoteRequestParams) error
	VisitCreateStaff(*CreateStaffParams) error
	VisitCreateTransferRule(*CreateTransferRuleParams) error
	VisitDeleteAgentProcedure(*DeleteAgentProcedureParams) error
	VisitDeleteAgentRule(*DeleteAgentRuleParams) error
	VisitDeleteAsset(*DeleteAssetParams) error
	VisitDeleteCalendarEvent(*DeleteCalendarEventParams) error
	VisitDeleteClient(*DeleteClientParams) error
	VisitDeleteClientInteraction(*DeleteClientInteractionParams) error
	VisitDeleteGroupSession(*DeleteGroupSessionParams) error
	VisitDeleteHoliday(*DeleteHolidayParams) error
	VisitDeleteLocation(*DeleteLocationParams) error
	VisitDeleteProduct(*DeleteProductParams) error
	VisitDeleteService(*DeleteServiceParams) error
	VisitDeleteStaff(*DeleteStaffParams) error
	VisitDeleteTransferRule(*DeleteTransferRuleParams) error
	VisitGetAgentConfig(*GetAgentConfigParams) error
	VisitGetAnalyticsSummary(*GetAnalyticsSummaryParams) error
	VisitGetAppointmentByConfirmationNumber(*GetAppointmentByConfirmationNumberParams) error
	VisitGetBookingPageSettings(*GetBookingPageSettingsParams) error
	VisitGetBookingSlugStatus(*GetBookingSlugStatusParams) error
	VisitGetClientAppointments(*GetClientAppointmentsParams) error
	VisitGetClientByPhone(*GetClientByPhoneParams) error
	VisitGetClientOrders(*GetClientOrdersParams) error
	VisitGetOrderByConfirmationNumber(*GetOrderByConfirmationNumberParams) error
	VisitGetSchedule(*GetScheduleParams) error
	VisitGetWorkspacePlan(*GetWorkspacePlanParams) error
	VisitLeaveMessage(*LeaveMessageParams) error
	VisitListAgentProcedures(*ListAgentProceduresParams) error
	VisitListAgentReferences(*ListAgentReferencesParams) error
	VisitListAgentRules(*ListAgentRulesParams) error
	VisitListAssets(*ListAssetsParams) error
	VisitListCalendarEvents(*ListCalendarEventsParams) error
	VisitListClientInteractions(*ListClientInteractionsParams) error
	VisitListClients(*ListClientsParams) error
	VisitListCustomerFacingAgents(*ListCustomerFacingAgentsParams) error
	VisitListGroupSessions(*ListGroupSessionsParams) error
	VisitListHolidays(*ListHolidaysParams) error
	VisitListLocations(*ListLocationsParams) error
	VisitListProducts(*ListProductsParams) error
	VisitListServices(*ListServicesParams) error
	VisitListStaff(*ListStaffParams) error
	VisitListTransferRules(*ListTransferRulesParams) error
	VisitOptInSmsReminder(*OptInSmsReminderParams) error
	VisitOptOutSmsReminder(*OptOutSmsReminderParams) error
	VisitOwnerGetClientAppointments(*OwnerGetClientAppointmentsParams) error
	VisitOwnerListAssets(*OwnerListAssetsParams) error
	VisitOwnerListLocations(*OwnerListLocationsParams) error
	VisitOwnerListProducts(*OwnerListProductsParams) error
	VisitOwnerListServices(*OwnerListServicesParams) error
	VisitOwnerListStaff(*OwnerListStaffParams) error
	VisitRegisterForGroupSession(*RegisterForGroupSessionParams) error
	VisitReportKnowledgeGap(*ReportKnowledgeGapParams) error
	VisitRestoreCalendarEvent(*RestoreCalendarEventParams) error
	VisitScheduleGroupSession(*ScheduleGroupSessionParams) error
	VisitSearchClients(*SearchClientsParams) error
	VisitSendCustomEmail(*SendCustomEmailParams) error
	VisitSetBookingSlug(*SetBookingSlugParams) error
	VisitSubmitBusinessInfo(*SubmitBusinessInfoParams) error
	VisitTriggerUserVerification(*TriggerUserVerificationParams) error
	VisitUpdateAgentConfig(*UpdateAgentConfigParams) error
	VisitUpdateAgentProcedure(*UpdateAgentProcedureParams) error
	VisitUpdateAgentRule(*UpdateAgentRuleParams) error
	VisitUpdateAsset(*UpdateAssetParams) error
	VisitUpdateBookingPageAppearance(*UpdateBookingPageAppearanceParams) error
	VisitUpdateBookingPageSettings(*UpdateBookingPageSettingsParams) error
	VisitUpdateBusinessInfo(*UpdateBusinessInfoParams) error
	VisitUpdateCalendarEvent(*UpdateCalendarEventParams) error
	VisitUpdateClient(*UpdateClientParams) error
	VisitUpdateGroupSessionSeats(*UpdateGroupSessionSeatsParams) error
	VisitUpdateHoliday(*UpdateHolidayParams) error
	VisitUpdateLocation(*UpdateLocationParams) error
	VisitUpdateOrder(*UpdateOrderParams) error
	VisitUpdateProduct(*UpdateProductParams) error
	VisitUpdateService(*UpdateServiceParams) error
	VisitUpdateStaff(*UpdateStaffParams) error
	VisitUpdateTransferRule(*UpdateTransferRuleParams) error
	VisitValidateUserVerificationCode(*ValidateUserVerificationCodeParams) error
}

func (s *SmbToolConfigParams) Accept(visitor SmbToolConfigParamsVisitor) error {
	if s.CancelCalendarEvent != nil {
		return visitor.VisitCancelCalendarEvent(s.CancelCalendarEvent)
	}
	if s.CancelGroupSessionForAll != nil {
		return visitor.VisitCancelGroupSessionForAll(s.CancelGroupSessionForAll)
	}
	if s.CancelGroupSessionRegistration != nil {
		return visitor.VisitCancelGroupSessionRegistration(s.CancelGroupSessionRegistration)
	}
	if s.CancelOrder != nil {
		return visitor.VisitCancelOrder(s.CancelOrder)
	}
	if s.CheckServiceAvailability != nil {
		return visitor.VisitCheckServiceAvailability(s.CheckServiceAvailability)
	}
	if s.CreateAgentProcedure != nil {
		return visitor.VisitCreateAgentProcedure(s.CreateAgentProcedure)
	}
	if s.CreateAgentRule != nil {
		return visitor.VisitCreateAgentRule(s.CreateAgentRule)
	}
	if s.CreateAsset != nil {
		return visitor.VisitCreateAsset(s.CreateAsset)
	}
	if s.CreateClient != nil {
		return visitor.VisitCreateClient(s.CreateClient)
	}
	if s.CreateClientAppointment != nil {
		return visitor.VisitCreateClientAppointment(s.CreateClientAppointment)
	}
	if s.CreateClientInteraction != nil {
		return visitor.VisitCreateClientInteraction(s.CreateClientInteraction)
	}
	if s.CreateHoliday != nil {
		return visitor.VisitCreateHoliday(s.CreateHoliday)
	}
	if s.CreateLocation != nil {
		return visitor.VisitCreateLocation(s.CreateLocation)
	}
	if s.CreateOrder != nil {
		return visitor.VisitCreateOrder(s.CreateOrder)
	}
	if s.CreateProduct != nil {
		return visitor.VisitCreateProduct(s.CreateProduct)
	}
	if s.CreateProductQuoteRequest != nil {
		return visitor.VisitCreateProductQuoteRequest(s.CreateProductQuoteRequest)
	}
	if s.CreateService != nil {
		return visitor.VisitCreateService(s.CreateService)
	}
	if s.CreateServiceQuoteRequest != nil {
		return visitor.VisitCreateServiceQuoteRequest(s.CreateServiceQuoteRequest)
	}
	if s.CreateStaff != nil {
		return visitor.VisitCreateStaff(s.CreateStaff)
	}
	if s.CreateTransferRule != nil {
		return visitor.VisitCreateTransferRule(s.CreateTransferRule)
	}
	if s.DeleteAgentProcedure != nil {
		return visitor.VisitDeleteAgentProcedure(s.DeleteAgentProcedure)
	}
	if s.DeleteAgentRule != nil {
		return visitor.VisitDeleteAgentRule(s.DeleteAgentRule)
	}
	if s.DeleteAsset != nil {
		return visitor.VisitDeleteAsset(s.DeleteAsset)
	}
	if s.DeleteCalendarEvent != nil {
		return visitor.VisitDeleteCalendarEvent(s.DeleteCalendarEvent)
	}
	if s.DeleteClient != nil {
		return visitor.VisitDeleteClient(s.DeleteClient)
	}
	if s.DeleteClientInteraction != nil {
		return visitor.VisitDeleteClientInteraction(s.DeleteClientInteraction)
	}
	if s.DeleteGroupSession != nil {
		return visitor.VisitDeleteGroupSession(s.DeleteGroupSession)
	}
	if s.DeleteHoliday != nil {
		return visitor.VisitDeleteHoliday(s.DeleteHoliday)
	}
	if s.DeleteLocation != nil {
		return visitor.VisitDeleteLocation(s.DeleteLocation)
	}
	if s.DeleteProduct != nil {
		return visitor.VisitDeleteProduct(s.DeleteProduct)
	}
	if s.DeleteService != nil {
		return visitor.VisitDeleteService(s.DeleteService)
	}
	if s.DeleteStaff != nil {
		return visitor.VisitDeleteStaff(s.DeleteStaff)
	}
	if s.DeleteTransferRule != nil {
		return visitor.VisitDeleteTransferRule(s.DeleteTransferRule)
	}
	if s.GetAgentConfig != nil {
		return visitor.VisitGetAgentConfig(s.GetAgentConfig)
	}
	if s.GetAnalyticsSummary != nil {
		return visitor.VisitGetAnalyticsSummary(s.GetAnalyticsSummary)
	}
	if s.GetAppointmentByConfirmationNumber != nil {
		return visitor.VisitGetAppointmentByConfirmationNumber(s.GetAppointmentByConfirmationNumber)
	}
	if s.GetBookingPageSettings != nil {
		return visitor.VisitGetBookingPageSettings(s.GetBookingPageSettings)
	}
	if s.GetBookingSlugStatus != nil {
		return visitor.VisitGetBookingSlugStatus(s.GetBookingSlugStatus)
	}
	if s.GetClientAppointments != nil {
		return visitor.VisitGetClientAppointments(s.GetClientAppointments)
	}
	if s.GetClientByPhone != nil {
		return visitor.VisitGetClientByPhone(s.GetClientByPhone)
	}
	if s.GetClientOrders != nil {
		return visitor.VisitGetClientOrders(s.GetClientOrders)
	}
	if s.GetOrderByConfirmationNumber != nil {
		return visitor.VisitGetOrderByConfirmationNumber(s.GetOrderByConfirmationNumber)
	}
	if s.GetSchedule != nil {
		return visitor.VisitGetSchedule(s.GetSchedule)
	}
	if s.GetWorkspacePlan != nil {
		return visitor.VisitGetWorkspacePlan(s.GetWorkspacePlan)
	}
	if s.LeaveMessage != nil {
		return visitor.VisitLeaveMessage(s.LeaveMessage)
	}
	if s.ListAgentProcedures != nil {
		return visitor.VisitListAgentProcedures(s.ListAgentProcedures)
	}
	if s.ListAgentReferences != nil {
		return visitor.VisitListAgentReferences(s.ListAgentReferences)
	}
	if s.ListAgentRules != nil {
		return visitor.VisitListAgentRules(s.ListAgentRules)
	}
	if s.ListAssets != nil {
		return visitor.VisitListAssets(s.ListAssets)
	}
	if s.ListCalendarEvents != nil {
		return visitor.VisitListCalendarEvents(s.ListCalendarEvents)
	}
	if s.ListClientInteractions != nil {
		return visitor.VisitListClientInteractions(s.ListClientInteractions)
	}
	if s.ListClients != nil {
		return visitor.VisitListClients(s.ListClients)
	}
	if s.ListCustomerFacingAgents != nil {
		return visitor.VisitListCustomerFacingAgents(s.ListCustomerFacingAgents)
	}
	if s.ListGroupSessions != nil {
		return visitor.VisitListGroupSessions(s.ListGroupSessions)
	}
	if s.ListHolidays != nil {
		return visitor.VisitListHolidays(s.ListHolidays)
	}
	if s.ListLocations != nil {
		return visitor.VisitListLocations(s.ListLocations)
	}
	if s.ListProducts != nil {
		return visitor.VisitListProducts(s.ListProducts)
	}
	if s.ListServices != nil {
		return visitor.VisitListServices(s.ListServices)
	}
	if s.ListStaff != nil {
		return visitor.VisitListStaff(s.ListStaff)
	}
	if s.ListTransferRules != nil {
		return visitor.VisitListTransferRules(s.ListTransferRules)
	}
	if s.OptInSmsReminder != nil {
		return visitor.VisitOptInSmsReminder(s.OptInSmsReminder)
	}
	if s.OptOutSmsReminder != nil {
		return visitor.VisitOptOutSmsReminder(s.OptOutSmsReminder)
	}
	if s.OwnerGetClientAppointments != nil {
		return visitor.VisitOwnerGetClientAppointments(s.OwnerGetClientAppointments)
	}
	if s.OwnerListAssets != nil {
		return visitor.VisitOwnerListAssets(s.OwnerListAssets)
	}
	if s.OwnerListLocations != nil {
		return visitor.VisitOwnerListLocations(s.OwnerListLocations)
	}
	if s.OwnerListProducts != nil {
		return visitor.VisitOwnerListProducts(s.OwnerListProducts)
	}
	if s.OwnerListServices != nil {
		return visitor.VisitOwnerListServices(s.OwnerListServices)
	}
	if s.OwnerListStaff != nil {
		return visitor.VisitOwnerListStaff(s.OwnerListStaff)
	}
	if s.RegisterForGroupSession != nil {
		return visitor.VisitRegisterForGroupSession(s.RegisterForGroupSession)
	}
	if s.ReportKnowledgeGap != nil {
		return visitor.VisitReportKnowledgeGap(s.ReportKnowledgeGap)
	}
	if s.RestoreCalendarEvent != nil {
		return visitor.VisitRestoreCalendarEvent(s.RestoreCalendarEvent)
	}
	if s.ScheduleGroupSession != nil {
		return visitor.VisitScheduleGroupSession(s.ScheduleGroupSession)
	}
	if s.SearchClients != nil {
		return visitor.VisitSearchClients(s.SearchClients)
	}
	if s.SendCustomEmail != nil {
		return visitor.VisitSendCustomEmail(s.SendCustomEmail)
	}
	if s.SetBookingSlug != nil {
		return visitor.VisitSetBookingSlug(s.SetBookingSlug)
	}
	if s.SubmitBusinessInfo != nil {
		return visitor.VisitSubmitBusinessInfo(s.SubmitBusinessInfo)
	}
	if s.TriggerUserVerification != nil {
		return visitor.VisitTriggerUserVerification(s.TriggerUserVerification)
	}
	if s.UpdateAgentConfig != nil {
		return visitor.VisitUpdateAgentConfig(s.UpdateAgentConfig)
	}
	if s.UpdateAgentProcedure != nil {
		return visitor.VisitUpdateAgentProcedure(s.UpdateAgentProcedure)
	}
	if s.UpdateAgentRule != nil {
		return visitor.VisitUpdateAgentRule(s.UpdateAgentRule)
	}
	if s.UpdateAsset != nil {
		return visitor.VisitUpdateAsset(s.UpdateAsset)
	}
	if s.UpdateBookingPageAppearance != nil {
		return visitor.VisitUpdateBookingPageAppearance(s.UpdateBookingPageAppearance)
	}
	if s.UpdateBookingPageSettings != nil {
		return visitor.VisitUpdateBookingPageSettings(s.UpdateBookingPageSettings)
	}
	if s.UpdateBusinessInfo != nil {
		return visitor.VisitUpdateBusinessInfo(s.UpdateBusinessInfo)
	}
	if s.UpdateCalendarEvent != nil {
		return visitor.VisitUpdateCalendarEvent(s.UpdateCalendarEvent)
	}
	if s.UpdateClient != nil {
		return visitor.VisitUpdateClient(s.UpdateClient)
	}
	if s.UpdateGroupSessionSeats != nil {
		return visitor.VisitUpdateGroupSessionSeats(s.UpdateGroupSessionSeats)
	}
	if s.UpdateHoliday != nil {
		return visitor.VisitUpdateHoliday(s.UpdateHoliday)
	}
	if s.UpdateLocation != nil {
		return visitor.VisitUpdateLocation(s.UpdateLocation)
	}
	if s.UpdateOrder != nil {
		return visitor.VisitUpdateOrder(s.UpdateOrder)
	}
	if s.UpdateProduct != nil {
		return visitor.VisitUpdateProduct(s.UpdateProduct)
	}
	if s.UpdateService != nil {
		return visitor.VisitUpdateService(s.UpdateService)
	}
	if s.UpdateStaff != nil {
		return visitor.VisitUpdateStaff(s.UpdateStaff)
	}
	if s.UpdateTransferRule != nil {
		return visitor.VisitUpdateTransferRule(s.UpdateTransferRule)
	}
	if s.ValidateUserVerificationCode != nil {
		return visitor.VisitValidateUserVerificationCode(s.ValidateUserVerificationCode)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SmbToolConfigParams) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.CancelCalendarEvent != nil {
		fields = append(fields, "cancel_calendar_event")
	}
	if s.CancelGroupSessionForAll != nil {
		fields = append(fields, "cancel_group_session_for_all")
	}
	if s.CancelGroupSessionRegistration != nil {
		fields = append(fields, "cancel_group_session_registration")
	}
	if s.CancelOrder != nil {
		fields = append(fields, "cancel_order")
	}
	if s.CheckServiceAvailability != nil {
		fields = append(fields, "check_service_availability")
	}
	if s.CreateAgentProcedure != nil {
		fields = append(fields, "create_agent_procedure")
	}
	if s.CreateAgentRule != nil {
		fields = append(fields, "create_agent_rule")
	}
	if s.CreateAsset != nil {
		fields = append(fields, "create_asset")
	}
	if s.CreateClient != nil {
		fields = append(fields, "create_client")
	}
	if s.CreateClientAppointment != nil {
		fields = append(fields, "create_client_appointment")
	}
	if s.CreateClientInteraction != nil {
		fields = append(fields, "create_client_interaction")
	}
	if s.CreateHoliday != nil {
		fields = append(fields, "create_holiday")
	}
	if s.CreateLocation != nil {
		fields = append(fields, "create_location")
	}
	if s.CreateOrder != nil {
		fields = append(fields, "create_order")
	}
	if s.CreateProduct != nil {
		fields = append(fields, "create_product")
	}
	if s.CreateProductQuoteRequest != nil {
		fields = append(fields, "create_product_quote_request")
	}
	if s.CreateService != nil {
		fields = append(fields, "create_service")
	}
	if s.CreateServiceQuoteRequest != nil {
		fields = append(fields, "create_service_quote_request")
	}
	if s.CreateStaff != nil {
		fields = append(fields, "create_staff")
	}
	if s.CreateTransferRule != nil {
		fields = append(fields, "create_transfer_rule")
	}
	if s.DeleteAgentProcedure != nil {
		fields = append(fields, "delete_agent_procedure")
	}
	if s.DeleteAgentRule != nil {
		fields = append(fields, "delete_agent_rule")
	}
	if s.DeleteAsset != nil {
		fields = append(fields, "delete_asset")
	}
	if s.DeleteCalendarEvent != nil {
		fields = append(fields, "delete_calendar_event")
	}
	if s.DeleteClient != nil {
		fields = append(fields, "delete_client")
	}
	if s.DeleteClientInteraction != nil {
		fields = append(fields, "delete_client_interaction")
	}
	if s.DeleteGroupSession != nil {
		fields = append(fields, "delete_group_session")
	}
	if s.DeleteHoliday != nil {
		fields = append(fields, "delete_holiday")
	}
	if s.DeleteLocation != nil {
		fields = append(fields, "delete_location")
	}
	if s.DeleteProduct != nil {
		fields = append(fields, "delete_product")
	}
	if s.DeleteService != nil {
		fields = append(fields, "delete_service")
	}
	if s.DeleteStaff != nil {
		fields = append(fields, "delete_staff")
	}
	if s.DeleteTransferRule != nil {
		fields = append(fields, "delete_transfer_rule")
	}
	if s.GetAgentConfig != nil {
		fields = append(fields, "get_agent_config")
	}
	if s.GetAnalyticsSummary != nil {
		fields = append(fields, "get_analytics_summary")
	}
	if s.GetAppointmentByConfirmationNumber != nil {
		fields = append(fields, "get_appointment_by_confirmation_number")
	}
	if s.GetBookingPageSettings != nil {
		fields = append(fields, "get_booking_page_settings")
	}
	if s.GetBookingSlugStatus != nil {
		fields = append(fields, "get_booking_slug_status")
	}
	if s.GetClientAppointments != nil {
		fields = append(fields, "get_client_appointments")
	}
	if s.GetClientByPhone != nil {
		fields = append(fields, "get_client_by_phone")
	}
	if s.GetClientOrders != nil {
		fields = append(fields, "get_client_orders")
	}
	if s.GetOrderByConfirmationNumber != nil {
		fields = append(fields, "get_order_by_confirmation_number")
	}
	if s.GetSchedule != nil {
		fields = append(fields, "get_schedule")
	}
	if s.GetWorkspacePlan != nil {
		fields = append(fields, "get_workspace_plan")
	}
	if s.LeaveMessage != nil {
		fields = append(fields, "leave_message")
	}
	if s.ListAgentProcedures != nil {
		fields = append(fields, "list_agent_procedures")
	}
	if s.ListAgentReferences != nil {
		fields = append(fields, "list_agent_references")
	}
	if s.ListAgentRules != nil {
		fields = append(fields, "list_agent_rules")
	}
	if s.ListAssets != nil {
		fields = append(fields, "list_assets")
	}
	if s.ListCalendarEvents != nil {
		fields = append(fields, "list_calendar_events")
	}
	if s.ListClientInteractions != nil {
		fields = append(fields, "list_client_interactions")
	}
	if s.ListClients != nil {
		fields = append(fields, "list_clients")
	}
	if s.ListCustomerFacingAgents != nil {
		fields = append(fields, "list_customer_facing_agents")
	}
	if s.ListGroupSessions != nil {
		fields = append(fields, "list_group_sessions")
	}
	if s.ListHolidays != nil {
		fields = append(fields, "list_holidays")
	}
	if s.ListLocations != nil {
		fields = append(fields, "list_locations")
	}
	if s.ListProducts != nil {
		fields = append(fields, "list_products")
	}
	if s.ListServices != nil {
		fields = append(fields, "list_services")
	}
	if s.ListStaff != nil {
		fields = append(fields, "list_staff")
	}
	if s.ListTransferRules != nil {
		fields = append(fields, "list_transfer_rules")
	}
	if s.OptInSmsReminder != nil {
		fields = append(fields, "opt_in_sms_reminder")
	}
	if s.OptOutSmsReminder != nil {
		fields = append(fields, "opt_out_sms_reminder")
	}
	if s.OwnerGetClientAppointments != nil {
		fields = append(fields, "owner_get_client_appointments")
	}
	if s.OwnerListAssets != nil {
		fields = append(fields, "owner_list_assets")
	}
	if s.OwnerListLocations != nil {
		fields = append(fields, "owner_list_locations")
	}
	if s.OwnerListProducts != nil {
		fields = append(fields, "owner_list_products")
	}
	if s.OwnerListServices != nil {
		fields = append(fields, "owner_list_services")
	}
	if s.OwnerListStaff != nil {
		fields = append(fields, "owner_list_staff")
	}
	if s.RegisterForGroupSession != nil {
		fields = append(fields, "register_for_group_session")
	}
	if s.ReportKnowledgeGap != nil {
		fields = append(fields, "report_knowledge_gap")
	}
	if s.RestoreCalendarEvent != nil {
		fields = append(fields, "restore_calendar_event")
	}
	if s.ScheduleGroupSession != nil {
		fields = append(fields, "schedule_group_session")
	}
	if s.SearchClients != nil {
		fields = append(fields, "search_clients")
	}
	if s.SendCustomEmail != nil {
		fields = append(fields, "send_custom_email")
	}
	if s.SetBookingSlug != nil {
		fields = append(fields, "set_booking_slug")
	}
	if s.SubmitBusinessInfo != nil {
		fields = append(fields, "submit_business_info")
	}
	if s.TriggerUserVerification != nil {
		fields = append(fields, "trigger_user_verification")
	}
	if s.UpdateAgentConfig != nil {
		fields = append(fields, "update_agent_config")
	}
	if s.UpdateAgentProcedure != nil {
		fields = append(fields, "update_agent_procedure")
	}
	if s.UpdateAgentRule != nil {
		fields = append(fields, "update_agent_rule")
	}
	if s.UpdateAsset != nil {
		fields = append(fields, "update_asset")
	}
	if s.UpdateBookingPageAppearance != nil {
		fields = append(fields, "update_booking_page_appearance")
	}
	if s.UpdateBookingPageSettings != nil {
		fields = append(fields, "update_booking_page_settings")
	}
	if s.UpdateBusinessInfo != nil {
		fields = append(fields, "update_business_info")
	}
	if s.UpdateCalendarEvent != nil {
		fields = append(fields, "update_calendar_event")
	}
	if s.UpdateClient != nil {
		fields = append(fields, "update_client")
	}
	if s.UpdateGroupSessionSeats != nil {
		fields = append(fields, "update_group_session_seats")
	}
	if s.UpdateHoliday != nil {
		fields = append(fields, "update_holiday")
	}
	if s.UpdateLocation != nil {
		fields = append(fields, "update_location")
	}
	if s.UpdateOrder != nil {
		fields = append(fields, "update_order")
	}
	if s.UpdateProduct != nil {
		fields = append(fields, "update_product")
	}
	if s.UpdateService != nil {
		fields = append(fields, "update_service")
	}
	if s.UpdateStaff != nil {
		fields = append(fields, "update_staff")
	}
	if s.UpdateTransferRule != nil {
		fields = append(fields, "update_transfer_rule")
	}
	if s.ValidateUserVerificationCode != nil {
		fields = append(fields, "validate_user_verification_code")
	}
	if len(fields) == 0 {
		if s.SmbToolType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.SmbToolType)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.SmbToolType != "" {
		field := fields[0]
		if s.SmbToolType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.SmbToolType,
				s,
			)
		}
	}
	return nil
}
