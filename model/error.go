package model

import (
	"errors"
)

type SessionNotFoundError struct {
	msg string
}

func (e *SessionNotFoundError) Error() string {
	return e.msg
}

type SessionExpiresError struct {
	msg string
}

func (e *SessionExpiresError) Error() string {
	return e.msg
}

type SessionInvalidError struct {
	msg string
}

func (e *SessionInvalidError) Error() string {
	return e.msg
}

// session
var ErrSessionNotFound = errors.New("Session not found")
var ErrSessionExpires = errors.New("Session expiress")
var ErrSessionInvalid = errors.New("The token you're using is invalid")

//register
type UserAlreadyExistError struct {
	msg string
}

func (e *UserAlreadyExistError) Error() string {
	return e.msg
}

type TooWeakPasswordError struct {
	msg string
}

func (e *TooWeakPasswordError) Error() string {
	return e.msg
}

type InvalidMailError struct {
	msg string
}

func (e *InvalidMailError) Error() string {
	return e.msg
}

var ErrUserAlreadyExist = errors.New("Requested user already exists")
var ErrTooWeakPassword = errors.New("Requested password is too weak")
var ErrInvalidMail = errors.New("Requested email address is invalid")

// login
var ErrUserNotFound = errors.New("User not found")
var ErrIncorrectPassword = errors.New("Password is incorrect")
var ErrAuthenticationFailed = errors.New("Authentication failed")
var ErrAlreadyLogin = errors.New("User already logged in")

type UserNotFoundError struct {
	msg string
}

func (e *UserNotFoundError) Error() string {
	return e.msg
}

type IncorrectPassword struct {
	msg string
}

func (e *IncorrectPassword) Error() string {
	return e.msg
}

type AuthenticationFailedError struct {
	msg string
}

func (e *AuthenticationFailedError) Error() string {
	return e.msg
}

type AlreadyLoginError struct {
	msg string
}

func (e *AlreadyLoginError) Error() string {
	return e.msg
}

// fetch
var ErrCommandNotFound = errors.New("Requested Command not found")

// other error
var ErrServerSystem = errors.New("Comstock server has internal error")
