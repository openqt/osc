package ui_test

import (
	"testing"

	"github.com/openqt/osc/internal/config"
	"github.com/openqt/osc/internal/ui"
	"github.com/stretchr/testify/assert"
)

func TestNewSplash(t *testing.T) {
	s := ui.NewSplash(config.NewStyles(), "bozo")

	x, y, w, h := s.GetRect()
	assert.Equal(t, 0, x)
	assert.Equal(t, 0, y)
	assert.Equal(t, 15, w)
	assert.Equal(t, 10, h)
}
