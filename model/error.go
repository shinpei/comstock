package model

import (
	"errors"
)

var ErrSessionNotFound = errors.New("Session not found")
var ErrSessionExpires = errors.New("Session expiress")
var ErrSessionInvalid = errors.New("The token you're using is invalid")
