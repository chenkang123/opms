package customer

import (
	"github.com/astaxie/beego/orm"
	"opms/models"
)

/********

	Contract bean

 */
type Customer struct {

	CustomerId int64 `orm:"pk;column(customer_id)"`
	RealName string `orm:"column(realNmae)"`
	Sex string `orm:"column(sex)"`
	Birth string `orm:"column(birth)"`
	Email string `orm:"column(email)"`
	Webchat string `orm:"column(webchat)"`
	QQ string `orm:"column(qq)"`
	Phone string `orm:"column(phone)"`
	Tel string `orm:"column(tel)"`
	Address string `orm:"column(address)"`
	Photo string `orm:"column(photo)"`
	CustomerType string `orm:"column(customerType)"`
	InsertTime string `orm:"column(insertTime)"`
	UpdateTime string `orm:"column(updateTime)"`
	IsActive bool `orm:"column(isActive)"`
}

func (this *Customer) TableName() string {
	return models.TableName("customer")
}

func init() {
	orm.RegisterModel(new(Customer))
}

