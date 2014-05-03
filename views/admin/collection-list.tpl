{{template "admin/header.tpl" .}}

<div class="collection-new">
  <a href="/admin/collections/new">New</a>
</div>

<table class="collection-list">
  <thead>
    <th width="50%">Title</th>
    <th>Author</th>
    <th>Edit</th>
    <th>Delete</th>
  </thead>
  <tbody>
    {{range $index, $c := .Collections}}
    <tr>
      <td class="collection-title">{{$c.Title}}</td>
      <td>{{$c.Author}}</td>
      <td><a href="/admin/collections/{{$c.Id}}">Edit</a></td>
      <td><a href="#">Delete</a></td>
    </tr>
    {{end}}
  </tbody>
</table>

{{template "admin/footer.tpl" .}}