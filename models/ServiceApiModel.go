package models

import "github.com/astaxie/beego/orm"

type ServiceApi struct {
	Id          int
	ServiceName string
	Method      string
	ApiAlias    string
	ApiPath     string
	InnerPath   string
}

/**
获取所有的数据
 */
func (*ServiceApi) List() []*ServiceApi {
	var serviceApi []*ServiceApi
	query := orm.NewOrm().QueryTable(TableName("service_api"))
	query.All(&serviceApi)
	return serviceApi
}
