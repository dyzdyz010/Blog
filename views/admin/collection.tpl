{{template "admin/header.tpl" .}}

<div id="body">
<form class="collection" action="/admin/collections/{{.PostId}}" method="POST">
  <input class="collection-title" type="text" name="title" value="{{.Collection.Title}}" placeholder="Title">
  <input class="collection-subtitle" type="text" name="subtitle" value="{{.Collection.Subtitle}}" placeholder="Subtitle">

  <div class="msg">{{.Message}}</div>
  <input class="post-button" type="submit" value="POST">
</form>
</div>

{{template "admin/footer.tpl" .}}