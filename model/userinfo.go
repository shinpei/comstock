package model

type UserInfo struct {
	authinfo string // need
	email    string
}

func CreateUserinfo(ai string, e string) *UserInfo {
	return &UserInfo{authinfo: ai, email: e}
}

func (u *UserInfo) AuthInfo() string {
	return u.authinfo
}

func (u *UserInfo) Email() string {
	return u.email
}
