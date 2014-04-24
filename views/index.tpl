{{template "header.tpl" .}}

{{range $index, $e := .Entries}}
<div class="row entry">
  <h2 class="entry-title"><a href="/entry/{{$e.Id}}">{{$e.Title}}</a></h2>
  <h3 class="entry-subtitle">{{$e.Subtitle}}</h3>

  <div class="entry-meta">
    <span class="entry-date"><a href=""><i class="icon-calendar"></i>{{$e.Date}}</a></span>
    <span class="delimiter">/</span>
    <span class="entry-author"><a href=""><i class="icon-author"></i>{{$e.Author}}</a></span>

    <div class="like-share">
      <span class="entry-like"><a href=""><i class="icon-like"></i>{{$e.Likes}} Likes</a></span>
      <span class="entry-share"><a href=""><i class="icon-share"></i>Share</a></span>
    </div>
  </div>
</div>
{{end}}

{{template "footer.tpl" .}}