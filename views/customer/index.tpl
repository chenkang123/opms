<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <form class="searchform" action="/contract/manage" method="get">

        <button type="submit" class="btn btn-primary">搜索</button>
      </form>
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 客户管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/contract/manage">客户管理</a> </li>
        <li class="active"> 客户 </li>
      </ul>
      <div class="pull-right"><a href="/customer/add" class="btn btn-success">+添加新客户</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 客户管理 / 总数：{{.countCustomer}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              <!--a href="javascript:;" class="fa fa-times"></a-->
              </span> </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="contract-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>姓名</th>
                        <th>性别</th>
                        <th>出生日期</th>
                         <th>邮箱</th>
						<th>手机号</th>
                      </tr>
                    </thead>
                    <tbody>
                     {{range $k,$v := .customers}}
                     <tr>
                        <td>{{$v.RealName}}</td>
                     	 <td>{{if eq "1" $v.Sex}}男{{else}}女{{end}}</td>
                     	 <td>{{ $v.Birth}}</td>
                     	 <td>{{ $v.Email}}</td>
                     	 <td>{{$v.Phone}}</td>
                     </tr>
                    {{end}}
                    </tbody>
                  </table>
                </form>
                {{template "inc/page.tpl" .}}
				 </section>
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
</body>
</html>
