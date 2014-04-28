{{template "header.tpl" .}}

<div class="row entry">
  <h2 class="entry-title">{{.Entry.Title}}</h2>
  <h3 class="entry-subtitle">{{.Entry.Subtitle}}</h3>

  <div class="entry-content">{{.Entry.Content}}</div>

  <div class="entry-meta">
    <span class="entry-date"><a href=""><i class="icon-calendar"></i>{{.Entry.Date}}</a></span>
    <span class="delimiter">/</span>
    <span class="entry-author"><a href=""><i class="icon-author"></i>{{.Entry.Author}}</a></span>

    <div class="like-share">
      <span class="entry-like"><a href=""><i class="icon-like"></i>{{.Entry.Likes}} Likes</a></span>
      <span class="entry-share"><a href=""><i class="icon-share"></i>Share</a></span>
    </div>
  </div>
</div>

{{template "footer.tpl" .}}