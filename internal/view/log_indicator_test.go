package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/config"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestLogIndicatorRefresh(t *testing.T) {
	defaults := config.NewStyles()
	v := view.NewLogIndicator(config.NewConfig(nil), defaults)
	v.Refresh()

	assert.Equal(t, "[::b]Autoscroll:On     [::b]FullScreen:Off     [::b]Timestamps:Off     [::b]Wrap:Off\n", v.GetText(false))
}
