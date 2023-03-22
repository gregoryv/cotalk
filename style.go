package cotalk

import . "github.com/gregoryv/web"

func themeOldstyle() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".slide",
		"padding: 10px 20%",
		"text-align: center",
		"height: calc( 100vh - 50px)",
	)
	css.Style(".slide ul, p, pre",
		"text-align: left",
	)

	// navbar
	css.Style(".slide nav",
		"text-align: center",
		"float: left",
		"clear: both",
		"width: 100%",
		"display: block",
		"margin-top: 2em",
	)
	css.Style(".slide nav ul",
		"list-style-type: none",
		"margin: 0 0",
		"padding: 0 0",
		"text-align: center",
	)
	css.Style(".slide nav ul li",
		"margin: 0 1px",
		"display: inline",
	)
	css.Style(".slide nav ul li.current a, .slide nav ul li:hover a",
		"background-color: #e2e2e2",
		"border-radius: 5px",
	)
	css.Style("nav a:link, nav a:visited",
		"color: #727272",
		"padding: 3px 5px",
		"text-decoration: none",
		"cursor: pointer",
	)

	// goish colors
	css.Style("a:link, a:visited",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	css.Style(".header",
		"width: 100%",
		"border-bottom: 1px solid #727272",
		"text-align: right",
		"margin-top: -2em",
		"margin-bottom: 1em",
	)
	css.Style("h1, h2, h3",
		"text-align: center",
	)
	css.Style(".srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"padding-left: 36px",
		"background-color: #fafafa",
		"tab-size: 4",
		"-moz-tab-size: 4",
	)

	css.Style(".filename",
		"float: left",
		"margin-right: 1.6em",
		"margin-top: -1.6em",
	)

	css.Style(".srcfile code",
		"padding: .6em 0 .6em 0",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"background-position: right top",
		"display: block",
		"text-align: left",
	)
	// toc
	css.Style("li.h3",
		"margin-left: 2em",
	)
	css.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
	)
	css.Style(".dark",
		"background-color: #2e2e34",
		"color: aliceblue",
	)
	css.Style(".light",
		"background-color: #ffffff",
		"color: #3b2616",
	)
	css.Style("img.center",
		"display: block",
		"margin-left: auto",
		"margin-right: auto",
	)
	css.Style("img.left",
		"float: left",
		"margin-right: 2em",
	)
	css.Style(".group",
		"float: left",
		"text-align: left",
		"margin-right: 3em",
	)
	css.Style("td",
		"vertical-align: top",
	)
	css.Style("td:nth-child(2)",
		"padding-left: 2em",
	)

	return css
}
