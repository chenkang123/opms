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

func (this *AddContractController) Get() {

	if !strings.Contains(this.GetSession("userPermission").(string), "permission-manage") {
		this.Abort("401")
	}

	this.TplName = "contract/contract-add.tpl"
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
