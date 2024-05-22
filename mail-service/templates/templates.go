package templates

import _ "embed"

//go:embed mail.html.gohtml
var MailHTML string

//go:embed mail.plain.gohtml
var MailPlain string
