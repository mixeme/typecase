package layout

// All is the layout registry for the current step.
func All() Spec {
	boxes := append([]Box{}, tickRowBoxes...)
	boxes = append(boxes, rowBoxes...)
	boxes = append(boxes, shellBoxes...)
	boxes = append(boxes, formBoxes...)
	boxes = append(boxes, restBoxes...)
	boxes = append(boxes, colsBoxes...)
	flows := append([]Flow{}, rowFlows...)
	flows = append(flows, shellFlows...)
	flows = append(flows, formFlows...)
	return Spec{
		Grids: []Grid{
			TickRow,
			RowExcerpt,
			RowAvatarCols,
			RowTickCols,
			RowTimeCols,
			RowTickMarkCols,
			RowFind,
			Shell,
			ShellHasSide,
			ShellNoRail,
			ShellNoRailHasSide,
			Doc,
			DocNoMeta,
			Fields,
			ColsHalves,
			ColsThirds,
			ColsSidebar,
		},
		Boxes: boxes,
		Flows: flows,
	}
}
