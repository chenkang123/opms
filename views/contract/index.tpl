<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->
      <form class="searchform" action="/message/manage" method="get">

      </form>
      <!--search end-->
      {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h3> 消息管理 </h3>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">OPMS</a> </li>
        <li> <a href="/message/manage">消息管理</a> </li>
        <li class="active"> 消息 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">

    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script>
$(function(){
	//全选　
	$('.checkboxbtn').click(function(){
		$(this).parent().find("input[type='checkbox']").prop('checked', $(this).is(':checked'));   							 
	});
	
	$('#deleteMsg').on('click', function(){
	
		var ck = $('.checked:checked');
		if(ck.length <= 0) { dialogInfo('至少选择一个复选框'); return false; }
		
		var str = '';
		$.each(ck, function(i, n){
			str += n['value']+',';
		});
		str = str.substring(0, str.length - 1)
		$.post('/message/ajax/delete', {ids:str},function(data){
			dialogInfo(data.message)
			if (data.code) {
				setTimeout(function(){ window.location.reload(); }, 2000);
			} else {
				setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
			}			
		},'json');
	});
})
</script>
</body>
</html>
