package route

import "entrance/models"

/**
 获取数据库设置的所有url配置数组
 */
func GetRouteConfig() map[string]interface{} {

	var url models.ServiceUrl
	var api models.ServiceApi
	urldata := url.List()
	apidata := api.List()

	return getHandleApi(apidata,urldata)
}

/**
 获取格式化后的api数据
 */
func getHandleApi(api []*models.ServiceApi,url []*models.ServiceUrl) map[string]interface{} {
	data := make(map[string]interface{})
	u := getHandleUrl(url)
	for _,a := range api{
		formatA := make(map[string]interface{})
		formatA["Method"] = a.Method
		formatA["ServiceUrl"] = u[a.ServiceName]
		formatA["ApiAlias"] = a.ApiAlias
		formatA["ApiPath"] = a.ApiPath
		formatA["InnerPath"] = a.InnerPath

		data["["+a.Method+"]"+a.ApiAlias] = formatA
	}
	return data
}

/**
 获取格式化的url数据
 */
func getHandleUrl(url []*models.ServiceUrl) map[string]string {
	data := make(map[string]string)
	for _,u := range url{
		data[u.ServiceName] = u.ServiceUrl
	}
	return data
}

/**
 解析请求路由
 */
func GetCurrentRoute() string {
	route := "asdf"
	return route
}

/**
 判断当前请求是否存在
 */
func ParseRoute(route string,config []interface{}) bool {
	return true
}
