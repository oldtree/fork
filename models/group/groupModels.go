package group

import (
	"encoding/json"
	"fork/models"
	"fork/tools/log"

	"github.com/astaxie/beego/orm"
)

func InitGroupModels() {

}

func init() {
	models.RegisterModels(new(Group))
}

type Group struct {
	ID         int64  `json:"uid" orm:"uid;pk;auto"`
	GroupID    string `json:"groupid" orm:"groupid"`
	GnroupName string `json:"groupname" orm:"groupname"`
	Desc       string `json:"desc" orm:"desc"`
	GroupType  int    `json:"grouptype" orm:"grouptype"`
	GroupStatu int    `json:"groupstatu" orm:"groupstatu"`
	Created    int64  `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated    int64  `json:"updated" orm:"auto_now;type(datetime)"`
}

func (g *Group) GetGroupByName() (bool, error) {
	object := orm.NewOrm()
	err := object.Read(g, "groupname")
	if err != nil {
		if err == orm.ErrNoRows {
			return true, err
		}
		return false, err
	}

	return true, nil
}

func (g *Group) GetGroupById() (bool, error) {
	object := orm.NewOrm()
	err := object.Read(g, "groupid")
	if err != nil {
		if err == orm.ErrNoRows {
			return true, err
		}
		return false, err
	}
	return true, nil
}

func QueyGroupListByCondition(sqlquery string) (result []*Group, err error) {
	ob := orm.NewOrm()
	_, err = ob.Raw(sqlquery).QueryRows(result)
	if err != nil {
		return nil, nil
	}
	return
}

func (g *Group) FormatToJson() ([]byte, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) GetGroup() (bool, error) {
	object := orm.NewOrm()
	err := object.Read(g)
	if err != nil {
		if err == orm.ErrNoRows {
			return true, err
		}
		return false, err
	}

	return true, nil
}

func (g *Group) InsertGroup() (bool, error) {
	object := orm.NewOrm()
	numb, err := object.Insert(g)
	log.Info(numb, err)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *Group) UpdateGroup() (bool, error) {
	object := orm.NewOrm()
	numb, err := object.Update(g)
	log.Info(numb, err)
	if err != nil {
		if err == orm.ErrNoRows {
			return true, err
		}
		return false, err
	}
	return true, nil
}

func (g *Group) DeleteGroup() (bool, error) {
	object := orm.NewOrm()
	num, err := object.Delete(g)
	log.Info(num, err)
	if err != nil {
		if err == orm.ErrNoRows {
			return true, err
		}
		return false, err
	}
	return true, nil
}
