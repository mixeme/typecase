// Maintained by hand; step 7 checks it against widget/*.go (go/parser).

package refs

// Widgets lists the zero value of every exported widget type with
// Render() string. Steps 3 and 7 work from it; families add their types.
var Widgets = []Renderer{}

// Parts lists the zero value of every exported part type: a piece of a
// widget that is not a widget of its own (no block, no Render), such as a
// field's hint. The pair check looks into them; families add theirs.
var Parts = []any{}
