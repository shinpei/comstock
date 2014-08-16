package model

type AuthInfo struct {
	token string // need
	mail  string
}

func CreateUserinfo(ai string, e string) *AuthInfo {
	return &AuthInfo{token: ai, mail: e}
}

func (u *AuthInfo) Token() string {
	return u.token
}

func (u *AuthInfo) Mail() string {
	return u.mail
}
