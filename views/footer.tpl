  </div> <!-- #body -->

  {{if .PageNav}}
  {{template "pagination.tpl" .}}
  {{end}}

  <footer id="footer">
    <div class="social">
      <a class="icon-social-github" href="#"></a>
      <a class="icon-social-dribbble" href="#"></a>
      <a class="icon-social-linkedin" href="#"></a>
    </div>
    <div class="links">
      <p>
        <span>© 2014 </span>
        <a href="#">blog.duyizhuo.com</a>
         | 
        <a href="#">About</a>
         | 
        <a href="mailto:dyzdyz010@sina.com">Contact</a>
         | 
        <a href="#">RSS</a>
      </p>
    </div>
  </footer>
</div> <!-- #site -->

<script type="text/javascript" src="/static/bower_components/jquery/dist/jquery.min.js"></script>
{{if .MarkdownEnabled}}
<script type="text/javascript" src="//cdn.staticfile.org/marked/0.3.2/marked.min.js"></script>
<script type="text/javascript" src="/static/js/highlight.pack.js"></script>
<script type="text/javascript" src="/static/js/markdown-front.js"></script>
{{end}}
</body>
</html>