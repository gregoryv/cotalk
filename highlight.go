package cotalk

import (
	"regexp"

	. "github.com/gregoryv/web"
)

// highlight go source code
func highlight(v string) string {
	v = keywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

// highlightGoDoc output
func highlightGoDoc(v string) string {
	v = docKeywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

var types = regexp.MustCompile(`(\W)(\w+\.\w+)(\)|\n)`)
var keywords = regexp.MustCompile(`(\W?)(^package|select|default|defer|import|for|func|range|return|go|var|switch|if|case|label|type|struct|interface)(\W)`)
var docKeywords = regexp.MustCompile(`(\W?)(^package|func|type|struct|interface)(\W)`)
var comments = regexp.MustCompile(`(//[^\n]*)`)

func highlightColors() *CSS {
	css := NewCSS()
	css.Style(".keyword", "color: darkviolet")
	css.Style(".type", "color: dodgerblue")
	css.Style(".comment, .comment>span", "color: darkgreen")
	return css
}
