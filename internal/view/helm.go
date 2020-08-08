package view

import (
	"context"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/render"
	"github.com/openqt/osc/internal/ui"
	"github.com/gdamore/tcell"
)

// Helm represents a helm chart view.
type Helm struct {
	ResourceViewer
}

// NewHelm returns a new alias view.
func NewHelm(gvr client.GVR) ResourceViewer {
	c := Helm{
		ResourceViewer: NewBrowser(gvr),
	}
	c.GetTable().SetColorerFn(render.Helm{}.ColorerFunc())
	c.GetTable().SetBorderFocusColor(tcell.ColorMediumSpringGreen)
	c.GetTable().SetSelectedStyle(tcell.ColorWhite, tcell.ColorMediumSpringGreen, tcell.AttrNone)
	c.SetBindKeysFn(c.bindKeys)
	c.SetContextFn(c.chartContext)

	return &c
}

func (c *Helm) chartContext(ctx context.Context) context.Context {
	return ctx
}

func (c *Helm) bindKeys(aa ui.KeyActions) {
	aa.Delete(ui.KeyShiftA, ui.KeyShiftN, tcell.KeyCtrlS, tcell.KeyCtrlSpace, ui.KeySpace)
	aa.Add(ui.KeyActions{
		ui.KeyShiftN: ui.NewKeyAction("Sort Name", c.GetTable().SortColCmd(nameCol, true), false),
		ui.KeyShiftS: ui.NewKeyAction("Sort Status", c.GetTable().SortColCmd(statusCol, true), false),
		ui.KeyShiftA: ui.NewKeyAction("Sort Age", c.GetTable().SortColCmd(ageCol, true), false),
	})
}
