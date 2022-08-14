package main

var tags = []string{
	"html",
	"base",
	"head",
	"style",
	"title",
	"address",
	"article",
	"footer",
	"header",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"hgroup",
	"nav",
	"section",
	"dd",
	"div",
	"dl",
	"dt",
	"figcaption",
	"figure",
	"hr",
	"li",
	"main",
	"ol",
	"p",
	"pre",
	"ul",
	"abbr",
	"b",
	"bdi",
	"bdo",
	"br",
	"cite",
	"code",
	"data",
	"dfn",
	"em",
	"i",
	"kbd",
	"mark",
	"q",
	"rp",
	"rt",
	"rtc",
	"ruby",
	"s",
	"samp",
	"small",
	"span",
	"strong",
	"sub",
	"sup",
	"time",
	"u",
	"var",
	"wbr",
	"area",
	"audio",
	"map",
	"track",
	"video",
	"embed",
	"object",
	"param",
	"source",
	"canvas",
	"noscript",
	"script",
	"del",
	"ins",
	"caption",
	"col",
	"colgroup",
	"table",
	"tbody",
	"td",
	"tfoot",
	"th",
	"thead",
	"tr",
	"button",
	"datalist",
	"fieldset",
	"form",
	"input",
	"keygen",
	"label",
	"legend",
	"meter",
	"optgroup",
	"option",
	"output",
	"progress",
	"select",
	"details",
	"dialog",
	"menu",
	"menuitem",
	"summary",
	"content",
	"element",
	"shadow",
	"template",
	"acronym",
	"applet",
	"basefont",
	"big",
	"blink",
	"center",
	"dir",
	"frame",
	"frameset",
	"isindex",
	"listing",
	"noembed",
	"plaintext",
	"spacer",
	"strike",
	"tt",
	"xmp"}

var noStyle = []string{
	"head",
	"title",
	"meta",
	"script",
	"link",
	"style",
	"track",
	"param",
	"source",
	"optgroup"}

var deprecated map[string]string

func makeDeprecated() {
	deprecated = make(map[string]string)
	deprecated["applet"] = "Use the object tag instead."
	deprecated["basefont"] = "No alternative."
	deprecated["big"] = "Use css styles instead."
	deprecated["blink"] = "Use CSS animation with the display property instead."
	deprecated["center"] = "Use the text-align property instead."
	deprecated["dir"] = "Use the ul or ol tag and the list-style property."
	deprecated["embed"] = "Use the object tag instead."
	deprecated["font"] = "Use the font-family property instead."
	deprecated["frame"] = "Use the iframe tag instead."
	deprecated["frameset"] = "No alternative."
	deprecated["isindex"] = "No alternative."
	deprecated["noframes"] = "No alternative."
	deprecated["marquee"] = "It's not the 90's anymore (Use CSS animation with transform properties instead)."
	deprecated["menu"] = "Use the ul tag instead."
	deprecated["plaintext"] = "Use the pre tag instead."
	deprecated["s"] = "Use the text-decoration property if applicable, or use the del tag instead."
	deprecated["strike"] = "Use the text-decoration property if applicable, or use the del tag instead."
	deprecated["tt"] = "Use the font-family property or the kbd, code, or spam tag instead."
	deprecated["u"] = "Use the text-decoration property instead."
	deprecated["xmp"] = "Use the pre tag instead."
}
