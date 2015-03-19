$(document).ready(function () {
	// Markdown
	marked.setOptions({
		highlight: function (code, lang) {
			if (lang && hljs.getLanguage(lang)) {
				return hljs.highlight(lang, code).value;
			} else {
				return hljs.highlightAuto(code).value;
			}
		}
	});
	markdown();
	$(".entry-content-edit").keyup(markdown);

	// Autosize
	autosize($('.entry-content-edit'));

	// Image upload
	imgUploadInit();

	// MathJax
	MathJax.Hub.Config({tex2jax: {inlineMath: [['$','$']]}});
});

function markdown () {
	// console.log($(".entry-content-edit").val());
	var markedStr = marked($(".entry-content-edit").val());
	var markedDom = $(markedStr);
	$(".entry-content").html(markedDom);
	// console.log(markedDom)
	MathJax.Hub.Typeset();
}