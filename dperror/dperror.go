package dperror

import (
	"strings"

	"github.com/pkg/errors"
)

func Is(err error, targets ...error) bool {
	if err == nil {
		return false
	}

	for _, target := range targets {
		if strings.Contains(errors.Cause(err).Error(), target.Error()) {
			return true
		}
	}
	return false
}

var NOT_FOUND = errors.New("not found")
var FORBIDDEN = errors.New("invoking member lacks authority")
var INVALID_RESET_TOKEN = errors.New("invalid password reset token")
var INVALID_PASSWORD = errors.New("invalid password")
var ALREADY_IN_USE = errors.New("already in use")
var DEACTIVATED = errors.New("deactivated")
var INVALID_DATA = errors.New("invalid data")
var DB_ERROR = errors.New("database error")
var EXPIRED = errors.New("expired")
var UNKNOWN_DEVICE = errors.New("unknown device")
var FUNC_NOT_ALLOWED = errors.New("function is not allowed on this contract")
var CONTRACT_NOT_ALLOWED = errors.New("contract not allowed by platform on network")
