{{template "admin/header.tpl" .}}

<div id="body">
<form class="login" action="/admin/login" method="POST">
	<input type="text" class="login-name" placeholder="Name">
	<input type="password" class="login-password" placeholder="Password">
	<input class="post-button" type="submit" value="POST">
	<div class="msg">{{.Message}}</div>
</form>
</div>

{{template "admin/footer.tpl" .}}