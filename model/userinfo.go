package model

type UserInfo struct {
	authinfo string // need
}

func (u *UserInfo) AuthInfo() string {
	return u.authinfo
}
