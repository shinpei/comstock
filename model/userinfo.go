package model

type UserInfo struct {
	authinfo string // need
	mail     string
}

func CreateUserinfo(ai string, e string) *UserInfo {
	return &UserInfo{authinfo: ai, mail: e}
}

func (u *UserInfo) AuthInfo() string {
	return u.authinfo
}

func (u *UserInfo) Mail() string {
	return u.mail
}
