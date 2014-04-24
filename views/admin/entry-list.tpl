{{template "admin/header.tpl" .}}

<table class="entry-list">
	<thead>
		<th width="50%">Title</th>
		<th>Date</th>
		<th>Author</th>
		<th>Edit</th>
		<th>Delete</th>
	</thead>
	<tbody>
		{{range $index, $e := .Entries}}
		<tr>
			<td>{{$e.Title}}</td>
			<td>{{$e.Date}}</td>
			<td>{{$e.Author}}</td>
			<td><a href="/admin/entries/{{$e.Id}}">Edit</a></td>
			<td><a href="#">Delete</a></td>
		</tr>
		{{end}}
	</tbody>
</table>

{{template "admin/footer.tpl" .}}