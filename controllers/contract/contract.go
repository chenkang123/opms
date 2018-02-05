package contract

import (
	"opms/controllers"
	"strings"
)

/***
	合同管理controller
 */
type ManagerContractController struct {
	controllers.BaseController
}

func (this *ManagerContractController) Get() {

	if !strings.Contains(this.GetSession("userPermission").(string), "permission-manage") {
		this.Abort("401")
	}

	this.TplName = "contract/index.tpl"
}

func (this *ManagerContractController) Post() {

}

/***
	合同添加
 */
type AddContractController struct {
	controllers.BaseController
}

/*
	跳转方法
 */
func (this *AddContractController) Get() {
	this.TplName = "contract/contract-add.tpl"
}
/*
	method提交方法
 */
func (this *AddContractController) Post() {

	this.TplName = "contract/index.tpl"
}

/***
	合同编辑
 */
type EditContractController struct {
	controllers.BaseController
}

/***
	合同删除
 */
type DelContractController struct {
	controllers.BaseController
}
