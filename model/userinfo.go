package model

type UserInfo struct {
	authinfo string // need
}

func CreateUserinfo(ai string) *UserInfo {
	return &UserInfo{authinfo: ai}
}

func (u *UserInfo) AuthInfo() string {
	return u.authinfo
}
