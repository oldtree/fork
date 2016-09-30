package user

import (
	"encoding/json"
	"fork/models"
	"fork/tools/log"

	"github.com/astaxie/beego/orm"
)

func InitUserModels() {
	models.RegisterModels(new(Person))
}

func init() {
	models.RegisterModels(new(Person))
}

//Person : common person models
type Person struct {
	ID       int64  `json:"id" orm:"id;pk;auto"`
	UID      string `json:"uid" orm:"uid"`
	Email    string `json:"email" orm:"email"`
	Password string `json:"password" orm:"password"`
	Phone    string `json:"phone" orm:"phone"`
	Age      int    `json:"age" orm:"age"`
	Location string `json:"location" orm:"location"`

	Delete    bool   `json:"delete" orm:"del"`
	Disable   bool   `json:"disable" orm:"disable"`
	Education string `json:"edu" orm:"edu"`

	Nickname string `json:"nickname" orm:"nickname"`
	RealName string `json:"realname" orm:"realname"`

	Created int64 `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated int64 `json:"updated" orm:"auto_now;type(datetime)"`
}

//Format : format to json form
func (p *Person) FormatToJson() ([]byte, error) {
	content, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return content, nil
}

//NewPerson : new person object
func NewPerson() *Person {
	return &Person{}
}

//InsertPerson : insert person
func (p *Person) InsertPerson() (bool, error) {
	object := orm.NewOrm()
	statu, err := object.Insert(p)
	if err != nil {
		log.Error(statu, err)
		return false, err
	}
	return true, nil
}

//UpdatePerson : update person
func (p *Person) UpdatePerson() (bool, error) {
	object := orm.NewOrm()
	statu, err := object.Update(p)
	if err != nil {
		log.Error(statu, err)
		return false, err
	}
	return true, nil
}

//DeletePerson : delete person
func (p *Person) DeletePerson() (bool, error) {
	object := orm.NewOrm()
	statu, err := object.Delete(p)
	if err != nil {
		log.Error(statu, err)
		return false, err
	}
	return true, nil
}

//GetPerson : get person
func (p *Person) GetPerson() (bool, error) {
	object := orm.NewOrm()
	err := object.Read(p)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Info(err)
			return false, err
		}
	}
	if p.Delete == true {
		return false, nil
	}
	return true, nil
}
