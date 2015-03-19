function markdown () {
	// console.log($(".entry-content").html());
	var markedStr = marked($(".entry-content").text());
	var markedDom = $(markedStr);
	$(".entry-content").html(markedDom);
	// console.log(markedDom);
	MathJax.Hub.Typeset();
}

$(document).ready(function () {
	marked.setOptions({
		highlight: function (code, lang) {
			if (lang && hljs.getLanguage(lang)) {
				return hljs.highlight(lang, code).value;
			} else {
				return hljs.highlightAuto(code).value;
			}
		}
	});

	// MathJax config
	MathJax.Hub.Config({tex2jax: {inlineMath: [['$','$']]}});

	markdown();
});
