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
              <form class="form-horizontal adminex-form" id="contract-form" type="post" action="" >
                <header><b> 合同信息 </b></header>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">合同名称</label>
                  <div class="col-sm-4">
                    <input type="text" name="contractName"  value="" class="form-control" required placeholder="请填写合同名称!">
                  </div>
                </div>
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label">合同内容</label>
                  <div class="col-sm-4">
                    <textarea required rows="5" cols="4000" name="contractContent" style="resize:none;" value="" class="form-control" placeholder="请填写合同内容"></textarea>
                  </div>
                </div>
                <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label" >合同开始时间</label>
                    <div class="col-sm-4">
                       <input type="text" name="startTime"  required value="" class="form-control default-date-picker" placeholder="请输入合同开始时间">
                    </div>
                 </div>
                <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label">合同结束时间</label>
                    <div class="col-sm-4">
                       <input required type="text" name="endTime"  value="" class="form-control default-date-picker" placeholder="请输入合同开始时间">
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
