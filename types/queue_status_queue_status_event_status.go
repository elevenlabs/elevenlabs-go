package types

import (
	"fmt"
)

type QueueStatusQueueStatusEventStatus string

const (
	QueueStatusQueueStatusEventStatusWaiting  QueueStatusQueueStatusEventStatus = "waiting"
	QueueStatusQueueStatusEventStatusAdmitted QueueStatusQueueStatusEventStatus = "admitted"
	QueueStatusQueueStatusEventStatusTimedOut QueueStatusQueueStatusEventStatus = "timed_out"
)

func NewQueueStatusQueueStatusEventStatusFromString(s string) (QueueStatusQueueStatusEventStatus, error) {
	switch s {
	case "waiting":
		return QueueStatusQueueStatusEventStatusWaiting, nil
	case "admitted":
		return QueueStatusQueueStatusEventStatusAdmitted, nil
	case "timed_out":
		return QueueStatusQueueStatusEventStatusTimedOut, nil
	}
	var t QueueStatusQueueStatusEventStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e QueueStatusQueueStatusEventStatus) Ptr() *QueueStatusQueueStatusEventStatus {
	return &e
}
