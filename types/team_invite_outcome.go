package types

import (
	"fmt"
)

type TeamInviteOutcome string

const (
	TeamInviteOutcomeInvited TeamInviteOutcome = "invited"
	TeamInviteOutcomeFailed  TeamInviteOutcome = "failed"
)

func NewTeamInviteOutcomeFromString(s string) (TeamInviteOutcome, error) {
	switch s {
	case "invited":
		return TeamInviteOutcomeInvited, nil
	case "failed":
		return TeamInviteOutcomeFailed, nil
	}
	var t TeamInviteOutcome
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TeamInviteOutcome) Ptr() *TeamInviteOutcome {
	return &e
}
