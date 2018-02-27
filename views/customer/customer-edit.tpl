<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <form class="form-horizontal adminex-form" id="contract-form" method="post" action="/customer/submitData" >
                <header><b> 客户信息 </b></header>
                <!--客户真实姓名 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户真实姓名</label>
                  <div class="col-sm-4">
                    <input type="text" name="realName"  value="{{.customer.RealName}}" class="form-control" required placeholder="请填写客户真实姓名!">
                  </div>
                </div>
                <!--客户性别 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户性别</label>
                  <div class="col-sm-4">
                    <select id="sex" class="form-control"  name="sex" required >
                      <option value="">请选择客户性别</option>
                       {{range .sexArray}}
                          <option {{if eq .customer.Sex .Code }}selected{{end}}  value="{{ .Code}}">{{.Desc}}</option>
                        {{end}}
                    </select>
                  </div>
                </div>
                <!--客户出生日期 -->
                <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label" >客户出生日期</label>
                    <div class="col-sm-4">
                       <input type="text" name="birth"  required value="{{.customer.Birth}}" class="form-control default-date-picker" placeholder="请输入客户出生日期">
                    </div>
                 </div>
                  <!--客户邮箱 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户邮箱</label>
                  <div class="col-sm-4">
                    <input type="text" name="email"  value="{{.customer.Email}}" class="form-control" required placeholder="请填写客户邮箱!">
                  </div>
                </div>
                 <!--客户微信 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户微信</label>
                  <div class="col-sm-4">
                    <input type="text" name="wechat"  value="{{.customer.Webchat}}" class="form-control" required placeholder="请填写客户微信!">
                  </div>
                </div>
                <!--客户qq -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户QQ</label>
                  <div class="col-sm-4">
                    <input type="text" name="qq"  value="{{.customer.QQ}}" class="form-control" required placeholder="请填写客户qq!">
                  </div>
                </div>
                 <!--客户类型 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户类型</label>
                  <div class="col-sm-4">
                    <select id="customerType" class="form-control"  name="customerType" required >
                      <option value="">请选择客户类型</option>

                    </select>
                  </div>
                </div>
                <!--客户手机号 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户手机号</label>
                  <div class="col-sm-4">
                    <input type="text" name="phone"  value="" class="form-control" required placeholder="请填写客户手机号!">
                  </div>
                </div>
                 <!--客户固定电话 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户固定电话</label>
                  <div class="col-sm-4">
                    <input type="text" name="tel"  value="" class="form-control" required placeholder="请填写客户固定电话!">
                  </div>
                </div>

                <!--客户地址 -->
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">客户地址</label>
                  <div class="col-sm-4">
                    <input type="text" name="address"  value="" class="form-control" required placeholder="请填写客户地址!">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script src="/static/js/datepicker-zh-CN.js"></script>
<script>
$(function(){
	$('.default-date-picker').datepicker('option', $.datepicker.regional['zh-CN']);
	$('.default-date-picker').datepicker({
        dateFormat: 'yy-mm-dd',
		changeMonth: true,
		changeYear: true,
		yearRange:'-60:+0'
    });
})
</script>
</body>
</html>
