package model

import (
	"errors"
)

// session
var ErrSessionNotFound = errors.New("Session not found")
var ErrSessionExpires = errors.New("Session expiress")
var ErrSessionInvalid = errors.New("The token you're using is invalid")

// login
var ErrUserNotFound = errors.New("User not found")
var ErrIncorrectPassword = errors.New("Password is incorrect")
var ErrAuthenticationFailed = errors.New("Authentication failed")
var ErrAlreadyLogin = errors.New("User already logged in")

// other error
var ErrServerSystem = errors.New("Comstock server has internal error")
