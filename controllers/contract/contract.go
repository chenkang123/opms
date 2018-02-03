package contract


import (
	"opms/controllers"
)

/***
	合同管理controller
 */
type ManagerContractController struct {
	controllers.BaseController
}


func (this *ManagerContractController) Post() {

}

/***
	合同添加
 */
type AddContractController struct {
	controllers.BaseController
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
