package rendering

import (
	"html/template"
	"strings"

	"github.com/yznts/kyoto/v3/htmx"

	"github.com/yznts/kyoto/v3/component"
	"github.com/yznts/zen/v3/errorsx"
	"github.com/yznts/zen/v3/mapx"
)

// FuncMap holds a library predefined template functions.
// You have to include it in your template building to use kyoto properly.
var FuncMap = template.FuncMap{
	// Inline render function.
	// Allows to avoid explicit template syntax
	// and customize render behavior.
	"render": func(f component.Future) template.HTML {
		// Await future
		state := f()
		// Check if state implements render
		if r, ok := state.(Renderer); ok {
			// Render
			var out strings.Builder
			errorsx.Must(0, r.Render(state, &out))
			// Pack and return
			return template.HTML(out.String())
		}
		// Panic if state does not implement render
		panic("state does not implement render")
	},
}

// FuncMapAll holds all funcmap instances of kyoto library.
var FuncMapAll = mapx.Merge(
	FuncMap,
	htmx.FuncMap,
	component.FuncMap,
)
