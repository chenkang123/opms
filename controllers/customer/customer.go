package customer

import (
	"opms/controllers"
	"strings"
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
