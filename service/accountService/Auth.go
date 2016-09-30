package accountService

import (
	"fork/tools/log"

	"fork/models/user"

	"errors"

	"github.com/boltdb/bolt"
)

var CacheDB *bolt.DB
var defaultCache string

func Init() {
	dbFileName := "/tmp/bolt.db"
	var err error
	CacheDB, err = bolt.Open(dbFileName, 0666, nil)
	if err != nil {
		log.Error("bolt db open eror : ", err.Error())
	}
	defaultCache = "cache"
}

func CloseCache() {
	err := CacheDB.Close()
	if err != nil {
		log.Error(err.Error())
	}
}

//Auth :
type Auth struct {
	Uid       string
	Account   string
	Password  string
	Nickname  string
	AuthToken string
}

//CheckAuth : login statu check
func (u *Auth) CheckAuth(sessionid string) (bool, int, error) {
	if u.AuthToken != "" {
		return false, -1, errors.New("AuthToken is empty")
	}
	return true, 0, nil
}

//Login : 登陆
func (u *Auth) Login(accont string, password string) (bool, int, error) {
	log.Info(accont, password)
	logon := new(user.LoginUser)
	logon.Account = accont
	logon.Password = password
	ok, statucode, err := logon.PasswordLogin()
	if ok == false {
		return false, statucode, err
	}
	u.Uid = logon.UID
	return true, 0, nil
}

func (u *Auth) Register(accont string, password string, nickname string) (bool, int, error) {
	log.Info(accont, password)
	/****************注册登录表****************/
	logon := new(user.LoginUser)

	logon.Account = accont
	logon.Password = password
	ok, statucode, err := logon.RegisterUser()
	if ok == false {
		return false, statucode, err
	}
	/****************注册用户信息表****************/
	user := new(user.Person)
	user.Email = accont
	user.Password = password
	user.UID = logon.UID
	user.Disable = false
	ok, err = user.InsertPerson()
	if !ok {
		log.Error("user register error : ", user, err.Error())
	}
	return true, 0, nil
}
