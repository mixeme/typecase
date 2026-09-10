package layout

// All is the layout registry for the current step.
func All() Spec {
	boxes := append([]Box{}, tickRowBoxes...)
	boxes = append(boxes, rowBoxes...)
	return Spec{
		Grids: []Grid{
			TickRow,
			RowExcerpt,
			RowAvatarCols,
			RowTickCols,
			RowTimeCols,
			RowTickMarkCols,
			RowFind,
		},
		Boxes: boxes,
		Flows: rowFlows,
	}
}
