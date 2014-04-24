{{template "admin/header.tpl" .}}

<div id="body">
<form class="entry" action="/admin/entries/{{.Entry.Id}}" method="POST">
	<input class="entry-title" type="text" name="title" value="{{.Entry.Title}}" placeholder="Title">
	<input class="entry-subtitle" type="text" name="subtitle" value="{{.Entry.Subtitle}}" placeholder="Subtitle">
	<textarea class="entry-content-edit" type="text" name="content" placeholder="Content">{{.Entry.Content}}</textarea>
	<div class="entry-content"></div>

	<input type="submit" value="POST">
</form>
</div>

{{template "admin/footer.tpl" .}}