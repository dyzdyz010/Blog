<!doctype html>
<html>
<head>
  <title>{{.Title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />

  <!-- Web Fonts -->
  <link href='http://fonts.googleapis.com/css?family=Cabin|Open+Sans:400,300' rel='stylesheet' type='text/css'>

  <!-- Style Sheets -->
  <link rel="stylesheet" type="text/css" href="/static/css/app.css">
  {{if .MarkdownEnabled}}
  <link rel="stylesheet" type="text/css" href="/static/css/monokai_sublime.css">
  {{end}}
</head>
<body>
<div id="site">
  <header>
    <h1 class="blog-title">{{.Title}}</h1>
    <p class="blog-subtitle">{{.Subtitle}}</p>

    <nav class="main-nav">
      <ul>
        <li><a class="{{if .DashboardActive}}  active {{end}}" href="/admin">Dashboard</a></li>
        <li><a class="{{if .EntryActive}}      active {{end}}" href="/admin/entries">Entries</a></li>
        <li><a class="{{if .CollectionActive}} active {{end}}" href="#">Collections</a></li>
        <li><a class="{{if .AboutActive}}      active {{end}}" href="#">About</a></li>
      </ul>
    </nav>
  </header>

  <div id="admin">