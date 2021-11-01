package models

type AuthorizeLimits struct {
	LimitAttemptsForLogin, LimitAttemptsForPassword, LimitAttemptsForIP int64
}
