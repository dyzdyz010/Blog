{{template "admin/header.tpl" .}}

<div class="collection-new">
  <a href="/admin/collections/new">New</a>
</div>

<table class="entry-list">
  <thead>
    <th width="50%">Title</th>
    <th>Edit</th>
    <th>Delete</th>
  </thead>
  <tbody>
    {{range $index, $c := .Collections}}
    <tr>
      <td>{{$c.Title}}</td>
      <td><a href="/admin/collections/{{$c.Id}}">Edit</a></td>
      <td><a href="#">Delete</a></td>
    </tr>
    {{end}}
  </tbody>
</table>

{{template "admin/footer.tpl" .}}