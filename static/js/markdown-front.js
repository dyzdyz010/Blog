function markdown () {
	// console.log($(".entry-content").html());
	var markedStr = marked($(".entry-content").html());
	var markedDom = $(markedStr);
	$(".entry-content").html(markedDom);
	// console.log(markedDom);
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
	markdown();
});
