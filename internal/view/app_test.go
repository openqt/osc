package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/config"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestAppNew(t *testing.T) {
	a := view.NewApp(config.NewConfig(ks{}))
	a.Init("blee", 10)

	assert.Equal(t, 9, len(a.GetActions()))
}
