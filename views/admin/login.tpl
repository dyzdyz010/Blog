{{template "admin/header.tpl" .}}

<div id="body">
<form class="login" action="/admin/login" method="POST">
	<input type="text" class="login-name" placeholder="Name" name="name">
	<input type="password" class="login-password" placeholder="Password" name="password">
	<input class="post-button" type="submit" value="LOGIN">
	<div class="msg">{{.Message}}</div>
</form>
</div>

{{template "admin/footer.tpl" .}}