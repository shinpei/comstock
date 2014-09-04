package model

// session
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

// login

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
type CommandNotFoundError struct {
	msg string
}

func (e *CommandNotFoundError) Error() string {
	return e.msg
}

// other error
type ServerSystemError struct {
	msg string
}

func (e *ServerSystemError) Error() string {
	return e.msg
}
