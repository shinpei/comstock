package model

type SessionExpiresError struct {
	msg string
}

func (e *SessionExpiresError) Error() string {
	return e.msg
}

func (e *SessionExpiresError) SetError(msg string) *SessionExpiresError {
	e.msg = msg
	return e
}

type SessionNotFoundError struct {
	msg string
}

func (e *SessionNotFoundError) Error() string {
	return e.msg
}

func (e *SessionNotFoundError) SetError(msg string) *SessionNotFoundError {
	e.msg = msg
	return e
}

type SessionInvalidError struct {
	msg string
}

func (e *SessionInvalidError) Error() string {
	return e.msg
}

func (e *SessionInvalidError) SetError(msg string) *SessionInvalidError {
	e.msg = msg
	return e
}

type UserAlreadyExistError struct {
	msg string
}

func (e *UserAlreadyExistError) Error() string {
	return e.msg
}

func (e *UserAlreadyExistError) SetError(msg string) *UserAlreadyExistError {
	e.msg = msg
	return e
}

type TooWeakPasswordError struct {
	msg string
}

func (e *TooWeakPasswordError) Error() string {
	return e.msg
}

func (e *TooWeakPasswordError) SetError(msg string) *TooWeakPasswordError {
	e.msg = msg
	return e
}

type InvalidMailError struct {
	msg string
}

func (e *InvalidMailError) Error() string {
	return e.msg
}

func (e *InvalidMailError) SetError(msg string) *InvalidMailError {
	e.msg = msg
	return e
}

type UserNotFoundError struct {
	msg string
}

func (e *UserNotFoundError) Error() string {
	return e.msg
}

func (e *UserNotFoundError) SetError(msg string) *UserNotFoundError {
	e.msg = msg
	return e
}

type IncorrectPasswordError struct {
	msg string
}

func (e *IncorrectPasswordError) Error() string {
	return e.msg
}

func (e *IncorrectPasswordError) SetError(msg string) *IncorrectPasswordError {
	e.msg = msg
	return e
}

type AuthenticationFailedError struct {
	msg string
}

func (e *AuthenticationFailedError) Error() string {
	return e.msg
}

func (e *AuthenticationFailedError) SetError(msg string) *AuthenticationFailedError {
	e.msg = msg
	return e
}

type CommandNotFoundError struct {
	msg string
}

func (e *CommandNotFoundError) Error() string {
	return e.msg
}

func (e *CommandNotFoundError) SetError(msg string) *CommandNotFoundError {
	e.msg = msg
	return e
}

type ServerSystemError struct {
	msg string
}

func (e *ServerSystemError) Error() string {
	return e.msg
}

func (e *ServerSystemError) SetError(msg string) *ServerSystemError {
	e.msg = msg
	return e
}
