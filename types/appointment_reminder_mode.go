package types

import (
	"fmt"
)

type AppointmentReminderMode string

const (
	AppointmentReminderModeReminder     AppointmentReminderMode = "reminder"
	AppointmentReminderModeConfirmation AppointmentReminderMode = "confirmation"
)

func NewAppointmentReminderModeFromString(s string) (AppointmentReminderMode, error) {
	switch s {
	case "reminder":
		return AppointmentReminderModeReminder, nil
	case "confirmation":
		return AppointmentReminderModeConfirmation, nil
	}
	var t AppointmentReminderMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AppointmentReminderMode) Ptr() *AppointmentReminderMode {
	return &e
}
