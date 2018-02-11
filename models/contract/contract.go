package contract

import (
	"github.com/astaxie/beego/orm"
	"opms/models"
)

/********

	Contract bean

 */
type Contract struct {
	ContractId      int64  `orm:"pk;column(contract_id)"`
	ContractName    string `orm:"column(contract_name)"`
	ContractContent string `orm:"column(contract_content)"`
	Photo           string `orm:"column(photo)"`
	StartTime       string `orm:"column(start_time)"`
	EndTime         string `orm:"column(end_time)"`
	InsertTime      string `orm:"column(insert_time)"`
	UpdateTime      string `orm:"column(update_time)"`
	IsActive        bool   `orm:"column(is_active)"`
}

func (this *Contract) TableName() string {
	return models.TableName("contract")
}

func init() {
	orm.RegisterModel(new(Contract))
}

func AddContract(pmscontract Contract) error {

	o := orm.NewOrm()
	contract := new(Contract)
	contract.ContractId = pmscontract. ContractId
	contract.ContractName = pmscontract. ContractName
	contract.ContractContent = pmscontract. ContractContent
	contract.StartTime = pmscontract. StartTime
	contract.EndTime = pmscontract. EndTime
	contract.Photo = pmscontract. Photo
	contract.InsertTime = pmscontract.InsertTime
	contract.UpdateTime = pmscontract.UpdateTime
	contract.IsActive = pmscontract.IsActive
	_, err := o.Insert(contract)
	return err

}
