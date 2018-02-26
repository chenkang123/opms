package customer

import (
	"opms/controllers"
	"strings"
	. "opms/models/customer"
	"opms/utils"
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
	客户添加
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
	sexArray := GetSex()
	this.Data["sexArray"] = sexArray
	//获取客户类型列表
	customerTypeArray := GetCustomerType()
	this.Data["customerTypeArray"] = customerTypeArray
	this.TplName = "customer/customer-add.tpl"
}

type SubmitDataCustomerController struct {
	controllers.BaseController
}

func (this *SubmitDataCustomerController) Post() {
	if !strings.Contains(this.GetSession("userPermission").(string), "permission-manage") {
		this.Abort("401")
	}
	//真实姓名
	realName := this.GetString("realName")
	if "" == realName {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorRealName}
		this.ServeJSON()
	}
	//客户性别
	sex := this.GetString("sex")
	if "" == sex {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorSex}
		this.ServeJSON()
	}
	//客户出生日期
	birth := this.GetString("birth")
	if "" == birth {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorBirth}
		this.ServeJSON()
	}
	//客户邮箱
	email := this.GetString("email")
	if "" == email {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorEmail}
		this.ServeJSON()
	}

	//客户微信
	wechat := this.GetString("wechat")
	if "" == wechat {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorWeChat}
		this.ServeJSON()
	}
	//客户QQ
	qq := this.GetString("qq")
	if "" == qq {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorQQ}
		this.ServeJSON()
	}

	//客户类型
	customerType := this.GetString("customerType")
	if "" == customerType {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorCustomerType}
		this.ServeJSON()
	}
	//客户手机号
	phone := this.GetString("phone")
	if "" == phone {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorPhone}
		this.ServeJSON()
	}
	//客户固定电话
	tel := this.GetString("tel")
	if "" == tel {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorTel}
		this.ServeJSON()
	}
	///客户地址
	address := this.GetString("address")
	if "" == address {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorAddress}
		this.ServeJSON()
	}

	var customer Customer

	customer.CustomerId = utils.SnowFlakeId()
	customer.UpdateTime = utils.GetDateString()
	customer.QQ = qq
	customer.Phone = phone
	customer.Email = email
	customer.CustomerType = customerType
	customer.Birth = birth
	customer.Address = address
	customer.Photo = address
	customer.InsertTime = utils.GetDateString()
	customer.IsActive = true
	customer.Sex = sex
	customer.RealName = realName
	customer.Tel = tel
	customer.Webchat = wechat
	AddCustomer(customer)
	this.TplName = "contract/index.tpl"

}

const (
	errorRealName string = "请填写真实姓名 !"

	errorSex string = "请选择客户性别 !"

	errorBirth string = "请填写客户出生日期 !"

	errorEmail string = "请填写客户邮箱 !"

	errorWeChat string = "请填写客户微信 !"

	errorQQ string = "请填写客户QQ !"

	errorCustomerType string = "请填写客户微信 !"

	errorPhone string = "请填写客户电话 !"

	errorTel string = "请填写客户固定电话 !"

	errorAddress string = "请填写客户固定电话 !"
)
