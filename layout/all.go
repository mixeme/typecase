package layout

// All is the layout registry for the current step.
func All() Spec {
	return Spec{
		Grids: []Grid{TickRow, RowExcerpt},
		Boxes: tickRowBoxes,
	}
}
