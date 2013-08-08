package googleanalytics

import (
	"github.com/bborbe/server/renderer"
	"github.com/bborbe/server/renderer/content"
	"io"
)

const CONTENT string = `<script language="javascript" type="text/javascript">
var _gaq = _gaq || [];
_gaq.push(['_setAccount', 'UA-325476-3']);
_gaq.push(['_trackPageview']);
(function() {
var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
})();
</script>`

type GoogleanalyticsRenderer interface {
	renderer.Renderer
}

type googleanalyticsRenderer struct {
	renderer renderer.Renderer
}

func NewGoogleanalyticsRenderer() *googleanalyticsRenderer {
	v := new(googleanalyticsRenderer)
	contentRenderer := content.NewContentRenderer()
	contentRenderer.SetContent(CONTENT)
	v.renderer = contentRenderer
	return v
}

func (v *googleanalyticsRenderer) Render(writer io.Writer) error {
	return v.renderer.Render(writer)
}
