package model

import (
	"errors"
)

var ErrSessionExpires = errors.New("Session expiress")
var ErrSessionInvalid = errors.New("The token you're using is invalid")
