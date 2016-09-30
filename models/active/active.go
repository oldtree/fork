package active

import (
	"fork/models"
	"fork/tools/log"
)

func initActiveModels() {
	log.Info("init active models")
}

func init() {
	models.RegisterModels(new(Active))
}

type Active struct {
	ID        int64  `json:"id" orm:"uid;pk;auto"`
	AID       string `json:"actid" orm:"actid"`
	Name      string `json:"name" orm:"name"`
	Desc      string `json:"desc" orm:"desc"`
	ActType   int    `json:"type" orm:"acttype"`
	ActStatus int    `json:"status" orm:"actstatus"`
	Created   int64  `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated   int64  `json:"updated" orm:"auto_now;type(datetime)"`
}

func (a *Active) GetActiveByAid() (bool, int, error) {
	return false, -1, nil
}
