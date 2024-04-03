package constants

const (
	HttpsFormatString       = "https://%s"
	Oauth2TcsFormatString   = "https://%s/oauth2/v1/authorize?client_id=%s&response_type=code&state=%s"
	RedirectUrlFormatString = "http://%s:%s/api/v1/auth"

	HtmlTemplatePostMsg = "<!DOCTYPE html>" +
		"<html>" +
		"<head>" +
		"	<meta charset=\"UTF-8\">" +
		"	<title>TrueConf Standard client registration</title>" +
		"</head>" +
		"<body id=\"wl_box\">" +
		"	<script type=\"text/javascript\">" +
		"		window.opener.postMessage(`%s`, \"%s\");" +
		"	</script>  " +
		"</body>" +
		"</html>"
)
