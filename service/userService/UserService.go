package userService

import (
	"fork/models/user"
	"fork/tools/log"
)

//UserService : 用户的登陆，注册，个人设置取得与更改，环境切换与增加，
type UserService struct {
	UID          int64
	UserNickname string
	Account      string
	Password     string
}

//RegisterAction : 注册
func (u *UserService) UpdateProfile(nickname string, account string, password string) (bool, int, error) {
	p := new(user.Person)
	p.Nickname = nickname
	p.Email = account
	p.Password = password
	ok, err := p.GetPerson()
	if ok {
		return false, 2, err
	}
	ok, err = p.InsertPerson()
	if !ok {
		return false, 2, err
	}
	return true, 0, nil
}

//GetUserSetting : get setting profile
func (u *UserService) GetUserProfile() (bool, int, error) {
	log.Info("GetUserSetting")
	return true, 0, nil
}
