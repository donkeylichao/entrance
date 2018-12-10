package models

import "github.com/astaxie/beego/orm"

type ServiceUrl struct {
	Id          int
	ServiceName string
	ServiceUrl  string
}

/**
获取所有列表
 */
func (*ServiceUrl) List() []*ServiceUrl {
	var serviceUrl []*ServiceUrl
	orm.NewOrm().QueryTable(TableName("service_url")).All(&serviceUrl)
	return serviceUrl
}
