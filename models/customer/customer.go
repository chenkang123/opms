package customer

import (
	"github.com/astaxie/beego/orm"
	"opms/models"
)

/********

	Contract bean

 */
type Customer struct {
	CustomerId   int64  `orm:"pk;column(customer_id)"`
	RealName     string `orm:"column(realNmae)"`
	Sex          string `orm:"column(sex)"`
	Birth        string `orm:"column(birth)"`
	Email        string `orm:"column(email)"`
	Webchat      string `orm:"column(webchat)"`
	QQ           string `orm:"column(qq)"`
	Phone        string `orm:"column(phone)"`
	Tel          string `orm:"column(tel)"`
	Address      string `orm:"column(address)"`
	Photo        string `orm:"column(photo)"`
	CustomerType string `orm:"column(customerType)"`
	InsertTime   string `orm:"column(insertTime)"`
	UpdateTime   string `orm:"column(updateTime)"`
	IsActive     bool   `orm:"column(isActive)"`
}

type Sex struct {
	Code string
	Desc string
}

type CustomerType struct {
	Code string
	Desc string
}

/****
	获取性别列表
 */
func GetSex() ([2]Sex) {
	var male Sex
	var female Sex
	male.Code = "1"
	male.Desc = "男"
	female.Code = "2"
	female.Desc = "女"
	var result [2]Sex
	result[0] = male
	result[1] = female
	return result
}

/****
获取客户类型列表
 */
func GetCustomerType() ([3]CustomerType) {
	//客户类型，1-卖家，2-买家,3-中介
	var buy CustomerType
	var sell CustomerType
	var agency CustomerType
	sell.Code = "1"
	sell.Desc = "卖家"
	buy.Code = "2"
	buy.Desc = "买家"

	agency.Code = "3"
	agency.Desc = "中介"

	var result [3]CustomerType
	result[0] = buy
	result[1] = sell
	result[2] = agency
	return result
}

func (this *Customer) TableName() string {
	return models.TableName("customer")
}

func init() {
	orm.RegisterModel(new(Customer))
}
