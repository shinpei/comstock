package model

import (
	"errors"
)

// session
var ErrSessionNotFound = errors.New("Session not found")
var ErrSessionExpires = errors.New("Session expiress")
var ErrSessionInvalid = errors.New("The token you're using is invalid")

//register
var ErrUserAlreadyExist = errors.New("Requested user already exists")
var ErrTooWeakPassword = errors.New("Requested password is too weak")
var ErrInvalidMail = errors.New("Requested email address is invalid")

// login
var ErrUserNotFound = errors.New("User not found")
var ErrIncorrectPassword = errors.New("Password is incorrect")
var ErrAuthenticationFailed = errors.New("Authentication failed")
var ErrAlreadyLogin = errors.New("User already logged in")

// fetch
var ErrCommandNotFound = errors.New("Requested Command not found")

// other error
var ErrServerSystem = errors.New("Comstock server has internal error")
