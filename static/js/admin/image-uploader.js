var editArea = $('.entry-content-edit');
function imgUploadInit () {
	editArea.on('dragenter', function (e) 
	{
		e.stopPropagation();
		e.preventDefault();
		editArea.css('border', '2px dotted #22B861');
	});
	editArea.on('dragover', function (e) 
	{
		e.stopPropagation();
		e.preventDefault();
	});
	editArea.on('drop', function (e) 
	{
		editArea.css('border', '1px solid #E6E6E6');
		e.preventDefault();
		var files = e.originalEvent.dataTransfer.files;
		// console.log(files);
		handleFileUpload(files,editArea);
	});

	$(document).on('dragenter', function (e) 
	{
		e.stopPropagation();
		e.preventDefault();
	});
	$(document).on('dragover', function (e) 
	{
		e.stopPropagation();
		e.preventDefault();
		editArea.css('border', '2px dotted #DD5757');
	});
	$(document).on('drop', function (e) 
	{
		e.stopPropagation();
		e.preventDefault();
		editArea.css('border', '1px solid #E6E6E6');
	});
}

function handleFileUpload (files, element) {
	$.getJSON('/admin/qiniu/tokens', function (token) {
		console.log(token);
		for (var i = 0; i < files.length; i++) {
			var fd = new FormData();
			fd.append('file', files[i])
			sendFileToServer(fd, token);
		}
	});
}

function sendFileToServer (fd, token) {
	console.log(fd);
	fd.append('token', token);

	var uploadURL = 'http://upload.qiniu.com';
	var jqXHR = $.ajax({
		url: uploadURL,
		type: 'POST',
		processData: false,
		contentType: false,
		cache: false,
		data: fd,
		success: function (data) {
			var imgBaseURL = 'http://7vztwe.com1.z0.glb.clouddn.com/'
			console.log(data);
			var imgURL = imgBaseURL + data.key;
			insertImgURL(imgURL);

			var ta = document.querySelector('textarea');
			var evt = document.createEvent('Event');
			evt.initEvent('autosize.update', true, false);
			ta.dispatchEvent(evt);

			markdown();
		}
	})
}

function insertImgURL (url) {
	var cursorPos = editArea.prop('selectionStart');
	var content = editArea.val();
	var textBefore = content.substring(0, cursorPos);
	var textAfter = content.substring(cursorPos, content.length);
	content = textBefore + '\r\n\r\n![](' + url + ')\r\n\r\n' + textAfter;
	editArea.val(content);
}