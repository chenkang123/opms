package contract

import (
	"opms/controllers"
	"strings"
	"opms/utils"
	. "opms/models/contract"
)

const (
	errorContractName string = "请填写合同名称"

	errorContractContent string = "请填写合同内容"

	errorStartTime string = "请填写合同开始时间"

	errorEndTime string = "请填写合同结束时间"
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

type SubmitDataContractController struct {
	controllers.BaseController
}

/*
	method提交方法
 */
func (this *SubmitDataContractController) Post() {
	/***
		合同名称
	 */
	contractName := this.GetString("contractName")
	if "" == contractName {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorContractName}
		this.ServeJSON()
	}

	/***
	合同内容
	*/
	contractContent := this.GetString("contractContent")

	if "" == contractContent {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorContractContent}
		this.ServeJSON()
	}
	/***
		合同开始时间
		*/
	startTime := this.GetString("startTime")

	if "" == startTime {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorStartTime}
		this.ServeJSON()
	}
	/***
			合同结束时间
			*/
	endTime := this.GetString("endTime")

	if "" == endTime {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": errorEndTime}
		this.ServeJSON()
	}

	//插入db
	var contract Contract
	contractId := utils.SnowFlakeId()
	contract.ContractId = contractId
	contract.ContractName = contractName
	contract.ContractContent = contractContent
	contract.StartTime = startTime
	contract.EndTime = endTime
	contract.Photo = endTime
	contract.InsertTime = endTime




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
