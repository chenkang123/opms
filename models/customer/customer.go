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
	code string
	desc string
}

type CustomerType struct {
	code string
	desc string
}

/****
	获取性别列表
 */
func GetSex() ([]Sex) {
	var male Sex
	var female Sex
	male.code = "1"
	male.desc = "男"
	female.code = "2"
	female.desc = "女"
	var result []Sex
	result[0] = male
	result[1] = female
	return result
}

/****
获取客户类型列表
 */
func GetCustomerType() ([]CustomerType) {
	//客户类型，1-卖家，2-买家,3-中介
	var buy CustomerType
	var sell CustomerType
	var agency CustomerType
	sell.code = "1"
	sell.desc = "卖家"
	buy.code = "2"
	buy.desc = "买家"

	agency.code = "3"
	agency.desc = "中介"

	var result []CustomerType
	result[0] = buy
	result[1] = sell
	result[1] = agency
	return result
}

func (this *Customer) TableName() string {
	return models.TableName("customer")
}

func init() {
	orm.RegisterModel(new(Customer))
}