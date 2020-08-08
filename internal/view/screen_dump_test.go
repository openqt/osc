package view_test

import (
	"testing"

	"github.com/openqt/osc/internal/client"
	"github.com/openqt/osc/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestScreenDumpNew(t *testing.T) {
	po := view.NewScreenDump(client.NewGVR("screendumps"))

	assert.Nil(t, po.Init(makeCtx()))
	assert.Equal(t, "ScreenDumps", po.Name())
	assert.Equal(t, 5, len(po.Hints()))
}
