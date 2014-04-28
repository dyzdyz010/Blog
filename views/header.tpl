<!doctype html>
<html>
<head>
  <title>{{.Title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />

  <!-- Web Fonts -->
 <link href='http://fonts.googleapis.com/css?family=Cabin|Open+Sans:400italic,400,600,300' rel='stylesheet' type='text/css'>

  <!-- Style Sheets -->
  <link rel="stylesheet" type="text/css" href="/static/css/app.css">
</head>
<body>
<div id="site">
  <header>
    <h1 class="blog-title">{{.Title}}</h1>
    <p class="blog-subtitle">{{.Subtitle}}</p>

    <nav class="main-nav">
      <ul>
        <li><a class="{{if .HomeActive}}       active {{end}}" href="/">Home</a></li>
        <li><a class="{{if .CollectionActive}} active {{end}}" href="#">Collections</a></li>
        <li><a class="{{if .AboutActive}}      active {{end}}" href="#">About</a></li>
      </ul>
    </nav>
  </header>

  <div id="body">