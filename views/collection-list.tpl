{{template "header.tpl" .}}

{{range $index, $c := .Collections}}
<div class="row entry">
  <h2 class="entry-title"><a href="/collection/{{$c.Id}}">{{$c.Title}}</a></h2>
  <h3 class="entry-subtitle">{{$c.Subtitle}}</h3>

  <div class="entry-meta">
    <span class="entry-date"><a href=""><i class="icon-calendar"></i>{{$c.Date}}</a></span>
    <span class="delimiter">/</span>
    <span class="entry-author"><a href=""><i class="icon-author"></i>{{$c.Author}}</a></span>
  </div>
</div>
{{end}}

{{template "footer.tpl" .}}