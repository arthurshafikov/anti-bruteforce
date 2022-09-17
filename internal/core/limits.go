package core

type AuthorizeLimits struct {
	LimitAttemptsForLogin, LimitAttemptsForPassword, LimitAttemptsForIP int64
}
