package user

import (
	"fmt"
	"fork/tools/log"

	"errors"
	"fork/models"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

const (
	LoginPassword = 1
	LoginWeixin   = 2
	LoginQQ       = 3
	LogingWeibo   = 4
)

var loginorm = orm.NewOrm()

func InitLoginModels() {
	log.Info("init login models")
}

func init() {
	models.RegisterModels(new(LoginUser))
}

func GetUUID() string {
	var uidgen = uuid.NewV4()
	return uidgen.String()
}

//LoginUser : used fo login info pass by
type LoginUser struct {
	ID       int64  `orm:"id;pk;auto" json:"id"`
	UID      string `form:"uid" json:"uid"`
	NickName string `from:"nickname" json:"nickname" orm:"nickname"`
	Account  string `form:"account" json:"account" binding:"account"  orm:"nickname"`
	Password string `form:"password" json:"-" binding:"required"  orm:"nickname"`

	LoginType int    `orm:"logintype" json:"logintype"`
	ThirdId   string `orm:"thirdid" json:"thirdid"`

	Created int64 `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated int64 `json:"updated" orm:"auto_now;type(datetime)"`
}

//NewLoginUser : new loginmodels
func NewLoginUser() *LoginUser {
	return &LoginUser{}
}

const (
	ErrorCodeNotSet                    = 10000
	ErrorCodeUserNotExist              = 10001
	ErrorCodeUserHasExist              = 10002
	ErrorCodeUserUnActived             = 10003
	ErrorCodeUserPasswrodNotRight      = 10004
	ErrorCodeUserPasswordConfirmFailed = 10005
	ErrorCodeUserInternelError         = -1
)

//CheckLogin :check user login
func (l *LoginUser) PasswordLogin() (bool, int, error) {
	sql := "select * from loginuser where account = %s;"
	sql = fmt.Sprintf(sql, l.Account)
	temp := l.Password
	ok, statucode, err := l.QueryFunc(sql)
	if !ok {
		if statucode == ErrorCodeUserInternelError {
			return false, ErrorCodeUserInternelError, err
		}
	}
	if err != nil {
		return false, ErrorCodeUserNotExist, errors.New("user not exist")
	}
	if l.Password != temp {
		return false, ErrorCodeUserPasswrodNotRight, errors.New("password not right")
	}
	return true, ErrorCodeNotSet, nil
}

func (l *LoginUser) ThirdPartLogin() (bool, int, error) {
	sql := "select * from loginuser where account = %s and password = %s ;"
	sql = fmt.Sprintf(sql, l.Account, l.Password)

	return l.QueryFunc(sql)
}

func (l *LoginUser) RegisterUser() (bool, int, error) {
	sql := "insert into loginuser (uid,account,password) values (%s,%s ,%s) ;"
	l.UID = GetUUID()
	sql = fmt.Sprintf(sql, l.UID, l.Account, l.Password)
	ok, statucode, err := l.AddLoginFunc(sql)
	if !ok {
		if statucode == ErrorCodeUserHasExist {
			return false, ErrorCodeUserHasExist, err
		} else {
			return false, ErrorCodeUserInternelError, err
		}
	}
	return true, ErrorCodeNotSet, nil
}

func (l *LoginUser) ThrirdRegister() (bool, int, error) {
	sql := "insert into loginuser (account,password) values (%s ,%s) ;"
	sql = fmt.Sprintf(sql, l.Account, l.Password)
	return l.AddLoginFunc(sql)
}
func (l *LoginUser) QueryFunc(sqlquery string) (bool, int, error) {
	err := loginorm.Raw(sqlquery).QueryRow(l)
	if err != nil {
		if err == orm.ErrNoRows {
			return false, 1011, nil
		} else {
			return false, -1, err
		}
	}
	return true, 1000, nil
}
func (l *LoginUser) AddLoginFunc(sqlquery string) (bool, int, error) {
	re, err := loginorm.Raw(sqlquery).Exec()
	log.Info(re)
	if err != nil {
		return false, -1, errors.New("add user failed")
	}
	return true, 1000, nil
}
