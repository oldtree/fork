//service 管理者，提供可用的service信息
package service

import "fork/tools/log"
import "errors"

var sm *ServiceManage

func init() {
	log.Info("service package init")
	sm = new(ServiceManage)
	sm.ServiceList = make(map[string]string)

}

type ServiceManage struct {
	ServiceList map[string]string
}

func (s *ServiceManage) AddService(srvname string) (bool, error) {
	if _, ok := s.ServiceList[srvname]; !ok {
		return false, errors.New("service is exist")
	}
	s.ServiceList[srvname] = srvname
	return true, nil
}
