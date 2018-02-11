package customer

import (
	"opms/controllers"
	"strings"
	."opms/models/customer"
)

type ManagerCustomerController struct {
	controllers.BaseController
}

func (this *ManagerCustomerController) Get() {

	if !strings.Contains(this.GetSession("userPermission").(string), "permission-manage") {
		this.Abort("401")
	}
	this.TplName = "customer/index.tpl"
}

/***
	合同添加
 */
type AddCustomerController struct {
	controllers.BaseController
}

/*
	跳转方法
 */
func (this *AddCustomerController) Get() {

	if !strings.Contains(this.GetSession("userPermission").(string), "permission-manage") {
		this.Abort("401")
	}
	//获取性别列表

	//sexArray = GetSex()
	//获取客户类型列表







	this.TplName = "customer/customer-add.tpl"
}
